package logic

import (
	"context"
	"main/app/common/log"
	"main/app/service/question/dao/model"
	"main/app/service/question/rpc/crud/internal/svc"
	"main/app/service/question/rpc/crud/pb"
	"main/app/utils/net/ip"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishQuestionLogic {
	return &PublishQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishQuestionLogic) PublishQuestion(in *pb.PublishQuestionReq) (res *pb.PublishQuestionRes, err error) {
	logger := log.GetSugaredLogger()
	logger.Debugf("recv message: %v", in.String())
	j := gjson.Parse(in.UserDetails)

	questionSubjectModel := l.svcCtx.QuestionModel.QuestionSubject
	questionContentModel := l.svcCtx.QuestionModel.QuestionContent

	err = questionSubjectModel.WithContext(l.ctx).Create(&model.QuestionSubject{
		UserID: j.Get("user_id").Int(),
	})
	if err != nil {
		logger.Errorf("publish question failed, err: mysql err")
		res = &pb.PublishQuestionRes{
			Code: http.StatusInternalServerError,
			Mag:  "publish question failed, err: internal err",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return nil, err
	}

	err = questionContentModel.WithContext(l.ctx).Create(&model.QuestionContent{
		Title:   in.Title,
		Topic:   in.Topic,
		Tag:     in.Tag,
		Content: in.Content,
		IPLoc:   ip.GetIpLocFromApi(j.Get("last_ip").String()),
	})
	if err != nil {
		logger.Errorf("publish question failed, err: mysql err")
		res = &pb.PublishQuestionRes{
			Code: http.StatusInternalServerError,
			Mag:  "publish question failed, err: internal err",
			Ok:   false,
		}
		logger.Debugf("send message: %v", res.String())
		return nil, err
	}

	res = &pb.PublishQuestionRes{
		Code: http.StatusOK,
		Mag:  "publish question successfully",
		Ok:   true,
	}
	logger.Debugf("send message: %v", res.String())
	return res, nil
}