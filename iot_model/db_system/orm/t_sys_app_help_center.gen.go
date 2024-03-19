// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_system/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTSysAppHelpCenter(db *gorm.DB) tSysAppHelpCenter {
	_tSysAppHelpCenter := tSysAppHelpCenter{}

	_tSysAppHelpCenter.tSysAppHelpCenterDo.UseDB(db)
	_tSysAppHelpCenter.tSysAppHelpCenterDo.UseModel(&model.TSysAppHelpCenter{})

	tableName := _tSysAppHelpCenter.tSysAppHelpCenterDo.TableName()
	_tSysAppHelpCenter.ALL = field.NewField(tableName, "*")
	_tSysAppHelpCenter.Id = field.NewInt64(tableName, "id")
	_tSysAppHelpCenter.Name = field.NewString(tableName, "name")
	_tSysAppHelpCenter.TemplateType = field.NewInt32(tableName, "template_type")
	_tSysAppHelpCenter.Version = field.NewString(tableName, "version")
	_tSysAppHelpCenter.CreatedBy = field.NewInt64(tableName, "created_by")
	_tSysAppHelpCenter.CreatedAt = field.NewTime(tableName, "created_at")
	_tSysAppHelpCenter.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tSysAppHelpCenter.UpdatedAt = field.NewTime(tableName, "updated_at")

	_tSysAppHelpCenter.fillFieldMap()

	return _tSysAppHelpCenter
}

type tSysAppHelpCenter struct {
	tSysAppHelpCenterDo tSysAppHelpCenterDo

	ALL          field.Field
	Id           field.Int64
	Name         field.String
	TemplateType field.Int32
	Version      field.String
	CreatedBy    field.Int64
	CreatedAt    field.Time
	UpdatedBy    field.Int64
	UpdatedAt    field.Time

	fieldMap map[string]field.Expr
}

func (t tSysAppHelpCenter) Table(newTableName string) *tSysAppHelpCenter {
	t.tSysAppHelpCenterDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tSysAppHelpCenter) As(alias string) *tSysAppHelpCenter {
	t.tSysAppHelpCenterDo.DO = *(t.tSysAppHelpCenterDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tSysAppHelpCenter) updateTableName(table string) *tSysAppHelpCenter {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.Name = field.NewString(table, "name")
	t.TemplateType = field.NewInt32(table, "template_type")
	t.Version = field.NewString(table, "version")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.UpdatedAt = field.NewTime(table, "updated_at")

	t.fillFieldMap()

	return t
}

func (t *tSysAppHelpCenter) WithContext(ctx context.Context) *tSysAppHelpCenterDo {
	return t.tSysAppHelpCenterDo.WithContext(ctx)
}

func (t tSysAppHelpCenter) TableName() string { return t.tSysAppHelpCenterDo.TableName() }

func (t tSysAppHelpCenter) Alias() string { return t.tSysAppHelpCenterDo.Alias() }

func (t *tSysAppHelpCenter) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tSysAppHelpCenter) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["id"] = t.Id
	t.fieldMap["name"] = t.Name
	t.fieldMap["template_type"] = t.TemplateType
	t.fieldMap["version"] = t.Version
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["updated_at"] = t.UpdatedAt
}

func (t tSysAppHelpCenter) clone(db *gorm.DB) tSysAppHelpCenter {
	t.tSysAppHelpCenterDo.ReplaceDB(db)
	return t
}

type tSysAppHelpCenterDo struct{ gen.DO }

func (t tSysAppHelpCenterDo) Debug() *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Debug())
}

func (t tSysAppHelpCenterDo) WithContext(ctx context.Context) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tSysAppHelpCenterDo) Clauses(conds ...clause.Expression) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tSysAppHelpCenterDo) Returning(value interface{}, columns ...string) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tSysAppHelpCenterDo) Not(conds ...gen.Condition) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tSysAppHelpCenterDo) Or(conds ...gen.Condition) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tSysAppHelpCenterDo) Select(conds ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tSysAppHelpCenterDo) Where(conds ...gen.Condition) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tSysAppHelpCenterDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tSysAppHelpCenterDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tSysAppHelpCenterDo) Order(conds ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tSysAppHelpCenterDo) Distinct(cols ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tSysAppHelpCenterDo) Omit(cols ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tSysAppHelpCenterDo) Join(table schema.Tabler, on ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tSysAppHelpCenterDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tSysAppHelpCenterDo) RightJoin(table schema.Tabler, on ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tSysAppHelpCenterDo) Group(cols ...field.Expr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tSysAppHelpCenterDo) Having(conds ...gen.Condition) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tSysAppHelpCenterDo) Limit(limit int) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tSysAppHelpCenterDo) Offset(offset int) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tSysAppHelpCenterDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tSysAppHelpCenterDo) Unscoped() *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tSysAppHelpCenterDo) Create(values ...*model.TSysAppHelpCenter) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tSysAppHelpCenterDo) CreateInBatches(values []*model.TSysAppHelpCenter, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tSysAppHelpCenterDo) Save(values ...*model.TSysAppHelpCenter) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tSysAppHelpCenterDo) First() (*model.TSysAppHelpCenter, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAppHelpCenter), nil
	}
}

func (t tSysAppHelpCenterDo) Take() (*model.TSysAppHelpCenter, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAppHelpCenter), nil
	}
}

func (t tSysAppHelpCenterDo) Last() (*model.TSysAppHelpCenter, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAppHelpCenter), nil
	}
}

func (t tSysAppHelpCenterDo) Find() ([]*model.TSysAppHelpCenter, error) {
	result, err := t.DO.Find()
	return result.([]*model.TSysAppHelpCenter), err
}

func (t tSysAppHelpCenterDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSysAppHelpCenter, err error) {
	buf := make([]*model.TSysAppHelpCenter, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tSysAppHelpCenterDo) FindInBatches(result *[]*model.TSysAppHelpCenter, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tSysAppHelpCenterDo) Attrs(attrs ...field.AssignExpr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tSysAppHelpCenterDo) Assign(attrs ...field.AssignExpr) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tSysAppHelpCenterDo) Joins(field field.RelationField) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tSysAppHelpCenterDo) Preload(field field.RelationField) *tSysAppHelpCenterDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tSysAppHelpCenterDo) FirstOrInit() (*model.TSysAppHelpCenter, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAppHelpCenter), nil
	}
}

func (t tSysAppHelpCenterDo) FirstOrCreate() (*model.TSysAppHelpCenter, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAppHelpCenter), nil
	}
}

func (t tSysAppHelpCenterDo) FindByPage(offset int, limit int) (result []*model.TSysAppHelpCenter, count int64, err error) {
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

func (t tSysAppHelpCenterDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tSysAppHelpCenterDo) withDO(do gen.Dao) *tSysAppHelpCenterDo {
	t.DO = *do.(*gen.DO)
	return t
}
