# ccwc

Unix wc tool implemented from scratch with Go


This Repo hosts my solution to the wc coding challenge on https://codingchallenges.fyi/challenges/challenge-wc/ by John Cricket

## Usage
```
  -c	Count bytes
  -l	Count lines
  -m	Count characters
  -w	Count words
```

It supports piping from standard input as well

```bash
cat test.txt | go run main.go
```

## Try it out

1. Clone this repo with git.
2. You will find test.txt inside this repo.
3. Run ```go run main.go -l test.txt``` to count the lines in a file.
