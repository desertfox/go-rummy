package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

type RummyBash struct {
	bashAliases string
	overwrite   *bool
}

type RummyVim struct {
	vimrc, vimPlugUrl, vimPlugDestFile string
	overwrite                          *bool
}

type RummyConfig struct {
	cwd, dotfilesName string
	rBash             RummyBash
	rVim              RummyVim
}

var (
	Config    RummyConfig
	waitGroup = sync.WaitGroup{}
)

func main() {
	fmt.Println("Main")

	tasks := map[string]func(){
		"installBash": installBash,
		"installVim":  installVim,
	}

	for name, method := range tasks {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println(name)
			method()
		}()
	}

	waitGroup.Wait()
}

func init() {
	fmt.Println("Init")

	cwd, err := os.Getwd()
	check(err)

	Config = RummyConfig{
		cwd:          cwd,
		dotfilesName: "dot-files",
		rBash: RummyBash{
			bashAliases: ".bash_aliases",
			overwrite:   flag.Bool("oba", false, "Overwrite .bash_aliases file")},
		rVim: RummyVim{
			vimrc:           ".vimrc",
			vimPlugUrl:      "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
			vimPlugDestFile: ".vim/autoload/plug.vim",
			overwrite:       flag.Bool("ovrc", false, "Overwrite .vimrc file")}}

	if _, err := os.Stat(strings.Join([]string{Config.cwd, Config.dotfilesName}, "/")); os.IsNotExist(err) {
		fmt.Println("Unable to find dotfiles")
	}

	flag.Parse()
}

func installVim() {
	installVimrc()
	installVimPlug()
}

func installVimrc() {
	fmt.Println("installVimrc")

	initConfigFile(Config.rVim.vimrc, Config.rVim.vimrc, Config.rVim.overwrite)
}

func installVimPlug() {
	var vimAutoloadPath = strings.Join([]string{os.Getenv("HOME"), Config.rVim.vimPlugDestFile}, "/")

	if _, err := os.Stat(vimAutoloadPath); os.IsNotExist(err) {
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

func installBash() {
	installBashAliases(Config.rBash.overwrite)
}

func installBashAliases(overwriteBashAliases *bool) {
	fmt.Println("installBashAliases")

	initConfigFile(Config.rBash.bashAliases, Config.rBash.bashAliases, Config.rBash.overwrite)
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
