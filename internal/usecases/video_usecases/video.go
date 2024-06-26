package video_usecases

import (
	"github.com/inview-team/sadko_indexer/internal/entities"
	"github.com/inview-team/sadko_indexer/internal/usecases/video_usecases/commands"
)

type VideoUsecases struct {
	Commands
}

type Commands struct {
	IndexVideo commands.IndexVideoCommand
	AddVectors commands.AddVectorsCommand
}

func NewVideoUsecases(vRepo entities.VideoRepository, tService entities.TagService) VideoUsecases {
	return VideoUsecases{
		Commands: Commands{
			IndexVideo: commands.NewIndexVideoCommand(vRepo, tService),
			AddVectors: commands.NewAddVectorCommand(vRepo),
		},
	}
}
