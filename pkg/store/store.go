package store

import (
	_ "pokered/pkg/data/statik"

	"github.com/rakyll/statik/fs"
)

// FS statik filesystem
var FS, _ = fs.New()

// DelayFrames VBlank以外を拒否
var DelayFrames uint

// FrameCounter VBlankごとにデクリメント
// used to control letter print speed
var FrameCounter uint = 0

// DecFrameCounter decrement FrameCounter
// this function is called at every vBlank
func DecFrameCounter() {
	if FrameCounter > 0 {
		FrameCounter--
	}
}

// TMName wcf4b
var TMName = ""

// CD60 :
// bit 0: `TrainerEngage`でプレイヤーがトレーナーに見つかったか (複数のトレーナーに同時に見つかった時は発見されてないことにする)
// bit 1: かいりきのアニメーション再生を待機中
// bit 2: ???
// bit 3: 普通のPCを使っているか (0ならマサキのパソコンを使用している)
// bit 4: 1 -> .skipMovingSprites
// bit 5: 1なら menu で A/Bボタンが押された時にサウンドをならさない
// bit 6: 一度かいりきの岩を押してみた状態か (you need to push twice before it will move)
var CD60 byte

// D72C :
// bit 0: if not set, the 3 minimum steps between random battles have passed
// bit 1: セットされているならオーディオのフェードアウトを防ぐ
var D72C byte

// D72D :
// この変数は一時的なフラグの格納に使用されたり、トレードセンターまたはコロシアムにワープするときdestination mapとして使用される
// bit 0: トレードセンターでスプライトの方向が初期化されているときに立つフラグ
// bit 3: scripted warpを行うか（ポケモンタワーの上部からシオンタウンにワープするときに使用されます）
// bit 4: ダンジョンワープ中か
// bit 5: NPCが話しかけられたときにプレイヤーのほうを向かないようにするフラグ
// bit 6: ストーリー上で主要なバトルの開始時にセットされるが特になんの効果もないように思われる 任意のバトル終了時にリセットされる
// bit 7: トレーナーとのバトルの開始時にセットされるが特になんの効果もないように思われる バトル終了時にリセットされる
var D72D byte

// D730 :
// bit 0: NPCスプライトがスクリプトによって動かされているか(scripted NPC)
// bit 1: ???
// bit 2: 方向キーが押されたかの判定に OverworldLoop で使われている
// bit 5: キー入力を無視する
// bit 6: 1なら テキスト出力時に文字ごとに遅延を生じない
// bit 7: キー入力がゲーム内で勝手に入れられているか(simulated joypad)
var D730 byte

// D736 :
// bit 0: check if the player is standing on a door and make him walk down a step if so
// bit 1: 1ならプレイヤーは今、ドアから下に向かって歩いている状態である
// bit 2: (0なら)プレイヤーが warpマスの上に立っている
// bit 6: 段差をジャンプしているモーション中 / 釣りのモーション中に立つフラグ
// bit 7: player sprite spinning due to spin tiles (Rocket hideout / Viridian Gym)
var D736 byte

// BagItems items in bag
// [A@1, B@2, ...]
var BagItems = []string{}

// EventMap event ID -> flag
var EventMap = map[uint]bool{}
