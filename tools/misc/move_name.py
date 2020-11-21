f = open('move_name.txt', 'r')

output = ""
constlist = f.readlines()
for constdata in constlist:
    movename = constdata.replace(" ", "").replace("\t", "").replace("\n", "")
    output += movename + ':"' + movename.replace("_", " ") + '",' + "\n"

f.close()


with open("output.txt", mode='w') as f:
    f.write(output)
