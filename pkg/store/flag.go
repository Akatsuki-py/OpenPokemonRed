package store

type GameFlag struct {
	CD60   CD60
	D72C   D72C
	D72D   D72D
	D730   D730
	D736   D736
	Enable Enable
}

type Enable struct {
	NormalWarp bool
}

// CD60 wram.asm:wcd60
type CD60 struct {
	// bit 0: is player engaged by trainer (to avoid being engaged by multiple trainers simultaneously)
	EngagedByTrainer bool

	// bit 1: かいりきのアニメーション再生を待機中
	WaitBoulderAnim bool

	Bit2 bool

	// bit 3: 普通のPCを使っているか (0ならマサキのパソコンを使用している)
	UseGenericPC bool

	// bit 4: 1 -> .skipMovingSprites
	Bit4 bool

	// bit 5: 1なら menu で A/Bボタンが押された時にサウンドをならさない
	DontPlayMenuSound bool

	// bit 6: tried pushing against boulder once (you need to push twice before it will move)
	Bit6 bool
}

// D72C wram.asm:wd72c
type D72C struct {
	// bit 0: if not set, the 3 minimum steps between random battles have passed
	Bit0 bool

	// bit 1: セットされているならオーディオのフェードアウトを防ぐ
	DisturbAudioFadeout bool
}

// D72D wram.asm:wd72d
type D72D struct {
	// bit 5: don't make NPCs face the player when spoken to
	DontFacePlayer bool
}

// D730 wram.asm:wd730
type D730 struct {
	// bit 0: NPCスプライトがスクリプトによって動かされているか(scripted NPC)
	IsNPCScripted bool

	// bit 5: キー入力を無視する
	IgnoreKeyInput bool

	// bit 6: 1なら テキスト出力時に文字ごとに遅延を生じない
	DelayText bool

	// bit 7: キー入力がゲーム内で勝手に入れられているか(simulated joypad)
	IsSimulatedJoypad bool
}

// D736 wram.asm:wd72d
type D736 struct {
	// bit 0: check if the player is standing on a door and make him walk down a step if so
	Bit0 bool

	// bit 2: standing on a warp
	OnWarp bool

	// bit 6: jumping down a ledge / fishing animation
	InLedgeOrFishingAnim bool

	// bit 7: player sprite spinning due to spin tiles (Rocket hideout / Viridian Gym)
	Bit7 bool
}

var Flag = GameFlag{
	Enable: Enable{
		NormalWarp: true,
	},
}
