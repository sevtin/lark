package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"lark/apps/apis/upload/internal/dto"
	"lark/pkg/common/xants"
	"lark/pkg/common/xlog"
	"lark/pkg/common/xminio"
	"lark/pkg/common/xresize"
	"lark/pkg/constant"
	"lark/pkg/proto/pb_chat"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
	"lark/pkg/xhttp"
	"mime/multipart"
)

func (s *uploadService) UploadAvatar(ctx *gin.Context, params *dto.UploadAvatarReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		fileHeader *multipart.FileHeader
		file       multipart.File
		err        error
	)
	fileHeader, err = ctx.FormFile(constant.UPLOAD_PART_NAME_PHOTO)
	if err != nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED, err.Error())
		return
	}
	file, err = fileHeader.Open()
	if err != nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_OPEN_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_OPEN_UPLOAD_FILE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_OPEN_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_OPEN_UPLOAD_FILE_FAILED, err.Error())
		return
	}
	defer func() {
		file.Close()
	}()

	var (
		photos     *xresize.Photos
		resultList *xminio.PutResultList
		pr         *xminio.PutResult
		avatarReq  = &dto.UploadAvatar{
			OwnerId:   params.OwnerId,
			OwnerType: pb_enum.AVATAR_OWNER(params.OwnerType),
		}
	)
	photos = xresize.CropAvatar(file, s.cfg.Minio.PhotoDirectory)
	if photos.Error != nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_CROP_PHOTO_FAILED, xhttp.ERROR_HTTP_CROP_PHOTO_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_CROP_PHOTO_FAILED, xhttp.ERROR_HTTP_CROP_PHOTO_FAILED, photos.Error.Error())
		return
	}
	resultList = xminio.FPutPhotoListToMinio(photos)
	if resultList.Err != nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_READ_UPLOAD_FILE_FAILED, xhttp.ERROR_HTTP_READ_UPLOAD_FILE_FAILED, resultList.Err.Error())
		return
	}
	var (
		pi   *xresize.PhotoInfo
		host = "http://" + xminio.GetEndpoint() + "/photos/"
	)
	for _, pr = range resultList.List {
		pi = photos.Maps[pr.Info.Key]
		switch pi.Tag {
		case xresize.PhotoTagSmall:
			avatarReq.AvatarSmall = host + pi.Key
		case xresize.PhotoTagMedium:
			avatarReq.AvatarMedium = host + pi.Key
		case xresize.PhotoTagLarge:
			avatarReq.AvatarLarge = host + pi.Key
		}
		path := photos.Maps[pr.Info.Key].Path
		xants.Submit(func() {
			utils.Remove(path)
		})
	}
	switch avatarReq.OwnerType {
	case pb_enum.AVATAR_OWNER_USER_AVATAR:
		var req = new(pb_user.UploadAvatarReq)
		copier.Copy(req, avatarReq)
		reply := s.userClient.UploadAvatar(req)
		if reply == nil {
			resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
			xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
			return
		}
		if reply.Code > 0 {
			resp.SetResult(reply.Code, reply.Msg)
			xlog.Warn(reply.Code, reply.Msg)
			return
		}
		resp.Data = reply.Avatar
	case pb_enum.AVATAR_OWNER_CHAT_AVATAR:
		var req = new(pb_chat.UploadAvatarReq)
		copier.Copy(req, avatarReq)
		reply := s.chatClient.UploadAvatar(req)
		if reply == nil {
			resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
			xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
			return
		}
		if reply.Code > 0 {
			resp.SetResult(reply.Code, reply.Msg)
			xlog.Warn(reply.Code, reply.Msg)
			return
		}
		resp.Data = reply.Avatar
	}
	return
}
