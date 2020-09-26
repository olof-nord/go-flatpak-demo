package main

import (
    "flag"
    "github.com/olof-nord/go-flatpak-demo/helloworld"
    "github.com/olof-nord/go-flatpak-demo/hellofile"
)

func main() {
    extra := flag.Bool("extra", false, "This flag is for displaying extra output")

    // Parse command line arguments (if any)
    flag.Parse()

	helloworld.Say()

    if(*extra) {
        hellofile.Say()
    }
}
