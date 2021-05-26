package apiChat

import (
	"Open_IM/src/common/config"
	"Open_IM/src/common/log"
	pbChat "Open_IM/src/proto/chat"
	"Open_IM/src/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/skiffer-git/grpc-etcdv3/getcdv3"
	"net/http"
	"strings"
)

type paramsUserSendMsg struct {
	ReqIdentifier int32  `json:"reqIdentifier" binding:"required"`
	PlatformID    int32  `json:"platformID" binding:"required"`
	SendID        string `json:"sendID" binding:"required"`
	OperationID   string `json:"operationID" binding:"required"`
	MsgIncr       int32  `json:"msgIncr" binding:"required"`
	Data          struct {
		SessionType int32                  `json:"sessionType" binding:"required"`
		MsgFrom     int32                  `json:"msgFrom" binding:"required"`
		ContentType int32                  `json:"contentType" binding:"required"`
		RecvID      string                 `json:"recvID" binding:"required"`
		ForceList   []string               `json:"forceList" binding:"required"`
		Content     string                 `json:"content" binding:"required"`
		Options     map[string]interface{} `json:"options" binding:"required"`
		ClientMsgID string                 `json:"clientMsgID" binding:"required"`
		OffLineInfo map[string]interface{} `json:"offlineInfo" binding:"required"`
		Ex          map[string]interface{} `json:"ext"`
	}
}

func newUserSendMsgReq(token string, params *paramsUserSendMsg) *pbChat.UserSendMsgReq {
	pbData := pbChat.UserSendMsgReq{
		ReqIdentifier: params.ReqIdentifier,
		Token:         token,
		SendID:        params.SendID,
		OperationID:   params.OperationID,
		MsgIncr:       params.MsgIncr,
		PlatformID:    params.PlatformID,
		SessionType:   params.Data.SessionType,
		MsgFrom:       params.Data.MsgFrom,
		ContentType:   params.Data.ContentType,
		RecvID:        params.Data.RecvID,
		ForceList:     params.Data.ForceList,
		Content:       params.Data.Content,
		Options:       utils.MapToJsonString(params.Data.Options),
		ClientMsgID:   params.Data.ClientMsgID,
		OffLineInfo:   utils.MapToJsonString(params.Data.OffLineInfo),
		Ex:            utils.MapToJsonString(params.Data.Ex),
	}
	return &pbData
}

func UserSendMsg(c *gin.Context) {
	params := paramsUserSendMsg{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		log.ErrorByKv("json unmarshal err", "", "err", err.Error(), "data", c.PostForm("data"))
		return
	}

	token := c.Request.Header.Get("token")
	if !utils.VerifyToken(token, params.SendID) {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "token validate err"})
		return
	}

	log.InfoByKv("Ws call success to sendMsgReq", params.OperationID, "Parameters", params)

	pbData := newUserSendMsgReq(token, &params)
	log.Info("", "", "api UserSendMsg call start..., [data: %s]", pbData.String())

	etcdConn := getcdv3.GetConn(config.Config.Etcd.EtcdSchema, strings.Join(config.Config.Etcd.EtcdAddr, ","), config.Config.RpcRegisterName.OpenImOfflineMessageName)
	client := pbChat.NewChatClient(etcdConn)

	log.Info("", "", "api UserSendMsg call, api call rpc...")

	reply, _ := client.UserSendMsg(context.Background(), pbData)
	log.Info("", "", "api UserSendMsg call end..., [data: %s] [reply: %s]", pbData.String(), reply.String())

	c.JSON(http.StatusOK, gin.H{
		"errCode":       0,
		"errMsg":        "",
		"msgIncr":       reply.MsgIncr,
		"reqIdentifier": reply.ReqIdentifier,
		"data": gin.H{
			"clientMsgID": reply.ClientMsgID,
			"serverMsgID": reply.ServerMsgID,
		},
	})

}
