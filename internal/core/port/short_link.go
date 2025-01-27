package port

import "short-link/internal/core/domain"

type ShortLinkService interface {
	Create(link string) (string, error)
	GetByShortPath(shortPath string) (string, error)
}

type ShortLinkRepository interface {
	Save(shortLink *domain.ShortLink) error
	GetByShortPath(shortPath string) (*domain.ShortLink, error)
}
