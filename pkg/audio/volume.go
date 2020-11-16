package audio

const (
	maxVol    = 7
	widgetVol = 3
	minVol    = 0
)

// internal volue(NR50): [0, 7]
var volume uint

// Volume returns current game volume
func Volume() uint {
	return volume
}

func setVolume(v uint) {
	if v > maxVol {
		v = maxVol
	}
	volume = v
}

func setVolumeMax() {
	setVolume(maxVol)
}

func incrementVolume() {
	if volume < maxVol {
		setVolume(volume + 1)
	}
}

func decrementVolume() {
	if volume > minVol {
		setVolume(volume - 1)
	}
}

// ReduceVolume pokedex, status,
func ReduceVolume() {
	setVolume(widgetVol)
}

func offVolume() {
	setVolume(minVol)
}
