import os
import subprocess

tilecolls = [
    "cavern",
   	"cemetery",
   	"club",
   	"facility",
   	"forest",
   	"gate",
   	"gym",
   	"house",
   	"interior",
   	"lab",
   	"lobby",
   	"mansion",
   	"overworld",
   	"plateau",
   	"pokecenter",
   	"reds_house",
   	"ship_port",
   	"ship",
   	"underground",
]

tilecolls_path = "../../../PokemonRedAsset/tilecoll"

for tilecoll in tilecolls:
    path = os.path.join(tilecolls_path, tilecoll) + ".tilecoll"
    subprocess.run(['python3', 'tilecoll.py', path])
