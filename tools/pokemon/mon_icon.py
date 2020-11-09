import cv2
import os
import shutil

if os.path.exists("result"):
    shutil.rmtree("result")
os.mkdir("result")
# https://www.spriters-resource.com/fullview/64342/
img = cv2.imread("mon_icon.png", -1)

names = [
    "mon",
    "bird",
    "water",
    "fairy",
    "grass",
    "bug",
    "snake",
    "quadruped",
    "ball",
    "helix",
]

for i in range(10):
    name = names[i]
    if i < 8:
        for j in range(2):
            x0 = j*16
            y0 = i*16
            x1 = x0 + 16
            y1 = y0 + 16

            tile = img[y0:y1, x0:x1]
            cv2.imwrite("./result/{}_mon_icon_{}.png".format(name, j), tile)
    else:
        x0 = 0
        y0 = i*16
        x1 = x0 + 16
        y1 = y0 + 16

        tile = img[y0:y1, x0:x1]
        cv2.imwrite("./result/{}_mon_icon_{}.png".format(name, 0), tile)

        x0 = 0
        y0 = i*16-1
        x1 = x0 + 16
        y1 = y0 + 16

        tile = img[y0:y1, x0:x1]
        cv2.imwrite("./result/{}_mon_icon_{}.png".format(name, 1), tile)
