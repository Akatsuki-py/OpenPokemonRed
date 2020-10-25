package song

import (
	"pokered/pkg/audio"
	"pokered/pkg/data/worldmap"
)

// MapMusics MapID -> MusicID
// ref: MapSongBanks
var MapMusics = map[int]int{
	worldmap.PALLET_TOWN:     audio.MUSIC_PALLET_TOWN,
	worldmap.VIRIDIAN_CITY:   audio.MUSIC_CITIES1,
	worldmap.PEWTER_CITY:     audio.MUSIC_CITIES1,
	worldmap.CERULEAN_CITY:   audio.MUSIC_CITIES2,
	worldmap.LAVENDER_TOWN:   audio.MUSIC_LAVENDER,
	worldmap.VERMILION_CITY:  audio.MUSIC_VERMILION,
	worldmap.CELADON_CITY:    audio.MUSIC_CELADON,
	worldmap.FUCHSIA_CITY:    audio.MUSIC_CITIES2,
	worldmap.CINNABAR_ISLAND: audio.MUSIC_CINNABAR,
	worldmap.INDIGO_PLATEAU:  audio.MUSIC_INDIGO_PLATEAU,
	worldmap.SAFFRON_CITY:    audio.MUSIC_CITIES1,
	worldmap.ROUTE_1:         audio.MUSIC_ROUTES1,
	worldmap.ROUTE_2:         audio.MUSIC_ROUTES1,
	worldmap.ROUTE_3:         audio.MUSIC_ROUTES3,
	worldmap.ROUTE_4:         audio.MUSIC_ROUTES3,
	worldmap.ROUTE_5:         audio.MUSIC_ROUTES3,
	worldmap.ROUTE_6:         audio.MUSIC_ROUTES3,
	worldmap.ROUTE_7:         audio.MUSIC_ROUTES3,
	worldmap.ROUTE_8:         audio.MUSIC_ROUTES3,
	worldmap.ROUTE_9:         audio.MUSIC_ROUTES3,
	worldmap.ROUTE_10:        audio.MUSIC_ROUTES3,
}
