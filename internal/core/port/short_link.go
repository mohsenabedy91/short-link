package port

type ShortLinkService interface {
	Create(link string) (string, error)
	GetByShortPath(shortPath string) (string, error)
}

type ShortLinkRepository interface {
	Save(url string) (uint64, error)
	GetByID(id uint64) (string, error)
	GetByUrl(url string) (uint64, error)
}
