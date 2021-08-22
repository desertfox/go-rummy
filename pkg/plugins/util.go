package plugins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

//Make private?
func (p *PluginData) BuildDestWithFile(file string) string {
	return filepath.Join(p.DestFilesDir, file)
}

func (p *PluginData) AddConfigToCreate(source *string, to string, overwrite bool) {
	p.Configs = append(p.Configs, ConfigToCreate{
		Source:    source,
		To:        p.BuildDestWithFile(to),
		Overwrite: overwrite,
	})
}

func (p *PluginData) CreateConfigs() {
	for _, ctc := range p.Configs {
		ctc.Create()
	}
}

func (ctc ConfigToCreate) Create() {
	if _, err := os.Stat(ctc.To); err == nil {
		fmt.Printf("%v already exists\n", ctc.To)

		if ctc.Overwrite == false {
			fmt.Printf("%v Overwrite not set, bailing.\n", ctc.To)
			return
		}
	}

	f, err := os.Create(ctc.To)
	Check(err)
	defer f.Close()

	bytesCopied, err := f.WriteString(*ctc.Source)
	Check(err)

	fmt.Printf("Installed %v bytes:%v\n", ctc.To, bytesCopied)
}

func DownloadFile(url string) []byte {
	fmt.Printf("url:%v\n", url)

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	return body
}

func Check(e error) {
	if e != nil {
		fmt.Printf("%vn", e)
		panic(e)
	}
}
