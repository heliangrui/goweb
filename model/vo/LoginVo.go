package vo

type LoginFrom struct {
	Username   string `form:"username" json:"username" binding:"required" errorMessage:"用户名填写错误！"`
	Password   string `form:"password" json:"password" binding:"required" errorMessage:"密码填写错误！"`
	RememberMe bool   `form:"rememberMe" json:"rememberMe"`
	captcha    string `form:"captcha" json:"captcha"`
}

type Point struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Type      int     `json:"type"`
}
