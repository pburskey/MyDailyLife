package utility

type Configuration struct {
	Env   string
	Redis RedisConfiguration
	MySQL map[string]interface{}
}

type RedisConfiguration struct {
	Password string
}

func LoadConfiguration(path string) *Configuration {

	var config *Configuration
	return config

}
