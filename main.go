package main

import (
    "github.com/gin-gonic/gin"
    "github.com/Horizontal-org/tella-feedback/pkg/common/db"
    "github.com/Horizontal-org/tella-feedback/pkg/common/middleware"
    "github.com/Horizontal-org/tella-feedback/pkg/common/email"
		"github.com/Horizontal-org/tella-feedback/pkg/opinions"
    "github.com/spf13/viper"    
)

//	@title			Tella feedback
//	@version		1.0
//	@description	microservice for receiving feedback about tella in all platforms
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	https://wearehorizontal.org

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.feedback.tella-app.org
// @schemes https
// @BasePath /

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)
    m := email.Init()

    r.Use(middleware.CheckHeader())

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "port": port,
            "dbUrl": dbUrl,
        })
    })

    // REGISTER ROUTES 
		opinions.RegisterRoutes(r, h, m)

    r.Run(port)
}