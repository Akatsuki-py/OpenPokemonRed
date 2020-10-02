package menu

type SelectMenu struct {
	Elm     []string
	z       uint // zindex 0:hide
	Wrap    bool
	Current uint
}

func (s *SelectMenu) Z() uint {
	return s.z
}

var CurSelectMenus = []SelectMenu{}
