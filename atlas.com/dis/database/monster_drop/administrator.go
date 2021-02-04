package monster_drop

import (
	"atlas-dis/domain"
	"gorm.io/gorm"
)

func BulkCreateMonsterDrop(db *gorm.DB, monsterDrops []domain.MonsterDrop) error {
	//db.Begin()
	for _, md := range monsterDrops {
		m := &monsterDrop{
			MonsterId:       md.MonsterId(),
			ItemId:          md.ItemId(),
			MinimumQuantity: md.MinimumQuantity(),
			MaximumQuantity: md.MaximumQuantity(),
			Chance:          md.Chance(),
		}

		err := db.Create(m).Error
		if err != nil {
			//db.Rollback()
			return err
		}
	}
	//db.Commit()
	return nil
}
