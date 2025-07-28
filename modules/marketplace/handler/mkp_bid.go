package handler

import (
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/base/event"
	"github.com/Rivalz-ai/framework-be/framework/log"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/service"
	"github.com/gofiber/fiber/v2"
)

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
// @Summary GetBids
// @Description GetBids
// @Param ragent_id path string true "Ragent ID"
// @Param wallet_address query string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.BidResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/bids/{ragent_id} [get]
func (s *MarketplaceHandler) GetBids(c *fiber.Ctx) error {
	ctx := c.Context()
	ragent_id := c.Params("ragent_id")
	wallet_address := strings.ToLower(c.Query("wallet_address"))
	if wallet_address == "0x" {
		wallet_address = ""
	}

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return c.JSON(fiber.Map{"code": 0, "message": "ragent does not have any bid", "data": []dto.BidResponse{}})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	bid, err := marketService.GetBids(ctx, ragent_id, wallet_address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if bid == nil {
		return c.JSON(fiber.Map{"code": 0, "message": "GetBids", "data": []dto.BidResponse{}})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetBids", "data": bid})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
// @Summary GetMyActiveBids
// @Description GetMyActiveBids
// @Param wallet_address path string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.MyActiveBidResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/my-active-bids/{wallet_address} [get]
func (s *MarketplaceHandler) GetMyActiveBids(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := strings.ToLower(c.Params("wallet_address"))
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	bid, err := marketService.GetMyActiveBids(ctx, wallet_address)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return c.JSON(fiber.Map{"code": 0, "message": "user does not have any active bid", "data": []dto.MyActiveBidResponse{}})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if bid == nil {
		return c.JSON(fiber.Map{"code": 0, "message": "GetMyActiveBids", "data": []dto.MyActiveBidResponse{}})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetMyActiveBids", "data": bid})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
// @Summary GetMyHistoricalBids
// @Description GetMyHistoricalBids
// @Param wallet_address path string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.BidHistoryResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/my-historical-bids/{wallet_address} [get]
func (s *MarketplaceHandler) GetMyHistoricalBids(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := strings.ToLower(c.Params("wallet_address"))
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	bid, err := marketService.GetMyHistoricalBids(ctx, wallet_address)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return c.JSON(fiber.Map{"code": 0, "message": "user does not have any historical bid", "data": []dto.BidHistoryResponse{}})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if bid == nil {
		return c.JSON(fiber.Map{"code": 0, "message": "GetMyHistoricalBids", "data": []dto.BidHistoryResponse{}})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetMyHistoricalBids", "data": bid})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
// @Summary GetMyVestingTokens
// @Description GetMyVestingTokens
// @Param wallet_address path string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.VestingTokenResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/my-vesting-tokens/{wallet_address} [get]
func (s *MarketplaceHandler) GetMyVestingTokens(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := strings.ToLower(c.Params("wallet_address"))
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	bid, err := marketService.GetMyVestingTokens(ctx, wallet_address)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return c.JSON(fiber.Map{"code": 0, "message": "user does not have any vesting token", "data": []dto.VestingTokenResponse{}})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if bid == nil {
		return c.JSON(fiber.Map{"code": 0, "message": "GetMyVestingTokens", "data": []dto.VestingTokenResponse{}})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetMyVestingTokens", "data": bid})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Summary VerifyAgreeBid
// @Description VerifyAgreeBid
// @Param ragent_id path string true "Ragent ID"
// @Param hash path string true "Hash"
// @Success 200 {object} dto.APIResponseSuccess "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/verify-agree-bid/{ragent_id}/{hash} [post]
func (s *MarketplaceHandler) VerifyAgreeBid(c *fiber.Ctx) error {
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

	ctx := c.Context()
	wallet_address := c.Locals("wallet_address").(string)
	ragent_id := c.Params("ragent_id")
	hash := c.Params("hash")
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	err, isRetry := marketService.VerifyAgreeBid(ctx, wallet_address, ragent_id, hash)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if isRetry {
		is_push_kafka := true
		jwt := c.Locals("jwt")
		event1 := event.Event{
			EventHeader: map[string]interface{}{
				"counter": counter,
				"error":   "",
			},
			EventName: "marketplace|verify-agree-bid",
			EventData: map[string]interface{}{
				"jwt":            jwt,
				"wallet_address": wallet_address,
				"ragent_id":      ragent_id,
				"hash":           hash,
			},
		}
		if s.sv.Pub["task"] != nil && is_push_kafka {
			errp := s.sv.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyAgreeBid-pushToKafka", event1)
			}
		} else {
			log.Error("Error when push to kafka: pub task is nil", "MarketplaceVerifyAgreeBid-pushToKafka", event1)
		}

		//alert tele
		title := "Marketplace Verify Agree Bid"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "error",
				"title": title,
				"in": map[string]interface{}{
					"wallet_address": wallet_address,
					"ragent_id":      ragent_id,
					"hash":           hash,
				},
				"out": map[string]string{
					"error": err.Error(),
				},
			},
		}
		if s.sv.Pub["tele-internal-alert"] != nil {
			errp := s.sv.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyAgreeBid-pushToKafka", event2)
			}
		}

		return c.JSON(fiber.Map{"code": 0, "message": "success"})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "success"})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Summary VerifyReleaseBid
// @Description VerifyReleaseBid
// @Param ragent_id path string true "Ragent ID"
// @Param hash path string true "Hash"
// @Success 200 {object} dto.APIResponseSuccess "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/verify-release-bid/{ragent_id}/{hash} [post]
func (s *MarketplaceHandler) VerifyReleaseBid(c *fiber.Ctx) error {
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

	ctx := c.Context()
	wallet_address := c.Locals("wallet_address").(string)
	ragent_id := c.Params("ragent_id")
	hash := c.Params("hash")
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	err, isRetry := marketService.VerifyReleaseBid(ctx, wallet_address, ragent_id, hash)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if isRetry {
		is_push_kafka := true
		jwt := c.Locals("jwt")
		event1 := event.Event{
			EventHeader: map[string]interface{}{
				"counter": counter,
				"error":   "",
			},
			EventName: "marketplace|verify-release-bid",
			EventData: map[string]interface{}{
				"jwt":            jwt,
				"wallet_address": wallet_address,
				"ragent_id":      ragent_id,
				"hash":           hash,
			},
		}
		if s.sv.Pub["task"] != nil && is_push_kafka {
			errp := s.sv.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyReleaseBid-pushToKafka", event1)
			}
		} else {
			log.Error("Error when push to kafka: pub task is nil", "MarketplaceVerifyReleaseBid-pushToKafka", event1)
		}

		//alert tele
		title := "Marketplace Verify Release Bid"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "error",
				"title": title,
				"in": map[string]interface{}{
					"wallet_address": wallet_address,
					"ragent_id":      ragent_id,
					"hash":           hash,
				},
				"out": map[string]string{
					"error": err.Error(),
				},
			},
		}
		if s.sv.Pub["tele-internal-alert"] != nil {
			errp := s.sv.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyReleaseBid-pushToKafka", event2)
			}
		}

		return c.JSON(fiber.Map{"code": 0, "message": "success"})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "success"})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Summary VerifyRemoveBid
// @Description VerifyRemoveBid
// @Param ragent_id path string true "Ragent ID"
// @Param hash path string true "Hash"
// @Success 200 {object} dto.APIResponseSuccess "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/verify-remove-bid/{ragent_id}/{hash} [post]
func (s *MarketplaceHandler) VerifyRemoveBid(c *fiber.Ctx) error {
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

	ctx := c.Context()
	wallet_address := c.Locals("wallet_address").(string)
	ragent_id := c.Params("ragent_id")
	hash := c.Params("hash")
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	err, isRetry := marketService.VerifyRemoveBid(ctx, wallet_address, ragent_id, hash)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if isRetry {
		is_push_kafka := true
		jwt := c.Locals("jwt")
		event1 := event.Event{
			EventHeader: map[string]interface{}{
				"counter": counter,
				"error":   "",
			},
			EventName: "marketplace|verify-remove-bid",
			EventData: map[string]interface{}{
				"jwt":            jwt,
				"wallet_address": wallet_address,
				"ragent_id":      ragent_id,
				"hash":           hash,
			},
		}
		if s.sv.Pub["task"] != nil && is_push_kafka {
			errp := s.sv.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyRemoveBid-pushToKafka", event1)
			}
		} else {
			log.Error("Error when push to kafka: pub task is nil", "MarketplaceVerifyRemoveBid-pushToKafka", event1)
		}

		//alert tele
		title := "Marketplace Verify Remove Bid"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "error",
				"title": title,
				"in": map[string]interface{}{
					"wallet_address": wallet_address,
					"ragent_id":      ragent_id,
					"hash":           hash,
				},
				"out": map[string]string{
					"error": err.Error(),
				},
			},
		}
		if s.sv.Pub["tele-internal-alert"] != nil {
			errp := s.sv.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyRemoveBid-pushToKafka", event2)
			}
		}

		return c.JSON(fiber.Map{"code": 0, "message": "success"})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "success"})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
