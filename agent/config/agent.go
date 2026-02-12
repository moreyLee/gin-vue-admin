package config

type Agent struct {
	Listen   string   `mapstructure:"listen" json:"listen" yaml:"listen"`
	Token    string   `mapstructure:"token" json:"token" yaml:"token"`
	AllowIPs []string `mapstructure:"allow_ips" json:"allow_ips" yaml:"allow_ips"`
}
