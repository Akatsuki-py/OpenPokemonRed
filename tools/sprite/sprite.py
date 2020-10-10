import cv2
import os
import shutil

if os.path.exists("result"):
    shutil.rmtree("result")
os.mkdir("result")
# https://www.spriters-resource.com/game_boy_gbc/pokemonredblue/sheet/8728/
img = cv2.imread("sprite.png")

a = [[9, 34], [9, 51], [9, 978], [9, 1071]]

# char
names_num = [
    ["red", 10],
    ["red_cycling", 10],
    ["blue", 10],
    ["oak", 10],
    ["mom", 4],
    ["daisy", 10],
    ["girl", 10],
    ["fisher2", 10],
    ["oak_aide", 10],
    ["bug_catcher", 10],
    ["black_hair_boy1", 10],
    ["gentleman", 10],
    ["nurse", 4],
    ["cable_club_woman", 4],
    ["mart_guy", 4],
    ["brunette_girl", 10],
    ["lass", 10],
    ["balding_guy", 4],
    ["little_girl", 10],
    ["gambler", 10],
    ["old_person", 4],
    ["black_hair_boy2", 10],
    ["fat_bald_guy", 10],
    ["young_boy", 4],
    ["gym_helper", 4],
    ["hiker", 10],
    ["mr_fuji", 10],
    ["rocket", 10],
    ["bike_shop_guy", 4],
    ["mom_geisha", 10],
    ["old_medium_woman", 4],
    ["guard", 4],
    ["fisher", 4],
    ["swimmer", 10],
    ["foulard_woman", 10],
    ["sailor", 10],
    ["waiter", 10],
    ["cook", 10],
    ["ss_captain", 4],
    ["rocker", 10],
    ["gameboy_kid", 4],
    ["mr_masterball", 4],
    ["medium", 10],
    ["erika", 10],
    ["giovanni", 10],
    ["biker", 10],
    ["blackbelt", 10],
    ["warden", 4],
    ["white_player", 4],
    ["lapras_giver", 4],
    ["lorelei", 10],
    ["bruno", 10],
    ["agatha", 10],
    ["lance", 10]
]
for h, name_num in enumerate(names_num):
    name = name_num[0]
    width = name_num[1]
    for w in range(width):
        x0 = 9+(17*w)
        x1 = x0 + 16
        y0 = 34 + (17*h)
        if h >= 9:
            y0 += 1
        y1 = y0 + 16

        tile = img[y0:y1, x0:x1]
        cv2.imwrite("./result/{}_{}.png".format(name, w), tile)

# poke
names = ["bird", "clefairy", "slowbro", "seel"]
for h in range(4):
    for w in range(9):
        x0 = 9+(17*w)
        x1 = x0 + 16
        y0 = 978+(17*h)
        y1 = y0+16

        tile = img[y0:y1, x0:x1]
        cv2.imwrite("./result/{}_{}.png".format(names[h], w), tile)

# misc
names = ["book_map_dex", "ball", "omanyte", "old_amber",
         "lying_old_man", "snorlax", "boulder", "clipboard", "paper_sheet"]
for w in range(9):
    x0 = 9+(17*w)
    x1 = x0 + 16
    y0 = 1071
    y1 = y0 + 16

    tile = img[y0:y1, x0:x1]
    cv2.imwrite("./result/{}.png".format(names[w]), tile)
