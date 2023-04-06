package logic

import (
	"github.com/gin-gonic/gin"
	"go-chat/internal/dao/mysql/query"
	"go-chat/internal/dao/mysql/tx"
	"go-chat/internal/model/reply"
	"go-chat/internal/myerr"
	"go-chat/internal/pkg/app/errcode"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type application struct {
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
	applicationTX := tx.NewApplicationTX()
	err = applicationTX.AcceptApplicationWithTX(applicantID, receiverID)
	if err != nil {
		zap.S().Errorf("applicationTX.AcceptApplication failed, err:%v", err)
		return errcode.ErrNotFound
	}

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
	return nil
}

func (application) GetApplicationsList(c *gin.Context, accountID, limit, offset uint) (*reply.ApplicationsList, errcode.Err) {
	qApplication := query.NewQueryApplication()
	applicationList, totalCount, err := qApplication.GetApplicationsList(accountID, limit, offset)
	if err != nil {
		zap.S().Errorf("qApplication.GetApplicationList failes,err:%v", err)
		return &reply.ApplicationsList{}, errcode.ErrServer
	}
	if totalCount == 0 {
		return &reply.ApplicationsList{}, nil
	}
	qAccount := query.NewQueryAccount()
	selfAccount, err := qAccount.GetAccountByID(accountID)
	if err != nil {
		zap.S().Errorf("GetApplicationList,qAccount.GetAccountByID failes,err:%v", err)
		return nil, errcode.ErrServer
	}
	selfEasyAccount := &reply.EasyAccount{
		AccountID: selfAccount.ID,
		Name:      selfAccount.Name,
		Avatar:    selfAccount.Avatar,
	}
	// accountID1是申请者，accountID2是接受者
	replyApplicationList := make([]*reply.GetApplication, 0, len(applicationList))
	for _, v := range applicationList {
		applicantInfo, receiverInfo, err := InquireApplicantAndReceiver(uint(v.AccountID1), uint(v.AccountID2), accountID, selfEasyAccount)
		if err != nil {
			zap.S().Errorf("InquireApplicantAndReceiver failed,err:%v", err)
			return &reply.ApplicationsList{}, errcode.ErrServer
		}
		replyApplicationList = append(replyApplicationList, &reply.GetApplication{
			Applicant: applicantInfo,
			Receiver:  receiverInfo,
			ApplyMsg:  v.ApplyMsg,
			RefuseMsg: v.RefuseMsg,
			Status:    v.Status,
		})
	}
	return &reply.ApplicationsList{
		ApplicationList: replyApplicationList,
		Total:           totalCount,
	}, nil
}

func InquireApplicantAndReceiver(applicantID, receiverID, accountID uint, selfEasyAccount *reply.EasyAccount) (*reply.EasyAccount, *reply.EasyAccount, errcode.Err) {
	var applicant, receiver *reply.EasyAccount
	qAccount := query.NewQueryAccount()

	if applicantID == accountID {
		applicant = selfEasyAccount
		receiverInfo, err := qAccount.GetAccountByID(receiverID)
		if err != nil {
			return nil, nil, errcode.ErrServer
		}
		receiver = &reply.EasyAccount{
			AccountID: receiverInfo.ID,
			Name:      receiverInfo.Name,
			Avatar:    receiverInfo.Avatar,
		}
	} else if receiverID == accountID {
		receiver = selfEasyAccount
		applicantInfo, err := qAccount.GetAccountByID(applicantID)
		if err != nil {
			return nil, nil, errcode.ErrServer
		}
		applicant = &reply.EasyAccount{
			AccountID: applicantInfo.ID,
			Name:      applicantInfo.Name,
			Avatar:    applicantInfo.Avatar,
		}
	}
	return receiver, applicant, nil
}
