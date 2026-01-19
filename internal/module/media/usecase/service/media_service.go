package service

import (
	"context"
	"os"

	"github.com/arfanxn/welding/internal/infrastructure/id"
	mediaRepository "github.com/arfanxn/welding/internal/module/media/domain/repository"
	"github.com/arfanxn/welding/internal/module/media/usecase/dto"
	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
	"github.com/arfanxn/welding/pkg/fileutil"
	"github.com/arfanxn/welding/pkg/query"
	"go.uber.org/fx"
)

type MediaService interface {
	CreateFromMultipartFiles(ctx context.Context, _dto dto.CreateFromMultipartFiles) ([]*entity.Media, error)
	DestroyByQuery(ctx context.Context, q *query.Query) error
}

type mediaService struct {
	idService       id.IdService
	mediaRepository mediaRepository.MediaRepository
}

type NewMediaServiceParams struct {
	fx.In

	IdService       id.IdService
	MediaRepository mediaRepository.MediaRepository
}

func NewMediaService(params NewMediaServiceParams) MediaService {
	return &mediaService{
		idService:       params.IdService,
		mediaRepository: params.MediaRepository,
	}
}

// CreateFromMultipartFiles handles the creation of multiple media entries from uploaded files.
// It processes each multipart file, creates corresponding media entities, saves the files to storage,
// and persists the media metadata to the database.
// Returns the created media entities or an error if any operation fails.
func (s *mediaService) CreateFromMultipartFiles(ctx context.Context, _dto dto.CreateFromMultipartFiles) (medias []*entity.Media, err error) {
	// Process each file in the DTO
	for _, m := range _dto {
		// Extract file information from the multipart file
		multipartFileInfo, err := fileutil.NewMultipartFileInfo(m.File)
		if err != nil {
			return nil, err
		}

		// Initialize new media entity with basic information
		media := entity.NewMedia()
		media.Id = s.idService.Generate()
		media.ModelType = m.ModelType
		media.ModelId = m.ModelId
		media.Ulid = m.Ulid
		media.CollectionName = m.CollectionName
		media.Name = m.Name

		// Set file-specific properties
		media.FileName = multipartFileInfo.FileName
		media.MimeType = &multipartFileInfo.MimeType
		media.Size = int64(multipartFileInfo.Size)

		// Set media processing and storage properties
		media.Disk = m.Disk
		media.ConversionsDisk = m.ConversionsDisk
		media.Manipulations = m.Manipulations
		media.CustomProperties = m.CustomProperties
		media.GeneratedConversions = m.GeneratedConversions
		media.ResponsiveImages = m.ResponsiveImages
		media.OrderColumn = m.OrderColumn

		// Save the uploaded file to the local filesystem
		// TODO: implement abstract filesystem for better flexibility
		_path := "./storage/medias/" + media.Id
		_filePath := _path + "/" + media.FileName
		err = fileutil.SaveMultipartFile(m.File, _filePath)
		if err != nil {
			return nil, err
		}

		// Add the media to the results slice
		medias = append(medias, media)
	}

	// Persist all media entities to the database in a single transaction
	err = s.mediaRepository.SaveMany(medias)
	if err != nil {
		return nil, err
	}

	return medias, nil
}

func (s *mediaService) DestroyByQuery(ctx context.Context, q *query.Query) (err error) {
	medias, err := s.mediaRepository.Get(q)
	if err != nil {
		return err
	}

	// TODO: implement abstract filesystem for better flexibility
	for _, media := range medias {
		_path := "./storage/medias/" + media.Id
		err = os.RemoveAll(_path)
		if err != nil {
			return err
		}
	}

	err = s.mediaRepository.DestroyMany(medias)
	if err != nil {
		return err
	}

	return
}
