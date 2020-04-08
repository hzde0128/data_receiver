package errcode

import (
	"errors"
)

type ECode int

// 错误码
const (
	SUCCESS       ECode = 0
	FAILURE       ECode = -1
	TOKEN_INVAILD ECode = 401
	SYSTEM_ERROR  ECode = 500

	// 参数错误
	REQUEST_ERROR            ECode = 10000
	REQUEST_PARAM_INVAILD    ECode = 10001
	REQUEST_PARAM_IS_NULL    ECode = 10002
	REQUEST_PARAM_TYPE_ERROR ECode = 10003
	REQUEST_PARAM_MISS       ECode = 10004

	// 用户错误
	USER_NOT_LOGIN               ECode = 20001
	USER_NOEXIST_OR_PASSWD_ERROR ECode = 20002
	USER_IS_DISABLED             ECode = 20003
	USER_NOT_EXIST               ECode = 20004
	USER_HAS_EXIST               ECode = 20005
	USER_AUTH_HAS_EXIST          ECode = 20006

	// 业务错误
	BUSINESS_CREATE_ERROR ECode = 30001

	// 系统错误
	SYSTEM_BUSY ECode = 40001

	// 数据错误
	DATA_NOT_FOUND ECode = 50001
	DATA_ERROR     ECode = 50002
	DATA_HAS_EXIST ECode = 50003

	// 服务错误
	SERVICE_INTER_CALL_ERROR ECode = 60001
	SERVICE_OUTER_CALL_ERROR ECode = 60002
	SERVICE_ACCESS_DENY      ECode = 60003
	SERVICE_INVALID          ECode = 60004
	SERVICE_REQUEST_TIMEOUT  ECode = 60005
	SERVICE_FULLLOAD         ECode = 60006

	// 权限错误
	PERMISSION_ONLY_OWNER    ECode = 70001
	PERMISSION_CANNOT_DELETE ECode = 70002
)

// 错误描述
var (
	ErrorDesc_SUCCESS       = errors.New("成功")
	ErrorDesc_FAILURE       = errors.New("失败")
	ErrorDesc_TOKEN_INVAILD = errors.New("令牌失效")
	ErrorDesc_SYSTEM_ERROR  = errors.New("服务器错误")

	// 参数错误
	ErrorDesc_REQUEST_ERROR            = errors.New("请求错误")
	ErrorDesc_REQUEST_PARAM_INVAILD    = errors.New("参数无效")
	ErrorDesc_REQUEST_PARAM_IS_NULL    = errors.New("参数为空")
	ErrorDesc_REQUEST_PARAM_TYPE_ERROR = errors.New("参数类型错误")
	ErrorDesc_REQUEST_PARAM_MISS       = errors.New("参数缺失")

	// 用户错误
	ErrorDesc_USER_NOT_LOGIN               = errors.New("用户未登录")
	ErrorDesc_USER_NOEXIST_OR_PASSWD_ERROR = errors.New("账号不存在或密码错误")
	ErrorDesc_USER_IS_DISABLED             = errors.New("账号已被禁用")
	ErrorDesc_USER_NOT_EXIST               = errors.New("用户不存在")
	ErrorDesc_USER_HAS_EXIST               = errors.New("用户已存在")
	ErrorDesc_USER_AUTH_HAS_EXIST          = errors.New("认证已存在")

	// 业务错误
	ErrorDesc_BUSINESS_CREATE_ERROR = errors.New("创建失败")

	// 系统错误
	ErrorDesc_SYSTEM_BUSY = errors.New("系统繁忙，请稍后重试")

	// 数据错误
	ErrorDesc_DATA_NOT_FOUND = errors.New("数据未找到")
	ErrorDesc_DATA_ERROR     = errors.New("数据有误")
	ErrorDesc_DATA_HAS_EXIST = errors.New("数据已存在")

	// 服务错误
	ErrorDesc_SERVICE_INTER_CALL_ERROR = errors.New("内部系统接口调用异常")
	ErrorDesc_SERVICE_OUTER_CALL_ERROR = errors.New("外部系统接口调用异常")
	ErrorDesc_SERVICE_ACCESS_DENY      = errors.New("该接口禁止访问")
	ErrorDesc_SERVICE_INVALID          = errors.New("接口地址无效")
	ErrorDesc_SERVICE_REQUEST_TIMEOUT  = errors.New("接口请求超时")
	ErrorDesc_SERVICE_FULLLOAD         = errors.New("接口负载过高")

	// 权限错误
	ErrorDesc_PERMISSION_ONLY_OWNER    = errors.New("只有标签 Owner ,才具备删除权限")
	ErrorDesc_PERMISSION_CANNOT_DELETE = errors.New("此认证标签已有员工认证，不可以进行删除")
)

