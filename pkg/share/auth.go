package share

import (
	"context"
	"fmt"
	"github.com/aswancen/libs1.0/pkg/xmw"
)

// GetRpcToken 获取Token
func GetRpcToken(ctx context.Context) (*xmw.AuthInfo, error) {
	value := ctx.Value(xmw.CtxAuthKey)
	if token, IsOk := value.(*xmw.AuthInfo); IsOk {
		return token, nil
	}
	return nil, fmt.Errorf("token not found")
}
