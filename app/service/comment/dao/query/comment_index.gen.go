// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"main/app/service/comment/dao/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newCommentIndex(db *gorm.DB) commentIndex {
	_commentIndex := commentIndex{}

	_commentIndex.commentIndexDo.UseDB(db)
	_commentIndex.commentIndexDo.UseModel(&model.CommentIndex{})

	tableName := _commentIndex.commentIndexDo.TableName()
	_commentIndex.ALL = field.NewField(tableName, "*")
	_commentIndex.ID = field.NewInt64(tableName, "id")
	_commentIndex.SubjectID = field.NewInt64(tableName, "subject_id")
	_commentIndex.UserID = field.NewInt64(tableName, "user_id")
	_commentIndex.IPLoc = field.NewString(tableName, "ip_loc")
	_commentIndex.RootID = field.NewInt64(tableName, "root_id")
	_commentIndex.CommentFloor = field.NewInt32(tableName, "comment_floor")
	_commentIndex.CommentID = field.NewInt64(tableName, "comment_id")
	_commentIndex.ReplyFloor = field.NewInt32(tableName, "reply_floor")
	_commentIndex.ApproveCount = field.NewInt32(tableName, "approve_count")
	_commentIndex.State = field.NewInt32(tableName, "state")
	_commentIndex.Attrs = field.NewInt32(tableName, "attrs")
	_commentIndex.CreateTime = field.NewTime(tableName, "create_time")
	_commentIndex.UpdateTime = field.NewTime(tableName, "update_time")

	_commentIndex.fillFieldMap()

	return _commentIndex
}

type commentIndex struct {
	commentIndexDo commentIndexDo

	ALL          field.Field
	ID           field.Int64
	SubjectID    field.Int64
	UserID       field.Int64
	IPLoc        field.String
	RootID       field.Int64
	CommentFloor field.Int32
	CommentID    field.Int64
	ReplyFloor   field.Int32
	ApproveCount field.Int32
	State        field.Int32
	Attrs        field.Int32
	CreateTime   field.Time
	UpdateTime   field.Time

	fieldMap map[string]field.Expr
}

func (c commentIndex) Table(newTableName string) *commentIndex {
	c.commentIndexDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentIndex) As(alias string) *commentIndex {
	c.commentIndexDo.DO = *(c.commentIndexDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentIndex) updateTableName(table string) *commentIndex {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt64(table, "id")
	c.SubjectID = field.NewInt64(table, "subject_id")
	c.UserID = field.NewInt64(table, "user_id")
	c.IPLoc = field.NewString(table, "ip_loc")
	c.RootID = field.NewInt64(table, "root_id")
	c.CommentFloor = field.NewInt32(table, "comment_floor")
	c.CommentID = field.NewInt64(table, "comment_id")
	c.ReplyFloor = field.NewInt32(table, "reply_floor")
	c.ApproveCount = field.NewInt32(table, "approve_count")
	c.State = field.NewInt32(table, "state")
	c.Attrs = field.NewInt32(table, "attrs")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")

	c.fillFieldMap()

	return c
}

func (c *commentIndex) WithContext(ctx context.Context) *commentIndexDo {
	return c.commentIndexDo.WithContext(ctx)
}

func (c commentIndex) TableName() string { return c.commentIndexDo.TableName() }

func (c commentIndex) Alias() string { return c.commentIndexDo.Alias() }

func (c *commentIndex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentIndex) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 13)
	c.fieldMap["id"] = c.ID
	c.fieldMap["subject_id"] = c.SubjectID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["ip_loc"] = c.IPLoc
	c.fieldMap["root_id"] = c.RootID
	c.fieldMap["comment_floor"] = c.CommentFloor
	c.fieldMap["comment_id"] = c.CommentID
	c.fieldMap["reply_floor"] = c.ReplyFloor
	c.fieldMap["approve_count"] = c.ApproveCount
	c.fieldMap["state"] = c.State
	c.fieldMap["attrs"] = c.Attrs
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
}

func (c commentIndex) clone(db *gorm.DB) commentIndex {
	c.commentIndexDo.ReplaceDB(db)
	return c
}

type commentIndexDo struct{ gen.DO }

func (c commentIndexDo) Debug() *commentIndexDo {
	return c.withDO(c.DO.Debug())
}

func (c commentIndexDo) WithContext(ctx context.Context) *commentIndexDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentIndexDo) ReadDB() *commentIndexDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentIndexDo) WriteDB() *commentIndexDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentIndexDo) Clauses(conds ...clause.Expression) *commentIndexDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentIndexDo) Returning(value interface{}, columns ...string) *commentIndexDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentIndexDo) Not(conds ...gen.Condition) *commentIndexDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentIndexDo) Or(conds ...gen.Condition) *commentIndexDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentIndexDo) Select(conds ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentIndexDo) Where(conds ...gen.Condition) *commentIndexDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentIndexDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *commentIndexDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentIndexDo) Order(conds ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentIndexDo) Distinct(cols ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentIndexDo) Omit(cols ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentIndexDo) Join(table schema.Tabler, on ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentIndexDo) LeftJoin(table schema.Tabler, on ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentIndexDo) RightJoin(table schema.Tabler, on ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentIndexDo) Group(cols ...field.Expr) *commentIndexDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentIndexDo) Having(conds ...gen.Condition) *commentIndexDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentIndexDo) Limit(limit int) *commentIndexDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentIndexDo) Offset(offset int) *commentIndexDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentIndexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *commentIndexDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentIndexDo) Unscoped() *commentIndexDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentIndexDo) Create(values ...*model.CommentIndex) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentIndexDo) CreateInBatches(values []*model.CommentIndex, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentIndexDo) Save(values ...*model.CommentIndex) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentIndexDo) First() (*model.CommentIndex, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentIndex), nil
	}
}

func (c commentIndexDo) Take() (*model.CommentIndex, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentIndex), nil
	}
}

func (c commentIndexDo) Last() (*model.CommentIndex, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentIndex), nil
	}
}

func (c commentIndexDo) Find() ([]*model.CommentIndex, error) {
	result, err := c.DO.Find()
	return result.([]*model.CommentIndex), err
}

func (c commentIndexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommentIndex, err error) {
	buf := make([]*model.CommentIndex, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentIndexDo) FindInBatches(result *[]*model.CommentIndex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentIndexDo) Attrs(attrs ...field.AssignExpr) *commentIndexDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentIndexDo) Assign(attrs ...field.AssignExpr) *commentIndexDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentIndexDo) Joins(fields ...field.RelationField) *commentIndexDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentIndexDo) Preload(fields ...field.RelationField) *commentIndexDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentIndexDo) FirstOrInit() (*model.CommentIndex, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentIndex), nil
	}
}

func (c commentIndexDo) FirstOrCreate() (*model.CommentIndex, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentIndex), nil
	}
}

func (c commentIndexDo) FindByPage(offset int, limit int) (result []*model.CommentIndex, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentIndexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentIndexDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c *commentIndexDo) withDO(do gen.Dao) *commentIndexDo {
	c.DO = *do.(*gen.DO)
	return c
}
