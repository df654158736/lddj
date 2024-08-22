package server

import (
	"context"
	v1 "driver/api/helloworld/v1"
	"driver/internal/conf"
	"driver/internal/service"

	d1 "driver/api/driver/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, driver *service.DriverService, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//跨域中间件
			selector.Server(
				corsMiddleware(),
			).Match(func(ctx context.Context, operation string) bool {
				return true
			}).Build(),
			//jwt中间件
			selector.Server(
				jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
					return []byte("secret"), nil
				}),
				jwtMiddleware(driver),
			).Match(func(ctx context.Context, operation string) bool {

				whileOder := map[string]bool{
					"/driver.v1.Driver/Login":         true,
					"/driver.v1.Driver/GetVerifyCode": true,
				}
				flag := whileOder[operation]
				println("operation: ", operation, "-flag: ", flag)
				return !flag
			}).Build(),
		),
	}
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
	v1.RegisterGreeterHTTPServer(srv, greeter)
	d1.RegisterDriverHTTPServer(srv, driver)
	return srv
}
