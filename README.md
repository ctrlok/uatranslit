uatranslit is a go library for transliteration from ukrainian to english based on [55-2010-п](http://zakon5.rada.gov.ua/laws/show/55-2010-п). 

There is two methods: ReplaceUARune and ReplaceUARunes, more into source code

# Usage

``` go
import "github.com/ctrlok/uatranslit/uatranslit"

func some(){
  letter := uatranslit.ReplaceUARunes([]rune("це якийсь текст"))
}
```

#TODO:
* add methods with io
* add documentation to godoc
* add cmd for converting
