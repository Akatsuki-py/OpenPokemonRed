import cv2
import numpy as np
import os

d = "result"
files = os.listdir(d)

green = [0x55, 0xc5, 0x25, 0xff]  # BGR
orange = [0x27, 0x7f, 0xff, 0xff]  # BGR

for f in files:
    name = os.path.join(d, f)
    img = cv2.imread(name, -1)

    width = img.shape[0]
    height = img.shape[1]
    out = np.zeros([16, 16, 4], np.uint8)
    
    for i in range(width):
        for j in range(height):
            b = img[i][j][0]
            g = img[i][j][1]
            r = img[i][j][2]
            if b == orange[0] and g == orange[1] and r == orange[2]:
                out[i][j][0] = 255
                out[i][j][1] = 255
                out[i][j][2] = 255
                out[i][j][3] = 0
            elif b == green[0] and g == green[1] and r == green[2]:
                out[i][j][0] = 255
                out[i][j][1] = 255
                out[i][j][2] = 255
                out[i][j][3] = 0
            else:
                out[i][j][0] = b
                out[i][j][1] = g
                out[i][j][2] = r
                out[i][j][3] = 0xff

    cv2.imwrite(name, out)