//
//	@Param			Authorization	header		string				true	"Bearer Token"
//
// @Summary VerifyCreateBid
// @Description VerifyCreateBid
// @Param ragent_id path string true "Ragent ID"
// @Param hash path string true "Hash"
// @Param is_premium query string true "Is Premium"
// @Success 200 {object} dto.APIResponseSuccess "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/verify-create-bid/{ragent_id}/{hash} [post]
func (s *MarketplaceHandler) VerifyCreateBid(c *fiber.Ctx) error {
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

	ctx := c.Context()
	wallet_address := c.Locals("wallet_address").(string)
	user_id := c.Locals("user_id").(string)
	ragent_id := c.Params("ragent_id")
	hash := c.Params("hash")
	is_premium := c.Query("is_premium")
	is_premium_bool := false
	if is_premium != "" {
		is_premium_bool = is_premium == "true"
	}
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	err, isRetry := marketService.VerifyCreateBid(ctx, wallet_address, user_id, ragent_id, hash, is_retry, is_premium_bool)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if isRetry {
		is_push_kafka := true
		jwt := c.Locals("jwt")
		event1 := event.Event{
			EventHeader: map[string]interface{}{
				"counter": counter,
				"error":   "",
			},
			EventName: "marketplace|verify-create-bid",
			EventData: map[string]interface{}{
				"jwt":            jwt,
				"wallet_address": wallet_address,
				"ragent_id":      ragent_id,
				"hash":           hash,
			},
		}
		if s.sv.Pub["task"] != nil && is_push_kafka {
			errp := s.sv.Pub["task"].Publish(event1)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyCreateBid-pushToKafka", event1)
			}
		} else {
			log.Error("Error when push to kafka: pub task is nil", "MarketplaceVerifyCreateBid-pushToKafka", event1)
		}

		//alert tele
		title := "Marketplace Verify Create Bid"
		event2 := event.Event{
			EventName: "alert-internal-tele",
			EventData: map[string]interface{}{
				"level": "error",
				"title": title,
				"in": map[string]interface{}{
					"wallet_address": wallet_address,
					"ragent_id":      ragent_id,
					"hash":           hash,
				},
				"out": map[string]string{
					"error": err.Error(),
				},
			},
		}
		if s.sv.Pub["tele-internal-alert"] != nil {
			errp := s.sv.Pub["tele-internal-alert"].Publish(event2)
			if errp != nil {
				log.Error("Error when push to kafka: "+errp.Error(), "MarketplaceVerifyCreateBid-pushToKafka", event2)
			}
		}

		return c.JSON(fiber.Map{"code": 0, "message": "success"})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "success"})
}

// @BasePath	/marketplace
// @Tags			marketplace
// @Accept			json
// @Produce		json
// @Summary AcceptBidQuantity
// @Description AcceptBidQuantity
// @Param ragent_id path string true "Ragent ID"
// @Param bid_id path string true "Bid ID"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.AcceptBidQuantityResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/bid-quantity/{ragent_id}/{bid_id} [get]
func (s *MarketplaceHandler) AcceptBidQuantity(c *fiber.Ctx) error {
	ctx := c.Context()
	bid_id := c.Params("bid_id")
	ragent_id := c.Params("ragent_id")
	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	acceptBidQuantity, err := marketService.AcceptBidQuantity(ctx, bid_id, ragent_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "success", "data": acceptBidQuantity})
}
