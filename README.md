# mgo

够浪用 [globalsign/mgo]("https://github.com/globalsign/mgo") 连接 mongodb 的实例

## usage

[go-example]("https://github.com/CARVIN94/go-example")

```go
package main

import (
	"os"

	"github.com/CARVIN94/mgo"
)

func init() {
	mgo.Connect(&mgo.Config{
		Hosts:    "localhost:27017",
		Database: "mqtt",
		// Timeout:  time.Second * 60,
	})
}

func main() {
	quit := make(chan os.Signal)
	<-quit
}
```
