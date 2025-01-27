package requests

type ShortLink struct {
	Url string `json:"url" binding:"required,url" example:"https://github.com/mohsenabedy91"`
}

type Redirect struct {
	ShortPath string `uri:"shortPath" binding:"required,len=8"`
}
