# Unscramble a scrambled, encoded URL

## Problem statement

The problem statement follows.
I copied it literally because there's no good way to summarize and retain
any meaning.

Found at [https://dev.to/0shuvo0/can-you-solve-this-interview-problem-4gaa](https://dev.to/0shuvo0/can-you-solve-this-interview-problem-4gaa)

---

So few days ago I applied to a MNC and I had some interesting problems to
solve in given amount of time.
So let me share one with you, see if you can solve this.

### Problem Description

So let's say we have a URL something like this:
```
let url = "https://dev.to/0shuvo0"
```
Now they have converted the url to a base64 string.
So the URL have become something like this:
```
let url = "aHR0cHM6Ly9kZXYudG8vMHNodXZvMA=="
```
Now what they did was split the sting to multiple parts and convert into an array
```
let url = ["aHR0cH", "M6Ly9kZX", "YudG8vMHN", "odXZvMA=="]
```
But of course the madnesses doesn't stop here. Then the shuffled the array so it became something like this:
```
let url = ["M6Ly9kZX", "aHR0cH", "odXZvMA==", "YudG8vMHN"]
```
And lastly they have converted that array to a string.
So here is your input

let url = `[ "6Ly9kZXYudG", "9jb21tZW5", "8vMHNodXZvMC", "aHR0cHM", "0LzFqZTFt" ]`

Use the input to find the original URL programmatically you have 45 Minutes to do it.
Useful JavaScript functions that can help you

You can convert your array to sting by calling join method on it. Eg.
```
let urlStr = url.join("")
```

You can use atob function to decode the base64 string.

```
let decoded = atob(urlStr)
```

Now go and and see if you can solve this. Best of luck 

---

## Analysis

The problem statement assumes the use of JavaScript.
I'll be using Go, because I'm not good at JavaScript.

There's three parts to this:

1. Permutation of the re-arranged substrings
2. Base64 decoding of the concatenation of each permutation
3. Deciding whether or not the decoded string represents a URL

It took me far longer than 45 minutes to get a permutation-generator working.
My first cut was a bizarre iterative thing.
My second attempt is much clearer, recursive, using a Go channel to return
concatenated permuted substrings.

I used the Go `base64` standard package to do the decoding.
Here's where a problem shows up.

`let url = "aHR0cHM6Ly9kZXYudG8vMHNodXZvMA=="` gets decoded correcty
by the `base64` package, which claims to do Base64 "as specified by RFC 4648".
The `base64` package also has "standard encoding", "URL encoding",
"raw standard encoding" and "raw URL encoding".
The example encoded URL has "==" on the end, indicating that padding was encoded into it.
But URLs of length that's a multiple of 3 won't have an "=" or "==" on the end,
whether the encoder used padding or not.

Should a decoder program look for an "=" in the permuted substrings and decode
based on that?
I chose to assume that URLs were originally Base64-encoded with padding.

Deciding whether or not the decoded string represents a URL causes another problem.
URLs can have all kinds of sizes and shapes, beginning with the "scheme".
Should a decoder program only consider `https://` URLs,
or should it also consider `http://` URLs?
What about `ftp://` URLs, or `file://` URLs?
I chose to consider only URLs with `http` or `https` schemes.

There's just no good way to say "this is a good URL" or "this isn't a URL"
for some arbitrary URL.
For example, take "https://axxxmercialy.com"

In the problem statement's argot, this could be given to a decoder like:

```
let url = `[ "Y29t", "eHh4", "aHR0cHM6Ly9h", "bWVyY2lhbHku" ]
```

This could plausibly decode to all these valid-format URLs:

```
https://acomxxxmercialy.
https://acommercialy.xxx
https://axxxcommercialy.
https://axxxmercialy.com
https://amercialy.comxxx
https://amercialy.xxxcom
```

Even if you decide on a small list of allowable top-level domains,
you can't legitimately distinguish between 2 of them.
I assume that https://acommercialy.xxx` is a very different web site than
`https://axxxmercialy.com`.
I chose to output all decoded strings that had no unprintable characters in them.

Just to note some things that might be considered as "optimizations":
It would be possible to find the first substring of the original
Base64 encoded URL by looking for substrings with 'aHR0cDov' or 'aHR0cHM6',
but there's no minimum length of a substring given, so probably best to ignore this.
Similarly, a substring ending in "=" could be taken as the final substring
of the original encoded string.

## Interview Analysis

This is a terrible interview question.

- It's absolutely disconnected from what anyone does at work.
- It has a very short time limit,
encouraging hacks you wouldn't want to see in production.
- "Base64" is not a single encoding,
which limits what a candidate's code can acheive.
- It relies on using a lot of library code,
which means that what the candidate can demonstrate is limited.
- The only real coding that a candidate can do is writing a permutor.
- It relies on everyone assuming the same things,
like which Base64 variant you're using,
and what a URL looks like.

Worst of all, there's no good way to decide if you've got the correct
permutation given the loose format(s) that a URL can take.

Don't use this problem in your interviews, if you're a hiring manager.
Point out all the flaws if you're a candidate,
and further, consider it a red flag for the potential employer.

I'll note again that it took me considerably longer than 45 minutes
to get a working decoder, and longer than that to get a decent,
non-hacky permutor going, plus some time to experiment with Base64 encoding
and decoding.
I would have failed this job interview, so take my analysis with a grain of salt.
