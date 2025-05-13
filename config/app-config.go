package config

// JWTConfig defines JWT configuration properties
type JWTConfig struct {
	Secret                   string `mapstructure:"secret"`
	AccessTokenExpiryMinutes int    `mapstructure:"access_token_expiry_minutes"`
	RefreshTokenExpiryDays   int    `mapstructure:"refresh_token_expiry_days"`
}

// AppConfig holds application configuration
type AppConfigs struct {
	Environment string `mapstructure:"environment"`
	Server      struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	JWT JWTConfig `mapstructure:"jwt"`
}

// Default configuration values
var AppConfig = AppConfigs{
	Environment: "development",
	Server: struct {
		Port int `mapstructure:"port"`
	}{
		Port: 8080,
	},
	JWT: JWTConfig{
		Secret:                   "your-secret-key-replace-in-production",
		AccessTokenExpiryMinutes: 15,
		RefreshTokenExpiryDays:   7,
	},
}
