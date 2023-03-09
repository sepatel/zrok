package metrics

import (
	"fmt"
	"github.com/openziti/zrok/util"
	"time"
)

type Usage struct {
	ProcessedStamp time.Time
	IntervalStart  time.Time
	ZitiServiceId  string
	ZitiCircuitId  string
	ShareToken     string
	FrontendTx     int64
	FrontendRx     int64
	BackendTx      int64
	BackendRx      int64
}

func (u Usage) String() string {
	out := "Usage {"
	out += fmt.Sprintf("processed '%v'", u.ProcessedStamp)
	out += ", " + fmt.Sprintf("interval '%v'", u.IntervalStart)
	out += ", " + fmt.Sprintf("service '%v'", u.ZitiServiceId)
	out += ", " + fmt.Sprintf("circuit '%v'", u.ZitiCircuitId)
	out += ", " + fmt.Sprintf("share '%v'", u.ShareToken)
	out += ", " + fmt.Sprintf("fe {rx %v, tx %v}", util.BytesToSize(u.FrontendRx), util.BytesToSize(u.FrontendTx))
	out += ", " + fmt.Sprintf("be {rx %v, tx %v}", util.BytesToSize(u.BackendRx), util.BytesToSize(u.BackendTx))
	out += "}"
	return out
}

type Source interface {
	Start(chan map[string]interface{}) (chan struct{}, error)
	Stop()
}

type Ingester interface {
	Ingest(msg map[string]interface{}) error
}
