package text

// Delay text print
type Delay struct {
	Option       uint // 設定の文字の速さ
	AtOnce       bool // trueなら一気にテキストを表示
	Frame        uint // 1文字ごとに待機するフレーム
	FrameCounter uint // Vblankごとにデクリメント
}
