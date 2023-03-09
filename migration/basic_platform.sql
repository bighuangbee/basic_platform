-- auto-generated definition
create table operation_log
(
    id               bigint unsigned auto_increment comment '主键id'
        primary key,
    corp_id          int                    null comment '企业id',
    corp_name        varchar(50)            null comment '企业名称',
    user_id          bigint                 null comment '操作人id',
    user_name        varchar(30) default '' null comment '操作人姓名',
    operation_name   varchar(50)            null comment '操作的名称',
    operation_module varchar(50)            null comment '操作的模块',
    operation_type   tinyint                null comment '操作类型，0=未知，1=创建，2=查看，3=编辑，4=删除，5=登陆，6=登出，7=导出，8=导入，9=保存',
    detail           varchar(2000)          null comment '详细内容',
    timestamp        timestamp              null comment '操作时间',
    created_at       timestamp              null comment '创建时间',
    status           tinyint                null comment '状态',
    reason           varchar(200)           null comment '请求失败的原因'
)
    comment '操作日志' charset = utf8mb3;

