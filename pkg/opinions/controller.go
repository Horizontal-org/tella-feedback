package opinions

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gopkg.in/gomail.v2"
)

type handler struct {
    DB *gorm.DB
    MailDialer *gomail.Dialer
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB, mail *gomail.Dialer) {
    h := &handler{
        DB: db,
        MailDialer: mail,
    }

    routes := r.Group("/opinions")
    routes.POST("/", h.AddOpinion)
    routes.GET("/", h.GetOpinion)
}