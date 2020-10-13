import sys

args = sys.argv[1:]

f = open(args[0], 'r')

def to_camelcase(snakecase: str) -> str:
    words = snakecase.lower().split('_')
    for i in range(len(words)):
        word = words[i].capitalize()
        words[i] = word
    return ''.join(words)

def encloseQuotes(target: str) -> str:
    return '"' + target.replace("\n", "") + '",' + "\n"

output = ""
constlist = f.readlines()
for constdata in constlist:
    snakecase = constdata.replace(" ", "").replace("\t", "")
    camelcase = encloseQuotes(to_camelcase(snakecase))
    print(camelcase)
    output += camelcase

f.close()

with open("output.txt", mode='w') as f:
    f.write(output)
