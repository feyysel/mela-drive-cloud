package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadDB() (DBConfig, error) {
	host, err := get("DB_HOST")
	if err != nil {
		return DBConfig{}, err
	}

	port, err := get("DB_PORT")
	if err != nil {
		return DBConfig{}, err
	}

	user, err := get("DB_USER")
	if err != nil {
		return DBConfig{}, err
	}

	password, err := get("DB_PASSWORD")
	if err != nil {
		return DBConfig{}, err
	}

	name, err := get("DB_NAME")
	if err != nil {
		return DBConfig{}, err
	}

	return DBConfig{Host: host, Port: port, User: user, Password: password, Name: name}, nil
}
