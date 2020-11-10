import cv2
import os
import shutil

if os.path.exists("result"):
    shutil.rmtree("result")
os.mkdir("result")
# https://www.spriters-resource.com/fullview/55182/
img = cv2.imread("oak_lecture.png")

names = [
    ["oak_lecture", 8],
    ["nidorino_lecture", 3],
    ["red_lecture", 3],
    ["rival_lecture", 8],
    ["red_shrink_lecture", 3],
    ["red_sprite_lecture", 3],
]

for i in range(6):
    name = names[i][0]
    num = names[i][1]

    for j in range(num):
        x0 = 8 + (j + (8-num))*64
        y0 = 24 + i*64
        x1 = x0 + 56
        y1 = y0 + 56

        print(x0, y0)
        tile = img[y0:y1, x0:x1]
        print("./result/{}_{}.png".format(name, j))
        cv2.imwrite("./result/{}_{}.png".format(name, j), tile)
