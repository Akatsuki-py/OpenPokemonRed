package sprdata

type SpriteID uint

const (
	SPRITE_RED SpriteID = iota + 1
	SPRITE_BLUE
	SPRITE_OAK
	SPRITE_BUG_CATCHER
	SPRITE_SLOWBRO
	SPRITE_LASS
	SPRITE_BLACK_HAIR_BOY_1
	SPRITE_LITTLE_GIRL
	SPRITE_BIRD
	SPRITE_FAT_BALD_GUY
	SPRITE_GAMBLER
	SPRITE_BLACK_HAIR_BOY_2
	SPRITE_GIRL
	SPRITE_HIKER
	SPRITE_FOULARD_WOMAN
	SPRITE_GENTLEMAN
	SPRITE_DAISY
	SPRITE_BIKER
	SPRITE_SAILOR
	SPRITE_COOK
	SPRITE_BIKE_SHOP_GUY
	SPRITE_MR_FUJI
	SPRITE_GIOVANNI
	SPRITE_ROCKET
	SPRITE_MEDIUM
	SPRITE_WAITER
	SPRITE_ERIKA
	SPRITE_MOM_GEISHA
	SPRITE_BRUNETTE_GIRL
	SPRITE_LANCE
	SPRITE_OAK_SCIENTIST_AIDE
	SPRITE_OAK_AIDE
	SPRITE_ROCKER
	SPRITE_SWIMMER
	SPRITE_WHITE_PLAYER
	SPRITE_GYM_HELPER
	SPRITE_OLD_PERSON
	SPRITE_MART_GUY
	SPRITE_FISHER
	SPRITE_OLD_MEDIUM_WOMAN
	SPRITE_NURSE
	SPRITE_CABLE_CLUB_WOMAN
	SPRITE_MR_MASTERBALL
	SPRITE_LAPRAS_GIVER
	SPRITE_WARDEN
	SPRITE_SS_CAPTAIN
	SPRITE_FISHER2
	SPRITE_BLACKBELT
	SPRITE_GUARD
	SPRITE_COP_GUARD
	SPRITE_MOM
	SPRITE_BALDING_GUY
	SPRITE_YOUNG_BOY
	SPRITE_GAMEBOY_KID
	SPRITE_GAMEBOY_KID_COPY
	SPRITE_CLEFAIRY
	SPRITE_AGATHA
	SPRITE_BRUNO
	SPRITE_LORELEI
	SPRITE_SEEL
	SPRITE_BALL
	SPRITE_OMANYTE
	SPRITE_BOULDER
	SPRITE_PAPER_SHEET
	SPRITE_BOOK_MAP_DEX
	SPRITE_CLIPBOARD
	SPRITE_SNORLAX
	SPRITE_OLD_AMBER_COPY
	SPRITE_OLD_AMBER
	SPRITE_LYING_OLD_MAN_UNUSED_1
	SPRITE_LYING_OLD_MAN_UNUSED_2
	SPRITE_LYING_OLD_MAN
)

