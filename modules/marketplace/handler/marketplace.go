package handler

import (
	"os"
	"strconv"
	"strings"

	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/service"
	"github.com/Rivalz-ai/framework-be/server"
	"github.com/gofiber/fiber/v2"
)

type MarketplaceHandler struct {
	sv *server.Server
}

func (s *MarketplaceHandler) Init(sv *server.Server) {
	s.sv = sv
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get token price
// @Description Get token price
// @Accept json
// @Produce json
// @Header 200 {string} Authorization "Bearer {jwt}"
// @Param token_address query string true "Token address"
// @Param decimals query int true "Decimals"
// @Param chain_id query int true "Chain ID"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.TokenPriceResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/token-price [get]
func (s *MarketplaceHandler) GetTokenPrice(c *fiber.Ctx) error {
	tokenAddress := c.Query("token_address")
	decimals := c.Query("decimals")
	decimalsInt := utils.ItoInt(decimals)
	chainId := c.Query("chain_id")
	chainIdInt := utils.ItoInt(chainId)

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	if os.Getenv("ENV") == "dev" {
		return c.JSON(fiber.Map{"code": 0, "message": "GetTokenPrice", "data": dto.TokenPriceResponse{Price: 1}})
	}

	price, err := marketService.GetTokenPrice(c.Context(), tokenAddress, decimalsInt, chainIdInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 0, "message": "GetTokenPrice", "data": dto.TokenPriceResponse{Price: price}})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get mkp ragent
// @Description Get mkp ragent
// @Accept json
// @Produce json
// @Param wallet_address path string false "Wallet address" default(0x)
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10000)
// @Success 200 {object} dto.APIResponseSuccess{data=[]dto.MkpRagent} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/ragent/{wallet_address} [get]
func (s *MarketplaceHandler) GetMkpRAgent(c *fiber.Ctx) error {
	ctx := c.Context()

	wallet_address := c.Params("wallet_address")
	if wallet_address == "" {
		wallet_address = "0x"
	}
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10000)

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	mkpRAgent, err := marketService.GetMkpRAgent(ctx, wallet_address, page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	// Ensure we return an empty array instead of null if mkpRAgent is nil
	if mkpRAgent == nil {
		mkpRAgent = []*dto.MkpRagent{}
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetMkpRAgent", "data": mkpRAgent})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get mkp my agent
// @Description Get mkp my agent
// @Accept json
// @Produce json
// @Param wallet_address path string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=[]dto.MkpRagentUserBuyResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/my-agent/{wallet_address} [get]
func (s *MarketplaceHandler) GetMkpMyAgent(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := c.Params("wallet_address")

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	mkpMyAgent, _, err := marketService.GetMkpMyAgent(ctx, wallet_address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	// Ensure we return an empty array instead of null if mkpMyAgent is nil
	if mkpMyAgent == nil {
		mkpMyAgent = []*dto.MkpRagentUserBuyResponse{}
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetMkpMyAgent", "data": mkpMyAgent})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get buy status
// @Description Get buy status
// @Accept json
// @Produce json
// @Param tracking_id path string true "Tracking ID"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.BuyStatusResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/status/{tracking_id} [get]
func (s *MarketplaceHandler) GetStatus(c *fiber.Ctx) error {
	ctx := c.Context()
	tracking_id := c.Params("tracking_id")

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	buyStatus, err := marketService.GetStatus(ctx, tracking_id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetStatus", "data": dto.BuyStatusResponse{Status: buyStatus}})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get swarm owner
// @Description Get swarm owner
// @Param wallet_address path string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.SwarmOwnerResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/swarm-owner/{wallet_address} [get]
func (s *MarketplaceHandler) SwarmOwner(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := c.Params("wallet_address")

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	swarmOwner, err := marketService.SwarmOwner(ctx, wallet_address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "SwarmOwner", "data": dto.SwarmOwnerResponse{IsSwarmOwner: swarmOwner}})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get quantity token
// @Description Get quantity token
// @Param ragent_id path string true "Ragent ID"
// @Param quantity query int64 true "Quantity"
// @Param side query string true "Side"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.GetQuantityTokenResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/quantity-token/{ragent_id} [get]
func (s *MarketplaceHandler) GetQuantityToken(c *fiber.Ctx) error {
	ctx := c.Context()
	ragent_id := c.Params("ragent_id")
	quantity := c.Query("quantity")
	side := c.Query("side")
	if side == "" {
		side = "buy"
	}
	wallet_address := c.Locals("wallet_address").(string)

	quantityInt, err := strconv.ParseInt(quantity, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	quantityResp, _, err := marketService.GetQuantityToken(ctx, ragent_id, quantityInt, side, wallet_address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetQuantityToken", "data": quantityResp})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get total value by wallet
// @Description Get total value by wallet
// @Param wallet_address path string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.GetTotalValueByWalletResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/total-value/{wallet_address} [get]
func (s *MarketplaceHandler) GetTotalValueByWallet(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := c.Params("wallet_address")

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	totalValue, err := marketService.GetTotalValueByWallet(ctx, wallet_address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetTotalValueByWallet", "data": dto.GetTotalValueByWalletResponse{TotalValue: totalValue}})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get ragent detail
// @Description Get ragent detail
// @Param ragent_id path string true "Ragent ID"
// @Param wallet_address path string true "Wallet address"
// @Success 200 {object} dto.APIResponseSuccess{data=dto.RagentDetailResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/ragent-detail/{ragent_id}/{wallet_address} [get]
func (s *MarketplaceHandler) GetRagentDetail(c *fiber.Ctx) error {
	ctx := c.Context()
	ragent_id := c.Params("ragent_id")
	wallet_address := c.Params("wallet_address")

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	ragentDetail, err := marketService.GetRagentDetail(ctx, ragent_id, wallet_address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetRagentDetail", "data": ragentDetail})
}

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Get ragent history
// @Description Get ragent history
// @Accept json
// @Produce json
// @Param wallet_address path string false "Wallet address" default(0x)
// @Success 200 {object} dto.APIResponseSuccess{data=[]dto.MkpRagentUserBuyResponse} "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/ragent-history/{wallet_address} [get]
func (s *MarketplaceHandler) GetRagentHistory(c *fiber.Ctx) error {
	ctx := c.Context()
	wallet_address := c.Params("wallet_address")

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	mkpMyAgent, err := marketService.GetRagentHistory(ctx, strings.ToLower(wallet_address))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	// Ensure we return an empty array instead of null if mkpMyAgent is nil
	if mkpMyAgent == nil {
		mkpMyAgent = []*dto.MkpRagentUserBuyResponse{}
	}

	return c.JSON(fiber.Map{"code": 0, "message": "GetRagentHistory", "data": mkpMyAgent})
}
