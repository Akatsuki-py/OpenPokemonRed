import os

sfxs = os.listdir("./assets")

for sfx in sfxs:
    new_name = sfx.lower().replace("sfx_", "")
    os.rename("./assets/" + sfx, "./assets/"  + new_name)
