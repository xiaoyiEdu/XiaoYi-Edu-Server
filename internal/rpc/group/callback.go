package group

import (
	"OpenIM/pkg/apistruct"
	"OpenIM/pkg/callbackstruct"
	"OpenIM/pkg/common/config"
	"OpenIM/pkg/common/constant"
	"OpenIM/pkg/common/db/table/relation"
	"OpenIM/pkg/common/http"
	"OpenIM/pkg/common/tracelog"
	"OpenIM/pkg/proto/group"
	"OpenIM/pkg/proto/wrapperspb"
	"OpenIM/pkg/utils"
	"context"
	"time"
)

func CallbackBeforeCreateGroup(ctx context.Context, req *group.CreateGroupReq) (err error) {
	if !config.Config.Callback.CallbackBeforeCreateGroup.Enable {
		return nil
	}
	defer func() {
		tracelog.SetCtxInfo(ctx, utils.GetFuncName(1), err, "req", req)
	}()
	cbReq := &callbackstruct.CallbackBeforeCreateGroupReq{
		CallbackCommand: constant.CallbackBeforeCreateGroupCommand,
		OperationID:     tracelog.GetOperationID(ctx),
		GroupInfo:       *req.GroupInfo,
	}
	cbReq.InitMemberList = append(cbReq.InitMemberList, &apistruct.GroupAddMemberInfo{
		UserID:    req.OwnerUserID,
		RoleLevel: constant.GroupOwner,
	})
	for _, userID := range req.AdminUserIDs {
		cbReq.InitMemberList = append(cbReq.InitMemberList, &apistruct.GroupAddMemberInfo{
			UserID:    userID,
			RoleLevel: constant.GroupAdmin,
		})
	}
	for _, userID := range req.AdminUserIDs {
		cbReq.InitMemberList = append(cbReq.InitMemberList, &apistruct.GroupAddMemberInfo{
			UserID:    userID,
			RoleLevel: constant.GroupOrdinaryUsers,
		})
	}
	resp := &callbackstruct.CallbackBeforeCreateGroupResp{}
	err = http.CallBackPostReturn(config.Config.Callback.CallbackUrl, cbReq, resp, config.Config.Callback.CallbackBeforeCreateGroup)
	if err != nil {
		return err
	}
	utils.NotNilReplace(&req.GroupInfo.GroupID, resp.GroupID)
	utils.NotNilReplace(&req.GroupInfo.GroupName, resp.GroupName)
	utils.NotNilReplace(&req.GroupInfo.Notification, resp.Notification)
	utils.NotNilReplace(&req.GroupInfo.Introduction, resp.Introduction)
	utils.NotNilReplace(&req.GroupInfo.FaceURL, resp.FaceURL)
	utils.NotNilReplace(&req.GroupInfo.OwnerUserID, resp.OwnerUserID)
	utils.NotNilReplace(&req.GroupInfo.Ex, resp.Ex)
	utils.NotNilReplace(&req.GroupInfo.Status, resp.Status)
	utils.NotNilReplace(&req.GroupInfo.CreatorUserID, resp.CreatorUserID)
	utils.NotNilReplace(&req.GroupInfo.GroupType, resp.GroupType)
	utils.NotNilReplace(&req.GroupInfo.NeedVerification, resp.NeedVerification)
	utils.NotNilReplace(&req.GroupInfo.LookMemberInfo, resp.LookMemberInfo)
	return nil
}

func CallbackBeforeMemberJoinGroup(ctx context.Context, groupMember *relation.GroupMemberModel, groupEx string) (err error) {
	if !config.Config.Callback.CallbackBeforeMemberJoinGroup.Enable {
		return nil
	}
	defer func() {
		tracelog.SetCtxInfo(ctx, utils.GetFuncName(1), err, "groupMember", *groupMember, "groupEx", groupEx)
	}()
	callbackReq := &callbackstruct.CallbackBeforeMemberJoinGroupReq{
		CallbackCommand: constant.CallbackBeforeMemberJoinGroupCommand,
		OperationID:     tracelog.GetOperationID(ctx),
		GroupID:         groupMember.GroupID,
		UserID:          groupMember.UserID,
		Ex:              groupMember.Ex,
		GroupEx:         groupEx,
	}
	resp := &callbackstruct.CallbackBeforeMemberJoinGroupResp{}
	err = http.CallBackPostReturn(config.Config.Callback.CallbackUrl, callbackReq, resp, config.Config.Callback.CallbackBeforeMemberJoinGroup)
	if err != nil {
		return err
	}
	if resp.MuteEndTime != nil {
		groupMember.MuteEndTime = time.UnixMilli(*resp.MuteEndTime)
	}
	utils.NotNilReplace(&groupMember.FaceURL, resp.FaceURL)
	utils.NotNilReplace(&groupMember.Ex, resp.Ex)
	utils.NotNilReplace(&groupMember.Nickname, resp.Nickname)
	utils.NotNilReplace(&groupMember.RoleLevel, resp.RoleLevel)
	return nil
}

func CallbackBeforeSetGroupMemberInfo(ctx context.Context, req *group.SetGroupMemberInfo) (err error) {
	if !config.Config.Callback.CallbackBeforeSetGroupMemberInfo.Enable {
		return nil
	}
	defer func() {
		tracelog.SetCtxInfo(ctx, utils.GetFuncName(1), err, "req", *req)
	}()
	callbackReq := callbackstruct.CallbackBeforeSetGroupMemberInfoReq{
		CallbackCommand: constant.CallbackBeforeSetGroupMemberInfoCommand,
		OperationID:     tracelog.GetOperationID(ctx),
		GroupID:         req.GroupID,
		UserID:          req.UserID,
	}
	if req.Nickname != nil {
		callbackReq.Nickname = &req.Nickname.Value
	}
	if req.FaceURL != nil {
		callbackReq.FaceURL = &req.FaceURL.Value
	}
	if req.RoleLevel != nil {
		callbackReq.RoleLevel = &req.RoleLevel.Value
	}
	if req.Ex != nil {
		callbackReq.Ex = &req.Ex.Value
	}
	resp := &callbackstruct.CallbackBeforeSetGroupMemberInfoResp{}
	err = http.CallBackPostReturn(config.Config.Callback.CallbackUrl, callbackReq, resp, config.Config.Callback.CallbackBeforeSetGroupMemberInfo)
	if err != nil {
		return err
	}
	if resp.FaceURL != nil {
		req.FaceURL = wrapperspb.String(*resp.FaceURL)
	}
	if resp.Nickname != nil {
		req.Nickname = wrapperspb.String(*resp.Nickname)
	}
	if resp.RoleLevel != nil {
		req.RoleLevel = wrapperspb.Int32(*resp.RoleLevel)
	}
	if resp.Ex != nil {
		req.Ex = wrapperspb.String(*resp.Ex)
	}
	return nil
}
