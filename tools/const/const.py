f = open('event.txt', 'r')

output = ""
constlist = f.readlines()
for constdata in constlist:
    output += constdata.replace("const", "").replace(" ", "").replace("\t", "").split(";")[0] + "\n"

f.close()


with open("output.txt", mode='w') as f:
    f.write(output)
