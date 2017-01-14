# Basic 4


https://www.hackthissite.org/missions/basic/4/


Viewed source and found the section in the HTML where the button's action was defined


```
<input type="hidden" name="to" value="sam@hackthissite.org" /><input type="submit" value="Send password to Sam" /></form></center><br /><br /><center><b>Password:</b><br />
```


We inspected the page and replaced `"sam@hackthissite.org"` with our own email. The email we received has the password.


`Password = f4ba2cc2`
