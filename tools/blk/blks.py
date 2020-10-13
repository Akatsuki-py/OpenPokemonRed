import os
import subprocess
import shutil

if os.path.exists("blk"):
    shutil.rmtree("blk")
os.mkdir("blk")

blks_path = "../../../PokemonRedAsset/blk"

blks = os.listdir(blks_path)

for blk in blks:
	if os.path.exists(blk):
		os.remove(blk)

for blk in blks:
	name = blk.replace(".blk", "")
	path = os.path.join(blks_path, blk)
	output = name + ".go"
	subprocess.run(['python3', 'blk.py', path])
	if os.path.exists(output):
		shutil.move(output, os.path.join("blk", output))
