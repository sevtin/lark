package xes

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
	"lark/pkg/common/xlog"
	"lark/pkg/conf"
)

var (
	cli *EsClient
)

type EsClient struct {
	cfg      *conf.Elasticsearch
	esClient *elasticsearch.Client
}

func NewElasticsearchClient(cfg *conf.Elasticsearch) {
	cli = &EsClient{cfg: cfg}
	cli.connectEs()
}

func (c *EsClient) connectEs() (esClient *elasticsearch.Client) {
	var (
		esCfg   elasticsearch.Config
		resp    *esapi.Response
		certBuf []byte
		respBuf []byte
		err     error
	)

	esCfg = elasticsearch.Config{
		Addresses: c.cfg.Addresses,
		Username:  c.cfg.Username,
		Password:  c.cfg.Password,
	}
	if c.cfg.TlsEnabled == true {
		certBuf, err = ioutil.ReadFile(c.cfg.CACert)
		if err != nil {
			xlog.Error(err.Error())
			return
		}
		esCfg.CACert = certBuf
	}
	esClient, err = elasticsearch.NewClient(esCfg)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	cli.esClient = esClient
	resp, err = esClient.Info()
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	if resp.StatusCode != 200 {
		respBuf, err = ioutil.ReadAll(resp.Body)
		xlog.Error(string(respBuf))
	}
	defer resp.Body.Close()
	return
}

func GetClient() *elasticsearch.Client {
	if cli.esClient == nil {
		cli.esClient = cli.connectEs()
	}
	return cli.esClient
}
