package bootstrap

import (
	"api/config"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kainonly/gin-extra/cors"
	"github.com/kainonly/gin-extra/validate"
	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
	"time"
)

var (
	LoadConfigurationNotExists = errors.New("the configuration file does not exist")
)

// LoadConfiguration application configuration
// reference config.example.yml
func LoadConfiguration() (cfg *config.Config, err error) {
	if _, err = os.Stat("./config.yml"); os.IsNotExist(err) {
		err = LoadConfigurationNotExists
		return
	}
	var buf []byte
	buf, err = ioutil.ReadFile("./config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return
	}
	return
}

// InitializeDatabase database configuration
// If it is another database, replace the driver here
// gorm.Open(mysql.Open(option.Dsn),...)
// reference https://gorm.io/docs/connecting_to_the_database.html
func InitializeDatabase(cfg *config.Config) (db *gorm.DB, err error) {
	option := cfg.Database
	db, err = gorm.Open(mysql.Open(option.Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   option.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	if option.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(option.MaxIdleConns)
	}
	if option.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(option.MaxOpenConns)
	}
	if option.ConnMaxLifetime != 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(option.ConnMaxLifetime))
	}
	return
}

// InitializeRedis the redis library configuration
// reference https://github.com/go-redis/redis
//func InitializeRedis(cfg *config.Config) *redis.Client {
//	option := cfg.Redis
//	return redis.NewClient(&redis.Options{
//		Addr:     option.Address,
//		Password: option.Password,
//		DB:       option.DB,
//	})
//}

// HttpServer Start http service
// https://gin-gonic.com/docs/examples/custom-http-config
func HttpServer(lc fx.Lifecycle, cfg *config.Config) (serve *gin.Engine) {
	serve = gin.New()
	serve.Use(gin.Logger())
	serve.Use(gin.Recovery())
	serve.Use(cors.Cors(cfg.Cors))
	validate.RegisterCustomValidate()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go serve.Run(cfg.Listen)
			return nil
		},
	})
	return
}
