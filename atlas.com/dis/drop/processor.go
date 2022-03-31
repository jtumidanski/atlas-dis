package drop

import (
	"atlas-dis/database"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetAll(_ logrus.FieldLogger, db *gorm.DB) ([]Model, error) {
	return database.ModelSliceProvider[Model, entity](db)(getAll(), makeDrop)()
}

func GetForMonster(_ logrus.FieldLogger, db *gorm.DB) func(monsterId uint32) ([]Model, error) {
	return func(monsterId uint32) ([]Model, error) {
		return database.ModelSliceProvider[Model, entity](db)(getByMonsterId(monsterId), makeDrop)()
	}
}
