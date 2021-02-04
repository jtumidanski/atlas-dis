package monster_drop

type MonsterDropDataContainer struct {
	Data MonsterDropData `json:"data"`
}

type MonsterDropListDataContainer struct {
	Data []MonsterDropData `json:"data"`
}

type MonsterDropData struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes MonsterDropAttributes `json:"attributes"`
}

type MonsterDropAttributes struct {
	MonsterId       uint32 `json:"monsterId"`
	ItemId          uint32 `json:"itemId"`
	MaximumQuantity uint32 `json:"maximumQuantity"`
	MinimumQuantity uint32 `json:"minimumQuantity"`
	Chance          uint32 `json:"chance"`
}
