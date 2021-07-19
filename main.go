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

var dotfilePath string

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
	//Set cwd, make sure dot-files exists
	fmt.Println("Init")
	//make configurable
	const Dotfiles = "dot-files"

	cwd, err := os.Getwd()
	check(err)

	dotfilePath = strings.Join([]string{cwd, Dotfiles}, "/")
}

func installBash(overwriteBashAlias *bool) {
	// Move alias and profile
	fmt.Println("installBash")

	installBashAliases(overwriteBashAlias)
}

func installBashAliases(overwriteBashAliases *bool) {
	fmt.Println("installBashAliases")

	const BashProfile = ".bash_aliases"

	initConfigFile(BashProfile, BashProfile, overwriteBashAliases)
}

func initConfigFile(source string, dest string, overwrite *bool) {
	source = strings.Join([]string{dotfilePath, source}, "/")

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
	const Vimrc = ".vimrc"

	initConfigFile(Vimrc, Vimrc, ovimrc)
}

func installVimPlug() {
	//Move to config
	const (
		VimPlugUrl      = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
		VimPlugDestFile = "/.vim/autoload/plug.vim"
	)
	var vimAutoloadPath = strings.Join([]string{os.Getenv("HOME"), VimPlugDestFile}, "/")

	if _, err := os.Stat(vimAutoloadPath); err == nil {
		fmt.Printf("vim plug.vim file already exists. bailing\n")
		return
	}

	resp, err := http.Get(VimPlugUrl)
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
