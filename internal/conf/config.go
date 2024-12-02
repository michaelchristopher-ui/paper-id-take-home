package conf

type Config struct {
	SQLite SQLite `yaml:"sqlite"`
	Server Server `yaml:"server"`
}

type SQLite struct {
	DBPath string `yaml:"db_path"`
}

type Server struct {
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
	Port         string `yaml:"port"`
}
