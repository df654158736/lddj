package server

import (
	"context"
	c1 "customer/api/customer/v1"
	v1 "customer/api/helloworld/v1"
	"customer/internal/conf"
	"customer/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, customerService *service.CustomerService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//自己的中间件
			selector.Server(jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
				return []byte("secret"), nil
			})).Match(func(ctx context.Context, operation string) bool {
				log.Info(operation)
				whileOder := map[string]bool{
					"/customer.v1.Customer/Login":         true,
					"/customer.v1.Customer/GetVerifyCode": true,
				}
				flag := whileOder[operation]
				return !flag
			}).Build(),
		)}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	c1.RegisterCustomerHTTPServer(srv, customerService)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
