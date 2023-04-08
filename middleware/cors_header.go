package middleware

import (
	"github.com/gin-contrib/cors"
	"os"
	"strconv"
)

func GetCorsConfigs() cors.Config{
	corsConfig := cors.Config{}
	corsConfig.AllowOrigins = []string{os.Getenv("ALLOWORIGINS")}
	corsConfig.AddAllowMethods(os.Getenv("ADDALLOWMETHODS"))
	allowCredentials, _ := strconv.ParseBool(os.Getenv("ALLOWCREDENTIALS"))
	corsConfig.AllowCredentials = allowCredentials
	return corsConfig
}
