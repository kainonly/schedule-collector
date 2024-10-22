package common

type Values struct {
	Mode    string  `yaml:"mode"`
	Elastic Elastic `yaml:"elastic"`
	Nats    Nats    `yaml:"nats"`
}

type Elastic struct {
	Hosts    []string `yaml:"hosts"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

type Nats struct {
	Hosts []string `yaml:"hosts"`
	Token string   `yaml:"token"`
}