import os
import subprocess

blocksets = [
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

blocksets_path = "../../../PokemonRedAsset/blocksets"

for blockset in blocksets:
    path = os.path.join(blocksets_path, blockset) + ".bst"
    subprocess.run(['python3', 'blockset.py', path])
