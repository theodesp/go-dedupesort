# Dedupesort
:wrench: A small function that takes an integer array 
removes duplicates and sorts the results

## Install
```console
go get github.com/theodesp/go-dedupesort/...
```

## Example Usage

```code
package main

import (
    "fmt"
    utils "github.com/theodesp/go-dedupesort"
)

func main() {
    array := []int{7, 2, 4, 2, 1, 2}
	fmt.Println(utils.DedupeSort(array)) // prints [1 2 4 7]
}
```

## Example
Run the example inside the cmd folder
```console
cd cmd
go run main.go
```

## License

MIT
