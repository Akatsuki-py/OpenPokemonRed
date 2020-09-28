package txt

import "pokered/pkg/store"

var RAM = map[string](func() string){
	"PLAYER": func() string { return store.PlayerName },
	"RIVAL":  func() string { return store.RivalName },
	"TMName": func() string { return store.TMName },
}
