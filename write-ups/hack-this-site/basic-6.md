# Basic 6


https://www.hackthissite.org/missions/basic/6/


Our first thought was to enter in some random strings and write their cyphered versions down, hoping to observe a pattern. The strings we tried are as follows

`abc => ace`

`aaaa => abcd`

`dog => dpi`

We noticed the first character wasn't changing. Then realized the second letter changed by one and the second by two. We wanted to figure out how the encryption handled symbols. Entering symbols, we were able to determine that the encryption was using the ordinal (numeric) value of ASCII to increment.

We'll call this cypher a "incremental off by one"

To reverse this, we'll decrement each character in a string by its index. As the question asks, we're to decrypt the string `c8e49=<@`. Instead of doing this by hand, we wrote the following script, which is in basic-y.py:

```
import sys


def main(cyphered):
    uncyphered = ""

    for i, character in enumerate(cyphered):
        uncyphered += chr(ord(character)-i)

    print("the flag is {}".format(uncyphered))


if __name__ == "__main__":
    main(sys.argv[1])
```


This returned us the key of `c7c15869`
