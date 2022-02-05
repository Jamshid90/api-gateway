package services

import (
	"fmt"

	importpb "github.com/Jamshid90/api-getawey/genproto/import"
	postpb "github.com/Jamshid90/api-getawey/genproto/post"
	"github.com/Jamshid90/api-getawey/internal/config"
	"google.golang.org/grpc"
)

type Service interface {
	PostService() postpb.PostServiceClient
	PostImportService() importpb.ImportServiceClient
}

type service struct {
	postService       postpb.PostServiceClient
	postImportService importpb.ImportServiceClient
}

func New(config *config.Config) (Service, error) {
	connPostService, err := grpc.Dial(
		fmt.Sprintf("%s%s", config.PostService.Host, config.PostService.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connPostImportService, err := grpc.Dial(
		fmt.Sprintf("%s%s", config.PostImportService.Host, config.PostImportService.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &service{
		postService:       postpb.NewPostServiceClient(connPostService),
		postImportService: importpb.NewImportServiceClient(connPostImportService),
	}, nil
}

func (s *service) PostService() postpb.PostServiceClient {
	return s.postService
}

func (s *service) PostImportService() importpb.ImportServiceClient {
	return s.postImportService
}
