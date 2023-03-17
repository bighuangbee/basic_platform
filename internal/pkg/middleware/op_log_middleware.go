package middleware

import (
	"context"
	"encoding/json"
	pbBasic "github.com/bighuangbee/basic-platform/api/basic/v1"
	pb "github.com/bighuangbee/gokit/api/common/v1"
	"github.com/bighuangbee/gokit/model"
	"github.com/bighuangbee/gokit/tools"
	"github.com/bighuangbee/gokit/userAccess"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	nhttp "net/http"
	"strconv"
	"strings"
)

//从proto文件解析router api=>title
var uriTitleMap = make(map[string]model.LogUrlInfoWithKey)

type OpLog struct {
	OpLogGrpcAddr 	string
	OpGrpcCli     	pbBasic.OperationLogClient
	ProtoPath		string
	Logger			log.Logger
}

func NewOpLog(protoPath, opLogGrpcAddr string, logger log.Logger) *OpLog {
	LoadOperationLogWithProto(protoPath, logger)

	return &OpLog{
		ProtoPath: protoPath,
		OpLogGrpcAddr: opLogGrpcAddr,
		Logger: logger,
		//OpGrpcCli: pbBasic.NewOperationLogClient(tools.GetGrpcClient(opLogGrpcAddr)),
	}
}

//后置，避免互依赖
func (r *OpLog) AfterStartOpLogGrpcConn()(func (ctx context.Context) error){
	return func(ctx context.Context) error {
		r.OpGrpcCli = pbBasic.NewOperationLogClient(tools.GetGrpcClient(r.OpLogGrpcAddr))
		r.Logger.Log(log.LevelInfo, "AfterStartOpLogGrpcConn", r.OpLogGrpcAddr)
		return nil
	}
}

func GetHostIp(transport interface{}) (ip string) {
	tr := transport.(*http.Transport)
	req := tr.Request()
	strs := strings.Split(req.Host, ":")
	if len(strs) > 0 {
		remoteIP := net.ParseIP(strs[0])
		if remoteIP == nil {
			return ""
		}
		return remoteIP.String()
	}

	return ""
}

func GetPathTemplateMethodIp(transport interface{}) (uri, rawQuery, method, ip string) {
	tr := transport.(*http.Transport)
	req := tr.Request()
	method = strings.ToUpper(req.Method)
	uri = tr.PathTemplate()
	rawQuery = tr.Request().URL.RawQuery
	ip = GetIP(req)
	return
}

func GetIP(r *nhttp.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}

	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return ""
	}
	remoteIP := net.ParseIP(ip)
	if remoteIP == nil {
		return ""
	}
	return remoteIP.String()
}

var accessToken = userAccess.NewToken()

func (r *OpLog) Save() middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {

				var (
					corpId   uint64
					userId   uint64
					userName string
					account  string
					bodyStr  string
				)
				uri, rawQuery, method, ip := GetPathTemplateMethodIp(tr)

				bs, err := json.Marshal(req)
				if err != nil {
					r.Logger.Log(log.LevelError, "OpLog Save, json.Marshal", err)
				} else {
					bodyStr = string(bs)
				}

				if uri == "/api/v1.0/user/login" {
					obj := make(map[string]interface{})
					json.Unmarshal([]byte(bodyStr), &obj)

					if _, ok := obj["password"]; ok {
						obj["password"] = "***"
					}
					if _, ok := obj["account"]; ok {
						account = obj["account"].(string)
					}
					bs, _ := json.Marshal(obj)
					bodyStr = string(bs)
				}

				if token := tr.RequestHeader().Get("jwtToken"); token != "" {
					userClaims, err :=  accessToken.Decode(token)
					if err != nil {
						r.Logger.Log(log.LevelError,"accessToken.Decode err", err)
					} else {
						corpId = uint64(userClaims.CorpId)
						id, _ := strconv.ParseInt(userClaims.JwtClaims.Id, 10, 64)
						userId = uint64(id)
						userName = userClaims.UserName
						account = userClaims.Account
					}
				}
				userAgent := tr.RequestHeader().Get("User-Agent")

				var sysLog model.SysLogRpc
				if v, ok := uriTitleMap[method+"_"+uri]; ok {
					sysLog = model.SysLogRpc{
						CorpId:    corpId,
						UserId:    uint32(userId),
						Name:      userName,
						LoginName: account,
						Type:      v.LogType,
						IP:        ip,
						Url:       uri+"?"+rawQuery,
						UrlTitle:  v.Title,
						Params:    bodyStr,
						UserAgent: userAgent,
					}
				}

				details, _ := json.Marshal(&sysLog)
				//operationName := getLogTypeStr(v.LogType)
				now := timestamppb.Now()
				if _, err := r.OpGrpcCli.Add(ctx, &pbBasic.AddRequest{
					Log: &pbBasic.Log{
						CorpId:        int32(corpId),
						UserId:        int32(userId),
						UserName:      userName,
						OperationModule: uri,
						OperationName: sysLog.UrlTitle,
						OperationType: int32(sysLog.Type),
						Detail:        string(details),
						Timestamp:     now,
						CreatedAt:     now,
					},
				}); err != nil {
					r.Logger.Log(log.LevelError, "创建操作日志失败 OpGrpcCli.Add", err)
				}
				r.Logger.Log(log.LevelInfo, "method", method, "uri", sysLog.Url, "title", sysLog.UrlTitle, "logType", sysLog.Type, "logType", getLogTypeStr(sysLog.Type), "corpId", corpId, "userId", userId, "userName", userName, "account", account, "ip", ip, "payload", bodyStr)
				return handler(ctx, req)
			}
			return nil, pb.ErrorUnauthenticated("CheckToken failed.")
		}
	}
}

//0=未知，1=创建，2=查看，3=编辑，4=删除，5=登陆，6=登出，7=导出，8=导入，9=保存
func getLogTypeStr(logType model.LogType) string {
	switch logType {
	case model.LogDefault:
		return "未知操作类型"
	case model.LogCreate:
		return "创建"
	case model.LogView:
		return "查看"
	case model.LogEdit:
		return "编辑"
	case model.LogDelete:
		return "删除"
	case model.LogLogin:
		return "登陆"
	case model.LogLogout:
		return "登出"
	case model.LogExport:
		return "导出"
	case model.LogImport:
		return "导入"
	case model.LogSave:
		return "保存"
	}
	return "未知操作类型"
}
