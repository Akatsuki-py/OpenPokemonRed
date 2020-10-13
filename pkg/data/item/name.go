package item

// ItemID -> Name
var nameMap = map[uint]string{
	MASTER_BALL: "MASTER BALL",
	ULTRA_BALL:  "ULTRA BALL",
	GREAT_BALL:  "GREAT BALL",
	POKE_BALL:   "POKé BALL",
	TOWN_MAP:    "TOWN MAP",
}

func Name(id uint) string {
	return nameMap[id]
}
