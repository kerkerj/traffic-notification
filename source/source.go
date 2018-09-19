package source

import (
	"encoding/xml"
	"fmt"
	"github.com/kerkerj/traffic-notification/utils"
	"go.uber.org/zap"
	
	"github.com/parnurzeal/gorequest"
)

const (
	IncidentDataURL = "http://tisvcloud.freeway.gov.tw/xml/1min_incident_data_1968.xml"
)

type Source struct {
	httpClient *gorequest.SuperAgent
	body string
	
	Result *XMLResult
	Errs []error
}

func New() *Source {
	return &Source{
		httpClient:gorequest.New(),
	}
}

func (s *Source) FetchXML(url string) *Source {
	utils.Logger.Info("start fetching")
	
	_, body, errs := s.httpClient.Get(url).End()
	if errs != nil {
		utils.Logger.Info("errors", zap.Errors("errs", errs))
		fmt.Printf("error: %v", errs)
		s.Errs = errs
	} else {
		s.body = body
	}
	return s
}

func (s *Source) ParseXML() *Source {
	utils.Logger.Info("start parsing")
	
	if s.body == "" {
		return s
	}
	
	x := &XMLResult{}
	err := xml.Unmarshal([]byte(s.body), x)
	if err != nil {
		fmt.Printf("error: %v", err)
	} else {
		s.Result = x
	}
	return s
}
