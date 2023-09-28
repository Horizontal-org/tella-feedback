package opinions

// @Description Create opinion request 
// @Description with text and platform
type AddOpinionRequestBody struct {
	Text       		string `json:"text" binding:"required"`
	Platform      string `json:"platform" binding:"required,oneof=IOS ANDROID WEB"`
}
