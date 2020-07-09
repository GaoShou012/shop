package main

import(
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/models"

	"fmt"
	"github.com/app"
	//"github.com/jinzhu/gorm"
	"github.com/models"
)

var context *gin.Context = nil

func main() {

	// 搜搜typeId是否存在
	goodsType := models.GoodsType{}

	res := models.DB.Table(models.GoodsTypeTableName).Where(&models.GoodsType{Model:gorm.Model{ID:13}}).Find(&goodsType)
	if res.Error != nil {
		if res.RecordNotFound() {
			app.ResponseError(context,fmt.Errorf("商品类型ID不存在"))
		}else{
			app.ResponseError(context,res.Error)
		}
		return
	}

	fmt.Printf("%v\n",goodsType)
	fmt.Printf("%d\n",*goodsType.Sort)
	fmt.Printf("%s\n",*goodsType.Desc)
	fmt.Printf("%s\n",*goodsType.Name)
}