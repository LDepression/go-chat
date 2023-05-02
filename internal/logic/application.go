package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	tx2 "go-chat/internal/dao/mysql/tx"
	reply2 "go-chat/internal/model/reply"
	"go-chat/internal/model/request"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type application struct{}

func (application) CreateApplication(ctx *gin.Context, req request.CreateApplicationReq, accountID uint) (*reply2.CreateApplicationRep, errcode.Err) {
	var rep = new(reply2.CreateApplicationRep)
	tx := tx2.NewApplicationTX()
	applicationID, err := tx.CreateApplicationWithTX(uint64(accountID), req.AccountID, req.ApplicationMsg)
	if err != nil {
		switch err {
		case tx2.ErrHasThisFriend:
			return nil, myerr.FriendHasAlreadyExists
		case tx2.ErrIsSelf:
			return nil, myerr.CanNotAddSelf
		default:
			return nil, errcode.ErrServer
		}
	}
	rep.ApplicationID = applicationID
	rep.Status = string(query.WAITING)

	//TODO:提示对方有新的消息
	return rep, nil
}

func (application) DeleteApplication(ctx *gin.Context, accountID uint64) errcode.Err {
	//
	q := query.NewApplication()
	if err := q.DeleteApplication(accountID); err != nil {
		return errcode.ErrServer
	}
	return nil
}

func (application) AcceptApplication(c *gin.Context, applicantID, receiverID uint) errcode.Err {
	qApplication := query.NewQueryApplication()
	_, err := qApplication.GetApplicationByID(applicantID, receiverID)
	if err != nil {
		zap.S().Errorf("qApplication.GetApplicationByID failed, err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return myerr.ApplicationNotFound
		}
		return errcode.ErrServer
	}
	applicationTX := tx2.NewApplicationTX()
	err = applicationTX.AcceptApplicationWithTX(applicantID, receiverID)
	if err != nil {
		zap.S().Errorf("applicationTX.AcceptApplication failed, err:%v", err)
		return errcode.ErrNotFound
	}
	//TODO:提示对方有新的消息

	return nil
}

func (application) RefuseApplication(c *gin.Context, applicantID, receiverID uint, refuseMsg string) errcode.Err {
	qApplication := query.NewQueryApplication()
	_, err := qApplication.GetApplicationByID(applicantID, receiverID)
	if err != nil {
		zap.S().Errorf("qApplication.GetApplicationByID failed, err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return myerr.ApplicationNotFound
		}
		return errcode.ErrServer
	}
	err = qApplication.RefuseApplication(applicantID, receiverID, refuseMsg)
	if err != nil {
		zap.S().Errorf("qApplication.RefuseApplication failed,err:%v", err)
		return errcode.ErrServer
	}
	//TODO:提示对方有新的消息

	return nil
}

func (application) GetApplicationsList(c *gin.Context, accountID, limit, offset uint) (*reply2.ApplicationsList, errcode.Err) {
	qApplication := query.NewQueryApplication()
	applicationList, totalCount, err := qApplication.GetApplicationsList(accountID, limit, offset)
	if err != nil {
		zap.S().Errorf("qApplication.GetApplicationList failes,err:%v", err)
		return &reply2.ApplicationsList{}, errcode.ErrServer
	}
	if totalCount == 0 {
		return &reply2.ApplicationsList{}, nil
	}
	qAccount := query.NewQueryAccount()
	selfAccount, err := qAccount.GetAccountByID(accountID)
	if err != nil {
		zap.S().Errorf("GetApplicationList,qAccount.GetAccountByID failes,err:%v", err)
		return nil, errcode.ErrServer
	}
	selfEasyAccount := &reply2.EasyAccount{
		AccountID: uint(selfAccount.ID),
		Name:      selfAccount.Name,
		Avatar:    selfAccount.Avatar,
	}
	// accountID1是申请者，accountID2是接受者
	reply2ApplicationList := make([]*reply2.GetApplication, 0, len(applicationList))
	for _, v := range applicationList {
		applicantInfo, receiverInfo, err := InquireApplicantAndReceiver(uint(v.AccountID1), uint(v.AccountID2), accountID, selfEasyAccount)
		if err != nil {
			zap.S().Errorf("InquireApplicantAndReceiver failed,err:%v", err)
			return &reply2.ApplicationsList{}, errcode.ErrServer
		}
		reply2ApplicationList = append(reply2ApplicationList, &reply2.GetApplication{
			Applicant: applicantInfo,
			Receiver:  receiverInfo,
			ApplyMsg:  v.ApplyMsg,
			RefuseMsg: v.RefuseMsg,
			Status:    v.Status,
		})
	}
	return &reply2.ApplicationsList{
		ApplicationList: reply2ApplicationList,
		Total:           totalCount,
	}, nil
}

func InquireApplicantAndReceiver(applicantID, receiverID, accountID uint, selfEasyAccount *reply2.EasyAccount) (*reply2.EasyAccount, *reply2.EasyAccount, errcode.Err) {
	var applicant, receiver *reply2.EasyAccount
	qAccount := query.NewQueryAccount()

	if applicantID == accountID {
		applicant = selfEasyAccount
		receiverInfo, err := qAccount.GetAccountByID(receiverID)
		if err != nil {
			return nil, nil, errcode.ErrServer
		}
		receiver = &reply2.EasyAccount{
			AccountID: uint(receiverInfo.ID),
			Name:      receiverInfo.Name,
			Avatar:    receiverInfo.Avatar,
		}
	} else if receiverID == accountID {
		receiver = selfEasyAccount
		applicantInfo, err := qAccount.GetAccountByID(applicantID)
		if err != nil {
			return nil, nil, errcode.ErrServer
		}
		applicant = &reply2.EasyAccount{
			AccountID: uint(applicantInfo.ID),
			Name:      applicantInfo.Name,
			Avatar:    applicantInfo.Avatar,
		}
	}
	return receiver, applicant, nil
}
