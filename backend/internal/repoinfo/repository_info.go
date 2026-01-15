package repoinfo

type RepositoryInfo struct {
	Owner        string   `yaml:"owner"`
	Name         string   `yaml:"name"`
	Domain       string   `yaml:"domain"`
	Dependencies []string `yaml:"dependencies"`
}
