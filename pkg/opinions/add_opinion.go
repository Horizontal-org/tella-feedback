package opinions

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/Horizontal-org/tella-feedback/pkg/common/models"
    "github.com/Horizontal-org/tella-feedback/pkg/common/email"
)

// Add opinion godoc
//	@Summary		Create opinion
//	@Description	Create new opinion and send email
//	@Tags			opinions
//	@Accept			json
//	@Produce		json
//	@Param			X-Tella-Platform	header		string							required	"query params"
//	@Param			request				body		opinions.AddOpinionRequestBody	true		"tella params"
//	@Success		200					{object}	models.Opinion
//	@Failure		400					
//	@Failure		404					
//	@Failure		500					
//	@Router			/opinions [post]
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

    // magic line
    go email.ReportFeedback(h.MailDialer, &opinion)

    c.JSON(http.StatusCreated, &opinion)
}