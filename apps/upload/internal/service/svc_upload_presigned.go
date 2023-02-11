package service

import (
	"github.com/gin-gonic/gin"
	"lark/apps/upload/internal/dto"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xminio"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
	"net/url"
)

func (s *uploadService) Presigned(ctx *gin.Context, req *dto.PresignedReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		url *url.URL
		err error
	)
	url, err = xminio.Presigned(req.FileType, utils.NewUUID())
	if err != nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_PRESIGNED_FAILED, xhttp.ERROR_HTTP_PRESIGNED_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_PRESIGNED_FAILED, xhttp.ERROR_HTTP_PRESIGNED_FAILED, err.Error())
		return
	}
	resp.Data = &dto.PresignedResp{Url: url.String()}
	return
}

/*
The key in this case is how the file is uploaded from the postman. While uploading the file,
you need to use Body > Binary > Select File, rather than using the Body > Form-Data.
*/
