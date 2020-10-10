import cv2
import os
import shutil

if os.path.exists("result"):
    shutil.rmtree("result")
os.mkdir("result")
# https://www.spriters-resource.com/fullview/8733/
img = cv2.imread("trainercard.png")


leader = [
    "brock",
    "misty",
    "lt_surge",
    "erika",
    "koga",
    "sabrina",
    "blaine",
    "giovanni"
]

width = 16
height = 16

face = [
    [31, 103],
    [31+32, 103],
    [31+32+32, 103],
    [31+32+32+32, 103],
    [31, 127],
    [31+32, 127],
    [31+32+32, 127],
    [31+32+32+32, 127],
]

badge = [
    [31, 168],
    [31+32, 168],
    [31+32+32, 168],
    [31+32+32+32, 168],
    [31, 192],
    [31+32, 192],
    [31+32+32, 192],
    [31+32+32+32, 192],
]

# face
for i in range(8):
    name = leader[i]
    x0 = face[i][0]
    y0 = face[i][1]
    x1 = x0 + width
    y1 = y0 + height

    tile = img[y0:y1, x0:x1]
    cv2.imwrite("./result/{}_face.png".format(name), tile)

# badge
for i in range(8):
    name = leader[i]
    x0 = badge[i][0]
    y0 = badge[i][1]
    x1 = x0 + width
    y1 = y0 + height

    tile = img[y0:y1, x0:x1]
    cv2.imwrite("./result/{}_badge.png".format(name), tile)
