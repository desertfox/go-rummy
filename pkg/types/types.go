package types

type Config struct {
	Cwd            string
	SourceFilesDir string
}

type PluginData struct {
	Name      string
	FileNames []string
	Overwrite bool
	Config    Config
}

type Installer interface {
	Install()
}

func NewConfig(path string, dotFiles string) *Config {
	return &Config{
		Cwd:            path,
		SourceFilesDir: dotFiles,
	}
}
