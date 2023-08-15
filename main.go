package main

import (
    "github.com/gin-gonic/gin"
    "github.com/Horizontal-org/tella-feedback/pkg/common/db"
    "github.com/Horizontal-org/tella-feedback/pkg/common/email"
		"github.com/Horizontal-org/tella-feedback/pkg/opinions"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)
    m := email.Init()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "port": port,
            "dbUrl": dbUrl,
        })
    })

		opinions.RegisterRoutes(r, h, m)

    r.Run(port)
}