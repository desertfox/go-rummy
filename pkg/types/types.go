package types

import "path/filepath"

type PluginData struct {
	Name           string
	Files          []FileToMove
	SourceFilesDir string
	DestFilesDir string
}

type FileToMove struct {
	From string
	To string
	Overwrite      bool
}

type Installer interface {
	Install()
}

func (p *PluginData) BuildSourceWithFile(file string) string {
	return filepath.Join(p.SourceFilesDir, file)
}

func (p *PluginData) BuildDestWithFile(file string) string {
	return filepath.Join(p.DestFilesDir, file)
}

func (p *PluginData) AddFileToMove(from string, to string, overwrite bool) {
	p.Files = append(p.Files, FileToMove{
		From: p.BuildSourceWithFile(from),
		To: p.BuildDestWithFile(to),
		Overwrite: overwrite,
	})
}
