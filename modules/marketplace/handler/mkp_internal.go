package handler

import (
	"github.com/Rivalz-ai/framework-be/framework/utils"
	_ "github.com/Rivalz-ai/framework-be/modules/marketplace/dto"
	"github.com/Rivalz-ai/framework-be/modules/marketplace/service"
	"github.com/gofiber/fiber/v2"
)

//	@BasePath	/marketplace
//	@Tags			marketplace
//	@Accept			json
//	@Produce		json
//
// @Summary Check balance of wallet
// @Description Check balance of wallet
// @Accept json
// @Produce json
// @Param wallet_address query string true "Wallet address"
// @Param consumed_token query string true "Consumed token"
// @Param consumed_token_decimals query string true "Consumed token decimals"
// @Success 200 {object} dto.APIResponseSuccess "Success Response"
// @Failure 400 {object} dto.APIResponseError "Fail Response"
// @Router /marketplace/internal/consume-swap-wallet [post]
func (s *MarketplaceHandler) ConsumeSwapWallet(c *fiber.Ctx) error {
	wallet_address := c.Query("wallet_address")
	if wallet_address == "" {
		wallet_address = "0x"
	}
	consumed_token := c.Query("consumed_token")
	if consumed_token == "" {
		consumed_token = "0x"
	}
	consumed_token_decimals := utils.ItoInt(c.Query("consumed_token_decimals"))

	secret := c.Query("secret")
	if secret == "" || secret != s.sv.ExtendConfig.InternalSecret {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": "secret is required"})
	}

	marketService, err := service.NewMarketplaceService(s.sv)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	err = marketService.CheckBalanceOfWallet(wallet_address, consumed_token, consumed_token_decimals)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 1, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"code": 0, "message": "CheckBalanceOfWallet"})
}
