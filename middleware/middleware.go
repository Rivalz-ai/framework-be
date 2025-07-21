package middleware

import (
	"fmt"
	"strings"

	"github.com/Rivalz-ai/framework-be/acl"
	"github.com/Rivalz-ai/framework-be/framework/utils"
	"github.com/Rivalz-ai/framework-be/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := ""
		authHeader_byte := c.Request().Header.Peek("Authorization")
		authHeader := string(authHeader_byte)

		//get token from url parameter
		if authHeader == "" {
			token := c.Query("token")
			if token == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code": 1,
					"msg":  "Authorization header is required",
				})
			}
			tokenString = token
		} else {
			arr := utils.Explode(authHeader, " ")
			if len(arr) != 2 {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"code": 1,
					"msg":  "Authorizationheader is invalid",
				})
			}
			tokenString = arr[1]
		}
		//
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})
		//
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 1,
				"msg":  "Invalid token",
			})
		}
		//
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 1,
				"msg":  "Invalid token claims",
			})
		}
		walletAddress := utils.ItoString(claims["walletAddress"])
		if len(walletAddress) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 1,
				"msg":  "Invalid wallet address",
			})
		}
		//
		roles, err := utils.ItoSliceString(claims["roles"])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 1,
				"msg":  "Invalid role",
			})
		}
		if len(roles) == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 1,
				"msg":  "Invalid role",
			})
		}
		//
		fmt.Println("method", c.Method())
		fmt.Println("route", c.Path())
		if acl.CheckACL(roles, c.Method(), c.Path()) == false {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": 1,
				"msg":  "Unauthorized access",
			})
		}
		//c.Locals("user_id", claims["user_id"])
		c.Locals("wallet_address", strings.ToLower(walletAddress))
		c.Locals("roles", roles)
		c.Locals("token", tokenString)
		c.Locals("user_id", utils.ItoString(claims["id"]))
		c.Locals("jwt", tokenString)
		return c.Next()
	}
}
func AllowCORS(sv *server.Server) {
	// use CORS
	sv.SVC.Use(cors.New())
	sv.SVC.Use(cors.New(cors.Config{
		AllowOrigins:     "*",                            // Chỉ cho phép origin này
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONAL", // Chỉ cho phép các method này
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false, // Cho phép gửi credentials (cookies, authorization headers, TLS client certs)
	}))
}
