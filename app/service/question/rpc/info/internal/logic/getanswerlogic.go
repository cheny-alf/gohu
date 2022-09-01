package logic

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"main/app/common/log"
	"net/http"
	"time"

	"main/app/service/question/rpc/info/internal/svc"
	"main/app/service/question/rpc/info/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAnswerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAnswerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAnswerLogic {
	return &GetAnswerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAnswerLogic) GetAnswer(in *pb.GetAnswerReq) (res *pb.GetAnswerRes, err error) {
	logger := log.GetSugaredLogger()
	logger.Debugf("recv message: %v", in.String())

	resData := &pb.GetAnswerRes_Data{}

	answerIndexBytes, err := l.svcCtx.Rdb.Get(l.ctx,
		fmt.Sprintf("anwerIndex_%d", in.AnswerId)).Bytes()
	if err != nil {
		logger.Errorf("get answerIndex cache failed, err: %v", err)

		answerIndexModel := l.svcCtx.QuestionModel.AnswerIndex

		answerIndex, err := answerIndexModel.WithContext(l.ctx).
			Where(answerIndexModel.ID.Eq(in.AnswerId)).
			First()
		if err != nil {
			logger.Errorf("get answerIndex failed, err: mysql err, %v", err)
			res = &pb.GetAnswerRes{
				Code: http.StatusInternalServerError,
				Msg:  "internal err",
				Ok:   false,
			}
			logger.Debugf("send message: %v", err)
			return res, nil
		}

		answerIndexProto := &pb.AnswerIndex{
			Id:           answerIndex.ID,
			QuestionId:   answerIndex.QuestionID,
			UserId:       answerIndex.UserID,
			IpLoc:        answerIndex.IPLoc,
			ApproveCount: answerIndex.ApproveCount,
			LikeCount:    answerIndex.LikeCount,
			CollectCount: answerIndex.CollectCount,
			State:        answerIndex.State,
			Attrs:        answerIndex.Attrs,
			CreateTime:   answerIndex.CreateTime.String(),
			UpdateTime:   answerIndex.UpdateTime.String(),
		}

		resData.AnswerIndex = answerIndexProto

		answerIndexBytes, err = proto.Marshal(answerIndexProto)
		if err != nil {
			logger.Errorf("marshal proto failed, err: %v")
		} else {
			l.svcCtx.Rdb.Set(l.ctx,
				fmt.Sprintf("answerIndex_%d", answerIndex.ID),
				answerIndexBytes,
				time.Second*86400)
		}
	}

	answerContentBytes, err := l.svcCtx.Rdb.Get(l.ctx,
		fmt.Sprintf("answerContent_%d", in.AnswerId)).Bytes()
	if err != nil {
		logger.Errorf("get answerContent cache failed, err: %v")

		answerContentModel := l.svcCtx.QuestionModel.AnswerContent

		answerContent, err := answerContentModel.WithContext(l.ctx).
			Where(answerContentModel.AnswerID.Eq(in.AnswerId)).
			First()
		if err != nil {
			logger.Errorf("get answerContent failed, err: mysql err, %v", err)
			res = &pb.GetAnswerRes{
				Code: http.StatusInternalServerError,
				Msg:  "internal err",
				Ok:   false,
			}
			logger.Debugf("send message: %v", res.String())
			return res, nil
		}

		answerContentProto := &pb.AnswerContent{
			AnswerId:   answerContent.AnswerID,
			Content:    answerContent.Content,
			Meta:       answerContent.Meta,
			CreateTime: answerContent.CreateTime.String(),
			UpdateTime: answerContent.UpdateTime.String(),
		}

		resData.AnswerContent = answerContentProto

		answerContentBytes, err = proto.Marshal(answerContentProto)
		if err != nil {
			logger.Errorf("marshal proto failed, err: %v", err)
		} else {
			l.svcCtx.Rdb.Set(l.ctx,
				fmt.Sprintf("answerContent_%d", answerContent.AnswerID),
				answerContentBytes,
				time.Second*86400)
		}
	}

	res = &pb.GetAnswerRes{
		Code: http.StatusOK,
		Msg:  "get answer successfully",
		Ok:   true,
		Data: resData,
	}
	logger.Debugf("send message: %v", res.String())
	return res, nil
}