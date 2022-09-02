package logic

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"main/app/common/log"
	"net/http"

	"main/app/service/user/rpc/info/internal/svc"
	"main/app/service/user/rpc/info/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetObjInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetObjInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetObjInfoLogic {
	return &GetObjInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetObjInfoLogic) GetObjInfo(in *pb.GetObjInfoReq) (res *pb.GetObjInfoRes, err error) {
	logger := log.GetSugaredLogger()
	logger.Debugf("recv message: %v", in.String())

	switch in.ObjType {
	case 0:
		// 问题
		questionIds := make([]int64, 0)

		questionIdsCache, err := l.svcCtx.Rdb.SMembers(l.ctx,
			fmt.Sprintf("question_%d", in.UserId)).Result()
		if err == nil {
			if len(questionIdsCache) > 1 {
				// 获取到缓存
				for _, questionIdCache := range questionIdsCache {
					if questionIdCache != "0" {
						questionIds = append(questionIds, cast.ToInt64(questionIdCache))
					}
				}
				res = &pb.GetObjInfoRes{
					Code: http.StatusOK,
					Msg:  "get question ids successfully",
					Ok:   true,
					Data: &pb.GetObjInfoRes_Data{Ids: questionIds},
				}
				logger.Debugf("send message: %v", err)
				return res, nil
			}
		} else {
			logger.Errorf("get [question] cache failed, err: %v", err)
		}

		questionSubjectModel := l.svcCtx.QuestionModel.QuestionSubject

		questionSubjects, err := questionSubjectModel.WithContext(l.ctx).
			Select(questionSubjectModel.ID, questionSubjectModel.UserID).
			Where(questionSubjectModel.UserID.Eq(in.UserId)).Find()
		switch err {
		case gorm.ErrRecordNotFound:
			//设置空缓存
			l.svcCtx.Rdb.SAdd(l.ctx,
				fmt.Sprintf("question_%d", in.UserId),
				0)
			res = &pb.GetObjInfoRes{
				Code: http.StatusForbidden,
				Msg:  "question not found",
				Ok:   false,
				Data: nil,
			}
			logger.Debugf("send message: %v", res.String())
			return res, nil
		case nil:

		default:
			logger.Debugf("get question info failed, err: mysql err, %v", err)
			res = &pb.GetObjInfoRes{
				Code: http.StatusInternalServerError,
				Msg:  "internal err",
				Ok:   false,
				Data: nil,
			}
			logger.Debugf("send message: %v", res.String())
			return res, nil
		}

		res = &pb.GetObjInfoRes{
			Code: http.StatusOK,
			Msg:  "get question successfully",
			Ok:   true,
			Data: &pb.GetObjInfoRes_Data{Ids: make([]int64, 0)},
		}

		for _, questionSubject := range questionSubjects {
			res.Data.Ids = append(res.Data.Ids, questionSubject.ID)
			// 设置缓存
			l.svcCtx.Rdb.SAdd(l.ctx,
				fmt.Sprintf("question_%d", in.UserId),
				questionSubject.ID)
		}

		logger.Debugf("send message: %v", res.String())
		return res, nil

	case 1:
		// 回答
		answerIds := make([]int64, 0)

		answerIdsCache, err := l.svcCtx.Rdb.SMembers(l.ctx,
			fmt.Sprintf("answer_%d", in.UserId)).Result()
		if err != nil {
			if len(answerIdsCache) > 1 {
				for _, answerIdCache := range answerIdsCache {
					if answerIdCache != "0" {
						answerIds = append(answerIds, cast.ToInt64(answerIdCache))
					}
				}
				res = &pb.GetObjInfoRes{
					Code: http.StatusOK,
					Msg:  "get answer ids successfully",
					Ok:   true,
					Data: &pb.GetObjInfoRes_Data{Ids: answerIds},
				}
			}
		}

		answerIndexModel := l.svcCtx.QuestionModel.AnswerIndex

		answerIndices, err := answerIndexModel.WithContext(l.ctx).
			Select(answerIndexModel.ID, answerIndexModel.UserID).
			Where(answerIndexModel.UserID.Eq(in.UserId)).Find()
		switch err {
		case gorm.ErrRecordNotFound:
			l.svcCtx.Rdb.SAdd(l.ctx,
				fmt.Sprintf("answer_%d", in.UserId),
				0)
			res = &pb.GetObjInfoRes{
				Code: http.StatusForbidden,
				Msg:  "question not found",
				Ok:   false,
				Data: nil,
			}
			logger.Debugf("send message: %v", res.String())
			return res, nil
		case nil:

		default:
			logger.Debugf("get question info failed, err: mysql err, %v", err)
			res = &pb.GetObjInfoRes{
				Code: http.StatusInternalServerError,
				Msg:  "internal err",
				Ok:   false,
				Data: nil,
			}
			logger.Debugf("send message: %v", res.String())
			return res, nil
		}
		res = &pb.GetObjInfoRes{
			Code: http.StatusOK,
			Msg:  "get question successfully",
			Ok:   true,
			Data: &pb.GetObjInfoRes_Data{Ids: make([]int64, 0)},
		}

		for _, answerIndex := range answerIndices {
			res.Data.Ids = append(res.Data.Ids, answerIndex.ID)
			l.svcCtx.Rdb.SAdd(l.ctx,
				fmt.Sprintf("answer_%d", in.UserId),
				answerIndex.ID)
		}

		logger.Debugf("send message: %v", res.String())
		return res, nil

	case 2:
		// 文章
		return nil, nil

	default:
		res = &pb.GetObjInfoRes{
			Code: http.StatusBadRequest,
			Msg:  `invalid param "obj_type"`,
			Ok:   false,
			Data: nil,
		}
		logger.Debugf("send message: %v", res.String())
		return res, nil
	}
}
