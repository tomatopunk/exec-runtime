package backend

import (
	"context"
	"errors"
	"exec-runtime/pkg/config"
)

type ImageInfo struct {
	ID   string
	Name string
	Size int64
}

type ImageFsInfo struct {
	UsedBytes uint64
}

type ImageBackend interface {
	PullImage(ctx context.Context, name string) (string, error)
}

func BuildImageBackend(config config.ImageConfig) (ImageBackend, error) {
	switch config.Mode {
	case "nexus":
		return NewNexusBackend(config), nil
	}
	return nil, errors.New("unsupported image mode")
}
