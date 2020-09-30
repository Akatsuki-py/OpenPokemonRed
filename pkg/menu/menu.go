package menu

const (
	// ChoseMenuItem Aボタンでアイテムを選択した or 2択menuの上
	ChoseMenuItem uint = iota + 1

	// CancelledMenu Bボタンでキャンセルした or 2択menuの下
	CancelledMenu
)

// MenuExitMethod プレイヤーが menu からどのように抜けたかを記録している
var MenuExitMethod uint

func maxZIndex() uint {
	return 1
}

// HandleMenuInput メニューでのキー入力に対処するハンドラ
// - - -
// INPUT: [wMenuWatchedKeys] = 反応する対象のキー入力 上下ボタンは必ず反応して選択オフセットを上下に移動させる
//
// OUTPUT:
// a = キー入力 [↓, ↑, ←, →, Start, Select, B, A]
// [wCurrentMenuItem] = 選択されたメニューアイテム
// [wMenuCursorLocation] = カーソルのあるタイルのアドレス
func HandleMenuInput() {}
