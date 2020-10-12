import cv2
import os
import shutil

if os.path.exists("result"):
    shutil.rmtree("result")
os.mkdir("result")
# https://www.spriters-resource.com/fullview/63033/
img = cv2.imread("tilesets.png")

x0 = 2
width = 16

def height(i: int) -> int:
    if i == 18:
        return 2
    return 6

def calc_y0(i: int) -> int:
    return 176 + 58*i

for i in range(19):
    x0 = 2
    y0 = calc_y0(i)
    x1 = x0 + width*8
    y1 = y0 + height(i)*8

    tile = img[y0:y1, x0:x1]
    cv2.imwrite("./result/tileset_{}.png".format(i), tile)

