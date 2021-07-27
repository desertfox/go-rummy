package rummy

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Config RummyConfig

func GoRummy() RummyConfig {
	fmt.Println("goRummy")

	Config = NewConfig()

	tasks := map[string]func(){
		"installBash": installBash,
		"installVim":  installVim,
	}

	for name, method := range tasks {
		fmt.Println(name)
		method()
	}

	return Config
}

func installVim() {
	installVimrc()
	installVimPlug()
}

func installVimrc() {
	fmt.Println("installVimrc")

	initConfigFile(Config.Vimrc, Config.Vimrc, Config.Voverwrite)
}

func installVimPlug() {
	var VimAutoloadPath = strings.Join([]string{os.Getenv("HOME"), Config.VimPlugDestFile}, "/")

	if _, err := os.Stat(VimAutoloadPath); os.IsNotExist(err) {
		fmt.Printf("vim plug.vim file already exists. bailing\n")
		return
	}

	resp, err := http.Get(Config.VimPlugUrl)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	err = ioutil.WriteFile(VimAutoloadPath, body, 0644)
	check(err)
}

func installBash() {
	installBashAliases(Config.Boverwrite)
}

func installBashAliases(OverwriteBashAliases bool) {
	fmt.Println("installBashAliases")

	initConfigFile(Config.BashAliases, Config.BashAliases, Config.Boverwrite)
}

func initConfigFile(source string, dest string, Overwrite bool) {
	fmt.Printf("source:%v, dest:%v, overwrite:%v\n", source, dest, Overwrite)

	source = strings.Join([]string{Config.DotfilesName, source}, "/")

	sourceFile, err := os.Open(source)
	check(err)
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
	check(err)
	defer destinationFile.Close()

	bytesCopied, err := io.Copy(destinationFile, sourceFile)
	check(err)

	fmt.Printf("Installed %v bytes:%v\n", dest, bytesCopied)
}

func check(e error) {
	if e != nil {
		fmt.Printf("%vn", e)
		panic(e)
	}
}
