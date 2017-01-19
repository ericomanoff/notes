# Basic 8


https://www.hackthissite.org/missions/basic/8/


We put in a random string and clicked `submit`, which brought us to a new page prompting us to click `here` to see our file. Doing so brings you to a page telling you how many characters your string input had.

What was interesting was the file type was of `shtml`. We looked into these file types and foundout that they are vulnerable to Server Side Includes, and allows us to arbitrarily execute code by using string like so `<!--#exec cmd="ls ../"-->`

To read more about Server Side Includes, see here https://en.wikipedia.org/wiki/Server_Side_Includes

By sending the above command, we got the response of

```
Hi, au12ha39vc.php index.php level8.php tmp!

Your name contains 39 characters.
```

Noticing that `au12ha39vc.php` looked odd, we can browse to that path and get back the flag which is `09c775a0`.
