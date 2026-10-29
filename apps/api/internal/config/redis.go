package config

type RedisConfig struct {
	URL string
}

func LoadRedis() (RedisConfig, error) {
	url, err := get("REDIS_URL")
	if err != nil {
		return RedisConfig{}, err
	}

	return RedisConfig{URL: url}, nil
}
