# hyde

`hyde` is an deeply opinionated static site generator written in `go`.

## What does it support

* Templating ([html/template][^1])
* Markdown ([blackfriday][^2])
* Assets (SCSS and simple Javascript)

## Why doesn't it support `$foo`

Because I don't care about `$foo` and I don't use `$foo`. I know that may sound really arogant to you, but I was simply fed up with other existing static site generators and their way of just adding tons of features, for everyone and her dog. I just want to able to create a site, add some stylesheets and content. 

[^1]: https://golang.org/pkg/html/template
[^2]: https://github.com/russross/blackfriday
