# Build Your Own WC

This is an implementation of the wc command line utility in golang

## How to build

Execute the following command from root of repository -

```
make build
```

## Commands

```

# Output number of bytes in file
> <PATH_TO_REPO>/bin/wcgo -c test.txt
        342190 test.txt

# Output number of lines in file
> <PATH_TO_REPO>/bin/wcgo -l test.txt
        7145 test.txt

# Output number of words in file
> <PATH_TO_REPO>/bin/wcgo -w test.txt
        58164 test.txt

# Output number of characters in file
> <PATH_TO_REPO>/bin/wcgo -m test.txt
        339292 test.txt

# Use pipe operator to pass input and run multiple commands at same time
> cat test.txt | <PATH_TO_REPO>/bin/wcgo -lcw
        1       52801   327900

# Run command on default params without invoking any flag
> <PATH_TO_REPO>/bin/wcgo test.txt
        7145    58164   342190 test.txt

```

## References

[StackOverflow Question Explaining Encoding, Unicode and UTF-8](https://stackoverflow.com/questions/643694/what-is-the-difference-between-utf-8-and-unicode)

[The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

[Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)

[Understanding Buffered I/O and bufio package in golang](https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762)
