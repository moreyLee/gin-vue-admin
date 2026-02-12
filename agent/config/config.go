package config

type Config struct {
	Agent Agent `mapstructure:"agent" json:"agent" yaml:"agent"`
	Mysql MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
