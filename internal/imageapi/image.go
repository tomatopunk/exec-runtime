package imageapi

import (
	"context"
	"exec-runtime/internal/backend"
	"exec-runtime/pkg/config"
	"fmt"
	"go.uber.org/zap"
	k8sCriApi "k8s.io/cri-api/pkg/apis/runtime/v1"
	"os"
	"path/filepath"
	"time"
)

type Service struct {
	config  *config.Runtime
	backend backend.ImageBackend
	logger  *zap.Logger
}

func (s *Service) ListImages(ctx context.Context, request *k8sCriApi.ListImagesRequest) (*k8sCriApi.ListImagesResponse, error) {
	var images []*k8sCriApi.Image
	entries, err := os.ReadDir(s.config.Image.ImageCacheDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read image cache dir: %w", err)
	}
	for _, entry := range entries {
		info, _ := entry.Info()

		images = append(images, &k8sCriApi.Image{
			Id:       entry.Name(),
			RepoTags: []string{entry.Name()},
			Size_:    uint64(info.Size()),
		})
	}
	return &k8sCriApi.ListImagesResponse{
		Images: images,
	}, nil
}

func (s *Service) ImageStatus(ctx context.Context, request *k8sCriApi.ImageStatusRequest) (*k8sCriApi.ImageStatusResponse, error) {
	imageName := request.Image.Image
	localPath := filepath.Join(s.config.Image.ImageCacheDir, imageName)
	fi, err := os.Stat(localPath)
	if err != nil {
		return &k8sCriApi.ImageStatusResponse{}, nil
	}
	return &k8sCriApi.ImageStatusResponse{
		Image: &k8sCriApi.Image{
			Id:       imageName,
			Size_:    uint64(fi.Size()),
			RepoTags: []string{imageName},
		},
	}, nil
}

func (s *Service) PullImage(ctx context.Context, request *k8sCriApi.PullImageRequest) (*k8sCriApi.PullImageResponse, error) {
	imageName := request.Image.Image
	image, err := s.backend.PullImage(ctx, imageName)
	if err != nil {
		s.logger.Error("Failed to pull image", zap.String("name", imageName), zap.Error(err))
	}
	return &k8sCriApi.PullImageResponse{
		ImageRef: image,
	}, nil
}

func (s *Service) RemoveImage(ctx context.Context, request *k8sCriApi.RemoveImageRequest) (*k8sCriApi.RemoveImageResponse, error) {
	imageName := request.Image.Image
	localPath := filepath.Join(s.config.Image.ImageCacheDir, imageName)
	if err := os.RemoveAll(localPath); err != nil {
		return nil, fmt.Errorf("pull image failed: %w", err)
	}
	return &k8sCriApi.RemoveImageResponse{}, nil
}

func (s *Service) ImageFsInfo(ctx context.Context, request *k8sCriApi.ImageFsInfoRequest) (*k8sCriApi.ImageFsInfoResponse, error) {
	used := uint64(0)
	filepath.Walk(s.config.Image.ImageCacheDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			used += uint64(info.Size())
		}
		return nil
	})
	return &k8sCriApi.ImageFsInfoResponse{
		ImageFilesystems: []*k8sCriApi.FilesystemUsage{
			{
				Timestamp: time.Now().UnixNano(),
				UsedBytes: &k8sCriApi.UInt64Value{
					Value: used,
				},
				FsId: &k8sCriApi.FilesystemIdentifier{
					Mountpoint: s.config.Image.ImageCacheDir,
				},
			},
		},
	}, nil
}

func NewService(config *config.Runtime, backend backend.ImageBackend, logger *zap.Logger) *Service {
	return &Service{
		config:  config,
		backend: backend,
		logger:  logger,
	}
}
