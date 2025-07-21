package base

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type MiddlewareFn = func() error
type HTTPMiddleWareFn = func(http.Handler) http.Handler

func LoadENV() {
	err := godotenv.Load(os.ExpandEnv("/config/.env"))
	if err != nil {
		err := godotenv.Load(os.ExpandEnv(".env"))
		if err != nil {
			panic(err)
		}
	}
}
