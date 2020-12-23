import os
import shutil
import subprocess

os.chdir("/Users/akatsuki/Dev/PokemonRed/tools")
if os.path.exists("public"):
    shutil.rmtree("public")
os.mkdir("public")

path = "../../PokemonRedAsset"

dirs = os.listdir(path)
for p in [".git", "public", "blk", "blocksets", ".DS_Store", "tools", "tilecoll"]:
    dirs.remove(p)

for d in dirs:
    files = os.listdir(os.path.join(path, d))

    for f in files:
        shutil.copyfile(os.path.join(path, d, f), os.path.join("public", f))

subprocess.run(['statik', '-src=public'])

if os.path.exists("../pkg/data/statik"):
    shutil.rmtree("../pkg/data/statik")
shutil.copytree("statik", "../pkg/data/statik")

shutil.rmtree("public")
shutil.rmtree("statik")
