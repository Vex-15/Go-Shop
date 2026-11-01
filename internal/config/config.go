package config

import "time"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}

Config.ServerConfig.Port

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret              string
	ExpiresIn           time.Duration
	RefreshTokenExpress time.Duration
}

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	S3Endpoint      string
}

type UploadConfig struct {
	Path        string
	MaxFileSize int64
}

fucn Load()(*Config, error){
	_ = godotenv.Load()

	jwtExpiresIn,_ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))
	refreshTokenExpiresIn,_ := time.ParseDuration(getEnv("JWT_REFRESH_TOKEN_EXPIRES_IN", "720h"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("UPLOAD_MAX_FILE_SIZE", "10485760"), 10, 64)

	return &Config{
		Server: ServerConfig{
			Port: 	getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:		 getEnv("DB_HOST", "localhost"),
			Port:		 getEnv("DB_PORT", "5432"),
			User:		 getEnv("DB_USER", "postgres"),
			Password:	 getEnv("DB_PASSWORD", "password"),
			Name:		 getEnv("DB_NAME", "ecommerce"),
			SSLMode:	 getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:			 getEnv("JWT_SECRET", "mysecret"),
			ExpiresIn:		 jwtExpiresIn,
			RefreshTokenExpress: refreshTokenExpiresIn,
	},
		AWS: AWSConfig{
			Region:			 getEnv("AWS_REGION", "us-east-1"),
			AccessKeyID:	 getEnv("AWS_ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", ""),
			S3Bucket:		 getEnv("AWS_S3_BUCKET", ""),
			S3Endpoint:		 getEnv("AWS_S3_ENDPOINT", ""),
		},
		Upload: UploadConfig{	 
			Path:		 getEnv("UPLOAD_PATH", "./uploads"),
			MaxFileSize: maxUploadSize,
	},
	}, nil
}

func getEnv(key, defaultValue string)string{
	value := os.Getenv(key)
	if value == ""{
		return defaultValue
	}
	return value
}
