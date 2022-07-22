package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/data/logs/app.log"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"foo": "bar"},
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		sugar.Infow("health check", "path", "/ping")
	})
	r.GET("/alert", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "alert",
		})
		sugar.Errorw("xxx failed", "error_msg", "id must be string.")
	})
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
