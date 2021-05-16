package common

import (
	"api/config"
	curd "github.com/kainonly/gin-curd"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Dependency struct {
	fx.In

	Config *config.Config
	Db     *gorm.DB
	//Redis  *redis.Model
	Curd *curd.Curd
}

func (c *Dependency) Inject(dependency interface{}) {
	dep := dependency.(Dependency)

	c.Config = dep.Config
	c.Db = dep.Db
	//c.Redis = dep.Redis
	c.Curd = dep.Curd
}

//var (
//	SystemCookie = typ.Cookie{
//		Name:     "system",
//		Secure:   true,
//		HttpOnly: true,
//		SameSite: http.SameSiteStrictMode,
//	}
//)
