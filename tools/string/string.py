f = open('sprite.txt', 'r')

output = ""
constlist = f.readlines()
for constdata in constlist:
    target = constdata.replace("const", "").replace(" ", "").replace("\t", "").split(";")[0]
    output += "case {0}:\n\t\treturn \"{1}\"".format(target[:-1], target.replace("SPRITE_", "").lower()[:-1]) + "\n"

f.close()


with open("output.txt", mode='w') as f:
    f.write(output)
