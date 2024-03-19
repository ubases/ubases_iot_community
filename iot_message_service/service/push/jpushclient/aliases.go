package jpushclient

import (
	"encoding/json"
)

type Aliases struct {
	//Platform        interface{}   `json:"platform"`
	RegistrationIds AliasesRemove `json:"registration_ids"`
}

type AliasesRemove struct {
	Remove []interface{} `json:"remove"`
}

func NewAliases() *Aliases {
	pl := &Aliases{}
	pl.RegistrationIds = AliasesRemove{
		Remove: make([]interface{}, 0),
	}
	return pl
}

func (s *Aliases) SetRegIds(regId interface{}) {
	s.RegistrationIds.Remove = append(s.RegistrationIds.Remove, regId)
}

func (this *Aliases) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}

///{"registration_ids":["101d855909a5d8b5b8d","121c83f760c7c8ec43d","13165ffa4ef545333a9","13165ffa4ef5452e703","18171adc03f0d1346cc","1517bfd3f72455443d9","121c83f760c7c8e6f4d","141fe1da9e4976d14db","161a3797c8e6a87b3c6","161a3797c8e7b9cfb56"]}
type AliasesResponse struct {
	Data []AliasesQuery `json:"data"`
}

type AliasesQuery struct {
	Platform       interface{} `json:"platform"`
	LastOnlineDate interface{} `json:"last_online_date"`
	RegistrationId interface{} `json:"registration_id"`
}

type AliasesOldResponse struct {
	Data []string `json:"registration_ids"`
}
