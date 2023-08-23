package cr_rad_env

import (
	"lark/domain/cache"
	"lark/domain/pdo"
	"lark/domain/repo"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
)

func GetRedEnvelopeStatus(redEnvCache cache.RedEnvelopeCache, redEnvRepo repo.RedEnvelopeRepository, envId int64) (status int, err error) {
	status, _ = redEnvCache.GetRedEnvelopeStatus(envId)
	if status > 0 {
		return
	}
	var (
		q = entity.NewMysqlQuery()
		s *pdo.RedEnvelopeStatus
	)
	q.SetFilter("env_id=?", envId)
	s, err = redEnvRepo.GetRedEnvelopeStatus(q)
	if err != nil {
		return
	}
	status = int(s.EnvStatus)
	if status == 0 {
		return
	}
	switch pb_enum.RED_ENVELOPE_STATUS(status) {
	case pb_enum.RED_ENVELOPE_STATUS_RECEIVED, pb_enum.RED_ENVELOPE_STATUS_EXPIRED:
		go redEnvCache.SetRedEnvelopeStatus(envId, int32(status))
	}
	return
}
