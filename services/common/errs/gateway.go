// Package errors 定义了项目全局统一的错误码、错误信息及错误对象创建工具
// 核心作用：统一错误码规范、简化错误使用、便于前端/客户端统一处理错误提示
package errs

import "github.com/zeromicro/x/errors"

// 错误码设计规范（严格遵循，便于扩展和维护）：
// 1. 成功码：0（唯一成功标识）
// 2. 错误码分段：按业务模块划分，每段1000个码值，避免冲突
//    - 通用错误：1000~1999（适用于所有模块的通用异常）
//    - 认证错误：2000~2999（身份认证相关异常）
//    - 后续扩展：3000~3999（用户模块）、4000~4999（订单模块）等，按业务新增

// 基础状态码

const (
	Success = 0 // 操作成功（所有成功响应的统一状态码）
)

// 通用错误码（1000 段）：适用于所有业务模块的通用异常场景
const (
	ErrCodeUnknown      = 1000 // 未知错误：未定义的异常场景
	ErrCodeParamInvalid = 1001 // 请求参数无效：参数格式/类型/范围不符合接口要求
	ErrCodeConfigError  = 1002 // 配置加载失败：系统启动时配置文件读取/解析失败
)

// 认证相关错误码（2000 段）：用户身份认证流程中的异常场景
const (
	ErrCodeAuthMissing     = 2000 // 缺少认证头：请求未携带 Authorization 头信息
	ErrCodeAuthInvalidForm = 2001 // 认证格式/解码错误：Authorization 格式错误或 Base64 解码失败
	ErrCodeAuthInvalid     = 2002 // 认证凭证无效：账号或密码错误，验证未通过
)

const (
	ErrCodeTemplateMissing        = 3000
	ErrCodeTemplateChannelMissing = 3001
)

const (
	ErrCodeDB = 4000
)

// errorMap 错误码-提示信息映射表
// 说明：
// 1. 严格与上方错误码常量一一对应，禁止出现无码的消息或无消息的码
// 2. 提示信息需简洁明了，用户可理解（必要时补充解决方案）
// 3. 后端日志可记录详细错误，前端仅展示此映射表中的消息，避免敏感信息泄露
var errorMap = map[int]string{
	// 基础状态
	Success: "操作成功",

	// 通用错误
	ErrCodeUnknown:      "未知错误，请联系管理员处理",
	ErrCodeParamInvalid: "请求参数无效，请检查参数格式是否正确",
	ErrCodeConfigError:  "系统配置加载失败，请联系管理员排查",

	// 认证错误
	ErrCodeAuthMissing:     "缺少Authorization认证头，请在请求头中携带认证信息",
	ErrCodeAuthInvalidForm: "认证信息不合法，正确格式：Basic <base64(账号:密码)>，需确保是「账号:密码」的Base64编码",
	ErrCodeAuthInvalid:     "账号或密码错误，认证失败，请核对后重试",

	//模版错误
	ErrCodeTemplateMissing:        "缺少模版code",
	ErrCodeTemplateChannelMissing: "模版没有配置通道",

	ErrCodeDB: "内部错误",
}

// 预定义错误对象：全局复用，避免重复创建
// 说明：所有错误对象统一通过 GetErr 函数创建，保证错误码与消息的一致性
var (
	// 通用错误对象
	ErrUnknown      = GetErr(ErrCodeUnknown)      // 未知错误
	ErrParamInvalid = GetErr(ErrCodeParamInvalid) // 请求参数无效
	ErrConfigError  = GetErr(ErrCodeConfigError)  // 配置加载失败

	// 认证错误对象
	ErrAuthMissing     = GetErr(ErrCodeAuthMissing)     // 缺少Authorization认证头
	ErrAuthInvalidForm = GetErr(ErrCodeAuthInvalidForm) // 认证信息格式/解码错误
	ErrAuthInvalid     = GetErr(ErrCodeAuthInvalid)     // 账号或密码错误

	ErrTemplateCodeMissing    = GetErr(ErrCodeTemplateMissing) // 缺少模版code
	ErrTemplateChannelMissing = GetErr(ErrCodeTemplateChannelMissing)
	ErrDB                     = GetErr(ErrCodeDB)
)

// GetErr 根据错误码获取对应的错误对象
// 参数：code - 预定义的错误码常量（如 ErrCodeParamInvalid）
// 返回：对应的错误对象（包含错误码和提示信息）
// 特性：
// 1. 若传入无效错误码（未在 errorMap 中定义），默认返回「未知错误」对象
// 2. 确保返回的错误对象始终包含合法的码和消息，避免空值或不一致
func GetErr(code int) error {
	msg, ok := errorMap[code]
	if !ok {
		// 可选：记录无效错误码日志，便于排查非法调用（建议在生产环境开启）
		// log.Printf("warning: use invalid error code: %d", code)
		return errors.New(ErrCodeUnknown, errorMap[ErrCodeUnknown])
	}
	return errors.New(code, msg)
}
