package opinions

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/Horizontal-org/tella-feedback/pkg/common/models"
)

func (h handler) GetOpinion(c *gin.Context) {
    var opinions []models.Opinion

    if result := h.DB.Find(&opinions); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusOK, &opinions)
}