package config

const configFile = "config.yaml"

type Config struct {
	Token string `yanl:"token"`
}