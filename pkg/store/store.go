package store

import "github.com/hajimehoshi/ebiten"

var PlayerName = "NINTEN"
var RivalName = "SONY"

var TileMap, _ = ebiten.NewImage(8*20*2, 8*18*2, ebiten.FilterDefault)

var TMName = ""

// D730 :
// bit 0: NPCスプライトがスクリプトによって動かされているか(scripted NPC)
// bit 1: ???
// bit 2: 方向キーが押されたかの判定に OverworldLoop で使われている
// bit 5: キー入力を無視する
// bit 6: 1なら テキスト出力時に文字ごとに遅延を生じない
// bit 7: キー入力がゲーム内で勝手に入れられているか(simulated joypad)
var D730 byte
