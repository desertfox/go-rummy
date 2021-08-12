package types

import "path/filepath"

type PluginData struct {
	Name           string
	FileNames      []string
	Overwrite      bool
	Dest           string
	SourceFilesDir string
}

type Installer interface {
	Install()
}

func NewPlugin(name string, fileNames []string, overwrite bool, dest string, sourceFilesDir string) *PluginData {
	return &PluginData{
		Name:           name,
		FileNames:      fileNames,
		Overwrite:      overwrite,
		Dest:           dest,
		SourceFilesDir: sourceFilesDir,
	}
}

func (p PluginData) BuildSourceWithFile(file string) string {
	return filepath.Join(p.SourceFilesDir, file)
}

func (p PluginData) BuildDestWithFile(file string) string {
	return filepath.Join(p.Dest, file)
}
