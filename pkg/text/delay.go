package text

const (
	fast   uint = 1
	normal uint = 3
	slow   uint = 5
)

// AtOnce trueなら一気にテキストを表示 最優先
var AtOnce bool = false

// Speed 設定の文字の速さ
var Speed uint = normal

// FrameCounter フレームごとにデクリメント
var FrameCounter uint = 0

func InDelay() bool {
	return true
}

func delay() {
	if AtOnce {
		FrameCounter = 0
	}
	FrameCounter = Speed
}
