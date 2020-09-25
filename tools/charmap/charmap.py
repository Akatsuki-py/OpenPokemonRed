import cv2
import os

os.mkdir("char")
# https://www.spriters-resource.com/fullview/8734/
img = cv2.imread("original.png")

# en
charmap = {
    2: 0x80, # "A"
    28: 0x9a,  # "("
    34: 0xa0, # "a...z"
    66: 0xe0, # "'"
    67: 0xe1, # "pk"
    68: 0xe2, # "mn"
    69: 0xe3, # "-"
    70: 0xd3, # "'r"
    71: 0xd2, # "'m"
    72: 0xe6, # "?"
    73: 0xe7, # "!"
    74: 0xe8, # "."
    88: 0xf6, # 0
    98: 0x7f, # " "
    37*6 + 6: 0xc0, # "AOUaou" 
    37*7 + 12: 0xe9, # "&" 
}

def get_charcode(count):
    if count >= 2 and count < 2+26:
        delta = count - 2
        return charmap[2] + delta
    elif count >= 28 and count < 34:
        delta = count - 28
        return charmap[28] + delta
    elif count >= 34 and count < 34+26+6:
        delta = count - 34
        return charmap[34] + delta
    elif count >= 74 and count <= 87:
        delta = count - 74
        return charmap[74] + delta
    elif count >= 88 and count <= 88 + 9:
        delta = count - 88
        return charmap[88] + delta
    elif count >= 37*6 + 6 and count <= 37*6 + 6 + 5:
        delta = count - 37*6 + 6
        return charmap[37*6 + 6] + delta
    elif not count in charmap:
        return -1
    else:
        return charmap[count]

count = 0
for h in range(8):
    for w in range(37):
        x0 = 29+(16*w)
        x1 = x0 + 8
        y0 = 27+16*h
        y1 = y0+8

        tile = img[y0:y1, x0:x1]

        code = get_charcode(count)
        if code == -1:
            count += 1
            continue
        cv2.imwrite("./char/{}.png".format(code), tile)
        count += 1

# border
def border(h, w, code):
    x0 = 31+(16*w)
    x1 = x0 + 8
    y0 = 16+16*h
    y1 = y0+8

    tile = img[y0:y1, x0:x1]
    cv2.imwrite("./char/{}.png".format(code), tile)

h, w, code = 10, 6, 0x79
border(h, w, code)
h, w, code = 10, 7, 0x7a
border(h, w, code)
h, w, code = 10, 8, 0x7b
border(h, w, code)
h, w, code = 11, 6, 0x7c
border(h, w, code)
h, w, code = 12, 6, 0x7d
border(h, w, code)
h, w, code = 12, 8, 0x7e
border(h, w, code)
