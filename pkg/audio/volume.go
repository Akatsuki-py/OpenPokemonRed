package audio

var Volume uint // internal volue(NR50): [0, 7]

func setVolume(v uint) {
	if v > 7 {
		v = 7
	}
	Volume = v
}

func setVolumeMax() {
	setVolume(7)
}

func incrementVolume() {
	if Volume < 7 {
		setVolume(Volume + 1)
	}
}

func decrementVolume() {
	if Volume > 0 {
		setVolume(Volume - 1)
	}
}

// pokedex, status,
func reduceVolume() {
	setVolume(3)
}

func offVolume() {
	setVolume(0)
}
