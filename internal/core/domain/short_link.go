package domain

type ShortLink struct {
	Base
	Modifier

	Path string
	Url  string
}
