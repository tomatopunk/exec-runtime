package backend

import (
	"context"
	"exec-runtime/pkg/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type NexusBackend struct {
	endpoint      string
	imageCacheDir string
}

func NewNexusBackend(cfg config.ImageConfig) *NexusBackend {
	return &NexusBackend{
		endpoint:      cfg.Endpoint,
		imageCacheDir: cfg.ImageCacheDir,
	}
}

func (n *NexusBackend) PullImage(ctx context.Context, name string) (string, error) {
	// todo need implement checksum
	url := strings.TrimRight(n.endpoint, "/")
	url += "/" + name
	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}
	cachePath := filepath.Join(n.imageCacheDir, name)
	if err := os.MkdirAll(filepath.Dir(cachePath), 0755); err != nil {
		return "", fmt.Errorf("failed to create cache dir: %w", err)
	}
	out, err := os.Create(cachePath)
	if err != nil {
		return "", fmt.Errorf("failed to create cache file: %w", err)
	}
	defer out.Close()
	if _, err = io.Copy(out, resp.Body); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}
	if err := os.Chmod(cachePath, 0755); err != nil {
		return "", fmt.Errorf("failed to chmod cache file: %w", err)
	}
	return cachePath, nil
}
