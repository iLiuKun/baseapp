package controller

import (
	"baseapp/service"
	"baseapp/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.foodCategory)
}
func (fcc *FoodCategoryController) foodCategory(context *gin.Context) {
	//调用service功能获取食品种类信息
	foodCategoryService := &service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(context, "食品种类数据获取失败")
		return
	}
	//转换格式
	//imgUrl: hello.png
	for _, category := range categories {
		if category.ImageUrl != "" { //图片url的拼接
			category.ImageUrl = "/" + category.ImageUrl
		}
	}
	tool.Success(context, categories)
}
