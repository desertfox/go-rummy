package plugins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type PluginMove struct {
	sourcedir, sourcefile, destdir, dest string
	overwrite                            bool
}

func (p *PluginMove) Move() {
	fmt.Printf("MovePluginFile: %v\n", p)

	sourceFile, err := os.Open(strings.Join([]string{p.sourcedir, p.sourcefile}, "/"))
	Check(err)
	defer sourceFile.Close()

	dest := strings.Join([]string{p.destdir, p.dest}, "/")
	if _, err := os.Stat(dest); err == nil {
		fmt.Printf("%v file already exists\n", dest)

		if p.overwrite == false {
			fmt.Printf("%v Overwrite not set, bailing.\n", dest)
			return
		}

		fmt.Printf("%v Overwrite set, vimrc will be overwritten.\n", dest)
	}

	destinationFile, err := os.Create(dest)
	Check(err)
	defer destinationFile.Close()

	bytesCopied, err := io.Copy(destinationFile, sourceFile)
	Check(err)

	fmt.Printf("Installed %v bytes:%v\n", p.dest, bytesCopied)
}

func Check(e error) {
	if e != nil {
		fmt.Printf("%vn", e)
		panic(e)
	}
}
