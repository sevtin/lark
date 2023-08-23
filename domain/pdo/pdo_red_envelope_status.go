package pdo

import "lark/pkg/utils"

var (
	field_tag_red_envelope_status string
)

type RedEnvelopeStatus struct {
	EnvId     int64 `json:"env_id" field:"env_id"`
	EnvStatus int32 `json:"env_status" field:"env_status"`
}

func (p *RedEnvelopeStatus) GetFields() string {
	if field_tag_red_envelope_status == "" {
		field_tag_red_envelope_status = utils.GetFields(*p)
	}
	return field_tag_red_envelope_status
}
