package drop

import (
	"atlas-dis/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func InitResource(router *mux.Router, l logrus.FieldLogger, db *gorm.DB) {
	eRouter := router.PathPrefix("/monsters").Subrouter()
	eRouter.HandleFunc("/drops", GetMonsterDrops(l, db)).Queries("monsterId", "{monsterId}").Methods(http.MethodGet)
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func GetMonsterDrops(_ logrus.FieldLogger, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var monsterDrops []Model
		if val, ok := vars["monsterId"]; ok {
			monsterId, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				json.ToJSON(&GenericError{Message: err.Error()}, rw)
				return
			}

			monsterDrops, err = GetDropsByMonsterId(db, uint32(monsterId))
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				json.ToJSON(&GenericError{Message: err.Error()}, rw)
				return
			}
		} else {
			var err2 error
			monsterDrops, err2 = GetAllMonsterDrops(db)
			if err2 != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				json.ToJSON(&GenericError{Message: err2.Error()}, rw)
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

		rw.WriteHeader(http.StatusOK)
		json.ToJSON(DataListContainer{Data: resultData}, rw)
	}
}
