package yamlConfig

const (
	YAMLDEFAULTPATH = "./config.yaml"
)

type YamlConfigure struct {
	Version string `yaml:"version"`
	Xlsx XlsxConfig `yaml:"xlsx"`
}

type XlsxConfig struct {
	AbsolutePath string `yaml:"absolutepath"`
	MatchFilePath string `yaml:"absoluteMatchPath"`
	SavedDirctory string `yaml:"saveddirctory"`
}