package config

type Database struct {
	Driver string 				`yaml:"driver"`
	Host string 				`yaml:"host"`
	Port int 					`yaml:"port"`
	Database string 			`yaml:"database"`
	UserName string 			`yaml:"username"`
	Password string 			`yaml:"password"`
	Charset string 				`yaml:"charset"`
	MaxIdleConns int 			`yaml:"max_idle_conns"`
	MaxOpenConns int 			`yaml:"max_open_conns"`
	LogMode string 				`yaml:"log_mode"`
	EnableFileLogWriter bool 	`yaml:"enable_file_log_writer"`
	LogFilename string 			`yaml:"log_filename"`
}


