// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"chihqiang/msgbox-go/services/common/errs"
	"chihqiang/msgbox-go/services/gateway/api/internal/types"
	"context"
	"encoding/base64"
	xhttp "github.com/zeromicro/x/http"
	"net/http"
	"strings"
)

type BasicAuthMiddleware struct {
}

func NewBasicAuthMiddleware() *BasicAuthMiddleware {
	return &BasicAuthMiddleware{}
}

func (m *BasicAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取请求上下文（用于传递账号密码和响应错误）
		ctx := r.Context()

		// 1. 提取 Authorization 请求头（从 types 常量中获取头名称，统一管理）
		authHeader := r.Header.Get(types.HeaderAuthorization)
		// 校验认证头是否为空
		if authHeader == "" {
			xhttp.JsonBaseResponseCtx(ctx, w, errs.ErrAuthMissing)
			return
		}

		// 2. 校验认证头格式（必须为 "Basic <base64编码串>" 格式）
		// SplitN 按空格分割为两部分，避免编码串中含空格导致分割异常
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != types.HeaderBasic {
			xhttp.JsonBaseResponseCtx(ctx, w, errs.ErrAuthInvalidForm)
			return
		}

		// 3. Base64 解码认证信息（parts[1] 为 base64(账号:密码) 编码串）
		decodedBytes, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			xhttp.JsonBaseResponseCtx(ctx, w, errs.ErrAuthInvalidForm)
			return
		}
		// 4. 拆分账号和密码（解码后格式为 "username:password"）
		authStr := string(decodedBytes)
		// SplitN 按 ":" 分割为两部分，避免密码中含 ":" 导致分割异常
		userPwdParts := strings.SplitN(authStr, ":", 2)
		// 校验分割结果（需拆分为账号和密码两部分）
		if len(userPwdParts) != 2 {
			xhttp.JsonBaseResponseCtx(ctx, w, errs.ErrAuthInvalidForm)
			return
		}
		// 5. 清洗账号密码（去除首尾空格，兼容客户端传入多余空格的场景）
		username := strings.TrimSpace(userPwdParts[0])
		password := strings.TrimSpace(userPwdParts[1])
		// 6. 账号密码存入上下文（使用 types 常量定义的 key，避免硬编码冲突）
		// 后续业务逻辑可通过 ctx.Value(types.BasicAuthUsername) 获取认证账号
		ctx = context.WithValue(ctx, types.BasicAuthUsername, username)
		ctx = context.WithValue(ctx, types.BasicAuthPassword, password)
		// 7. 校验通过，执行下一个处理函数（传递更新后的上下文）
		next(w, r.WithContext(ctx))
	}
}
