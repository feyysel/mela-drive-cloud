package config

type DBConfig struct {
	URL string
}

func LoadDB() (DBConfig, error) {
	url, err := get("DATABASE_URL")
	if err != nil {
		return DBConfig{}, err
	}

	return DBConfig{URL: url}, nil
}
