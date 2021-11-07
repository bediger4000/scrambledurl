# Unscramble a scrambled URL

Found at [https://dev.to/0shuvo0/can-you-solve-this-interview-problem-4gaa](https://dev.to/0shuvo0/can-you-solve-this-interview-problem-4gaa)

So few days ago I applied to a MNC and I had some interesting problems to
solve in given amount of time.
So let me share one with you, see if you can solve this.

## Problem Description

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
