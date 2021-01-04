# OpenPokémonRed

![CI](https://github.com/Akatsuki-py/PokemonRed/workflows/Go/badge.svg)
[![GitHub stars](https://img.shields.io/github/stars/Akatsuki-py/OpenPokemonRed)](https://github.com/Akatsuki-py/OpenPokemonRed/stargazers)
[![GitHub license](https://img.shields.io/github/license/Akatsuki-py/OpenPokemonRed)](https://github.com/Akatsuki-py/OpenPokemonRed/blob/master/LICENSE)

<img src="header.png" />

## 🔰 Description

**🚧 WARNING: This project is unofficial!! 🚧**

This project is open source re-implementation of [Pokémon Red](https://www.pokemon.com/us/pokemon-video-games/pokemon-red-version-and-pokemon-blue-version/).

The ROM version is not Japanese but English.

## 🏜 Roadmap

**This project is in the middle of development and is currently creating the underlying system part.**  

- [x] Text engine
- [x] Menu engine
- [x] Sprite engine
- [x] Statue screen
- [x] Trainer card
- [x] Title
- [x] Oak speech
- [x] Naming screen
- [ ] Combat system
- [ ] Wild pokemon battle
- [ ] Trainer battle
- [ ] Use item
- [ ] PC
- [ ] Pokecenter
- [ ] Pokemart
- [ ] Pokedex
- [ ] Save function
- [ ] Field move(Cut, Strength, Surf...)

## 🏞 Screenshots

**Screenshots as of v0.1.0**

<img src="./screenshots/title.png" width="360px" height="360px" /> &nbsp;&nbsp; <img src="./screenshots/oak_speech.png" width="360px" height="360px" />

<img src="./screenshots/overworld.png" width="360px" height="360px" /> &nbsp;&nbsp; <img src="./screenshots/status_screen.png" width="360px" height="360px" />

## 🎡 Try

**12/27: This cannot be played now! Since this repo uses Nintendo assets, I have temporarily suspended the publish of encrypted asset file. Please wait until I come up with a good solution ...**

### Download

Please download binary from [Release](https://github.com/pokemium/OpenPokemonRed/releases) page.

### Build

Requirements:
- Go 1.15
- Windows10 or MacOS(<= Big Sur)

```sh
gh repo clone pokemium/OpenPokemonRed
cd OpenPokemonRed && make build # For Mac OS. If your OS is windows, please `make build-windows`.
```

## 💻 Key

| keyboard             | game pad      |
| -------------------- | ------------- |
| <kbd>&larr;</kbd>    | &larr; button |
| <kbd>&uarr;</kbd>    | &uarr; button |
| <kbd>&darr;</kbd>    | &darr; button |
| <kbd>&rarr;</kbd>    | &rarr; button |
| <kbd>S</kbd>         | A button      |
| <kbd>A</kbd>         | B button      |
| <kbd>Enter</kbd>     | Start button  |
| <kbd>Right shift</kbd> | Select button |