package domain

type MonsterDrop struct {
	id              uint32
	monsterId       uint32
	itemId          uint32
	minimumQuantity uint32
	maximumQuantity uint32
	chance          uint32
}

func (d MonsterDrop) MonsterId() uint32 {
	return d.monsterId
}

func (d MonsterDrop) ItemId() uint32 {
	return d.itemId
}

func (d MonsterDrop) MinimumQuantity() uint32 {
	return d.minimumQuantity
}

func (d MonsterDrop) MaximumQuantity() uint32 {
	return d.maximumQuantity
}

func (d MonsterDrop) Chance() uint32 {
	return d.chance
}

func (d MonsterDrop) Id() uint32 {
	return d.id
}

type monsterDropBuilder struct {
	id              uint32
	monsterId       uint32
	itemId          uint32
	minimumQuantity uint32
	maximumQuantity uint32
	chance          uint32
}

func NewMonsterDropBuilder(id uint32) *monsterDropBuilder {
	return &monsterDropBuilder{id: id}
}

func (m *monsterDropBuilder) SetMonsterId(monsterId uint32) *monsterDropBuilder {
	m.monsterId = monsterId
	return m
}

func (m *monsterDropBuilder) SetItemId(itemId uint32) *monsterDropBuilder {
	m.itemId = itemId
	return m
}

func (m *monsterDropBuilder) SetMinimumQuantity(minimumQuantity uint32) *monsterDropBuilder {
	m.minimumQuantity = minimumQuantity
	return m
}

func (m *monsterDropBuilder) SetMaximumQuantity(maximumQuantity uint32) *monsterDropBuilder {
	m.maximumQuantity = maximumQuantity
	return m
}

func (m *monsterDropBuilder) SetChance(chance uint32) *monsterDropBuilder {
	m.chance = chance
	return m
}

func (m *monsterDropBuilder) Build() MonsterDrop {
	return MonsterDrop{
		id:              m.id,
		monsterId:       m.monsterId,
		itemId:          m.itemId,
		minimumQuantity: m.minimumQuantity,
		maximumQuantity: m.maximumQuantity,
		chance:          m.chance,
	}
}
