package rest

import (
	"atlas-dis/rest/monster_drop"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

func CreateRestService(l *logrus.Logger, db *gorm.DB, ctx context.Context, wg *sync.WaitGroup) {
	go NewServer(l, ctx, wg, ProduceRoutes(db))
}

func ProduceRoutes(db *gorm.DB) func(l logrus.FieldLogger) http.Handler {
	return func(l logrus.FieldLogger) http.Handler {
		router := mux.NewRouter().PathPrefix("/ms/dis").Subrouter()
		router.Use(CommonHeader)

		eRouter := router.PathPrefix("/monsters").Subrouter()
		eRouter.HandleFunc("/drops", monster_drop.GetMonsterDrops(l, db)).Queries("monsterId", "{monsterId}").Methods(http.MethodGet)

		return router
	}
}
