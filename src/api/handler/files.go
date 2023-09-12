package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/dto"
	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/Arshia-Izadyar/Car-sale-api/src/common"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/services"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{
		service: services.NewFileService(cfg),
	}
}

func (fh *FileHandler) CreateFile(ctx *gin.Context) {
	upload := dto.UploadFileRequest{}
	err := ctx.ShouldBind(&upload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.ValidationError), err.Error()))
		return
	}
	req := &dto.CreateFileRequest{}
	req.Description = upload.Description
	req.MineType = upload.File.Header.Get("Content-Type")
	req.Directory = "../../uploads"
	req.Name, err = common.SaveFile(upload.File, req.Directory)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	res, err := fh.service.Create(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, int(helper.Success), true))
}

func (fh *FileHandler) UpdateFile(ctx *gin.Context) {
	Update[dto.UpdateFileRequest, dto.FileResponse](ctx, fh.service.Update)
}

func (fh *FileHandler) GetFileById(ctx *gin.Context) {
	Get[dto.FileResponse](ctx, fh.service.GetById)
}

func (fh *FileHandler) DeleteFile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.ValidationError), "id  0 is invalid"))
		return
	}
	file, err := fh.service.GetById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.NotFoundError), err.Error()))
		return
	}
	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.ValidationError), err.Error()))
		return
	}
	err = fh.service.Delete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, int(helper.InternalError), err.Error()))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(nil, int(helper.Success), true))

}