func (s *SpriteID) String() string {
	switch *s {
	case SPRITE_RED:
		return "red"
	case SPRITE_BLUE:
		return "blue"
	case SPRITE_OAK:
		return "oak"
	case SPRITE_BUG_CATCHER:
		return "bug_catcher"
	case SPRITE_SLOWBRO:
		return "slowbro"
	case SPRITE_LASS:
		return "lass"
	case SPRITE_BLACK_HAIR_BOY_1:
		return "black_hair_boy_1"
	case SPRITE_LITTLE_GIRL:
		return "little_girl"
	case SPRITE_BIRD:
		return "bird"
	case SPRITE_FAT_BALD_GUY:
		return "fat_bald_guy"
	case SPRITE_GAMBLER:
		return "gambler"
	case SPRITE_BLACK_HAIR_BOY_2:
		return "black_hair_boy_2"
	case SPRITE_GIRL:
		return "girl"
	case SPRITE_HIKER:
		return "hiker"
	case SPRITE_FOULARD_WOMAN:
		return "foulard_woman"
	case SPRITE_GENTLEMAN:
		return "gentleman"
	case SPRITE_DAISY:
		return "daisy"
	case SPRITE_BIKER:
		return "biker"
	case SPRITE_SAILOR:
		return "sailor"
	case SPRITE_COOK:
		return "cook"
	case SPRITE_BIKE_SHOP_GUY:
		return "bike_shop_guy"
	case SPRITE_MR_FUJI:
		return "mr_fuji"
	case SPRITE_GIOVANNI:
		return "giovanni"
	case SPRITE_ROCKET:
		return "rocket"
	case SPRITE_MEDIUM:
		return "medium"
	case SPRITE_WAITER:
		return "waiter"
	case SPRITE_ERIKA:
		return "erika"
	case SPRITE_MOM_GEISHA:
		return "mom_geisha"
	case SPRITE_BRUNETTE_GIRL:
		return "brunette_girl"
	case SPRITE_LANCE:
		return "lance"
	case SPRITE_OAK_SCIENTIST_AIDE:
		return "oak_scientist_aide"
	case SPRITE_OAK_AIDE:
		return "oak_aide"
	case SPRITE_ROCKER:
		return "rocker"
	case SPRITE_SWIMMER:
		return "swimmer"
	case SPRITE_WHITE_PLAYER:
		return "white_player"
	case SPRITE_GYM_HELPER:
		return "gym_helper"
	case SPRITE_OLD_PERSON:
		return "old_person"
	case SPRITE_MART_GUY:
		return "mart_guy"
	case SPRITE_FISHER:
		return "fisher"
	case SPRITE_OLD_MEDIUM_WOMAN:
		return "old_medium_woman"
	case SPRITE_NURSE:
		return "nurse"
	case SPRITE_CABLE_CLUB_WOMAN:
		return "cable_club_woman"
	case SPRITE_MR_MASTERBALL:
		return "mr_masterball"
	case SPRITE_LAPRAS_GIVER:
		return "lapras_giver"
	case SPRITE_WARDEN:
		return "warden"
	case SPRITE_SS_CAPTAIN:
		return "ss_captain"
	case SPRITE_FISHER2:
		return "fisher2"
	case SPRITE_BLACKBELT:
		return "blackbelt"
	case SPRITE_GUARD:
		return "guard"
	case SPRITE_COP_GUARD:
		return "cop_guard"
	case SPRITE_MOM:
		return "mom"
	case SPRITE_BALDING_GUY:
		return "balding_guy"
	case SPRITE_YOUNG_BOY:
		return "young_boy"
	case SPRITE_GAMEBOY_KID:
		return "gameboy_kid"
	case SPRITE_GAMEBOY_KID_COPY:
		return "gameboy_kid_copy"
	case SPRITE_CLEFAIRY:
		return "clefairy"
	case SPRITE_AGATHA:
		return "agatha"
	case SPRITE_BRUNO:
		return "bruno"
	case SPRITE_LORELEI:
		return "lorelei"
	case SPRITE_SEEL:
		return "seel"
	case SPRITE_BALL:
		return "ball"
	case SPRITE_OMANYTE:
		return "omanyte"
	case SPRITE_BOULDER:
		return "boulder"
	case SPRITE_PAPER_SHEET:
		return "paper_sheet"
	case SPRITE_BOOK_MAP_DEX:
		return "book_map_dex"
	case SPRITE_CLIPBOARD:
		return "clipboard"
	case SPRITE_SNORLAX:
		return "snorlax"
	case SPRITE_OLD_AMBER_COPY:
		return "old_amber_copy"
	case SPRITE_OLD_AMBER:
		return "old_amber"
	case SPRITE_LYING_OLD_MAN_UNUSED_1:
		return "lying_old_man_unused_1"
	case SPRITE_LYING_OLD_MAN_UNUSED_2:
		return "lying_old_man_unused_2"
	case SPRITE_LYING_OLD_MAN:
		return "lying_old_man"
	}
	return "red"
}
