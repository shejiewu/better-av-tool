package scraper

import (
	myclient "better-av-tool/pkg/client"
	"better-av-tool/pkg/config"
	"better-av-tool/pkg/logger"
	"better-av-tool/third_party/dmm-go-sdk/api"
)

// Scraper is interface
type Scraper interface {
	FetchDoc(query string) (err error)
	GetPlot() string
	GetTitle() string
	GetDirector() string
	GetRuntime() string
	GetTags() []string
	GetMaker() string
	GetActors() []string
	GetLabel() string
	GetNumber() string
	GetFormatNumber() string
	GetCover() string
	GetWebsite() string
	GetPremiered() string
	GetYear() string
	GetSeries() string
	GetType() string
	NeedCut() bool
}

var (
	client            myclient.Client
	log               logger.Logger
	dmmProductService *api.ProductService
	needCut           bool
)

// Setup ...
func Setup(conf *config.Configs) {
	log = logger.New()
	client = myclient.New()
	if conf.Proxy.Enable {
		err := client.SetProxyUrl(conf.Proxy.Socket)
		if err != nil {
			log.Errorf("Error parse proxy url, %s, proxy disabled", err)
		}
	}
	if conf.DMMApi.ApiId != "" && conf.DMMApi.AffiliateId != "" {
		dmmProductService = api.NewProductService(conf.DMMApi.AffiliateId, conf.DMMApi.ApiId)
	}
	needCut = conf.Output.NeedCut
}
