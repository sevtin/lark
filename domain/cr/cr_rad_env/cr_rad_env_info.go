package cr_rad_env

import (
	"lark/domain/cache"
	"lark/domain/pdo"
	"lark/domain/repo"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_red_env"
	"lark/pkg/utils"
)

func GetRedEnvelopeInfo(redEnvCache cache.RedEnvelopeCache, redEnvRepo repo.RedEnvelopeRepository, envId int64) (info *pb_red_env.RedEnvelopeInfo, err error) {
	info, _ = redEnvCache.GetRedEnvelopeInfo(envId)
	if info != nil && info.EnvId > 0 {
		return
	}
	q := entity.NewMysqlQuery()
	q.SetFilter("env_id=?", envId)
	var ifo *pdo.RedEnvelopeInfo
	ifo, err = redEnvRepo.GetRedEnvelopeInfo(q)
	if err != nil {
		return
	}
	if ifo.EnvId == 0 {
		return
	}
	info = &pb_red_env.RedEnvelopeInfo{
		EnvId:          ifo.EnvId,
		EnvType:        pb_enum.RED_ENVELOPE_TYPE(ifo.EnvType),
		ReceiverType:   pb_enum.RECEIVER_TYPE(ifo.ReceiverType),
		TradeNo:        ifo.TradeNo,
		ChatId:         ifo.ChatId,
		SenderUid:      ifo.SenderUid,
		Total:          ifo.Total,
		Quantity:       ifo.Quantity,
		Message:        ifo.Message,
		ReceiverUids:   nil,
		ExpiredTs:      ifo.ExpiredTs,
		SenderPlatform: pb_enum.PLATFORM_TYPE(ifo.SenderPlatform),
	}
	if ifo.Receivers != "" {
		info.ReceiverUids = utils.SplitToInt64List(ifo.Receivers, ",")
	}
	switch pb_enum.RED_ENVELOPE_STATUS(ifo.EnvStatus) {
	case pb_enum.RED_ENVELOPE_STATUS_RECEIVED, pb_enum.RED_ENVELOPE_STATUS_EXPIRED:
		redEnvCache.SetRedEnvelope(envId, info, ifo.EnvStatus)
	default:
		redEnvCache.SetRedEnvelopeInfo(envId, info)
	}
	return
}
