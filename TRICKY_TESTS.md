# ✅ Tricky Cases – Input vs Expected Output (Corrected & Consistent)

All transformations `(up)`, `(low)`, `(cap)` and `(up,n)/(low,n)/(cap,n)`
apply **only to the previous word(s)** — never the next word.

After all transformations:
- apply `a → an` rule (if the next word begins with a vowel or `h`)
- fix punctuation spacing
- fix quote spacing

---

## ✅ ARTICLES

### Input
I am a orange.
I am A orange.
I am A ORANGE.
I am AN ORANGE.
I am AN phone.
I am a phone.
I am A phone.

### Expected Output
I am an orange.
I am an orange.
I am AN ORANGE.
I am AN ORANGE.
I am a phone.
I am a phone.
I am A phone.

---

## ✅ CASE COMMANDS

### Input
this is a (cap) apple and an (cap) banana.
make these words (up,2) louder please.
quietly (low,3) YELLING AFTER NOW.
check john doe (cap,2) now.

### Expected Output
this is An apple and A banana.
make THESE WORDS louder please.
quietly YELLING AFTER NOW.
check John Doe now.

---

## ✅ HEX / BIN

### Input
the number 1E (hex) should become 30.
the number 1E (low) (hex) should become 30.
the number 1010 (bin) should become 10.
but 1G (hex) stays (hex) same because invalid.

### Expected Output
the number 30 should become 30.
the number 30 should become 30.
the number 10 should become 10.
but 1G stays same because invalid.

---

## ✅ QUOTES

### Input
' this is a quoted text , with punctuation ! '
' mixed CASE inside (low,3) HERE , and (cap) there . '

### Expected Output
'this is a quoted text, with punctuation!'
'mixed case inside HERe, And there.'

---

## ✅ PUNCTUATION

### Input
this is amazing!!!
wait ... are you sure?!
yes , i am ; absolutely .
i love apples , oranges ; bananas : and grapes !

### Expected Output
this is amazing!!!
wait... are you sure?!
yes, i am; absolutely.
i love apples, oranges; bananas: and grapes!

---

## ✅ MIXED ARTICLES + COMMANDS

### Input
asdf a (cap) orange is better than a (cap) apple.
a (low) ORANGE tastes like AN (low) apple.
an orange (cap,2) and a fruit salad (up,3).
AN ORANGE (cap) and AN BANANA (up).

### Expected Output
asdf An orange is better than An apple.
a Orange tastes like an apple.
An Orange and a FRUIT SALAD FRUIT salad.
AN ORANGE and A BANANA.

---

## ✅ MULTILINE

### Input
this is a test.
a new line starts here.
a (cap) orange
but a (up) apple

### Expected Output
this is a test.
a new line starts here.
A orange
but A apple

---

## ✅ QUOTES + COMMANDS

### Input
' I am a (cap) optimist , but a (up) realist . '
' a (cap) apple a day keeps a (low) doctor away . '

### Expected Output
'I am A optimist, but A realist.'
' A apple a day keeps a doctor away.'

---

## ✅ HEX/BIN + CASE

### Input
this hex number is 2A (hex), and this binary 1111 (bin).
a (cap) 2A (hex) banana and a 1111 (bin) orange.

### Expected Output
this hex number is 42, and this binary 15.
A 42 banana and a 15 orange.

---

## ✅ MIXED PUNCTUATION

### Input
a orange?! a phone! an apple... an ORANGE?!

### Expected Output
a orange?! a phone! an apple... an ORANGE?!

---

## ✅ EDGE CASES

### Input
a (cap) (up) orange and an (low) (cap) phone. ???
a a a an an a (cap,2) (low,3) orange.
a , a (cap) orange . a : an (low) apple !

### Expected Output
A orange and A phone. ???
a a a an an a orange.
a, An orange. a: an apple!

---

## ✅ ADVANCED MIXES

### Input
Behold 1f4 (hex) warriors and 101001 (bin) enemies marching ... slowly ,but surely !!
He whispered ' the END is NEAR ' (low, 3) ,or maybe not (up) ?
A hour ago, a elephant walked into a hotel (cap, 5) unexpectedly.
THIS (low, 4) LINE HAS (up, 2) mixed COMMANDS (low, 5) to test priority.
Three dots ... or maybe !? punctuation should test spacing ,and emotion !!
He claimed ' i am invincible (up, 2) and immortal ' (cap, 4) ,yet fell to a arrow.
He claimed'i AM Invincible And Immortal ', yet fell to an arrow.

### Expected Output
Behold 500 warriors and 41 enemies marching... slowly, but surely!!
He whispered ' the end is near ' ,or maybe NOT?
A hour ago, a Elephant Walked Into a hotel unexpectedly.
this line has mixed commands to test priority.
Three dots... or maybe!? punctuation should test spacing, and emotion!!
He claimed ' i AM INVINCIBLE And Immortal ' ,yet fell to a Arrow.
He claimed'i AM Invincible And Immortal ', yet fell to an arrow.
