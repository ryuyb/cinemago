package config

// Config struct to store configuration information
type Config struct {
	Env      string   `mapstructure:"env" desc:"Environment (dev/prod)"`
	Server   Server   `mapstructure:"server" desc:"Server configuration"`
	Log      Log      `mapstructure:"log" desc:"Log configuration"`
	Database Database `mapstructure:"database" desc:"Database configuration"`
	Jwt      Jwt      `mapstructure:"jwt" desc:"JWT configuration"`
}

// Server struct to store server configuration
type Server struct {
	Addr string `mapstructure:"addr" desc:"Server address with port"`
}

type Database struct {
	Driver  string `mapstructure:"driver" desc:"Database driver (sqlite3/postgres)"`
	Source  string `mapstructure:"source" desc:"Database source url"`
	Migrate bool   `mapstructure:"migrate" desc:"Enable migrations"`
}

type Log struct {
	Level         string  `mapstructure:"level" desc:"Log level"`
	CallerSkip    int     `mapstructure:"callerSkip" desc:"Log caller skip"`
	EnableConsole bool    `mapstructure:"enableConsole" desc:"Whether to output to console"`
	File          LogFile `mapstructure:"file" desc:"Log file configuration"`
}

type LogFile struct {
	Enable     bool   `mapstructure:"enable" desc:"Whether to output to file"`
	JsonFormat bool   `mapstructure:"jsonFormat" desc:"Whether to output to file in JSON format"`
	FilePath   string `mapstructure:"filePath" desc:"Log file path"`
	MaxSize    int    `mapstructure:"maxSize" desc:"Maximum size of a single log file (MB)"`
	MaxBackups int    `mapstructure:"maxBackups" desc:"Maximum number of old log files to retain"`
	MaxAge     int    `mapstructure:"maxAge" desc:"Maximum number of days to retain old log files"`
	Compress   bool   `mapstructure:"compress" desc:"Whether to compress old log files"`
}

type Jwt struct {
	SigningKey         string `mapstructure:"signingKey" desc:"JWT signing key"`
	ValidWithinMinutes int    `mapstructure:"validWithinMinutes" desc:"Valid time within minutes"`
}
