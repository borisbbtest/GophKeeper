package models

import (
	"encoding/json"

	integrationcmdb "github.com/borisbbtest/GoMon/internal/fanout/clients/grpc/cmdb"
	"github.com/borisbbtest/GoMon/internal/fanout/utils"
)

type RequestGetCis struct {
	Item *[]string
}

type ResponseGetCis struct {
	Root *[]*integrationcmdb.Ci
}

type ResponseGetCi struct {
	Root *integrationcmdb.Ci
}

// UnmarshalJSON - функция переопределяющия правила анмаршалера для timestamp в Metric
func (hook *RequestGetCis) ParseRequest(data []byte) error {
	Req := &struct {
		ListCisName []string `json:"ListCisName"`
	}{}
	if err := json.Unmarshal(data, Req); err != nil {
		utils.Log.Error().Err(err).Msg("failed unmarshall json")
		return err
	}
	hook.Item = &Req.ListCisName
	return nil
}
