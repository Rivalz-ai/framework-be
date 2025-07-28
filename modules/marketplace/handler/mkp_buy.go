package handler

import (
	"strings"

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
// @Summary Generate signature
// @Description Generate signature for buy ragent
// @Accept json
// @Produce json
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Param ragent_id path string true "Ragent ID"
// @Param body body dto.SignBuyRagentMarketRequest true "Body"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.SignBuyRagentMarketResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/sign/{ragent_id} [post]
func (s *MarketplaceHandler) GenerateSignature(c *fiber.Ctx) error {
	wallet_address := c.Locals("wallet_address").(string)
	ragent_id := c.Params("ragent_id")
	var body dto.SignBuyRagentMarketRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if body.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "quantity must be greater than 0"})
	}

	if body.Quantity > 50 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "quantity must be less than 50"})
	}

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	signature, err := marketService.GenerateSignature(c.Context(), wallet_address, body.Quantity, ragent_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GenerateSignature", "data": signature})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Verify buy transaction
// @Description Verify buy transaction
// @Accept json
// @Produce json
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Param hash path string true "Hash"
// @Param ragent_id path string true "Ragent ID"
// @Param body body dto.VerifyBuyTransactionRequest true "Body"
// @Success 200 {object} dto.APIResponseSuccess "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/verify/{ragent_id}/{hash} [post]
func (s *MarketplaceHandler) VerifyBuyTransaction(c *fiber.Ctx) error {
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

	hash := c.Params("hash")
	ragent_id := c.Params("ragent_id")
	user_id := c.Locals("user_id").(string)
	wallet_address := c.Locals("wallet_address").(string)
	ctx := c.Context()

	var body dto.VerifyBuyTransactionRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	//
	lSv, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	errV, retry_job := lSv.VerifyBuyTransaction(ctx, ragent_id, user_id, hash, wallet_address, is_retry, body.Quantity, body.Nonce)
	//retry kafka
	//push to kafka
	success_msg := define.TX_SUCCESS
	jwt := c.Locals("jwt")
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
			EventName: "marketplace|verify",
			EventData: map[string]interface{}{
				"hash":           hash,
				"jwt":            jwt,
				"user_id":        user_id,
				"wallet_address": wallet_address,
				"quantity":       body.Quantity,
				"ragent_id":      ragent_id,
				"nonce":          body.Nonce,
				//retry lại như là lần đầu request vì chưa tạo dc db do lỗi db
				"first_create": first_create,
			},
		}
		if s.sv.Pub["task"] != nil && is_push_kafka {
			errp := s.sv.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "LendingVerify-pushToKafka", event1)
			}
		} else {
			log.Error("Error when push to kafka: pub task is nil", "LendingVerify-pushToKafka", event1)
		}
		//alert tele
		errStr := ""
		if errV != nil {
			errStr = errV.Error()
		}
		title := "Marketplace Verify"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "error",
				"title": title,
				"in": map[string]interface{}{
					"hash": hash,
					"jwt":  utils.ItoString(jwt),
				},
				"out": map[string]string{
					"error": errStr,
				},
			},
		}
		if s.sv.Pub["tele-internal-alert"] != nil {
			errp := s.sv.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "LendingVerify-pushToKafka", event2)
			}
		}
		//nếu có retry thì return success, vì đã push message lên topic cho worker retry, ko trả về lỗi http nữa để tránh loop
		return c.JSON(fiber.Map{"code": 0, "message": "RETRY"})
	}
	if errV != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": errV.Error()})
	}

	// if success -> push to kafka to contunue process internal swap token
	event3 := event.Event{
		EventName: "marketplace|internal-swap-token",
		EventData: map[string]interface{}{
			"hash":           hash,
			"jwt":            jwt,
			"user_id":        user_id,
			"wallet_address": wallet_address,
			"quantity":       body.Quantity,
			"ragent_id":      ragent_id,
			"nonce":          body.Nonce,
		},
	}
	if s.sv.Pub["task"] != nil {
		errp := s.sv.Pub["task"].Publish(event3)
		if errp != nil {
			log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerify-pushToKafka", event3)
		}
	}

	return c.JSON(fiber.Map{"code": 0, "message": success_msg})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Internal swap token
// @Description Internal swap token
// @Accept json
// @Produce json
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Param ragent_id path string true "Ragent ID"
// @Param hash path string true "Hash"
// @Param quantity query string true "Quantity"
// @Param is_retry query string true "is_retry"
// @Param secret query string true "secret"
// @Param counter query string true "counter"
// @Success 200 {object} dto.APIResponseSuccess{data=string} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/internal-swap-token/{ragent_id}/{hash} [post]
func (s *MarketplaceHandler) InternalSwapToken(c *fiber.Ctx) error {
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

	hash := c.Params("hash")
	ragent_id := c.Params("ragent_id")
	quantity := c.Query("quantity")
	wallet_address := c.Locals("wallet_address").(string)

	success_msg := define.TX_SUCCESS
	jwt := c.Locals("jwt").(string)
	// go func(jwt string, wallet_address string, ragent_id string, quantity string) {
	//
	lSv, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	// ctxDeadline, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	// defer cancel()

	errV, retry_job := lSv.InternalSwapToken(c.Context(), ragent_id, quantity, wallet_address, wallet_address, hash)
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

			if strings.Contains(errV.Error(), "mongo: no documents") {
				return c.JSON(fiber.Map{"code": 0, "message": success_msg})
			}
		}

		event1 := event.Event{
			EventHeader: map[string]interface{}{
				"counter": counter,
				"error":   "",
			},
			EventName: "marketplace|internal-swap-token",
			EventData: map[string]interface{}{
				"hash":           hash,
				"jwt":            jwt,
				"wallet_address": wallet_address,
				"quantity":       quantity,
				"ragent_id":      ragent_id,
				//retry lại như là lần đầu request vì chưa tạo dc db do lỗi db
				"first_create": first_create,
			},
		}
		if s.sv.Pub["task"] != nil && is_push_kafka {
			errp := s.sv.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceInternalSwapToken-pushToKafka", event1)
			}
		} else {
			log.Error("Error when push to kafka: pub task is nil", "MarketplaceInternalSwapToken-pushToKafka", event1)
		}
		//alert tele
		errStr := ""
		if errV != nil {
			errStr = errV.Error()
		}
		title := "Marketplace Internal Swap Token"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "error",
				"title": title,
				"in": map[string]interface{}{
					"hash":           hash,
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
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceInternalSwapToken-pushToKafka", event2)
			}
		}
		//nếu có retry thì return success, vì đã push message lên topic cho worker retry, ko trả về lỗi http nữa để tránh loop
		return c.JSON(fiber.Map{"code": 0, "message": success_msg})
	}
	if errV != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": errV.Error()})
	}

	// }(jwt, wallet_address, ragent_id, quantity)

	return c.JSON(fiber.Map{"code": 0, "message": success_msg})
}
