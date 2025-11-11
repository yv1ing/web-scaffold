package constant

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/11/11 14:31
// @Desc:	系统信息码定义

/*
# 信息码制定规范

基本格式：模块编码（2位10进制数）| 信息类型（1位10进制数）| 信息编码（3位10进制数）

## 模块编码划分

- 核心层：10
- 数据层：11
- 服务层：12
- 接口层：13
- 其它类：14

## 类型信息划分

- 基本通用信息：0
- 数据库类信息：1
- 编解码类信息：2
- 权限校验信息：3
*/

/* 核心层信息编码 */
const (
	CORE_INIT_CONF_ERROR   = 100001 // 初始化系统配置失败
	CORE_INIT_ENGINE_ERROR = 100002 // 初始化Web引擎失败
	CORE_INIT_USER_ERROR   = 100003 // 初始化系统用户失败
)

/* 数据层信息编码 */
const (
	DATA_INIT_DATABASE_ERROR   = 111001 // 初始化数据库失败
	DATA_INIT_REPOSITORY_ERROR = 111002 // 初始化仓储层失败
)

/* 服务层信息编码 */
const (
	SERVICE_CREATE_DATA_ERROR = 121001 // 新增数据失败
	SERVICE_UPDATE_DATA_ERROR = 121002 // 更新数据失败
	SERVICE_DELETE_DATA_ERROR = 121003 // 删除数据失败
	SERVICE_FIND_DATA_ERROR   = 121004 // 查询数据失败
)

/* 接口层信息编码 */
const (
	API_OPERATE_SUCCESS = 130000 // 接口操作成功
	API_INTERNAL_ERROR  = 130001 // 系统内部错误

	API_CREATE_DATA_ERROR = 131001 // 新增数据失败
	API_UPDATE_DATA_ERROR = 131002 // 更新数据失败
	API_DELETE_DATA_ERROR = 131003 // 删除数据失败
	API_FIND_DATA_ERROR   = 131004 // 查询数据失败

	API_INVALID_REQUEST_HEADER = 132001 // 非法请求头
	API_INVALID_REQUEST_PARAMS = 132002 // 非法请求参数

	API_PERMISSION_DENIED = 133001 // 权限验证失败
)
