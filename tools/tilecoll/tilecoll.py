import sys
import subprocess

package = "package tilecoll\n\n"


def output_tilecoll(name: str, tilecoll: bytes) -> str:
    var_start = "var {} = []byte{{\n".format(name.capitalize())
    var_content = ""
    for i, b in enumerate(tilecoll):
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
    if not path.endswith(".tilecoll"):
        print("InputError: input file must be .tilecoll file")
        return

    name = (path.replace(".tilecoll", "").split("/"))[-1]

    output = ""
    with open(path, mode='rb') as f:
        tilecoll = list(f.read())
        output = output_tilecoll(name, tilecoll)

    path = "{}.go".format(name)
    with open(path, mode='w') as f:
        f.write(package + output)

    subprocess.run(['go', 'fmt', path])


if __name__ == "__main__":
    main()
