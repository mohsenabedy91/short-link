package service

import (
	"errors"
	"fmt"
	"short-link/internal/core/config"
	"short-link/internal/core/domain"
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
	shortPath := helper.GenerateShortLink(url)

	_, getByShortPathErr := r.shortLinkRepository.GetByShortPath(shortPath)

	var se *serviceerror.ServiceError
	if errors.As(getByShortPathErr, &se) && se.GetErrorMessage() == serviceerror.RecordNotFound {
		shortLink := &domain.ShortLink{
			Path: shortPath,
			Url:  url,
		}

		if saveErr := r.shortLinkRepository.Save(shortLink); saveErr != nil {
			return "", saveErr
		}
	}

	return fmt.Sprintf("%s/%s", r.conf.Host, shortPath), nil
}

func (r *ShortLink) GetByShortPath(shortPath string) (string, error) {
	shortLink, err := r.shortLinkRepository.GetByShortPath(shortPath)
	if err != nil {
		return "", err
	}

	return shortLink.Url, nil
}
