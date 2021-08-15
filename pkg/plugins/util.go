package plugins

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (p *PluginData) BuildSourceWithFile(file string) string {
	return filepath.Join(p.SourceFilesDir, file)
}

func (p *PluginData) BuildDestWithFile(file string) string {
	return filepath.Join(p.DestFilesDir, file)
}

func (p *PluginData) AddFileToMove(from string, to string, overwrite bool) {
	fmt.Printf("%#v", p.Files)
	p.Files = append(p.Files, FileToMove{
		From:      p.BuildSourceWithFile(from),
		To:        p.BuildDestWithFile(to),
		Overwrite: overwrite,
	})
}

func (p *PluginData) MoveFiles() {
	for _, ftm := range p.Files {
		ftm.Move()
	}
}

func (file FileToMove) Move() {
	fmt.Printf("FileToMove:%#v\n", file)

	sf, err := os.Open(file.From)
	Check(err)
	defer sf.Close()

	if _, err := os.Stat(file.To); err == nil {
		fmt.Printf("%v file already exists\n", file.To)

		if file.Overwrite == false {
			fmt.Printf("%v Overwrite not set, bailing.\n", file.To)
			return
		}

		fmt.Printf("%v Overwrite set, will be overwritten.\n", file.To)
	}

	df, err := os.Create(file.To)
	Check(err)
	defer df.Close()

	bytesCopied, err := io.Copy(df, sf)
	Check(err)

	fmt.Printf("Installed %v bytes:%v\n", file.To, bytesCopied)
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
