package drop

import (
	"atlas-dis/json"
	"atlas-dis/rest"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const (
	getMonsterDrops = "get_monster_drops"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	eRouter := router.PathPrefix("/monsters").Subrouter()
	eRouter.HandleFunc("/drops", registerMonsterDrops(l, db)).Queries("monsterId", "{monsterId}").Methods(http.MethodGet)
}

func registerMonsterDrops(l logrus.FieldLogger, db *gorm.DB) http.HandlerFunc {
	return rest.RetrieveSpan(getMonsterDrops, func(span opentracing.Span) http.HandlerFunc {
		return handleMonsterDrops(l, db)(span)
	})
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func handleMonsterDrops(_ logrus.FieldLogger, db *gorm.DB) func(span opentracing.Span) http.HandlerFunc {
	return func(span opentracing.Span) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			var monsterDrops []Model
			if val, ok := vars["monsterId"]; ok {
				monsterId, err := strconv.ParseUint(val, 10, 32)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.ToJSON(&GenericError{Message: err.Error()}, w)
					return
				}

				monsterDrops, err = GetDropsByMonsterId(db, uint32(monsterId))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.ToJSON(&GenericError{Message: err.Error()}, w)
					return
				}
			} else {
				var err2 error
				monsterDrops, err2 = GetAllMonsterDrops(db)
				if err2 != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.ToJSON(&GenericError{Message: err2.Error()}, w)
					return
				}
			}

			var resultData []DataBody
			for _, md := range monsterDrops {
				result := DataBody{
					Id:   strconv.Itoa(int(md.Id())),
					Type: "com.atlas.dis.rest.attribute.MonsterDropAttributes",
					Attributes: Attributes{
						MonsterId:       md.MonsterId(),
						ItemId:          md.ItemId(),
						MaximumQuantity: md.MaximumQuantity(),
						MinimumQuantity: md.MinimumQuantity(),
						Chance:          md.Chance(),
					},
				}
				resultData = append(resultData, result)
			}

			w.WriteHeader(http.StatusOK)
			json.ToJSON(DataListContainer{Data: resultData}, w)
		}
	}
}
