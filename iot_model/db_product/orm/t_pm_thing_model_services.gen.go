// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_product/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTPmThingModelServices(db *gorm.DB) tPmThingModelServices {
	_tPmThingModelServices := tPmThingModelServices{}

	_tPmThingModelServices.tPmThingModelServicesDo.UseDB(db)
	_tPmThingModelServices.tPmThingModelServicesDo.UseModel(&model.TPmThingModelServices{})

	tableName := _tPmThingModelServices.tPmThingModelServicesDo.TableName()
	_tPmThingModelServices.ALL = field.NewField(tableName, "*")
	_tPmThingModelServices.Id = field.NewInt64(tableName, "id")
	_tPmThingModelServices.ModelId = field.NewInt64(tableName, "model_id")
	_tPmThingModelServices.ProductKey = field.NewString(tableName, "product_key")
	_tPmThingModelServices.Identifier = field.NewString(tableName, "identifier")
	_tPmThingModelServices.ServiceName = field.NewString(tableName, "service_name")
	_tPmThingModelServices.InputParams = field.NewString(tableName, "input_params")
	_tPmThingModelServices.OutputParams = field.NewString(tableName, "output_params")
	_tPmThingModelServices.Required = field.NewInt32(tableName, "required")
	_tPmThingModelServices.CallType = field.NewInt32(tableName, "call_type")
	_tPmThingModelServices.Custom = field.NewInt32(tableName, "custom")
	_tPmThingModelServices.Extension = field.NewString(tableName, "extension")
	_tPmThingModelServices.Desc = field.NewString(tableName, "desc")
	_tPmThingModelServices.TriggerCond = field.NewInt32(tableName, "trigger_cond")
	_tPmThingModelServices.ExecCond = field.NewInt32(tableName, "exec_cond")
	_tPmThingModelServices.CreatedBy = field.NewInt64(tableName, "created_by")
	_tPmThingModelServices.CreatedAt = field.NewTime(tableName, "created_at")
	_tPmThingModelServices.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tPmThingModelServices.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tPmThingModelServices.DeletedAt = field.NewField(tableName, "deleted_at")
	_tPmThingModelServices.StdId = field.NewInt64(tableName, "std_id")
	_tPmThingModelServices.Valid = field.NewInt32(tableName, "valid")
	_tPmThingModelServices.Dpid = field.NewInt32(tableName, "dpid")

	_tPmThingModelServices.fillFieldMap()

	return _tPmThingModelServices
}

type tPmThingModelServices struct {
	tPmThingModelServicesDo tPmThingModelServicesDo

	ALL          field.Field
	Id           field.Int64
	ModelId      field.Int64
	ProductKey   field.String
	Identifier   field.String
	ServiceName  field.String
	InputParams  field.String
	OutputParams field.String
	Required     field.Int32
	CallType     field.Int32
	Custom       field.Int32
	Extension    field.String
	Desc         field.String
	TriggerCond  field.Int32
	ExecCond     field.Int32
	CreatedBy    field.Int64
	CreatedAt    field.Time
	UpdatedBy    field.Int64
	UpdatedAt    field.Time
	DeletedAt    field.Field
	StdId        field.Int64
	Valid        field.Int32
	Dpid         field.Int32

	fieldMap map[string]field.Expr
}

func (t tPmThingModelServices) Table(newTableName string) *tPmThingModelServices {
	t.tPmThingModelServicesDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tPmThingModelServices) As(alias string) *tPmThingModelServices {
	t.tPmThingModelServicesDo.DO = *(t.tPmThingModelServicesDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tPmThingModelServices) updateTableName(table string) *tPmThingModelServices {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.ModelId = field.NewInt64(table, "model_id")
	t.ProductKey = field.NewString(table, "product_key")
	t.Identifier = field.NewString(table, "identifier")
	t.ServiceName = field.NewString(table, "service_name")
	t.InputParams = field.NewString(table, "input_params")
	t.OutputParams = field.NewString(table, "output_params")
	t.Required = field.NewInt32(table, "required")
	t.CallType = field.NewInt32(table, "call_type")
	t.Custom = field.NewInt32(table, "custom")
	t.Extension = field.NewString(table, "extension")
	t.Desc = field.NewString(table, "desc")
	t.TriggerCond = field.NewInt32(table, "trigger_cond")
	t.ExecCond = field.NewInt32(table, "exec_cond")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.StdId = field.NewInt64(table, "std_id")
	t.Valid = field.NewInt32(table, "valid")
	t.Dpid = field.NewInt32(table, "dpid")

	t.fillFieldMap()

	return t
}

func (t *tPmThingModelServices) WithContext(ctx context.Context) *tPmThingModelServicesDo {
	return t.tPmThingModelServicesDo.WithContext(ctx)
}

func (t tPmThingModelServices) TableName() string { return t.tPmThingModelServicesDo.TableName() }

func (t tPmThingModelServices) Alias() string { return t.tPmThingModelServicesDo.Alias() }

func (t *tPmThingModelServices) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tPmThingModelServices) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 22)
	t.fieldMap["id"] = t.Id
	t.fieldMap["model_id"] = t.ModelId
	t.fieldMap["product_key"] = t.ProductKey
	t.fieldMap["identifier"] = t.Identifier
	t.fieldMap["service_name"] = t.ServiceName
	t.fieldMap["input_params"] = t.InputParams
	t.fieldMap["output_params"] = t.OutputParams
	t.fieldMap["required"] = t.Required
	t.fieldMap["call_type"] = t.CallType
	t.fieldMap["custom"] = t.Custom
	t.fieldMap["extension"] = t.Extension
	t.fieldMap["desc"] = t.Desc
	t.fieldMap["trigger_cond"] = t.TriggerCond
	t.fieldMap["exec_cond"] = t.ExecCond
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["std_id"] = t.StdId
	t.fieldMap["valid"] = t.Valid
	t.fieldMap["dpid"] = t.Dpid
}

