package text

const (
	fast   uint = 1
	normal uint = 3
	slow   uint = 5
)

// AtOnce trueなら一気にテキストを表示 最優先
var AtOnce bool = false

// Option 設定の文字の速さ
var Option uint = normal

// Frame 1文字ごとに待機するフレーム
var Frame uint = 0

// FrameCounter フレームごとにデクリメント
var FrameCounter uint = 0

func IsDelay() bool {
	return true
}
