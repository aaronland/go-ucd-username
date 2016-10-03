# go-ucd-username

Generates ASCII URI-safe aliases for Unicode usernames containing non-ASCII characters.

## Important

This is work in progress. It should not be considered "production" ready yet.

## What is this thing?

### Example

Given the string `mr. 😁 / ../test 🚀 㐖`

```
2016/10/03 07:51:15 PARSE mr. 😁 / ../test 🚀 㐖
2016/10/03 07:51:15 RUNE 0 U+006D 'm'
2016/10/03 07:51:15 RUNE 1 U+0072 'r'
2016/10/03 07:51:15 RUNE 2 U+002E '.'
2016/10/03 07:51:15 RUNE 2 U+002E '.' is punctuation	SKIPPING
2016/10/03 07:51:15 RUNE 3 U+0020 ' '
2016/10/03 07:51:15 RUNE 3 U+0020 ' ' is space	SKIPPING
2016/10/03 07:51:15 RUNE 4 U+1F601 '😁'
2016/10/03 07:51:15 RUNE 4 U+1F601 '😁' is not whitelisted	PROCESSING
2016/10/03 07:51:15 RUNE 4 U+1F601 '😁' return string 'GRINNING FACE WITH SMILING EYES'	PROCESSING
2016/10/03 07:51:15 RUNE 4:0 U+0047 'G'
2016/10/03 07:51:15 RUNE 4:1 U+0052 'R'
2016/10/03 07:51:15 RUNE 4:2 U+0049 'I'
2016/10/03 07:51:15 RUNE 4:3 U+004E 'N'
2016/10/03 07:51:15 RUNE 4:4 U+004E 'N'
2016/10/03 07:51:15 RUNE 4:5 U+0049 'I'
2016/10/03 07:51:15 RUNE 4:6 U+004E 'N'
2016/10/03 07:51:15 RUNE 4:7 U+0047 'G'
2016/10/03 07:51:15 RUNE 4:8 U+0020 ' '
2016/10/03 07:51:15 RUNE 4:9 U+0046 'F'
2016/10/03 07:51:15 RUNE 4:10 U+0041 'A'
2016/10/03 07:51:15 RUNE 4:11 U+0043 'C'
2016/10/03 07:51:15 RUNE 4:12 U+0045 'E'
2016/10/03 07:51:15 RUNE 4:13 U+0020 ' '
2016/10/03 07:51:15 RUNE 4:14 U+0057 'W'
2016/10/03 07:51:15 RUNE 4:15 U+0049 'I'
2016/10/03 07:51:15 RUNE 4:16 U+0054 'T'
2016/10/03 07:51:15 RUNE 4:17 U+0048 'H'
2016/10/03 07:51:15 RUNE 4:18 U+0020 ' '
2016/10/03 07:51:15 RUNE 4:19 U+0053 'S'
2016/10/03 07:51:15 RUNE 4:20 U+004D 'M'
2016/10/03 07:51:15 RUNE 4:21 U+0049 'I'
2016/10/03 07:51:15 RUNE 4:22 U+004C 'L'
2016/10/03 07:51:15 RUNE 4:23 U+0049 'I'
2016/10/03 07:51:15 RUNE 4:24 U+004E 'N'
2016/10/03 07:51:15 RUNE 4:25 U+0047 'G'
2016/10/03 07:51:15 RUNE 4:26 U+0020 ' '
2016/10/03 07:51:15 RUNE 4:27 U+0045 'E'
2016/10/03 07:51:15 RUNE 4:28 U+0059 'Y'
2016/10/03 07:51:15 RUNE 4:29 U+0045 'E'
2016/10/03 07:51:15 RUNE 4:30 U+0053 'S'
2016/10/03 07:51:15 RUNE 8 U+0020 ' '
2016/10/03 07:51:15 RUNE 8 U+0020 ' ' is space	SKIPPING
2016/10/03 07:51:15 RUNE 9 U+002F '/'
2016/10/03 07:51:15 RUNE 9 U+002F '/' is punctuation	SKIPPING
2016/10/03 07:51:15 RUNE 10 U+0020 ' '
2016/10/03 07:51:15 RUNE 10 U+0020 ' ' is space	SKIPPING
2016/10/03 07:51:15 RUNE 11 U+002E '.'
2016/10/03 07:51:15 RUNE 11 U+002E '.' is punctuation	SKIPPING
2016/10/03 07:51:15 RUNE 12 U+002E '.'
2016/10/03 07:51:15 RUNE 12 U+002E '.' is punctuation	SKIPPING
2016/10/03 07:51:15 RUNE 13 U+002F '/'
2016/10/03 07:51:15 RUNE 13 U+002F '/' is punctuation	SKIPPING
2016/10/03 07:51:15 RUNE 14 U+0074 't'
2016/10/03 07:51:15 RUNE 15 U+0065 'e'
2016/10/03 07:51:15 RUNE 16 U+0073 's'
2016/10/03 07:51:15 RUNE 17 U+0074 't'
2016/10/03 07:51:15 RUNE 18 U+0020 ' '
2016/10/03 07:51:15 RUNE 18 U+0020 ' ' is space	SKIPPING
2016/10/03 07:51:15 RUNE 19 U+1F680 '🚀'
2016/10/03 07:51:15 RUNE 19 U+1F680 '🚀' is not whitelisted	PROCESSING
2016/10/03 07:51:15 RUNE 19 U+1F680 '🚀' return string 'ROCKET'	PROCESSING
2016/10/03 07:51:15 RUNE 19:0 U+0052 'R'
2016/10/03 07:51:15 RUNE 19:1 U+004F 'O'
2016/10/03 07:51:15 RUNE 19:2 U+0043 'C'
2016/10/03 07:51:15 RUNE 19:3 U+004B 'K'
2016/10/03 07:51:15 RUNE 19:4 U+0045 'E'
2016/10/03 07:51:15 RUNE 19:5 U+0054 'T'
2016/10/03 07:51:15 RUNE 23 U+0020 ' '
2016/10/03 07:51:15 RUNE 23 U+0020 ' ' is space	SKIPPING
2016/10/03 07:51:15 RUNE 24 U+3416 '㐖'
2016/10/03 07:51:15 RUNE 24 U+3416 '㐖' is not whitelisted	PROCESSING
2016/10/03 07:51:15 RUNE 24 U+3416 '㐖' return string '㐖毒, AN OLD NAME FOR INDIA'	PROCESSING
2016/10/03 07:51:15 RUNE 24:0 U+3416 '㐖'
2016/10/03 07:51:15 RUNE 24:3 U+6BD2 '毒'
2016/10/03 07:51:15 RUNE 24:6 U+002C ','
2016/10/03 07:51:15 RUNE 24:7 U+0020 ' '
2016/10/03 07:51:15 RUNE 24:8 U+0041 'A'
2016/10/03 07:51:15 RUNE 24:9 U+004E 'N'
2016/10/03 07:51:15 RUNE 24:10 U+0020 ' '
2016/10/03 07:51:15 RUNE 24:11 U+004F 'O'
2016/10/03 07:51:15 RUNE 24:12 U+004C 'L'
2016/10/03 07:51:15 RUNE 24:13 U+0044 'D'
2016/10/03 07:51:15 RUNE 24:14 U+0020 ' '
2016/10/03 07:51:15 RUNE 24:15 U+004E 'N'
2016/10/03 07:51:15 RUNE 24:16 U+0041 'A'
2016/10/03 07:51:15 RUNE 24:17 U+004D 'M'
2016/10/03 07:51:15 RUNE 24:18 U+0045 'E'
2016/10/03 07:51:15 RUNE 24:19 U+0020 ' '
2016/10/03 07:51:15 RUNE 24:20 U+0046 'F'
2016/10/03 07:51:15 RUNE 24:21 U+004F 'O'
2016/10/03 07:51:15 RUNE 24:22 U+0052 'R'
2016/10/03 07:51:15 RUNE 24:23 U+0020 ' '
2016/10/03 07:51:15 RUNE 24:24 U+0049 'I'
2016/10/03 07:51:15 RUNE 24:25 U+004E 'N'
2016/10/03 07:51:15 RUNE 24:26 U+0044 'D'
2016/10/03 07:51:15 RUNE 24:27 U+0049 'I'
2016/10/03 07:51:15 RUNE 24:28 U+0041 'A'
```

Resulting in the string `mrgrinningfacewithsmilingeyestestrocketanoldnameforindia`

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
