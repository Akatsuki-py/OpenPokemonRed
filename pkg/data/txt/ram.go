package txt

import (
	"pokered/pkg/store"
	"strings"
)

var RAM = map[string](func() string){
	"PLAYER":   func() string { return store.Player.Name },
	"RIVAL":    func() string { return store.Rival.Name },
	"TMName":   func() string { return store.TMName },
	"PStarter": func() string { return strings.ToUpper(store.Player.StarterName) },
	"RStarter": func() string { return strings.ToUpper(store.Rival.StarterName) },
}
