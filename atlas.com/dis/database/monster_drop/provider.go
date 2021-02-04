package monster_drop

import (
	"atlas-dis/domain"
	"gorm.io/gorm"
)

func GetAllMonsterDrops(db *gorm.DB) ([]domain.MonsterDrop, error) {
	var results []monsterDrop
	err := db.Find(&results).Error
	if err != nil {
		return nil, err
	}

	var monsterDrops []domain.MonsterDrop
	for _, r := range results {
		monsterDrops = append(monsterDrops, makeDrop(&r))
	}
	return monsterDrops, nil
}

func GetDropsByMonsterId(db *gorm.DB, monsterId uint32) ([]domain.MonsterDrop, error) {
	var results []monsterDrop
	err := db.Where(&monsterDrop{MonsterId: monsterId}).Find(&results).Error
	if err != nil {
		return nil, err
	}

	var monsterDrops []domain.MonsterDrop
	for _, r := range results {
		monsterDrops = append(monsterDrops, makeDrop(&r))
	}
	return monsterDrops, nil
}

func makeDrop(m *monsterDrop) domain.MonsterDrop {
	r := domain.NewMonsterDropBuilder(m.ID).
		SetMonsterId(m.MonsterId).
		SetItemId(m.ItemId).
		SetMinimumQuantity(m.MinimumQuantity).
		SetMaximumQuantity(m.MaximumQuantity).
		SetChance(m.Chance).
		Build()
	return r
}
