package config

type Config struct {
	Commands []Command `yaml:"commands"`
}

type Command struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Args        []string `yaml:"args"`
	Steps       []Step   `yaml:"steps"`
}

type Step struct {
	Script string `yaml:"script"`
}
