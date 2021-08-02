package types

type RummyConfig struct {
	Plugins []RummyPlugin
	RB      *RepoBase
}

type RepoBase struct {
	Cwd, SourceFilesDir string
}

type RummyPlugin interface {
	Install()
}

type PluginData struct {
	Name      string
	FileNames []string
	Overwrite bool
}

func (rb *RepoBase) NewRummyConfig(pd []RummyPlugin) *RummyConfig {
	return &RummyConfig{Plugins: pd, RB: rb}
}

func NewRepoBase(path string) *RepoBase {
	return &RepoBase{
		Cwd:            path,
		SourceFilesDir: "dot-files",
	}
}
