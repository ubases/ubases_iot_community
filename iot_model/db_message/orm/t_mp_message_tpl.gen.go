// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_message/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTMpMessageTpl(db *gorm.DB) tMpMessageTpl {
	_tMpMessageTpl := tMpMessageTpl{}

	_tMpMessageTpl.tMpMessageTplDo.UseDB(db)
	_tMpMessageTpl.tMpMessageTplDo.UseModel(&model.TMpMessageTpl{})

	tableName := _tMpMessageTpl.tMpMessageTplDo.TableName()
	_tMpMessageTpl.ALL = field.NewField(tableName, "*")
	_tMpMessageTpl.Id = field.NewInt64(tableName, "id")
	_tMpMessageTpl.TplCode = field.NewString(tableName, "tpl_code")
	_tMpMessageTpl.TplName = field.NewString(tableName, "tpl_name")
	_tMpMessageTpl.TplContent = field.NewString(tableName, "tpl_content")
	_tMpMessageTpl.PushType = field.NewInt32(tableName, "push_type")
	_tMpMessageTpl.MessageType = field.NewInt32(tableName, "message_type")
	_tMpMessageTpl.AgentType = field.NewInt32(tableName, "agent_type")
	_tMpMessageTpl.Lang = field.NewString(tableName, "lang")
	_tMpMessageTpl.ExpireHour = field.NewInt32(tableName, "expire_hour")
	_tMpMessageTpl.CreatedBy = field.NewInt64(tableName, "created_by")
	_tMpMessageTpl.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tMpMessageTpl.CreatedAt = field.NewTime(tableName, "created_at")
	_tMpMessageTpl.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tMpMessageTpl.DeletedAt = field.NewField(tableName, "deleted_at")

	_tMpMessageTpl.fillFieldMap()

	return _tMpMessageTpl
}

type tMpMessageTpl struct {
	tMpMessageTplDo tMpMessageTplDo

	ALL         field.Field
	Id          field.Int64
	TplCode     field.String
	TplName     field.String
	TplContent  field.String
	PushType    field.Int32
	MessageType field.Int32
	AgentType   field.Int32
	Lang        field.String
	ExpireHour  field.Int32
	CreatedBy   field.Int64
	UpdatedBy   field.Int64
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (t tMpMessageTpl) Table(newTableName string) *tMpMessageTpl {
	t.tMpMessageTplDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tMpMessageTpl) As(alias string) *tMpMessageTpl {
	t.tMpMessageTplDo.DO = *(t.tMpMessageTplDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tMpMessageTpl) updateTableName(table string) *tMpMessageTpl {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.TplCode = field.NewString(table, "tpl_code")
	t.TplName = field.NewString(table, "tpl_name")
	t.TplContent = field.NewString(table, "tpl_content")
	t.PushType = field.NewInt32(table, "push_type")
	t.MessageType = field.NewInt32(table, "message_type")
	t.AgentType = field.NewInt32(table, "agent_type")
	t.Lang = field.NewString(table, "lang")
	t.ExpireHour = field.NewInt32(table, "expire_hour")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tMpMessageTpl) WithContext(ctx context.Context) *tMpMessageTplDo {
	return t.tMpMessageTplDo.WithContext(ctx)
}

func (t tMpMessageTpl) TableName() string { return t.tMpMessageTplDo.TableName() }

func (t tMpMessageTpl) Alias() string { return t.tMpMessageTplDo.Alias() }

func (t *tMpMessageTpl) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tMpMessageTpl) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 14)
	t.fieldMap["id"] = t.Id
	t.fieldMap["tpl_code"] = t.TplCode
	t.fieldMap["tpl_name"] = t.TplName
	t.fieldMap["tpl_content"] = t.TplContent
	t.fieldMap["push_type"] = t.PushType
	t.fieldMap["message_type"] = t.MessageType
	t.fieldMap["agent_type"] = t.AgentType
	t.fieldMap["lang"] = t.Lang
	t.fieldMap["expire_hour"] = t.ExpireHour
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t tMpMessageTpl) clone(db *gorm.DB) tMpMessageTpl {
	t.tMpMessageTplDo.ReplaceDB(db)
	return t
}

type tMpMessageTplDo struct{ gen.DO }

func (t tMpMessageTplDo) Debug() *tMpMessageTplDo {
	return t.withDO(t.DO.Debug())
}

func (t tMpMessageTplDo) WithContext(ctx context.Context) *tMpMessageTplDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tMpMessageTplDo) Clauses(conds ...clause.Expression) *tMpMessageTplDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tMpMessageTplDo) Returning(value interface{}, columns ...string) *tMpMessageTplDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tMpMessageTplDo) Not(conds ...gen.Condition) *tMpMessageTplDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tMpMessageTplDo) Or(conds ...gen.Condition) *tMpMessageTplDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tMpMessageTplDo) Select(conds ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tMpMessageTplDo) Where(conds ...gen.Condition) *tMpMessageTplDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tMpMessageTplDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tMpMessageTplDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tMpMessageTplDo) Order(conds ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tMpMessageTplDo) Distinct(cols ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tMpMessageTplDo) Omit(cols ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tMpMessageTplDo) Join(table schema.Tabler, on ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tMpMessageTplDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tMpMessageTplDo) RightJoin(table schema.Tabler, on ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tMpMessageTplDo) Group(cols ...field.Expr) *tMpMessageTplDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tMpMessageTplDo) Having(conds ...gen.Condition) *tMpMessageTplDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tMpMessageTplDo) Limit(limit int) *tMpMessageTplDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tMpMessageTplDo) Offset(offset int) *tMpMessageTplDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tMpMessageTplDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tMpMessageTplDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tMpMessageTplDo) Unscoped() *tMpMessageTplDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tMpMessageTplDo) Create(values ...*model.TMpMessageTpl) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tMpMessageTplDo) CreateInBatches(values []*model.TMpMessageTpl, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tMpMessageTplDo) Save(values ...*model.TMpMessageTpl) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tMpMessageTplDo) First() (*model.TMpMessageTpl, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTpl), nil
	}
}

func (t tMpMessageTplDo) Take() (*model.TMpMessageTpl, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTpl), nil
	}
}

func (t tMpMessageTplDo) Last() (*model.TMpMessageTpl, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTpl), nil
	}
}

func (t tMpMessageTplDo) Find() ([]*model.TMpMessageTpl, error) {
	result, err := t.DO.Find()
	return result.([]*model.TMpMessageTpl), err
}

func (t tMpMessageTplDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TMpMessageTpl, err error) {
	buf := make([]*model.TMpMessageTpl, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tMpMessageTplDo) FindInBatches(result *[]*model.TMpMessageTpl, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tMpMessageTplDo) Attrs(attrs ...field.AssignExpr) *tMpMessageTplDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tMpMessageTplDo) Assign(attrs ...field.AssignExpr) *tMpMessageTplDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tMpMessageTplDo) Joins(field field.RelationField) *tMpMessageTplDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tMpMessageTplDo) Preload(field field.RelationField) *tMpMessageTplDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tMpMessageTplDo) FirstOrInit() (*model.TMpMessageTpl, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTpl), nil
	}
}

func (t tMpMessageTplDo) FirstOrCreate() (*model.TMpMessageTpl, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTpl), nil
	}
}

func (t tMpMessageTplDo) FindByPage(offset int, limit int) (result []*model.TMpMessageTpl, count int64, err error) {
	if limit <= 0 {
		count, err = t.Count()
		return
	}

	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tMpMessageTplDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tMpMessageTplDo) withDO(do gen.Dao) *tMpMessageTplDo {
	t.DO = *do.(*gen.DO)
	return t
}
