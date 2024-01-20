package App

import (
	grpcapp "github.com/Leleria/ServiceLoyalty/Internal/App/GRPc"
	l "github.com/Leleria/ServiceLoyalty/Internal/Service/Loyalty"
	sqlite "github.com/Leleria/ServiceLoyalty/Internal/Storage/SQLite"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	loyaltyService := l.New(log, storage)

	grpcApp := grpcapp.New(log, loyaltyService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
