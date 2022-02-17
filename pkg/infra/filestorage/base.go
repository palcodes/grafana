package filestorage

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/grafana/grafana/pkg/infra/log"
)

var (
	nameRegex = regexp.MustCompile("[A-Za-z0-9!\\-_.*'()]+")
)

func getPath(folderPath string, name string) string {
	return fmt.Sprintf("%s%s%s", folderPath, Delimiter, name)
}

func getParentFolderPath(path string) string {
	if path == Delimiter || path == "" {
		return Delimiter
	}

	split := strings.Split(path, Delimiter)
	splitWithoutLastPart := split[:len(split)-1]
	return strings.Join(splitWithoutLastPart, Delimiter)
}

func extractName(path string) string {
	if path == Delimiter || path == "" {
		return ""
	}

	split := strings.Split(path, Delimiter)
	return split[len(split)-1]
}

type baseFilestorageService struct {
	log     log.Logger
	wrapped FileStorage
}

func validateName(name string) error {
	matches := nameRegex.MatchString(name)

	if !matches {
		return ErrNameInvalid
	}

	return nil
}

func validatePath(path string) error {
	if !filepath.IsAbs(path) {
		return ErrRelativePath
	}

	if filepath.Clean(path) != path {
		return ErrNonCanonicalPath
	}

	if strings.HasSuffix(path, Delimiter) {
		return ErrPathEndsWithDelimiter
	}

	if len(path) > 1000 {
		return ErrPathTooLong
	}

	return nil
}

func (b baseFilestorageService) Get(ctx context.Context, filePath string) (*File, error) {
	if err := validatePath(filePath); err != nil {
		return nil, err
	}

	return b.wrapped.Get(ctx, filePath)
}

func (b baseFilestorageService) Delete(ctx context.Context, filePath string) error {
	if err := validatePath(filePath); err != nil {
		return err
	}

	return b.wrapped.Delete(ctx, filePath)
}

func (b baseFilestorageService) Upsert(ctx context.Context, file *UpsertFileCommand) error {
	if err := validatePath(file.Path); err != nil {
		return err
	}

	return b.wrapped.Upsert(ctx, file)
}

func (b baseFilestorageService) ListFiles(ctx context.Context, folderPath string, recursive bool, cursor *Cursor) (*ListFilesResponse, error) {
	if err := validatePath(folderPath); err != nil {
		return nil, err
	}

	return b.wrapped.ListFiles(ctx, folderPath, recursive, cursor)
}

func (b baseFilestorageService) ListFolders(ctx context.Context, folderPath string) (*[]Folder, error) {
	if err := validatePath(folderPath); err != nil {
		return nil, err
	}

	return b.wrapped.ListFolders(ctx, folderPath)
}

func (b baseFilestorageService) CreateFolder(ctx context.Context, folderPath string, folderName string) error {
	if err := validatePath(getPath(folderPath, folderName)); err != nil {
		return err
	}

	return b.wrapped.CreateFolder(ctx, folderPath, folderName)
}

func (b baseFilestorageService) DeleteFolder(ctx context.Context, folderPath string) error {
	if err := validatePath(folderPath); err != nil {
		return err
	}

	return b.wrapped.DeleteFolder(ctx, folderPath)
}