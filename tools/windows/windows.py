import cv2
import sys

args = sys.argv

target = args[1]

img = cv2.imread(target)

x0 = 1
y0 = 31
x1 = x0 + 160
y1 = y0 + 144

tile = img[y0:y1, x0:x1]
cv2.imwrite(target, tile)
