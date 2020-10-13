package move

// MoveID -> Name
var nameMap = map[uint]string{}

func Name(id uint) string {
	return nameMap[id]
}
