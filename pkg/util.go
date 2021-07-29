package rummy

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func GetWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

func MovePluginFile(source string, dest string, Overwrite bool) {
	fmt.Printf("source:%v, dest:%v, overwrite:%v\n", source, dest, Overwrite)

	source = strings.Join([]string{"dot-files", source}, "/")

	sourceFile, err := os.Open(source)
	Check(err)
	defer sourceFile.Close()

	dest = strings.Join([]string{os.Getenv("HOME"), dest}, "/")
	if _, err := os.Stat(dest); err == nil {
		fmt.Printf("%v file already exists\n", dest)

		if Overwrite == false {
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

	fmt.Printf("Installed %v bytes:%v\n", dest, bytesCopied)
}

func Check(e error) {
	if e != nil {
		fmt.Printf("%vn", e)
		panic(e)
	}
}
