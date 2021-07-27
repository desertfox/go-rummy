package rummy 

import "os"

type RummyBash struct {
	BashAliases string
	Boverwrite  bool
}

type RummyVim struct {
	Vimrc, VimPlugUrl, VimPlugDestFile string
	Voverwrite                         bool
}

type RummyConfig struct {
	Cwd, DotfilesName string
	RummyBash
	RummyVim
}

func NewConfig() RummyConfig {

	return RummyConfig{
		Cwd:          getWD(),
		DotfilesName: "dot-files",
		RummyBash: RummyBash{
			Boverwrite:  false,
			BashAliases: ".bash_aliases"},
		RummyVim: RummyVim{
			Voverwrite:      false,
			Vimrc:           ".vimrc",
			VimPlugUrl:      "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
			VimPlugDestFile: ".vim/autoload/plug.vim"}}
}

func getWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

/*
	config := &RummyConfig{
		Cwd:          cwd,
		DotfilesName: "dot-files",
		RummyBash: RummyBash{
			*Boverwrite:  false,
			BashAliases: ".bash_aliases"},
		RummyVim: RummyVim{
			*Voverwrite:      false,
			Vimrc:           ".vimrc",
			VimPlugUrl:      "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim",
			VimPlugDestFile: ".vim/autoload/plug.vim"}}
*/
