package server

import (
	"context"
	"customer/internal/service"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func jwtMiddleware(customerService *service.CustomerService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			//获取jwt用户id
			claims, ok := jwt.FromContext(ctx)
			if !ok {
				return nil, errors.Unauthorized("jwt token invalid", "claims not found")
			}

			//断言
			claimsMap := claims.(jwtv5.MapClaims)
			log.Info(claimsMap)
			id := claimsMap["jti"].(string)

			//查询数据库中的客户token
			token, err := customerService.CustomerRepo.GetToken(id)
			if err != nil {
				return nil, errors.Unauthorized("jwt token invalid", "customer not found")
			}

			//判断token是否一致
			header, _ := transport.FromServerContext(ctx)
			auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
			jwtToken := auths[1]
			if jwtToken != token {
				return nil, errors.Unauthorized("jwt token invalid", "token was updated")
			}

			return handler(ctx, req)
		}
	}
}
