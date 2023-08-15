package opinions

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/Horizontal-org/tella-feedback/pkg/common/models"
    "github.com/Horizontal-org/tella-feedback/pkg/common/email"
)

type AddOpinionRequestBody struct {
    Text       string `json:"text" binding:"required"`
    Platform      string `json:"platform" binding:"required,oneof=IOS ANDROID WEB"`
}

func (h handler) AddOpinion(c *gin.Context) {
    body := AddOpinionRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var opinion models.Opinion

    opinion.Text = body.Text
    opinion.Platform = body.Platform

    if result := h.DB.Create(&opinion); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }


    go email.ReportFeedback(h.MailDialer, &opinion)

    c.JSON(http.StatusCreated, &opinion)
}