package config

import "github.com/joho/godotenv"

type Config struct {
	App     AppConfig
	DB      DBConfig
	Redis   RedisConfig
	Storage StorageConfig
	Session SessionConfig
}

func Load() (Config, error) {
	_ = godotenv.Load()

	app, err := LoadApp()
	if err != nil {
		return Config{}, err
	}

	db, err := LoadDB()
	if err != nil {
		return Config{}, err
	}

	redis, err := LoadRedis()
	if err != nil {
		return Config{}, err
	}

	storage, err := LoadStorage()
	if err != nil {
		return Config{}, err
	}

	return Config{
		App:     app,
		DB:      db,
		Redis:   redis,
		Storage: storage,
		Session: NewSessionConfig(app),
	}, nil
}
