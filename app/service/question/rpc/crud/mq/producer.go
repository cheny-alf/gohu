package mq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	mynsq "main/app/common/mq/nsq"
)

type AnswerSubjectMessage struct {
	Action string            `json:"action"`
	Data   AnswerSubjectData `json:"data"`
}

type AnswerSubjectData struct {
	ObjType int32 `json:"obj_type"`
	ObjId   int64 `json:"obj_id"`
}

var producer *nsq.Producer

func InitProducer() (err error) {
	producer, err = mynsq.NewProducer()
	if err != nil {
		return err
	}
	return nil
}

func Publish(objType int32, objId int64, action string) (err error) {
	if producer == nil {
		return errors.New("empty producer")
	}
	message, err := json.Marshal(AnswerSubjectMessage{
		Action: action,
		Data: AnswerSubjectData{
			ObjType: objType,
			ObjId:   objId,
		},
	})
	if err != nil {
		return fmt.Errorf("marshal message failed, %v", err)
	}
	err = producer.Publish("comment-subject", message)
	if err != nil {
		return fmt.Errorf("publish msg to nsq failed, %v", err)
	}
	return nil
}