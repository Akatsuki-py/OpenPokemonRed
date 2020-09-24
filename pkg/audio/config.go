package audio

// Config audio config
type Config struct {
	volume uint // internal volue(NR50): [0, 7]
}

func (c *Config) setVolume(v uint) {
	if v > 7 {
		v = 7
	}
	c.volume = v
}

func (c *Config) setVolumeMax() {
	c.setVolume(7)
}

// pokedex, status,
func (c *Config) reduceVolume() {
	c.setVolume(3)
}

func (c *Config) offVolume() {
	c.setVolume(0)
}
