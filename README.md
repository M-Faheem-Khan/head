# HEAD

Go implementation of `head` program in linux(BSD). 

```bash
$ ./bin/main -l=10 ./src/main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)
```

### How to run?
`git clone https://github.com/M-Faheem-Khan/head.git`  
`cd head`  
`make run` - Build & Run the binary  

### How to replace the existing binary?
- Find the existing binaries
  - `find / -name head 2>/dev/null` - finds all files with the name head
  - most likely its at `/usr/bin/head` or `/bin/head`
- Replacing the binary `sudo cp ./head /usr/bin/head`

[//]: # (EOF)