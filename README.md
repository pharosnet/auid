#### Description
 
* AUID is a id generator base on UUID for Go, and it itself is completely stable.
* AUID = xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx-xxxxxxxxxxx 
* Support pool.
* Support UUIDs.
* Thread safety.

#### Example

The simplest way:

```go
package main

import (
	"github.com/pharosnet/auid"
	"fmt"
)

func main() {
    fmt.Println(auid.NewAuid())
}

```

Pool:

```go
package main

import (
	"github.com/pharosnet/auid"
	"fmt"
)

func main() {
    fmt.Println(auid.NewAuidWithPool())
}

```