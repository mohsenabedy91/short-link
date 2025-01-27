package service

import (
	"errors"
	"fmt"
	"short-link/internal/core/config"
	"short-link/internal/core/port"
	"short-link/pkg/helper"
	"short-link/pkg/serviceerror"
)

type ShortLink struct {
	conf                config.ShortLink
	shortLinkRepository port.ShortLinkRepository
}

func NewShortLinkService(config config.ShortLink, shortLinkRepository port.ShortLinkRepository) *ShortLink {
	return &ShortLink{
		conf:                config,
		shortLinkRepository: shortLinkRepository,
	}
}

func (r *ShortLink) Create(url string) (string, error) {
	id, getByShortPathErr := r.shortLinkRepository.GetByUrl(url)

	var se *serviceerror.ServiceError
	if errors.As(getByShortPathErr, &se) {
		if se.GetErrorMessage() == serviceerror.ServerError {
			return "", getByShortPathErr
		}
		if se.GetErrorMessage() == serviceerror.RecordNotFound {
			var saveErr error
			id, saveErr = r.shortLinkRepository.Save(url)
			if saveErr != nil {
				return "", saveErr
			}
		}
	}

	shortPath := helper.IdToShortURL(id)

	return fmt.Sprintf("%s/%s", r.conf.Host, shortPath), nil
}

func (r *ShortLink) GetByShortPath(shortPath string) (string, error) {
	id := helper.ShortURLToID(shortPath)

	shortLink, err := r.shortLinkRepository.GetByID(id)
	if err != nil {
		return "", err
	}

	return shortLink, nil
}
