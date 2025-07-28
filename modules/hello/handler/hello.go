package handler

import (
	"github.com/Rivalz-ai/framework-be/modules/hello/dto"
	"github.com/Rivalz-ai/framework-be/server"
	"github.com/gofiber/fiber/v2"
	//"github.com/Rivalz-ai/framework-be/framework/log"
	//"fmt"
	//"time"
	//"github.com/Rivalz-ai/framework-be/framework/apm/apmelk"
	//"context"
)

type HelloHandler struct {
	sv *server.Server
}

func (s *HelloHandler) Init(sv *server.Server) {
	s.sv = sv
}

type APIResponseError struct {
	Code    int    `json:"code" example:"0"`
	Message string `json:"message"`
}
type APIResponse struct {
	Code    int         `json:"code" example:"0"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//	@BasePath	/hello
//
// Rome hello godoc
//
//	@Summary	Rome hello
//	@Schemes
//	@Description	ROme hello
//	@Tags			hello
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Param			authen_key		query		string				true	"Authen Key"
//	@Param			x_id			query		int					true	"X ID"
//	@Param			request			body		XUserToken			true	"User Data Request"
//	@Success		200				{object}	APIResponse			"Success Response"
//	@Failure		400				{object}	APIResponseError	"Unauthorized"
//	@Router			/hello [get]
func (s *HelloHandler) GetHello(c *fiber.Ctx) error {
	hello := &dto.Hello{
		Message: "Hello, Rivalz Rome!",
	}
	return c.JSON(hello)
}
