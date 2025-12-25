package config

type Config struct {
	Agent    Agent      `mapstructure:"agent" json:"agent" yaml:"agent"`
	Mysql    MySQL      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	EsSyncMq []EsSyncMQ `mapstructure:"sync-mq" json:"sync-mq" yaml:"sync-mq"`
}

type Agent struct {
	Listen string `mapstructure:"listen" json:"agent" yaml:"agent"`
}

type MySQL struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

type EsSyncMQ struct {
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	WorkDir  string `mapstructure:"workdir" json:"workdir" yaml:"workdir"`
	Binary   string `mapstructure:"binary" json:"binary" yaml:"binary"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
}
