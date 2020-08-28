# Three problems in an hour

See how much progress I can make on three problems within an hour, allotting twenty minutes for each problem.

1. String to Integer
2. Container with Most Water
3. Integer to Roman

## 1. String to Integer

I didn't spec out a full algorithm before I ran out of time, but I fleshed out all the finnicky stuff before the actual conversion; handling whitespace, checking for signs, and grabbing all the valid numbers.

This one felt really tricky to try and create an algorithm for, though, because I kept feeling like `int(cleanedString)` was cheating. It's proooobably not, though? I don't know, I had a hard time trying to figure out what the heart of this problem was, doing it in a vacuum without someone to ask, _"hey, are you looking more for how to convert a string to an int, or how to handle a bunch of finnicky sanitization problems?"_ In retrospect, it seems now like the sanitization aspect is probably the more important, but I got bogged down in trying to assess how low-level the string-to-integer conversion was supposed to be.

The code doesn't compile because - not counting the incompleteness of the solution - of some syntactical errors around my understanding of `byte` and how to use it. I've gotten tripped up a few times on forgetting that `someString[index]` yields a `byte` and not a `string`, so I'll often try to compare the two different types. In VSCode, the editor jogs my memory pretty quickly, but that's not the case with leetcode, etc. editors. I'll need to solidify my understanding of those minutiae (`byte('a')` and not `byte{"a"}`, for example - it's hard to overwrite my python knowledge of being able to use `'` and `"` more or less interchangably).

## 2. Container with Most Water

This is the problem that I feel the best on. I got the process solidified in words and conception, but was unable to troubleshoot/bugfix/make-it-work before the twenty minutes were up.

## 3. Integer to Roman

This was a bad one. This problem felt so fiddly and the only solution I could think of to work towards was just a plethora of `if`-statements. If I have a solution for an interview question that consists mostly of just a lot of `if`-statements, I tend to interpret that as a sign that I'm approaching the problem the wrong way, so I burned a lot of time turning the question over in my mind trying to come up with something halfway elegant. In the end, I just started spewing out `if`s, but I'd already lost a lot of time, and they're time-consuming to write, so I didn't get very close to finishing.
