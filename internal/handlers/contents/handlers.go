package contents

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hun9k/gapi/biz"
	"github.com/hun9k/gapiDemo/internal/models"
)

// 创建一个资源
func Post(ctx *gin.Context) {
	// I. receive request
	pb := &postBody{}
	if err := ctx.ShouldBind(&pb); err != nil {
		ctx.JSON(http.StatusBadRequest, biz.Resp{
			Code:    1,
			Message: err.Error(),
			// Messages: map[string]string{},
		})
		return
	}

	// II. do
	model := postBodyToModel(pb)
	if err := models.ContentsInsert(&model); err != nil {
		ctx.JSON(http.StatusInternalServerError, biz.Resp{
			Code:    1,
			Message: "服务器内部错误，请稍后重试",
		})
		return
	}

	// III. send response
	rowModel, _ := models.ContentsRow(model.ID)
	ctx.JSON(http.StatusOK, biz.Resp{
		Code:    0,
		Message: "创建成功",
		Data:    modelToPostResp(rowModel),
	})
}

// 删除一个资源
func DeleteId(ctx *gin.Context) {

}

// 删除多个资源
func Delete(ctx *gin.Context) {

}

// 恢复一个资源
func RestoreID(ctx *gin.Context) {

}

// 恢复多个资源
func Restore(ctx *gin.Context) {

}

// 更新一个资源
func PutId(ctx *gin.Context) {

}

// 更新多个资源
func Put(ctx *gin.Context) {

}

// 更新一个资源的部分字段
func PatchId(ctx *gin.Context) {

}

// 更新多个资源的部分字段
func Patch(ctx *gin.Context) {

}

// 获取一个资源
func GetId(ctx *gin.Context) {

}

// 获取多个资源
func Get(ctx *gin.Context) {

}
