package shortlinkrepository

import (
	"database/sql"
	"errors"
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

func (r *ShortLinkRepository) Save(url string) (uint64, error) {
	var id uint64
	err := r.db.QueryRow("INSERT INTO short_links (url) VALUES ($1) RETURNING id;",
		url,
	).Scan(&id)
	if err != nil {
		return 0, serviceerror.NewServerError()
	}

	return id, nil
}

func (r *ShortLinkRepository) GetByID(id uint64) (string, error) {
	var url string

	err := r.db.QueryRow("SELECT url FROM short_links WHERE id = $1", id).Scan(&url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", serviceerror.New(serviceerror.RecordNotFound)
		}
		return "", serviceerror.NewServerError()
	}

	return url, nil
}

func (r *ShortLinkRepository) GetByUrl(url string) (uint64, error) {
	var id uint64

	err := r.db.QueryRow("SELECT id FROM short_links WHERE url = $1", url).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, serviceerror.New(serviceerror.RecordNotFound)
		}
		return 0, serviceerror.NewServerError()
	}

	return id, nil
}