// http 返回码，http 所有返回码必须从这里出
var (
	ResponseCode = map[ECode]string{
		SUCCESS:       ErrorDesc_SUCCESS.Error(),
		FAILURE:       ErrorDesc_FAILURE.Error(),
		TOKEN_INVAILD: ErrorDesc_TOKEN_INVAILD.Error(),
		SYSTEM_ERROR:  ErrorDesc_SYSTEM_ERROR.Error(),

		// 参数错误
		REQUEST_ERROR:            ErrorDesc_REQUEST_ERROR.Error(),
		REQUEST_PARAM_INVAILD:    ErrorDesc_REQUEST_PARAM_INVAILD.Error(),
		REQUEST_PARAM_IS_NULL:    ErrorDesc_REQUEST_PARAM_IS_NULL.Error(),
		REQUEST_PARAM_TYPE_ERROR: ErrorDesc_REQUEST_PARAM_TYPE_ERROR.Error(),
		REQUEST_PARAM_MISS:       ErrorDesc_REQUEST_PARAM_MISS.Error(),

		// 用户错误
		USER_NOT_LOGIN:               ErrorDesc_USER_NOT_LOGIN.Error(),
		USER_NOEXIST_OR_PASSWD_ERROR: ErrorDesc_USER_NOEXIST_OR_PASSWD_ERROR.Error(),
		USER_IS_DISABLED:             ErrorDesc_USER_IS_DISABLED.Error(),
		USER_NOT_EXIST:               ErrorDesc_USER_NOT_EXIST.Error(),
		USER_HAS_EXIST:               ErrorDesc_USER_HAS_EXIST.Error(),
		USER_AUTH_HAS_EXIST:          ErrorDesc_USER_AUTH_HAS_EXIST.Error(),

		// 业务错误
		BUSINESS_CREATE_ERROR: ErrorDesc_BUSINESS_CREATE_ERROR.Error(),

		// 系统错误
		SYSTEM_BUSY: ErrorDesc_SYSTEM_BUSY.Error(),

		// 数据错误
		DATA_NOT_FOUND: ErrorDesc_DATA_NOT_FOUND.Error(),
		DATA_ERROR:     ErrorDesc_DATA_ERROR.Error(),
		DATA_HAS_EXIST: ErrorDesc_DATA_HAS_EXIST.Error(),

		// 服务错误
		SERVICE_INTER_CALL_ERROR: ErrorDesc_SERVICE_INTER_CALL_ERROR.Error(),
		SERVICE_OUTER_CALL_ERROR: ErrorDesc_SERVICE_OUTER_CALL_ERROR.Error(),
		SERVICE_ACCESS_DENY:      ErrorDesc_SERVICE_ACCESS_DENY.Error(),
		SERVICE_INVALID:          ErrorDesc_SERVICE_INVALID.Error(),
		SERVICE_REQUEST_TIMEOUT:  ErrorDesc_SERVICE_REQUEST_TIMEOUT.Error(),
		SERVICE_FULLLOAD:         ErrorDesc_SERVICE_FULLLOAD.Error(),

		// 权限错误
		PERMISSION_ONLY_OWNER:    ErrorDesc_PERMISSION_ONLY_OWNER.Error(),
		PERMISSION_CANNOT_DELETE: ErrorDesc_PERMISSION_CANNOT_DELETE.Error(),
	}
)
