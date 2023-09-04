package domain

type Config struct {
	WebServer  ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	Encryption EncryptionConfig `yaml:"encryption"`
	Redis      RedisConfig      `yaml:"redis"`
	Google     GoogleConfig     `yaml:"google"`
	Mailgun    MailgunConfig    `yaml:"mailgun"`
}

type ServerConfig struct {
	Bind   string `yaml:"bind"`
	Port   int    `yaml:"port"`
	Domain string `yaml:"domain"`
	Debug  bool   `yaml:"debug"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	RootPassword string `yaml:"root_password"`
	Name         string `yaml:"name"`
	Debug        bool   `yaml:"debug"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

type EncryptionConfig struct {
	JwtSecret          string `yaml:"jwt_secret"`
	RefreshTokenSecret string `yaml:"refresh_token_secret"`
}

type MailgunConfig struct {
	ApiKey  string `yaml:"api_key"`
	BaseUrl string `yaml:"base_url"`
}

type GoogleConfig struct {
	AuthKey  string `yaml:"auth_key"`
	Audience string `yaml:"audience"`
}
