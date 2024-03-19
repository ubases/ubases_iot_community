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

func newTOpmThingModelServices(db *gorm.DB) tOpmThingModelServices {
	_tOpmThingModelServices := tOpmThingModelServices{}

	_tOpmThingModelServices.tOpmThingModelServicesDo.UseDB(db)
	_tOpmThingModelServices.tOpmThingModelServicesDo.UseModel(&model.TOpmThingModelServices{})

	tableName := _tOpmThingModelServices.tOpmThingModelServicesDo.TableName()
	_tOpmThingModelServices.ALL = field.NewField(tableName, "*")
	_tOpmThingModelServices.Id = field.NewInt64(tableName, "id")
	_tOpmThingModelServices.ModelId = field.NewInt64(tableName, "model_id")
	_tOpmThingModelServices.ProductId = field.NewString(tableName, "product_id")
	_tOpmThingModelServices.CreateTs = field.NewString(tableName, "create_ts")
	_tOpmThingModelServices.Identifier = field.NewString(tableName, "identifier")
	_tOpmThingModelServices.ServiceName = field.NewString(tableName, "service_name")
	_tOpmThingModelServices.InputParams = field.NewString(tableName, "input_params")
	_tOpmThingModelServices.OutputParams = field.NewString(tableName, "output_params")
	_tOpmThingModelServices.Required = field.NewInt32(tableName, "required")
	_tOpmThingModelServices.CallType = field.NewInt32(tableName, "call_type")
	_tOpmThingModelServices.Custom = field.NewInt32(tableName, "custom")
	_tOpmThingModelServices.Extension = field.NewString(tableName, "extension")
	_tOpmThingModelServices.CreatedBy = field.NewInt64(tableName, "created_by")
	_tOpmThingModelServices.CreatedAt = field.NewTime(tableName, "created_at")
	_tOpmThingModelServices.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tOpmThingModelServices.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tOpmThingModelServices.TenantId = field.NewString(tableName, "tenant_id")
	_tOpmThingModelServices.StdId = field.NewInt64(tableName, "std_id")
	_tOpmThingModelServices.Valid = field.NewInt32(tableName, "valid")
	_tOpmThingModelServices.ProductKey = field.NewString(tableName, "product_key")
	_tOpmThingModelServices.TriggerCond = field.NewInt32(tableName, "trigger_cond")
	_tOpmThingModelServices.ExecCond = field.NewInt32(tableName, "exec_cond")
	_tOpmThingModelServices.Desc = field.NewString(tableName, "desc")
	_tOpmThingModelServices.Dpid = field.NewInt32(tableName, "dpid")

	_tOpmThingModelServices.fillFieldMap()

	return _tOpmThingModelServices
}

type tOpmThingModelServices struct {
	tOpmThingModelServicesDo tOpmThingModelServicesDo

	ALL          field.Field
	Id           field.Int64
	ModelId      field.Int64
	ProductId    field.String
	CreateTs     field.String
	Identifier   field.String
	ServiceName  field.String
	InputParams  field.String
	OutputParams field.String
	Required     field.Int32
	CallType     field.Int32
	Custom       field.Int32
	Extension    field.String
	CreatedBy    field.Int64
	CreatedAt    field.Time
	UpdatedBy    field.Int64
	UpdatedAt    field.Time
	TenantId     field.String
	StdId        field.Int64
	Valid        field.Int32
	ProductKey   field.String
	TriggerCond  field.Int32
	ExecCond     field.Int32
	Desc         field.String
	Dpid         field.Int32

	fieldMap map[string]field.Expr
}

