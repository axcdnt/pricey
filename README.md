# pricey

Pricey is a pet project to explain crawling concepts with a "real" usecase. It was born after discussing Black Friday with a group of friends.
It's backed by the cool library [colly](https://github.com/gocolly/colly).

Pricey is not ambitious. It must be useful for monitoring prices on e-commerces on the web.

## Building
```
go build .
```
## Usage
```
$ ./pricey --help
Usage of ./pricey:
  -currency string
    	the currency (US$ (US), R$ (BR)...
  -interval duration
    	the time interval for crawling (default 5s)
  -pattern string
    	the corresponding css rule <required>
  -target int
    	the desired price <required>
  -url string
    	the url to be crawled <required>
```

## Limitations
There's a known limitation while the currency format. The *target* argument expects an int, so if you have a value like  $1,499.00, simply pass 149900. You know, floats and money :P

The cli interface is kind of *geeky*, asking a few params. This could be improved via configuration and defaults.

## Contributions
I try to implement only the minimal to be useful. If you want to leave a contribution, PRs are welcome. Try to be organized, even though the project doesn't have a guideline yet.

Be nice and help people! :P