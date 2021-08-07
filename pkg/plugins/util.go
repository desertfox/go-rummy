package plugins

import (
	"fmt"
	"io"
	"os"
)

func Move(sourceFile string, destFile string, overwrite bool) {
	fmt.Printf("sourceFile:%v, destFile:%v, overwrite:%v\n", sourceFile, destFile, overwrite)

	sf, err := os.Open(sourceFile)
	Check(err)
	defer sf.Close()

	if _, err := os.Stat(destFile); err == nil {
		fmt.Printf("%v file already exists\n", destFile)

		if overwrite == false {
			fmt.Printf("%v Overwrite not set, bailing.\n", destFile)
			return
		}

		fmt.Printf("%v Overwrite set, vimrc will be overwritten.\n", destFile)
	}

	df, err := os.Create(destFile)
	Check(err)
	defer df.Close()

	bytesCopied, err := io.Copy(df, sf)
	Check(err)

	fmt.Printf("Installed %v bytes:%v\n", destFile, bytesCopied)
}

func Check(e error) {
	if e != nil {
		fmt.Printf("%vn", e)
		panic(e)
	}
}
