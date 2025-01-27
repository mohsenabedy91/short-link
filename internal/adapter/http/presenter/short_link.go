package presenter

type ShortLink struct {
	ShortLink string `json:"shortLink" example:"bale-bank.ir/08c1b818"`
}

func ToResponseShortLink(shortLink string) *ShortLink {
	return &ShortLink{
		ShortLink: shortLink,
	}
}
