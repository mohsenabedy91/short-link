package shortlinkrepository

import (
	"database/sql"
	"errors"
	"short-link/internal/core/domain"
	"short-link/pkg/serviceerror"
)

type ShortLinkRepository struct {
	db *sql.DB
}

func NewShortLinkRepository(db *sql.DB) *ShortLinkRepository {
	return &ShortLinkRepository{
		db: db,
	}
}

func (r *ShortLinkRepository) Save(shortLink *domain.ShortLink) error {
	_, err := r.db.Exec("INSERT INTO short_links (path, url) VALUES ($1, $2);",
		shortLink.Path,
		shortLink.Url,
	)
	if err != nil {
		return serviceerror.NewServerError()
	}

	return nil
}

func (r *ShortLinkRepository) GetByShortPath(shortPath string) (*domain.ShortLink, error) {
	var shortLink domain.ShortLink

	err := r.db.QueryRow("SELECT url FROM short_links WHERE path = $1", shortPath).Scan(&shortLink.Url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, serviceerror.New(serviceerror.RecordNotFound)
		}
		return nil, serviceerror.NewServerError()
	}

	return &shortLink, nil
}
