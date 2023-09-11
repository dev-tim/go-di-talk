//go:build wireinject
// +build wireinject

package your_project

import (
	"gh.dev-tim/di-talk/config"
	"gh.dev-tim/di-talk/internal/server"
	"gh.dev-tim/di-talk/internal/storage"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, func(), error) {
	wire.Build(
		server.NewServer,
		storage.NewStorage,
		config.NewConfig,
		wire.Bind(new(storage.Storer), new(*storage.Storage)),
	)
	return nil, nil, nil
}

func InitializeMockServer() (*server.Server, func(), error) {
	wire.Build(
		server.NewServer,
		storage.NewMockStorage,
		config.NewConfig,
		wire.Bind(new(storage.Storer), new(*storage.MockStorage)),
	)
	return nil, nil, nil
}
