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

func newTMpMessageTemplate(db *gorm.DB) tMpMessageTemplate {
	_tMpMessageTemplate := tMpMessageTemplate{}

	_tMpMessageTemplate.tMpMessageTemplateDo.UseDB(db)
	_tMpMessageTemplate.tMpMessageTemplateDo.UseModel(&model.TMpMessageTemplate{})

	tableName := _tMpMessageTemplate.tMpMessageTemplateDo.TableName()
	_tMpMessageTemplate.ALL = field.NewField(tableName, "*")
	_tMpMessageTemplate.Id = field.NewInt64(tableName, "id")
	_tMpMessageTemplate.TplCode = field.NewString(tableName, "tpl_code")
	_tMpMessageTemplate.TplName = field.NewString(tableName, "tpl_name")
	_tMpMessageTemplate.TplContent = field.NewString(tableName, "tpl_content")
	_tMpMessageTemplate.TplParams = field.NewString(tableName, "tpl_params")
	_tMpMessageTemplate.PushType = field.NewInt32(tableName, "push_type")
	_tMpMessageTemplate.MessageType = field.NewInt32(tableName, "message_type")
	_tMpMessageTemplate.AgentType = field.NewInt32(tableName, "agent_type")
	_tMpMessageTemplate.Lang = field.NewString(tableName, "lang")
	_tMpMessageTemplate.ExpireHour = field.NewInt32(tableName, "expire_hour")
	_tMpMessageTemplate.CreatedBy = field.NewInt64(tableName, "created_by")
	_tMpMessageTemplate.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tMpMessageTemplate.CreatedAt = field.NewTime(tableName, "created_at")
	_tMpMessageTemplate.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tMpMessageTemplate.DeletedAt = field.NewField(tableName, "deleted_at")

	_tMpMessageTemplate.fillFieldMap()

	return _tMpMessageTemplate
}

type tMpMessageTemplate struct {
	tMpMessageTemplateDo tMpMessageTemplateDo

	ALL         field.Field
	Id          field.Int64
	TplCode     field.String
	TplName     field.String
	TplContent  field.String
	TplParams   field.String
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

func (t tMpMessageTemplate) Table(newTableName string) *tMpMessageTemplate {
	t.tMpMessageTemplateDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tMpMessageTemplate) As(alias string) *tMpMessageTemplate {
	t.tMpMessageTemplateDo.DO = *(t.tMpMessageTemplateDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tMpMessageTemplate) updateTableName(table string) *tMpMessageTemplate {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.TplCode = field.NewString(table, "tpl_code")
	t.TplName = field.NewString(table, "tpl_name")
	t.TplContent = field.NewString(table, "tpl_content")
	t.TplParams = field.NewString(table, "tpl_params")
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

func (t *tMpMessageTemplate) WithContext(ctx context.Context) *tMpMessageTemplateDo {
	return t.tMpMessageTemplateDo.WithContext(ctx)
}

func (t tMpMessageTemplate) TableName() string { return t.tMpMessageTemplateDo.TableName() }

func (t tMpMessageTemplate) Alias() string { return t.tMpMessageTemplateDo.Alias() }

func (t *tMpMessageTemplate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tMpMessageTemplate) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 15)
	t.fieldMap["id"] = t.Id
	t.fieldMap["tpl_code"] = t.TplCode
	t.fieldMap["tpl_name"] = t.TplName
	t.fieldMap["tpl_content"] = t.TplContent
	t.fieldMap["tpl_params"] = t.TplParams
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

func (t tMpMessageTemplate) clone(db *gorm.DB) tMpMessageTemplate {
	t.tMpMessageTemplateDo.ReplaceDB(db)
	return t
}

type tMpMessageTemplateDo struct{ gen.DO }

func (t tMpMessageTemplateDo) Debug() *tMpMessageTemplateDo {
	return t.withDO(t.DO.Debug())
}

func (t tMpMessageTemplateDo) WithContext(ctx context.Context) *tMpMessageTemplateDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tMpMessageTemplateDo) Clauses(conds ...clause.Expression) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tMpMessageTemplateDo) Returning(value interface{}, columns ...string) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tMpMessageTemplateDo) Not(conds ...gen.Condition) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tMpMessageTemplateDo) Or(conds ...gen.Condition) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tMpMessageTemplateDo) Select(conds ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tMpMessageTemplateDo) Where(conds ...gen.Condition) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tMpMessageTemplateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tMpMessageTemplateDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tMpMessageTemplateDo) Order(conds ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tMpMessageTemplateDo) Distinct(cols ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tMpMessageTemplateDo) Omit(cols ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tMpMessageTemplateDo) Join(table schema.Tabler, on ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tMpMessageTemplateDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tMpMessageTemplateDo) RightJoin(table schema.Tabler, on ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tMpMessageTemplateDo) Group(cols ...field.Expr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tMpMessageTemplateDo) Having(conds ...gen.Condition) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tMpMessageTemplateDo) Limit(limit int) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tMpMessageTemplateDo) Offset(offset int) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tMpMessageTemplateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tMpMessageTemplateDo) Unscoped() *tMpMessageTemplateDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tMpMessageTemplateDo) Create(values ...*model.TMpMessageTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tMpMessageTemplateDo) CreateInBatches(values []*model.TMpMessageTemplate, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tMpMessageTemplateDo) Save(values ...*model.TMpMessageTemplate) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tMpMessageTemplateDo) First() (*model.TMpMessageTemplate, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTemplate), nil
	}
}

func (t tMpMessageTemplateDo) Take() (*model.TMpMessageTemplate, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTemplate), nil
	}
}

func (t tMpMessageTemplateDo) Last() (*model.TMpMessageTemplate, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTemplate), nil
	}
}

func (t tMpMessageTemplateDo) Find() ([]*model.TMpMessageTemplate, error) {
	result, err := t.DO.Find()
	return result.([]*model.TMpMessageTemplate), err
}

func (t tMpMessageTemplateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TMpMessageTemplate, err error) {
	buf := make([]*model.TMpMessageTemplate, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tMpMessageTemplateDo) FindInBatches(result *[]*model.TMpMessageTemplate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tMpMessageTemplateDo) Attrs(attrs ...field.AssignExpr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tMpMessageTemplateDo) Assign(attrs ...field.AssignExpr) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tMpMessageTemplateDo) Joins(field field.RelationField) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tMpMessageTemplateDo) Preload(field field.RelationField) *tMpMessageTemplateDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tMpMessageTemplateDo) FirstOrInit() (*model.TMpMessageTemplate, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTemplate), nil
	}
}

func (t tMpMessageTemplateDo) FirstOrCreate() (*model.TMpMessageTemplate, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TMpMessageTemplate), nil
	}
}

func (t tMpMessageTemplateDo) FindByPage(offset int, limit int) (result []*model.TMpMessageTemplate, count int64, err error) {
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

func (t tMpMessageTemplateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tMpMessageTemplateDo) withDO(do gen.Dao) *tMpMessageTemplateDo {
	t.DO = *do.(*gen.DO)
	return t
}
