package controller

import (
	"time"

	"github.com/DemoHn/Catilina/app/service"
	"github.com/gin-gonic/gin"
)

type addRecordReq struct {
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	Debt      bool   `json:"debt"`
	IssueTime int    `json:"issueTime"`
}

func listRecordsHandler(c *gin.Context) {

}

func addRecordsHandler(c *gin.Context) {
	var req addRecordReq

	if err := c.ShouldBindJSON(&req); err != nil {
		handleRsp(c, nil, err)
		return
	}

	ts := time.Unix(int64(req.IssueTime), 0)
	record, err := service.AddRecord(req.Name, req.Amount, req.Debt, ts)

	recordData := map[string]interface{}{
		"id":     record.ID,
		"uid":    record.UID,
		"amount": record.Amount,
		"debt":   record.IsDebt,
	}
	handleRsp(c, recordData, err)

}

var recordRoutes = map[string]gin.HandlerFunc{
	"GET:/api/records/list": listRecordsHandler,
	"POST:/api/records/add": addRecordsHandler,
}
