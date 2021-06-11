# scramble
Example Go lib to scramble a string.

## Example
```
"oscilloscope" => "llcoocspoise"
```

## Install
```
go get github.com/zherner/scramble@v0.0.1
```

## Use
```
...
import (
	"fmt"

	"github.com/zherner/scramble"
)

func main() {
	fmt.Println(scramble.Scramble("oscilloscope"))
}
...
```