package config

type Config struct {
	loggerConfig *LoggerConfig
}

type LoggerConfig struct {
	Log_FILE_PATH string
	LOG_FILE_NAME string
}
