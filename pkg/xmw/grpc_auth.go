package xmw

import (
	"context"

	"github.com/aswancen/libs/api/xerr"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
)

var ErrAuthNotFound = xerr.ErrorAuth("auth not found")

const (
	UID        = "zgy-user-id"
	TID        = "zgy-tenant-id"
	CtxAuthKey = "userInfo"
)

type AuthInfo struct {
	UserId   string
	TenantId string
}

func ParseTenantUser() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if meta, ok := metadata.FromServerContext(ctx); ok {
				uid := meta.Get(UID)
				tid := meta.Get(TID)
				if tid == "" || uid == "" {
					return nil, ErrAuthNotFound
				}
				ctx = context.WithValue(ctx, CtxAuthKey, &AuthInfo{
					UserId:   uid,
					TenantId: tid,
				})
			}
			return handler(ctx, req)
		}
	}
}
