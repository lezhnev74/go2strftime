This library converts default [Go time format](https://go.dev/src/time/format.go)
to [strftime format](https://linux.die.net/man/3/strftime).

```go
layout := "Mon, 02 Jan 2006 15:04:05 -0700"
str2ftimeLayout, err := Go2Strftime(layout) // gives "%a, %d %b %Y %h:%M:%S %z"
```

As it seems the conversion cannot be correctly done as the features supported by both formats
are not equal. For example strftime cannot represent nanoseconds, but go cannot represent week numbers.

The package just replaces supported segments with strftime and does no any error checking otherwise.