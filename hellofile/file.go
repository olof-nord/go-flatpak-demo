package hellofile

import (
	"fmt"
	"github.com/dustin/go-humanize"
)

func Say() {
    // 82854982 bytes is 83 MB.
    fmt.Printf("The file is %s.", humanize.Bytes(82854982))
}
