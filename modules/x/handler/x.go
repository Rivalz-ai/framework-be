package handler

import (
	//"github.com/Rivalz-ai/framework-be/modules/reward/dto"
	"github.com/Rivalz-ai/framework-be/modules/x/dto"
	xService "github.com/Rivalz-ai/framework-be/modules/x/service"
	"github.com/Rivalz-ai/framework-be/server"
	"github.com/gofiber/fiber/v2"

	//"github.com/Rivalz-ai/framework-be/framework/log"
	//"fmt"
	"time"
	//"github.com/Rivalz-ai/framework-be/framework/apm/apmelk"
	//"context"
)

type XHandler struct {
	sv *server.Server
}

func (s *XHandler) Init(sv *server.Server) {
	s.sv = sv
}

type APIResponseError struct {
	Code    int    `json:"code" example:"1"`
	Message string `json:"message"`
}
type XInfo struct {
	XAccessToken  string     `json:"x_access_token"`
	XRefreshToken string     `json:"x_refresh_token"`
	ExpiresAt     *time.Time `json:"expires_at"`
}

type XConnectResponse struct {
	Code    int           `json:"code" example:"0"`
	Data    XInfoResponse `json:"data"`
	Message string        `json:"message"`
}
type XInfoResponse struct {
	XId          string     `json:"x_id"`
	XName        string     `json:"x_name"`
	XUser        string     `json:"x_user"`
	XAvatar      string     `json:"x_avatar"`
	XCreatedAt   *time.Time `json:"x_created_at"`
	XDescription string     `json:"x_description"`
	//
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
	LikeCount      int `json:"like_count"`
	MediaCount     int `json:"media_count"` //
}

//	@BasePath	/x
//
// XConnect godoc
//
//	@Summary	XConnect
//	@Schemes
//	@Description	XConnect
//	@Tags			x
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Param			request			body		dto.ConnectX				true	"XConnect Request"
//	@Success		200				{object}	XConnectResponse	"Success Response"
//	@Failure		400				{object}	APIResponseError	"Fail Response"
//	@Router			/x/user [post]
func (s *XHandler) Connect(c *fiber.Ctx) error {
	ctx := c.Context()
	var payload dto.ConnectX
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{"code": 1, "message": "Invalid request"})
	}
	if payload.XAccessToken == "" || payload.XRefreshToken == "" || payload.ExpiresAt == nil {
		return c.JSON(fiber.Map{"code": 1, "message": "Invalid request, missing field"})
	}
	xSV, err := xService.NewXUserService(s.sv)
	if err != nil {
		return c.JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	user_id := c.Locals("user_id").(string)
	xInfo, err := xSV.Connect(ctx, user_id, payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "message": "success", "data": xInfo})
}

//	@BasePath	/x
//
// XDisconnect godoc
//
//	@Summary	XDisconnect
//	@Schemes
//	@Description	XDisconnect
//	@Tags			x
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Router			/x/user [delete]
func (s *XHandler) Disconnect(c *fiber.Ctx) error {

	user_id := c.Locals("user_id").(string)
	xSV, err := xService.NewXUserService(s.sv)
	if err != nil {
		return c.JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	err = xSV.Disconnect(c.Context(), user_id)
	if err != nil {
		return c.JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "message": "success"})
}

//	@BasePath	/x
//
// GetApp godoc
//
//	@Summary	GetApp
//	@Schemes
//	@Description	GetApp
//	@Tags			x
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Success		200				{object}	dto.APIResponse{data=dto.App}	"Success Response"
//	@Failure		400				{object}	APIResponseError	"Fail Response"
//	@Router			/x/app [get]
func (s *XHandler) GetApp(c *fiber.Ctx) error {

	xSV, err := xService.NewXUserService(s.sv)
	if err != nil {
		return c.JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	app, err := xSV.GetApp(c.Context())
	if err != nil {
		return c.JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "success", "data": app})
}

//	@BasePath	/x
//
// MKPConnect godoc
//
//	@Summary	MKPConnect
//	@Schemes
//	@Description	MKPConnect
//	@Tags			x
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Param			request			body		dto.MKPConnect		true	"MKPConnect Request"
//	@Success		200				{object}	dto.APIResponse	"Success Response"
//	@Failure		400				{object}	APIResponseError	"Fail Response"
//	@Router			/x/mkp/connect [post]
func (s *XHandler) MKPConnect(c *fiber.Ctx) error {
	ctx := c.Context()
	user_id := c.Locals("user_id").(string)
	var payload dto.MKPConnect
	if err := c.BodyParser(&payload); err != nil {
		return c.JSON(fiber.Map{"code": 1, "message": "Invalid request"})
	}
	if payload.XId == "" || payload.XUsername == "" {
		return c.JSON(fiber.Map{"code": 1, "message": "Invalid request, missing field"})
	}
	xSVC, err := xService.NewXUserService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	err = xSVC.MKPConnect(ctx, user_id, payload.XId, payload.XUsername)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "message": "success"})
}

//	@BasePath	/x
//
// MKPDisconnect godoc
//
//	@Summary	MKPDisconnect
//	@Schemes
//	@Description	MKPDisconnect
//	@Tags			x
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer Token"
//	@Success		200				{object}	dto.APIResponse	"Success Response"
//	@Failure		400				{object}	APIResponseError	"Fail Response"
//	@Router			/x/mkp/disconnect [delete]
func (s *XHandler) MKPDisconnect(c *fiber.Ctx) error {
	ctx := c.Context()
	user_id := c.Locals("user_id").(string)
	xSVC, err := xService.NewXUserService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	err = xSVC.MKPDisconnect(ctx, user_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "message": "success"})
}
