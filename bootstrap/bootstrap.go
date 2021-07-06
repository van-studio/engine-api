package bootstrap

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kainonly/gin-helper/authx"
	"github.com/weplanx/api/config"
	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// LoadConfiguration application configuration
// reference config.example.yml
func LoadConfiguration() (cfg *config.Config) {
	if _, err := os.Stat("./config.yml"); os.IsNotExist(err) {
		return
	}
	buf, _ := ioutil.ReadFile("./config.yml")
	yaml.Unmarshal(buf, &cfg)
	return
}

// InitializeDatabase database configuration
// If it is another database, replace the driver here
// gorm.Open(mysql.Open(option.Dsn),...)
// reference https://gorm.io/docs/connecting_to_the_database.html
func InitializeDatabase(cfg *config.Config) (db *gorm.DB) {
	option := cfg.Database
	db, _ = gorm.Open(postgres.Open(option.Dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: option.TablePrefix,
		},
	})
	sqlDB, _ := db.DB()
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

func InitializeAuth(cfg *config.Config) *authx.Auth {
	option := cfg.Auth
	log.Println(option)
	return authx.Make(cfg.Auth, authx.Args{
		Method:    jwt.SigningMethodHS256,
		UseCookie: nil,
		RefreshFn: nil,
	})
}

// HttpServer Start http service
// https://gin-gonic.com/docs/examples/custom-http-config
func HttpServer(lc fx.Lifecycle, cfg *config.Config) (serve *gin.Engine) {
	serve = gin.New()
	serve.Use(gin.Logger())
	serve.Use(gin.Recovery())
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go serve.Run(cfg.Listen)
			return nil
		},
	})
	return
}
