package opinions

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := r.Group("/opinions")
    routes.POST("/", h.AddOpinion)
    routes.GET("/", h.GetOpinion)
}