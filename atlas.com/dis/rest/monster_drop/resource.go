package monster_drop

import (
	"atlas-dis/database/monster_drop"
	"atlas-dis/domain"
	"atlas-dis/rest/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func GetMonsterDrops(_ logrus.FieldLogger, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var monsterDrops []domain.MonsterDrop
		if val, ok := vars["monsterId"]; ok {
			monsterId, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				json.ToJSON(&GenericError{Message: err.Error()}, rw)
				return
			}

			monsterDrops, err = monster_drop.GetDropsByMonsterId(db, uint32(monsterId))
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				json.ToJSON(&GenericError{Message: err.Error()}, rw)
				return
			}
		} else {
			var err2 error
			monsterDrops, err2 = monster_drop.GetAllMonsterDrops(db)
			if err2 != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				json.ToJSON(&GenericError{Message: err2.Error()}, rw)
				return
			}
		}

		var resultData []MonsterDropData
		for _, md := range monsterDrops {
			result := MonsterDropData{
				Id:   strconv.Itoa(int(md.Id())),
				Type: "com.atlas.dis.rest.attribute.MonsterDropAttributes",
				Attributes: MonsterDropAttributes{
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
		json.ToJSON(MonsterDropListDataContainer{Data: resultData}, rw)
	}
}
