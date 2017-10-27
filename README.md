# anybase

Encode and decode numbers in any base you want.

Humans use base10 numbers lots. The _alphabet_ for base10 is `0123456789`. There are other base encoding schemes like hex (base16) the _alphabet_ for base16 is `0123456789abcdef`.

With anybase you can encode with any _alphabet_ you want!

## example

```go
package main

import (
	"fmt"

	"github.com/Luke-Vear/anybase"
)

func main() {
	base56Alphabet := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"

	gen, _ := anybase.NewGenerator(base56Alphabet)
	num := gen.NumberFromUint(1<<64 - 1)

	fmt.Println(num) // 36psTsTZwTUH
}
```
