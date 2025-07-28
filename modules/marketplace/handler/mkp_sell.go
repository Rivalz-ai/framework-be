package handler

import (
	"github.com/Rivalz-ai/framework-be/define"
	"github.com/Rivalz-ai/framework-be/framework/base/event"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/service"
	"github.com/gofiber/fiber/v2"
)

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Sell ragent
// @Description Sell ragent
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Param ragent_id path string true "Ragent ID"
// @Param body body dto.SellRagentRequest true "Body"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.SellRagentResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/sell-ragent/{ragent_id} [post]
func (s *MarketplaceHandler) SellRagent(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := c.Locals("wallet_address").(string)
	ragent_id := c.Params("ragent_id")
	// tracking_id := c.Params("tracking_id")
	if ragent_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "ragent_id are required"})
	}

	reqBody := new(dto.SellRagentRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	quantity, trackingId, err := marketService.SubmitSellUnwrapRagent(ctx, wallet_address, ragent_id, reqBody.Hash, true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	jwt := c.Locals("jwt")
	event3 := event.Event{
		EventName: "marketplace|internal-sell-ragent",
		EventData: map[string]interface{}{
			"jwt":            jwt,
			"wallet_address": wallet_address,
			"ragent_id":      ragent_id,
			"tracking_id":    trackingId,
			"hash":           reqBody.Hash,
			"quantity":       quantity,
		},
	}
	if s.sv.Pub["task"] != nil {
		errp := s.sv.Pub["task"].Publish(event3)
		if errp != nil {
			log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceSellRagent-pushToKafka", event3)
		}
	}

	return c.JSON(fiber.Map{"code": 0, "message": "SellRagent", "data": dto.SellRagentResponse{TrackingId: trackingId}})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
// @Summary InternalSellRagent
// @Description InternalSellRagent
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Param is_retry query bool true "Is retry"
// @Param secret query string true "Secret"
// @Param counter query int true "Counter"
// @Param ragent_id query string true "Ragent ID"
// @Param quantity query int64 true "Quantity"
// @Param wallet_address query string true "Wallet address"
// @Param tracking_id query string true "Tracking ID"
// @Success 200 {object} dto.APIResponseSuccess "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/internal-sell-ragent/{ragent_id} [post]
func (s *MarketplaceHandler) InternalSellRagent(c *fiber.Ctx) error {
	is_retry_p := c.Query("is_retry")
	secret := c.Query("secret")
	counter_p := c.Query("counter")
	is_retry := false
	if is_retry_p != "" {
		is_retry = is_retry_p == "true"
	}

	counter := 0
	if is_retry {
		if secret != s.sv.ExtendConfig.InternalSecret {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    1,
				"message": "Invalid internal secret",
			})
		}
		if counter_p != "" {
			counter = utils.ItoInt(counter_p)
			if counter < 0 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "invalid request counter"})
			}
		}
		counter++
	}

	ragent_id := c.Params("ragent_id")
	quantity := c.Query("quantity")
	wallet_address := c.Locals("wallet_address").(string)
	trackingId := c.Query("tracking_id")
	jwt := c.Locals("jwt").(string)

	success_msg := define.TX_SUCCESS

	// go func(jwt string, wallet_address string, ragent_id string, quantity string, trackingId string) {
	//
	lSv, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	// ctxDeadline, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	// defer cancel()

	errV, retry_job := lSv.InternalSellRagent(c.Context(), ragent_id, quantity, wallet_address, trackingId)
	//retry kafka
	//push to kafka

	if retry_job {
		success_msg = define.TX_PENDING //FE will call api check tx
		first_create := false
		is_push_kafka := true
		if errV != nil {
			if errV.Error() == "INSERT_TX" || errV.Error() == "FIND_TX" {
				first_create = true
			}
			if errV.Error() == "NO_RECEIPT" || errV.Error() == "NO_LOGS" {
				is_push_kafka = false
			}
		}
		event1 := event.Event{
			EventHeader: map[string]interface{}{
				"counter": counter,
				"error":   "",
			},
			EventName: "marketplace|internal-sell-ragent",
			EventData: map[string]interface{}{
				"jwt":            jwt,
				"wallet_address": wallet_address,
				"ragent_id":      ragent_id,
				"quantity":       quantity,
				"tracking_id":    trackingId,
				//retry lại như là lần đầu request vì chưa tạo dc db do lỗi db
				"first_create": first_create,
			},
		}
		if s.sv.Pub["task"] != nil && is_push_kafka {
			errp := s.sv.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceSellRagent-pushToKafka", event1)
			}
		} else {
			log.Error("Error when push to kafka: pub task is nil", "MarketplaceSellRagent-pushToKafka", event1)
		}
		//alert tele
		title := "Marketplace Sell Ragent"
		errStr := ""
		if errV != nil {
			errStr = errV.Error()
		}
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "error",
				"title": title,
				"in": map[string]interface{}{
					"wallet_address": wallet_address,
					"quantity":       quantity,
					"ragent_id":      ragent_id,
				},
				"out": map[string]string{
					"error": errStr,
				},
			},
		}
		if s.sv.Pub["tele-internal-alert"] != nil {
			errp := s.sv.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceSellRagent-pushToKafka", event2)
			}
		}
		//nếu có retry thì return success, vì đã push message lên topic cho worker retry, ko trả về lỗi http nữa để tránh loop
		return c.JSON(fiber.Map{"code": 0, "message": success_msg})
	}
	if errV != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": errV.Error()})
	}
	// }(jwt, wallet_address, ragent_id, quantity, trackingId)

	return c.JSON(fiber.Map{"code": 0, "message": success_msg})
}
