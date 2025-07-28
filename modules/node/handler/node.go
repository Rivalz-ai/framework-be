package handler

import (
	"github.com/Rivalz-ai/framework-be/server"
	//"github.com/Rivalz-ai/framework-be/framework/log"
	//"fmt"
	//"time"
	//"github.com/Rivalz-ai/framework-be/framework/apm/apmelk"
	//"context"
)

type NodeHandler struct {
	sv *server.Server
}

func (s *NodeHandler) Init(sv *server.Server) {
	s.sv = sv
}

//	@BasePath	/hello
// HelloExample godoc
//	@Summary	Hello example
//	@Schemes
//	@Description	Print Hello, Rivalz Rome!
//	@Tags			hello
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Hello,	Rivalz	Rome!
//	@Router			/hello [get]
/*
func (s *NodeHandler)GeTran(c *fiber.Ctx) error{
	hello:= &dto.Hello{
		Message: "Hello, Rivalz Rome!",
	}
	return c.JSON(hello)
}
*/
