# go-ucd-username

Generates ASCII URI-safe aliases for Unicode usernames containing non-ASCII characters.

## Important

This is work in progress. It should not be considered "production" ready yet.

## Install

So long as you already have [Go](http://www.golang.org) (and `make`) installed you should be able to simply type:

```
make bin
```

All of the dependencies are included in the [vendor](vendor) directory. If you don't have `make` installed you can get started by executin the relevant commands in the [Makefile](Makefile).

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

## See also

* https://github.com/cooperhewitt/go-ucd
* https://github.com/whosonfirst/go-sanitize
