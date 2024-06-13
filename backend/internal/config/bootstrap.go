package config

import (
	"1dv027/wt2/internal/dataaccess"
	authmiddleware "1dv027/wt2/internal/handler/middleware/auth"
	videogamesdatahandler "1dv027/wt2/internal/handler/video-games/data"
	videogamesparameterhandler "1dv027/wt2/internal/handler/video-games/parameters"
	videogamessearchhandler "1dv027/wt2/internal/handler/video-games/search"
	"1dv027/wt2/internal/model"
	"1dv027/wt2/internal/repository"
	dataservice "1dv027/wt2/internal/service/video-games/data"
	parametersservice "1dv027/wt2/internal/service/video-games/parameters"
	queryservice "1dv027/wt2/internal/service/video-games/query"
	searchservice "1dv027/wt2/internal/service/video-games/search"
	"net/http"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ContainerConfig struct {
	DbConnection       any
	ApiKey             string
	TypesenseSearchUrl string
	TypesenseApiKey    string
	HttpClient         *http.Client
}

// Setup for the IoC container
func SetupContainer(config ContainerConfig) *Container {
	c := NewContainer()

	c.ProvideTransient("AuthMiddleware", func() any {
		return authmiddleware.NewAuthMiddleware(config.ApiKey)
	})

	c.ProvideSingleton("VideoGamesDataAccess", func() any {
		return dataaccess.NewVideoGamesDataAccess(config.DbConnection.(clickhouse.Conn), config.TypesenseSearchUrl, config.TypesenseApiKey, config.HttpClient)
	})

	c.ProvideSingleton("VideoGamesRepository", func() any {
		dataaccess := c.Resolve("VideoGamesDataAccess", Singleton).(dataaccess.VideoGamesDataAccess)
		return repository.NewVideoGamesRepo(dataaccess)
	})

	c.ProvideSingleton("DataParameters", func() any {
		return model.NewDataParameters()
	})

	c.ProvideSingleton("QueryParamsValidator", func() any {
		dataParameters := c.Resolve("DataParameters", Singleton).(model.DataParameters)
		return queryservice.NewQueryParamValidator(dataParameters)
	})

	c.ProvideSingleton("VideoGamesDataService", func() any {
		paramValidator := c.Resolve("QueryParamsValidator", Singleton).(queryservice.QueryParamValidator)
		repo := c.Resolve("VideoGamesRepository", Singleton).(repository.VideoGamesRepo)
		return dataservice.NewVideoGamesDataService(paramValidator, repo)
	})

	c.ProvideSingleton("VideoGamesSearchService", func() any {
		repo := c.Resolve("VideoGamesRepository", Singleton).(repository.VideoGamesRepo)
		return searchservice.NewVideoGamesSearchService(repo)
	})

	c.ProvideSingleton("VideoGamesParameters", func() any {
		return model.NewDataParameters()
	})

	c.ProvideSingleton("VideoGamesParametersService", func() any {
		validParameters := c.Resolve("VideoGamesParameters", Singleton).(model.DataParameters)
		return parametersservice.NewVideoGamesParametersService(validParameters)
	})

	c.ProvideTransient("GetVideoGamesDataHandler", func() any {
		service := c.Resolve("VideoGamesDataService", Singleton).(dataservice.VideoGamesDataService)
		return videogamesdatahandler.NewGetVideoGamesDataHandler(service)
	})

	c.ProvideTransient("GetVideoGamesSearchHandler", func() any {
		service := c.Resolve("VideoGamesSearchService", Singleton).(searchservice.VideoGamesSearchService)
		return videogamessearchhandler.NewGetVideoGamesSearchHandler(service)
	})

	c.ProvideTransient("GetVideoGamesParameterHandler", func() any {
		service := c.Resolve("VideoGamesParametersService", Singleton).(parametersservice.VideoGamesParametersService)
		return videogamesparameterhandler.NewGetVideoGamesParametersHandler(service)
	})

	return c
}