func (t tPmThingModelServices) clone(db *gorm.DB) tPmThingModelServices {
	t.tPmThingModelServicesDo.ReplaceDB(db)
	return t
}

type tPmThingModelServicesDo struct{ gen.DO }

func (t tPmThingModelServicesDo) Debug() *tPmThingModelServicesDo {
	return t.withDO(t.DO.Debug())
}

func (t tPmThingModelServicesDo) WithContext(ctx context.Context) *tPmThingModelServicesDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tPmThingModelServicesDo) Clauses(conds ...clause.Expression) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tPmThingModelServicesDo) Returning(value interface{}, columns ...string) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tPmThingModelServicesDo) Not(conds ...gen.Condition) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tPmThingModelServicesDo) Or(conds ...gen.Condition) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tPmThingModelServicesDo) Select(conds ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tPmThingModelServicesDo) Where(conds ...gen.Condition) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tPmThingModelServicesDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tPmThingModelServicesDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tPmThingModelServicesDo) Order(conds ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tPmThingModelServicesDo) Distinct(cols ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tPmThingModelServicesDo) Omit(cols ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tPmThingModelServicesDo) Join(table schema.Tabler, on ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tPmThingModelServicesDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tPmThingModelServicesDo) RightJoin(table schema.Tabler, on ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tPmThingModelServicesDo) Group(cols ...field.Expr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tPmThingModelServicesDo) Having(conds ...gen.Condition) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tPmThingModelServicesDo) Limit(limit int) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tPmThingModelServicesDo) Offset(offset int) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tPmThingModelServicesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tPmThingModelServicesDo) Unscoped() *tPmThingModelServicesDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tPmThingModelServicesDo) Create(values ...*model.TPmThingModelServices) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tPmThingModelServicesDo) CreateInBatches(values []*model.TPmThingModelServices, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tPmThingModelServicesDo) Save(values ...*model.TPmThingModelServices) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tPmThingModelServicesDo) First() (*model.TPmThingModelServices, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmThingModelServices), nil
	}
}

func (t tPmThingModelServicesDo) Take() (*model.TPmThingModelServices, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmThingModelServices), nil
	}
}

func (t tPmThingModelServicesDo) Last() (*model.TPmThingModelServices, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmThingModelServices), nil
	}
}

func (t tPmThingModelServicesDo) Find() ([]*model.TPmThingModelServices, error) {
	result, err := t.DO.Find()
	return result.([]*model.TPmThingModelServices), err
}

func (t tPmThingModelServicesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPmThingModelServices, err error) {
	buf := make([]*model.TPmThingModelServices, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tPmThingModelServicesDo) FindInBatches(result *[]*model.TPmThingModelServices, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tPmThingModelServicesDo) Attrs(attrs ...field.AssignExpr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tPmThingModelServicesDo) Assign(attrs ...field.AssignExpr) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tPmThingModelServicesDo) Joins(field field.RelationField) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tPmThingModelServicesDo) Preload(field field.RelationField) *tPmThingModelServicesDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tPmThingModelServicesDo) FirstOrInit() (*model.TPmThingModelServices, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmThingModelServices), nil
	}
}

func (t tPmThingModelServicesDo) FirstOrCreate() (*model.TPmThingModelServices, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmThingModelServices), nil
	}
}

func (t tPmThingModelServicesDo) FindByPage(offset int, limit int) (result []*model.TPmThingModelServices, count int64, err error) {
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

func (t tPmThingModelServicesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tPmThingModelServicesDo) withDO(do gen.Dao) *tPmThingModelServicesDo {
	t.DO = *do.(*gen.DO)
	return t
}
