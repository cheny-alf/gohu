// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"main/app/service/question/dao/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newQuestionContent(db *gorm.DB) questionContent {
	_questionContent := questionContent{}

	_questionContent.questionContentDo.UseDB(db)
	_questionContent.questionContentDo.UseModel(&model.QuestionContent{})

	tableName := _questionContent.questionContentDo.TableName()
	_questionContent.ALL = field.NewField(tableName, "*")
	_questionContent.QuestionID = field.NewInt64(tableName, "question_id")
	_questionContent.Title = field.NewString(tableName, "title")
	_questionContent.Topic = field.NewString(tableName, "topic")
	_questionContent.Tag = field.NewString(tableName, "tag")
	_questionContent.Content = field.NewString(tableName, "content")
	_questionContent.IPLoc = field.NewString(tableName, "ip_loc")
	_questionContent.Meta = field.NewString(tableName, "meta")
	_questionContent.CreateTime = field.NewTime(tableName, "create_time")
	_questionContent.UpdateTime = field.NewTime(tableName, "update_time")

	_questionContent.fillFieldMap()

	return _questionContent
}

type questionContent struct {
	questionContentDo questionContentDo

	ALL        field.Field
	QuestionID field.Int64
	Title      field.String
	Topic      field.String
	Tag        field.String
	Content    field.String
	IPLoc      field.String
	Meta       field.String
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (q questionContent) Table(newTableName string) *questionContent {
	q.questionContentDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q questionContent) As(alias string) *questionContent {
	q.questionContentDo.DO = *(q.questionContentDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *questionContent) updateTableName(table string) *questionContent {
	q.ALL = field.NewField(table, "*")
	q.QuestionID = field.NewInt64(table, "question_id")
	q.Title = field.NewString(table, "title")
	q.Topic = field.NewString(table, "topic")
	q.Tag = field.NewString(table, "tag")
	q.Content = field.NewString(table, "content")
	q.IPLoc = field.NewString(table, "ip_loc")
	q.Meta = field.NewString(table, "meta")
	q.CreateTime = field.NewTime(table, "create_time")
	q.UpdateTime = field.NewTime(table, "update_time")

	q.fillFieldMap()

	return q
}

func (q *questionContent) WithContext(ctx context.Context) *questionContentDo {
	return q.questionContentDo.WithContext(ctx)
}

func (q questionContent) TableName() string { return q.questionContentDo.TableName() }

func (q questionContent) Alias() string { return q.questionContentDo.Alias() }

func (q *questionContent) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *questionContent) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 9)
	q.fieldMap["question_id"] = q.QuestionID
	q.fieldMap["title"] = q.Title
	q.fieldMap["topic"] = q.Topic
	q.fieldMap["tag"] = q.Tag
	q.fieldMap["content"] = q.Content
	q.fieldMap["ip_loc"] = q.IPLoc
	q.fieldMap["meta"] = q.Meta
	q.fieldMap["create_time"] = q.CreateTime
	q.fieldMap["update_time"] = q.UpdateTime
}

func (q questionContent) clone(db *gorm.DB) questionContent {
	q.questionContentDo.ReplaceDB(db)
	return q
}

type questionContentDo struct{ gen.DO }

func (q questionContentDo) Debug() *questionContentDo {
	return q.withDO(q.DO.Debug())
}

func (q questionContentDo) WithContext(ctx context.Context) *questionContentDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q questionContentDo) ReadDB() *questionContentDo {
	return q.Clauses(dbresolver.Read)
}

func (q questionContentDo) WriteDB() *questionContentDo {
	return q.Clauses(dbresolver.Write)
}

func (q questionContentDo) Clauses(conds ...clause.Expression) *questionContentDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q questionContentDo) Returning(value interface{}, columns ...string) *questionContentDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q questionContentDo) Not(conds ...gen.Condition) *questionContentDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q questionContentDo) Or(conds ...gen.Condition) *questionContentDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q questionContentDo) Select(conds ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q questionContentDo) Where(conds ...gen.Condition) *questionContentDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q questionContentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *questionContentDo {
	return q.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (q questionContentDo) Order(conds ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q questionContentDo) Distinct(cols ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q questionContentDo) Omit(cols ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q questionContentDo) Join(table schema.Tabler, on ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q questionContentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q questionContentDo) RightJoin(table schema.Tabler, on ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q questionContentDo) Group(cols ...field.Expr) *questionContentDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q questionContentDo) Having(conds ...gen.Condition) *questionContentDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q questionContentDo) Limit(limit int) *questionContentDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q questionContentDo) Offset(offset int) *questionContentDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q questionContentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *questionContentDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q questionContentDo) Unscoped() *questionContentDo {
	return q.withDO(q.DO.Unscoped())
}

func (q questionContentDo) Create(values ...*model.QuestionContent) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q questionContentDo) CreateInBatches(values []*model.QuestionContent, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q questionContentDo) Save(values ...*model.QuestionContent) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q questionContentDo) First() (*model.QuestionContent, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionContent), nil
	}
}

func (q questionContentDo) Take() (*model.QuestionContent, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionContent), nil
	}
}

func (q questionContentDo) Last() (*model.QuestionContent, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionContent), nil
	}
}

func (q questionContentDo) Find() ([]*model.QuestionContent, error) {
	result, err := q.DO.Find()
	return result.([]*model.QuestionContent), err
}

func (q questionContentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QuestionContent, err error) {
	buf := make([]*model.QuestionContent, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q questionContentDo) FindInBatches(result *[]*model.QuestionContent, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q questionContentDo) Attrs(attrs ...field.AssignExpr) *questionContentDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q questionContentDo) Assign(attrs ...field.AssignExpr) *questionContentDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q questionContentDo) Joins(fields ...field.RelationField) *questionContentDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q questionContentDo) Preload(fields ...field.RelationField) *questionContentDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q questionContentDo) FirstOrInit() (*model.QuestionContent, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionContent), nil
	}
}

func (q questionContentDo) FirstOrCreate() (*model.QuestionContent, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionContent), nil
	}
}

func (q questionContentDo) FindByPage(offset int, limit int) (result []*model.QuestionContent, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q questionContentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q questionContentDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q *questionContentDo) withDO(do gen.Dao) *questionContentDo {
	q.DO = *do.(*gen.DO)
	return q
}