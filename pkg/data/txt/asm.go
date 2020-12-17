package txt

var Asm = map[string](func() string){}

func RegisterAsmText(name string, fn func() string) {
	Asm[name] = fn
}
