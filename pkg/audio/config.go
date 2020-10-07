package audio

var volume uint // internal volue(NR50): [0, 7]

func setVolume(v uint) {
	if v > 7 {
		v = 7
	}
	volume = v
}

func setVolumeMax() {
	setVolume(7)
}

// pokedex, status,
func reduceVolume() {
	setVolume(3)
}

func offVolume() {
	setVolume(0)
}
