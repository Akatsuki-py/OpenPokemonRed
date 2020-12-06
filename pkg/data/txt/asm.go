package txt

var Asm = map[string](func()){}

func RegisterAsmText(name string, fn func()) {
	Asm[name] = fn
}
