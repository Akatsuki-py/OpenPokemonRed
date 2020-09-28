package text

import "pokered/pkg/store"

const (
	fast   uint = 1
	normal uint = 3
	slow   uint = 5
)

// AtOnce trueなら一気にテキストを表示 最優先
var AtOnce bool = false

// Speed 設定の文字の速さ
var Speed uint = normal

func printCharDelay() {
	if AtOnce {
		store.FrameCounter = 0
	}
	store.FrameCounter = Speed
}
