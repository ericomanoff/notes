import sys


def main(cyphered):
    uncyphered = ""

    for i, character in enumerate(cyphered):
        uncyphered += chr(ord(character)-i)

    print("the flag is {}".format(uncyphered))


if __name__ == "__main__":
    main(sys.argv[1])
