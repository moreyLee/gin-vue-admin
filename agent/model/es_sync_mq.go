package model

type EsSyncMQ struct {
	Name     string `mapstructure:"name" json:"name" yaml:"name"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	WorkDir  string `mapstructure:"workdir" json:"workdir" yaml:"workdir"`
	Binary   string `mapstructure:"binary" json:"binary" yaml:"binary"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
}
