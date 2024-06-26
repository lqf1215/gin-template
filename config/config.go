package config

var Config = struct {
	DB   DB     `yaml:"db"`
	Zap  Zap    `yaml:"zap"`
	Port string `yaml:"port"`
}{}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

//type NATS struct {
//	Servers    []string `yaml:"[nats://localhost:4222]"`
//	User       string   `yaml:"your_username"`
//	Password   string   `yaml:"your_password"`
//	Cluster    ClusterConfig
//	QueueGroup string `yaml:"my-queue-group"`
//}
//
//type ClusterConfig struct {
//	Enabled   bool     `yaml:"false"`
//	Name      string   `yaml:"my-cluster"`
//	Endpoints []string `yaml:"[nats://node1:4222]"`
//}
