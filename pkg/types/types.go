package types

type Config struct {
	Cwd            string
	SourceFilesDir string
}

type Installer interface {
	Install()
}

type PluginData struct {
	Name      string
	FileNames []string
	Overwrite bool
	Config    Config
}

func NewConfig(path string) *Config {
	return &Config{
		Cwd:            path,
		SourceFilesDir: "dot-files",
	}
}
