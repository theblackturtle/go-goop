# Go-Goop

This tool is a simple version of Goop (https://github.com/s0md3v/goop) rewriten in go.

Go-Goop can perform google searches without being blocked by the CAPTCHA or hitting any rate limits.

## How it works?
---
To make it work, set your Facebook Cookie as `FBCookie` environment variable.


## Usage
---
```
cat dorks.txt | go-goop
```

## Concurrency
---
You can set the concurrency level with the -c flag:
```
cat dorks.txt | go-goop -c 10
````

## Timeout
---
You can change the timeout by using the -t flag and specifying a timeout in seconds:
```
cat dorks.txt | go-goop -t 1
```

## Pages
---
You can set the number of pages to search (100 results per page) with -p flag:
```
cat dorks.txt | go-goop -p 10
```

## Unparsed results
---
Using -f flag if you want to get unparsed results
```
cat dorks.txt | go-goop -f
```
