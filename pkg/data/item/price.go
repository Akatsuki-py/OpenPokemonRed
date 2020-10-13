package item

// ItemID -> Price
var itemMap = map[uint]uint{
	MASTER_BALL: 0,
	ULTRA_BALL:  1200,
	GREAT_BALL:  600,
	POKE_BALL:   200,
	TOWN_MAP:    0,
}

// Price get price from Item ID
func Price(id uint) uint {
	return itemMap[id]
}
