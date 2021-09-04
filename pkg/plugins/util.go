package plugins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (p *PluginData) buildDestPath(destPath string, file string) string {
	return filepath.Join(destPath, file)
}

func (p *PluginData) AddConfigToCreate(source *string, to string, overwrite bool) {
	p.Configs = append(p.Configs, ConfigToCreate{
		Source:    source,
		To:        to,
		Overwrite: overwrite,
	})
}

func (p *PluginData) CreateConfigs() error {
	for _, ctc := range p.Configs {
		err := ctc.Create()
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctc ConfigToCreate) Create() error {
	if _, err := os.Stat(ctc.To); err == nil {
		fmt.Printf("%v already exists\n", ctc.To)
		if !ctc.Overwrite {
			fmt.Printf("%v Overwrite not set, bailing.\n", ctc.To)
			return nil
		}

		err := ctc.Backup()
		if err != nil {
			return err
		}
	}

	f, err := os.Create(ctc.To)
	if err != nil {
		return err
	}
	defer f.Close()

	bytesCopied, err := f.WriteString(*ctc.Source)
	if err != nil {
		return err
	}

	fmt.Printf("Installed %v bytes:%v\n", ctc.To, bytesCopied)

	return nil
}

func (ctc ConfigToCreate) Backup() error {
	return os.Rename(ctc.To, ctc.To+"-backup")
}

func DownloadFile(url string) ([]byte, error) {
	fmt.Printf("url:%v\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
