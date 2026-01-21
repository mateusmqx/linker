package main

type Service struct {
	Name       string
	Team       string
	Repository string
	Info       ServiceInfo
}

type ServiceInfo struct {
	Dependencies  []string `yaml:"dependencies"`
	Documentation string   `yaml:"documentation"`
}

func main() {

}
