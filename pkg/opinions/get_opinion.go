package opinions

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Horizontal-org/tella-feedback/pkg/common/models"
)


// Add opinion godoc
//	@Summary		Get opinions
//	@Description	Get list of all opinions
//	@Tags			opinions
//	@Produce		json
//	@Param			X-Tella-Feedback	header	string	required	"query params"
//	@Success		200					{array}	models.Opinion
//	@Failure		400					
//	@Failure		404					
//	@Failure		500					
//	@Router			/opinions [get]
func (h handler) GetOpinion(c *gin.Context) {
    var opinions []models.Opinion


    if result := h.DB.Find(&opinions); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &opinions)
}