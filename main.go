package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type RummyVim struct {
	vimrc, vimPlugUrl, vimPlugDestFile string
}

type RummyConfig struct {
	cwd, dotfilesName, bashAliases string
	rVim                           RummyVim
}

var Config RummyConfig

func main() {

	var (
		overwriteVimrc     = flag.Bool("ovimrc", false, "Overwrite .vimrc file")
		overwriteBashAlias = flag.Bool("obashalias", false, "Overwrite .bash_aliases file")
	)

	flag.Parse()

	installBash(overwriteBashAlias)

	installVimrc(overwriteVimrc)

	installVimPlug()
}

func init() {
	fmt.Println("Init")

	var (
		dotfiles        = "dot-files"
		bashAliases     = ".bash_aliases"
		vimrc           = ".vimrc"
		vimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
		vimPlugDestFile = "/.vim/autoload/plug.vim"
	)

	cwd, err := os.Getwd()
	check(err)

	Config = RummyConfig{cwd, dotfiles, bashAliases, RummyVim{vimrc, vimPlugUrl, vimPlugDestFile}}
}

func installBash(overwriteBashAlias *bool) {
	fmt.Println("installBash")

	installBashAliases(overwriteBashAlias)
}

func installBashAliases(overwriteBashAliases *bool) {
	fmt.Println("installBashAliases")

	initConfigFile(Config.bashAliases, Config.bashAliases, overwriteBashAliases)
}

func initConfigFile(source string, dest string, overwrite *bool) {
	source = strings.Join([]string{Config.dotfilesName, source}, "/")

	sourceFile, err := os.Open(source)
	check(err)
	defer sourceFile.Close()

	dest = strings.Join([]string{os.Getenv("HOME"), dest}, "/")
	if _, err := os.Stat(dest); err == nil {
		fmt.Printf("%v file already exists\n", dest)

		if *overwrite == false {
			fmt.Printf("%v overwrite not set, bailing.\n", dest)
			return
		}

		fmt.Printf("%v overwrite set, vimrc will be overwritten.\n", dest)
	}

	destinationFile, err := os.Create(dest)
	check(err)
	defer destinationFile.Close()

	bytesCopied, err := io.Copy(destinationFile, sourceFile)
	check(err)

	fmt.Printf("Installed %v bytes:%v\n", dest, bytesCopied)
}

func installVimrc(ovimrc *bool) {
	fmt.Println("installVimrc")

	initConfigFile(Config.rVim.vimrc, Config.rVim.vimrc, ovimrc)
}

func installVimPlug() {
	var vimAutoloadPath = strings.Join([]string{os.Getenv("HOME"), Config.rVim.vimPlugDestFile}, "/")

	if _, err := os.Stat(vimAutoloadPath); err == nil {
		fmt.Printf("vim plug.vim file already exists. bailing\n")
		return
	}

	resp, err := http.Get(Config.rVim.vimPlugUrl)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	err = ioutil.WriteFile(vimAutoloadPath, body, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
