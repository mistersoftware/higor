# Higor

Dataframe for Golang, simple but powerful

## Install

```Bash
go get -v -u github.com/osmandi/higor
```

## Hello Gorld

```Go
package main

import (
	"fmt"

	"github.com/osmandi/higor"
)

func main() {
	fmt.Println(higor.PrintHelloHigor())
}
```

`Output:`
```Bash
Hello from higor
```

## Read a CSV

```Go
package main

import (
	"github.com/osmandi/higor"
)

func main() {
    dfHigor := higor.NewDataFrame("examples/data/example1.csv")
    dfHigor.ReadCSV()
    dfHigor.Head(2)
}
```
`Output:`
```Bash
index   |id   |name     |work_remotely   |salary     |age   |country_code
0       |1    |Hamish   |false           |$4528.90   |96    |PE
1       |2    |Anson    |NaN             |$1418.86   |NaN   |NaN
```

Credits to [mockaroo](https://www.mockaroo.com/) by CSV generator content on `examples/data` folder