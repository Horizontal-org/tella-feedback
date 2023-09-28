package middleware

import "github.com/gin-gonic/gin"
import "log"
import "time"



func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()
		// Set example variable
		c.Set("example", "12345")

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {		
		if c.GetHeader("X-Tella-Platform") == "wearehorizontal" {
			c.Next()
		} else {
			c.AbortWithStatus(422)
		}
	}
}