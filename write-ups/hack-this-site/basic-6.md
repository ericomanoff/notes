Messed around with cyphering strings. Found

abc => ace

aaaa => abcd

dog => dpi

Noticed the first character wasn't changing. Then realized the second letter changed by one and the second by two...

We'll call this cypher a "incremental off by one"

To reverse this, we'll decrement each character in a string by its index.