func (t tOpmThingModelServices) Table(newTableName string) *tOpmThingModelServices {
	t.tOpmThingModelServicesDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOpmThingModelServices) As(alias string) *tOpmThingModelServices {
	t.tOpmThingModelServicesDo.DO = *(t.tOpmThingModelServicesDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOpmThingModelServices) updateTableName(table string) *tOpmThingModelServices {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.ModelId = field.NewInt64(table, "model_id")
	t.ProductId = field.NewString(table, "product_id")
	t.CreateTs = field.NewString(table, "create_ts")
	t.Identifier = field.NewString(table, "identifier")
	t.ServiceName = field.NewString(table, "service_name")
	t.InputParams = field.NewString(table, "input_params")
	t.OutputParams = field.NewString(table, "output_params")
	t.Required = field.NewInt32(table, "required")
	t.CallType = field.NewInt32(table, "call_type")
	t.Custom = field.NewInt32(table, "custom")
	t.Extension = field.NewString(table, "extension")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.TenantId = field.NewString(table, "tenant_id")
	t.StdId = field.NewInt64(table, "std_id")
	t.Valid = field.NewInt32(table, "valid")
	t.ProductKey = field.NewString(table, "product_key")
	t.TriggerCond = field.NewInt32(table, "trigger_cond")
	t.ExecCond = field.NewInt32(table, "exec_cond")
	t.Desc = field.NewString(table, "desc")
	t.Dpid = field.NewInt32(table, "dpid")

	t.fillFieldMap()

	return t
}

func (t *tOpmThingModelServices) WithContext(ctx context.Context) *tOpmThingModelServicesDo {
	return t.tOpmThingModelServicesDo.WithContext(ctx)
}

func (t tOpmThingModelServices) TableName() string { return t.tOpmThingModelServicesDo.TableName() }

func (t tOpmThingModelServices) Alias() string { return t.tOpmThingModelServicesDo.Alias() }

func (t *tOpmThingModelServices) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOpmThingModelServices) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 24)
	t.fieldMap["id"] = t.Id
	t.fieldMap["model_id"] = t.ModelId
	t.fieldMap["product_id"] = t.ProductId
	t.fieldMap["create_ts"] = t.CreateTs
	t.fieldMap["identifier"] = t.Identifier
	t.fieldMap["service_name"] = t.ServiceName
	t.fieldMap["input_params"] = t.InputParams
	t.fieldMap["output_params"] = t.OutputParams
	t.fieldMap["required"] = t.Required
	t.fieldMap["call_type"] = t.CallType
	t.fieldMap["custom"] = t.Custom
	t.fieldMap["extension"] = t.Extension
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["tenant_id"] = t.TenantId
	t.fieldMap["std_id"] = t.StdId
	t.fieldMap["valid"] = t.Valid
	t.fieldMap["product_key"] = t.ProductKey
	t.fieldMap["trigger_cond"] = t.TriggerCond
	t.fieldMap["exec_cond"] = t.ExecCond
	t.fieldMap["desc"] = t.Desc
	t.fieldMap["dpid"] = t.Dpid
}

func (t tOpmThingModelServices) clone(db *gorm.DB) tOpmThingModelServices {
	t.tOpmThingModelServicesDo.ReplaceDB(db)
	return t
}

type tOpmThingModelServicesDo struct{ gen.DO }

func (t tOpmThingModelServicesDo) Debug() *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Debug())
}

func (t tOpmThingModelServicesDo) WithContext(ctx context.Context) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOpmThingModelServicesDo) Clauses(conds ...clause.Expression) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOpmThingModelServicesDo) Returning(value interface{}, columns ...string) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOpmThingModelServicesDo) Not(conds ...gen.Condition) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOpmThingModelServicesDo) Or(conds ...gen.Condition) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOpmThingModelServicesDo) Select(conds ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOpmThingModelServicesDo) Where(conds ...gen.Condition) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOpmThingModelServicesDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOpmThingModelServicesDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOpmThingModelServicesDo) Order(conds ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOpmThingModelServicesDo) Distinct(cols ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOpmThingModelServicesDo) Omit(cols ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOpmThingModelServicesDo) Join(table schema.Tabler, on ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOpmThingModelServicesDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOpmThingModelServicesDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOpmThingModelServicesDo) Group(cols ...field.Expr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOpmThingModelServicesDo) Having(conds ...gen.Condition) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOpmThingModelServicesDo) Limit(limit int) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOpmThingModelServicesDo) Offset(offset int) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOpmThingModelServicesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOpmThingModelServicesDo) Unscoped() *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOpmThingModelServicesDo) Create(values ...*model.TOpmThingModelServices) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOpmThingModelServicesDo) CreateInBatches(values []*model.TOpmThingModelServices, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOpmThingModelServicesDo) Save(values ...*model.TOpmThingModelServices) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOpmThingModelServicesDo) First() (*model.TOpmThingModelServices, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmThingModelServices), nil
	}
}

func (t tOpmThingModelServicesDo) Take() (*model.TOpmThingModelServices, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmThingModelServices), nil
	}
}

func (t tOpmThingModelServicesDo) Last() (*model.TOpmThingModelServices, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmThingModelServices), nil
	}
}

func (t tOpmThingModelServicesDo) Find() ([]*model.TOpmThingModelServices, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOpmThingModelServices), err
}

func (t tOpmThingModelServicesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOpmThingModelServices, err error) {
	buf := make([]*model.TOpmThingModelServices, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOpmThingModelServicesDo) FindInBatches(result *[]*model.TOpmThingModelServices, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOpmThingModelServicesDo) Attrs(attrs ...field.AssignExpr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOpmThingModelServicesDo) Assign(attrs ...field.AssignExpr) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOpmThingModelServicesDo) Joins(field field.RelationField) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOpmThingModelServicesDo) Preload(field field.RelationField) *tOpmThingModelServicesDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOpmThingModelServicesDo) FirstOrInit() (*model.TOpmThingModelServices, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmThingModelServices), nil
	}
}

func (t tOpmThingModelServicesDo) FirstOrCreate() (*model.TOpmThingModelServices, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmThingModelServices), nil
	}
}

func (t tOpmThingModelServicesDo) FindByPage(offset int, limit int) (result []*model.TOpmThingModelServices, count int64, err error) {
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

func (t tOpmThingModelServicesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOpmThingModelServicesDo) withDO(do gen.Dao) *tOpmThingModelServicesDo {
	t.DO = *do.(*gen.DO)
	return t
}
