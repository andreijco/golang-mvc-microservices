package services

import (
	"golang/micro/src/api/config"
	"golang/micro/src/api/domain/github"
	"golang/micro/src/api/domain/repositories"
	"golang/micro/src/api/providers/github_provider"
	"golang/micro/src/api/utils/errors"
	"strings"
)


type repoService struct{}

type repoServiceInterface interface{
	CreateRepo(input repositories.CreateRepoRequest) ( *repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) ( *repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	} 

	request := github.CreateRepoRequest {
		Name: input.Name,
		Description: input.Description,
		Private: false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		 return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	return &repositories.CreateRepoResponse{Id: response.Id, Owner: response.Owner.Login, Name: response.Name}, nil
}