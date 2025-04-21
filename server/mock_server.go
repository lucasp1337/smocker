package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lucasp1337/smocker/server/config"
	"github.com/lucasp1337/smocker/server/handlers"
	"github.com/lucasp1337/smocker/server/services"
	log "github.com/sirupsen/logrus"
)

func NewMockServer(cfg config.Config) (*http.Server, services.Mocks) {
	mockServerEngine := echo.New()
	persistence := services.NewPersistence(cfg.PersistenceDirectory)
	sessions, err := persistence.LoadSessions()
	if err != nil {
		log.Error("Unable to load sessions: ", err)
	}
	mockServices := services.NewMocks(sessions, cfg.HistoryMaxRetention, persistence)

	mockServerEngine.HideBanner = true
	mockServerEngine.HidePort = true
	mockServerEngine.Use(recoverMiddleware(), loggerMiddleware(), HistoryMiddleware(mockServices))

	handler := handlers.NewMocks(mockServices)
	mockServerEngine.Any("/*", handler.GenericHandler)

	mockServerEngine.Server.Addr = ":" + strconv.Itoa(cfg.MockServerListenPort)
	return mockServerEngine.Server, mockServices
}
