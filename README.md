# ucd-username

## Important

Too soon. Move along.

## Install

```
make bin
```

## Usage

```
import (
	"flag"
	"fmt"		
	"log"
	"github.com/thisisaaronland/go-ucd-username"
	"strings"		
)

flag.Parse()
args := flag.Args()

username := strings.Join(args, " ")
	
safe, err := ucd.Username(username)

if err != nil {
   log.Fatal(err)
}
	
fmt.Println(safe)
```

## Tools

### ucd-username

```
./bin/ucd-username mr. 😁
mrgrinningfacewithsmilingeyes
```

