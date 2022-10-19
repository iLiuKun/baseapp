package service

import (
	"baseapp/dao"
	"baseapp/model"
	"baseapp/tool"
)

type FoodCategoryService struct {

}

/**
 * 获取美食类别
 */
func (fcs *FoodCategoryService)Categories()([]model.FoodCategory,error){
	//数据库操作层
	md := dao.FoodCategoryDao{tool.DbEngine}
	return md.QueryCategories()

}
