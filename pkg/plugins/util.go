package plugins

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-rummy/pkg/types"
)

func Move(ftm types.FileToMove) {
	fmt.Printf("FileToMove:%#v\n", ftm)

	sf, err := os.Open(ftm.From)
	Check(err)
	defer sf.Close()

	if _, err := os.Stat(ftm.To); err == nil {
		fmt.Printf("%v file already exists\n", ftm.To)

		if ftm.Overwrite == false {
			fmt.Printf("%v Overwrite not set, bailing.\n", ftm.To)
			return
		}

		fmt.Printf("%v Overwrite set, will be overwritten.\n", ftm.To)
	}

	df, err := os.Create(ftm.To)
	Check(err)
	defer df.Close()

	bytesCopied, err := io.Copy(df, sf)
	Check(err)

	fmt.Printf("Installed %v bytes:%v\n", ftm.To, bytesCopied)
}

func DownloadFile(url string, dest string) {
	fmt.Printf("url:%v, dest:%v\n", url, dest)
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		fmt.Printf("url:%v, dest:%v, File already exists. Bailing.\n", url, dest)
		return
	}

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	err = ioutil.WriteFile(dest, body, 0644)
	Check(err)
}

func Check(e error) {
	if e != nil {
		fmt.Printf("%vn", e)
		panic(e)
	}
}
