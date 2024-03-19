// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_statistics/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTAppUserActive30day(db *gorm.DB) tAppUserActive30day {
	_tAppUserActive30day := tAppUserActive30day{}

	_tAppUserActive30day.tAppUserActive30dayDo.UseDB(db)
	_tAppUserActive30day.tAppUserActive30dayDo.UseModel(&model.TAppUserActive30day{})

	tableName := _tAppUserActive30day.tAppUserActive30dayDo.TableName()
	_tAppUserActive30day.ALL = field.NewField(tableName, "*")
	_tAppUserActive30day.DataTime = field.NewTime(tableName, "data_time")
	_tAppUserActive30day.TenantId = field.NewString(tableName, "tenant_id")
	_tAppUserActive30day.AppKey = field.NewString(tableName, "app_key")
	_tAppUserActive30day.ActiveSum = field.NewInt64(tableName, "active_sum")
	_tAppUserActive30day.UpdatedAt = field.NewTime(tableName, "updated_at")

	_tAppUserActive30day.fillFieldMap()

	return _tAppUserActive30day
}

type tAppUserActive30day struct {
	tAppUserActive30dayDo tAppUserActive30dayDo

	ALL       field.Field
	DataTime  field.Time
	TenantId  field.String
	AppKey    field.String
	ActiveSum field.Int64
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (t tAppUserActive30day) Table(newTableName string) *tAppUserActive30day {
	t.tAppUserActive30dayDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tAppUserActive30day) As(alias string) *tAppUserActive30day {
	t.tAppUserActive30dayDo.DO = *(t.tAppUserActive30dayDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tAppUserActive30day) updateTableName(table string) *tAppUserActive30day {
	t.ALL = field.NewField(table, "*")
	t.DataTime = field.NewTime(table, "data_time")
	t.TenantId = field.NewString(table, "tenant_id")
	t.AppKey = field.NewString(table, "app_key")
	t.ActiveSum = field.NewInt64(table, "active_sum")
	t.UpdatedAt = field.NewTime(table, "updated_at")

	t.fillFieldMap()

	return t
}

func (t *tAppUserActive30day) WithContext(ctx context.Context) *tAppUserActive30dayDo {
	return t.tAppUserActive30dayDo.WithContext(ctx)
}

func (t tAppUserActive30day) TableName() string { return t.tAppUserActive30dayDo.TableName() }

func (t tAppUserActive30day) Alias() string { return t.tAppUserActive30dayDo.Alias() }

func (t *tAppUserActive30day) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tAppUserActive30day) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["data_time"] = t.DataTime
	t.fieldMap["tenant_id"] = t.TenantId
	t.fieldMap["app_key"] = t.AppKey
	t.fieldMap["active_sum"] = t.ActiveSum
	t.fieldMap["updated_at"] = t.UpdatedAt
}

func (t tAppUserActive30day) clone(db *gorm.DB) tAppUserActive30day {
	t.tAppUserActive30dayDo.ReplaceDB(db)
	return t
}

type tAppUserActive30dayDo struct{ gen.DO }

func (t tAppUserActive30dayDo) Debug() *tAppUserActive30dayDo {
	return t.withDO(t.DO.Debug())
}

func (t tAppUserActive30dayDo) WithContext(ctx context.Context) *tAppUserActive30dayDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tAppUserActive30dayDo) Clauses(conds ...clause.Expression) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tAppUserActive30dayDo) Returning(value interface{}, columns ...string) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tAppUserActive30dayDo) Not(conds ...gen.Condition) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tAppUserActive30dayDo) Or(conds ...gen.Condition) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tAppUserActive30dayDo) Select(conds ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tAppUserActive30dayDo) Where(conds ...gen.Condition) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tAppUserActive30dayDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tAppUserActive30dayDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tAppUserActive30dayDo) Order(conds ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tAppUserActive30dayDo) Distinct(cols ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tAppUserActive30dayDo) Omit(cols ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tAppUserActive30dayDo) Join(table schema.Tabler, on ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tAppUserActive30dayDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tAppUserActive30dayDo) RightJoin(table schema.Tabler, on ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tAppUserActive30dayDo) Group(cols ...field.Expr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tAppUserActive30dayDo) Having(conds ...gen.Condition) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tAppUserActive30dayDo) Limit(limit int) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tAppUserActive30dayDo) Offset(offset int) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tAppUserActive30dayDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tAppUserActive30dayDo) Unscoped() *tAppUserActive30dayDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tAppUserActive30dayDo) Create(values ...*model.TAppUserActive30day) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tAppUserActive30dayDo) CreateInBatches(values []*model.TAppUserActive30day, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tAppUserActive30dayDo) Save(values ...*model.TAppUserActive30day) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tAppUserActive30dayDo) First() (*model.TAppUserActive30day, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAppUserActive30day), nil
	}
}

func (t tAppUserActive30dayDo) Take() (*model.TAppUserActive30day, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAppUserActive30day), nil
	}
}

func (t tAppUserActive30dayDo) Last() (*model.TAppUserActive30day, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAppUserActive30day), nil
	}
}

func (t tAppUserActive30dayDo) Find() ([]*model.TAppUserActive30day, error) {
	result, err := t.DO.Find()
	return result.([]*model.TAppUserActive30day), err
}

func (t tAppUserActive30dayDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAppUserActive30day, err error) {
	buf := make([]*model.TAppUserActive30day, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tAppUserActive30dayDo) FindInBatches(result *[]*model.TAppUserActive30day, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tAppUserActive30dayDo) Attrs(attrs ...field.AssignExpr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tAppUserActive30dayDo) Assign(attrs ...field.AssignExpr) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tAppUserActive30dayDo) Joins(field field.RelationField) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tAppUserActive30dayDo) Preload(field field.RelationField) *tAppUserActive30dayDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tAppUserActive30dayDo) FirstOrInit() (*model.TAppUserActive30day, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAppUserActive30day), nil
	}
}

func (t tAppUserActive30dayDo) FirstOrCreate() (*model.TAppUserActive30day, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAppUserActive30day), nil
	}
}

func (t tAppUserActive30dayDo) FindByPage(offset int, limit int) (result []*model.TAppUserActive30day, count int64, err error) {
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

func (t tAppUserActive30dayDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tAppUserActive30dayDo) withDO(do gen.Dao) *tAppUserActive30dayDo {
	t.DO = *do.(*gen.DO)
	return t
}
