package tool

import (
	"baseapp/model"
	"github.com/go-xorm/xorm"
)

type Orm struct{
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(cfg *Config) (*Orm,error){

	database:= cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(database.ShowSql) 

	
	err = engine.Sync2(
		new(model.SmsCode),
		new(model.Member),
		new(model.FoodCategory),
		new(model.Shop),
		new(model.Service),
		new(model.ShopService),
		new(model.Goods),
	)
	if err != nil {
		return nil, err
	}

	orm := Orm{ engine}
	// new(Orm)
	// orm.Engine = engine
	// // 全局
	DbEngine = &orm

	return &orm, nil
}
