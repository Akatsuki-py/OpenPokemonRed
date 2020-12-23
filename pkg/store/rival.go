package store

type RivalState struct {
	Name        string
	Starter     uint
	StarterName string
	PartyMons   [6]PartyMon
}

var Rival = RivalState{
	Name: "SONY",
}

func RivalPartyMonLen() int {
	for i, mon := range Rival.PartyMons {
		if !mon.Initialized {
			return i
		}
	}
	return 6
}
