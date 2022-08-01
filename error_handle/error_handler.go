package error_handle

import (
	"github.com/pkg/errors"
)

// 1、自定义error结构体，并重写Error()方法
// 错误时返回自定义结构
type CustomError struct {
	Code       int    `json:"code"`    // 业务码
	TagMessage string `json:"message"` // 描述信息
}

func (e *CustomError) Error() string {
	return e.TagMessage
}

// 2、定义errorCode
const (
	// 服务级错误码
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	CallHTTPError      = 10105
	ResubmitError      = 10106
	ResubmitMsg        = 10107
	HashIdsDecodeError = 10108
	SignatureError     = 10109

	// 业务模块级错误码
	// 用户模块
	IllegalUserName = 20101
	UserCreateError = 20102
	UserUpdateError = 20103
	UserSearchError = 20104

	// 授权调用方
	AuthorizedCreateError    = 20201
	AuthorizedListError      = 20202
	AuthorizedDeleteError    = 20203
	AuthorizedUpdateError    = 20204
	AuthorizedDetailError    = 20205
	AuthorizedCreateAPIError = 20206
	AuthorizedListAPIError   = 20207
	AuthorizedDeleteAPIError = 20208

	// 管理员
	AdminCreateError             = 20301
	AdminListError               = 20302
	AdminDeleteError             = 20303
	AdminUpdateError             = 20304
	AdminResetPasswordError      = 20305
	AdminLoginError              = 20306
	AdminLogOutError             = 20307
	AdminModifyPasswordError     = 20308
	AdminModifyPersonalInfoError = 20309

	// 配置
	ConfigEmailError        = 20401
	ConfigSaveError         = 20402
	ConfigRedisConnectError = 20403
	ConfigMySQLConnectError = 20404
	ConfigMySQLInstallError = 20405
	ConfigGoVersionError    = 20406

	// 实用工具箱
	SearchRedisError = 20501
	ClearRedisError  = 20502
	SearchRedisEmpty = 20503
	SearchMySQLError = 20504

	// 菜单栏
	MenuCreateError = 20601
	MenuUpdateError = 20602
	MenuListError   = 20603
	MenuDeleteError = 20604
	MenuDetailError = 20605

	// 借书
	BookNotFoundError        = 20701
	BookHasBeenBorrowedError = 20702
)

// 3、定义errorCode对应的文本信息
var codeTag = map[int]string{
	ServerError:        "Internal Server Error",
	TooManyRequests:    "Too Many Requests",
	ParamBindError:     "参数信息有误",
	AuthorizationError: "签名信息有误",
	CallHTTPError:      "调用第三方 HTTP 接口失败",
	ResubmitError:      "Resubmit Error",
	ResubmitMsg:        "请勿重复提交",
	HashIdsDecodeError: "ID参数有误",
	SignatureError:     "SignatureError",

	IllegalUserName: "非法用户名",
	UserCreateError: "创建用户失败",
	UserUpdateError: "更新用户失败",
	UserSearchError: "查询用户失败",

	AuthorizedCreateError:    "创建调用方失败",
	AuthorizedListError:      "获取调用方列表页失败",
	AuthorizedDeleteError:    "删除调用方失败",
	AuthorizedUpdateError:    "更新调用方失败",
	AuthorizedDetailError:    "获取调用方详情失败",
	AuthorizedCreateAPIError: "创建调用方API地址失败",
	AuthorizedListAPIError:   "获取调用方API地址列表失败",
	AuthorizedDeleteAPIError: "删除调用方API地址失败",

	AdminCreateError:             "创建管理员失败",
	AdminListError:               "获取管理员列表页失败",
	AdminDeleteError:             "删除管理员失败",
	AdminUpdateError:             "更新管理员失败",
	AdminResetPasswordError:      "重置密码失败",
	AdminLoginError:              "登录失败",
	AdminLogOutError:             "退出失败",
	AdminModifyPasswordError:     "修改密码失败",
	AdminModifyPersonalInfoError: "修改个人信息失败",

	ConfigEmailError:        "修改邮箱配置失败",
	ConfigSaveError:         "写入配置文件失败",
	ConfigRedisConnectError: "Redis连接失败",
	ConfigMySQLConnectError: "MySQL连接失败",
	ConfigMySQLInstallError: "MySQL初始化数据失败",
	ConfigGoVersionError:    "GoVersion不满足要求",

	SearchRedisError: "查询RedisKey失败",
	ClearRedisError:  "清空RedisKey失败",
	SearchRedisEmpty: "查询的RedisKey不存在",
	SearchMySQLError: "查询mysql失败",

	MenuCreateError: "创建菜单失败",
	MenuUpdateError: "更新菜单失败",
	MenuDeleteError: "删除菜单失败",
	MenuListError:   "获取菜单列表页失败",
	MenuDetailError: "获取菜单详情失败",

	BookNotFoundError:        "书未找到",
	BookHasBeenBorrowedError: "书已经被借走了",
}

func Text(code int) string {
	return codeTag[code]
}

// 4、新建自定义error实例化
func NewCustomError(code int) error {
	// 初次调用得用Wrap方法，进行实例化
	return errors.Wrap(&CustomError{
		Code:       code,
		TagMessage: codeTag[code],
	}, "")
}
