package service

import (
	"baseapp/dao"
	"baseapp/model"
)

type GoodService struct {

}

func NewGoodService()*GoodService{
	return &GoodService{}
}

/**
 * 获取商家的食品列表
 */
func (gs *GoodService)GetFoods(shop_id int64)[]model.Goods{
	goodDao:=dao.NewGoodDao()
	return goodDao.QueryFoods(shop_id)
}
