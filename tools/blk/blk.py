import sys
import subprocess

package = "package blk\n\n"

def output_blk(name: str, blk: bytes) -> str:
    var_start = "var {} = [...]byte{{\n".format(name.capitalize())
    var_content = ""
    for i, b in enumerate(blk):
        var_content += str(b) + ", "
        if i > 0 and i % 16 == 0:
            var_content += "\n"
    var_end = "\n}"
    return var_start + var_content + var_end


def main():
    args = sys.argv
    if len(args) < 2:
        return

    path = args[1]
    if not path.endswith(".blk"):
        print("InputError: input file must be .blk file")
        return

    name = (path.replace(".blk", "").split("/"))[-1]

    output = ""
    with open(path, mode='rb') as f:
        blk = list(f.read())
        output = output_blk(name, blk)

    path = "{}.go".format(name)
    with open(path, mode='w') as f:
        f.write(package + output)

    subprocess.run(['go', 'fmt', path])


if __name__ == "__main__":
    main()
