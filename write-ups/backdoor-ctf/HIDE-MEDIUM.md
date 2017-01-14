# Backdoor Hidden Flag - Medium


https://backdoor.sdslabs.co/challenges/HIDE-MEDIUM


## Problem Introduction

The challenge provides a file, and its details are as follows:

```
n00b became depressed when 'Pro' found the flag in his binary in a matter of seconds. This time he hid the flag a little more securely. See if you can still find it:
```

To get this file to even run, we had to instal 32-bit architecture (as running 'file' command on it indicated that it was 32 bit).

Run the file -

Stupid message "It's not that easy as you think so"

GDB that shit:

What functions are there?

>info functions

Shows a couple of interesting functions: main, and print_flag - hmmm

Cant break at print flag since the application's execution is terminated before it hits.  Therefore, lets break where we KNOW
this program will go - the main() function.

>b main

Now run so we can hit our breakpoint:

>run

We will now be broken at main.  Now how do we get to that juicy print_flag function?  We jump.

> jump print_flag

We get this interesting tidbit:

(gdb) j print_flag
Continuing at 0x80484a2.
841f980abd04b26fe804ca0c207a574bef504cb6a3c3599a449e845ca993d2cf

Program received signal SIGSEGV, Segmentation fault.
0xffffd44c in ?? ()


Note the very key looking string, cause that's our key.
