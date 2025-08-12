package database

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSL      bool
}

type DbConfigBuilder interface {
	SetHost(host string) DbConfigBuilder
	SetPort(port int) DbConfigBuilder
	SetCredentials(user, password string) DbConfigBuilder
	SetDatabase(name string) DbConfigBuilder
	SetSSL(enable bool) DbConfigBuilder
	Build() (*DBConfig, error)
}

type ConfigBuilder struct {
	config *DBConfig
}

func NewDbConfigBuilder() DbConfigBuilder {
	return &ConfigBuilder{
		config: &DBConfig{
			Port: 5432,  // Default port
			SSL:  false, // Default SSL setting
		},
	}
}

func (cb *ConfigBuilder) SetHost(host string) DbConfigBuilder {
	cb.config.Host = host
	return cb
}

func (cb *ConfigBuilder) SetPort(port int) DbConfigBuilder {
	cb.config.Port = port
	return cb
}

func (cb *ConfigBuilder) SetCredentials(user, password string) DbConfigBuilder {
	cb.config.User = user
	cb.config.Password = password
	return cb
}

func (cb *ConfigBuilder) SetDatabase(name string) DbConfigBuilder {
	cb.config.Database = name
	return cb
}

func (cb *ConfigBuilder) SetSSL(enable bool) DbConfigBuilder {
	cb.config.SSL = enable
	return cb
}

func (cb *ConfigBuilder) Build() (*DBConfig, error) {
	return cb.config, nil
}
