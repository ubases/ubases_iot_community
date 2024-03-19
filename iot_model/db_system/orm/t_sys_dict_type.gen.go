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

func newTSysDictType(db *gorm.DB) tSysDictType {
	_tSysDictType := tSysDictType{}

	_tSysDictType.tSysDictTypeDo.UseDB(db)
	_tSysDictType.tSysDictTypeDo.UseModel(&model.TSysDictType{})

	tableName := _tSysDictType.tSysDictTypeDo.TableName()
	_tSysDictType.ALL = field.NewField(tableName, "*")
	_tSysDictType.DictId = field.NewInt64(tableName, "dict_id")
	_tSysDictType.DictName = field.NewString(tableName, "dict_name")
	_tSysDictType.DictType = field.NewString(tableName, "dict_type")
	_tSysDictType.Status = field.NewInt32(tableName, "status")
	_tSysDictType.CreateBy = field.NewInt32(tableName, "create_by")
	_tSysDictType.UpdateBy = field.NewInt32(tableName, "update_by")
	_tSysDictType.Remark = field.NewString(tableName, "remark")
	_tSysDictType.CreatedAt = field.NewTime(tableName, "created_at")
	_tSysDictType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tSysDictType.DeletedAt = field.NewField(tableName, "deleted_at")

	_tSysDictType.fillFieldMap()

	return _tSysDictType
}

type tSysDictType struct {
	tSysDictTypeDo tSysDictTypeDo

	ALL       field.Field
	DictId    field.Int64
	DictName  field.String
	DictType  field.String
	Status    field.Int32
	CreateBy  field.Int32
	UpdateBy  field.Int32
	Remark    field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (t tSysDictType) Table(newTableName string) *tSysDictType {
	t.tSysDictTypeDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tSysDictType) As(alias string) *tSysDictType {
	t.tSysDictTypeDo.DO = *(t.tSysDictTypeDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tSysDictType) updateTableName(table string) *tSysDictType {
	t.ALL = field.NewField(table, "*")
	t.DictId = field.NewInt64(table, "dict_id")
	t.DictName = field.NewString(table, "dict_name")
	t.DictType = field.NewString(table, "dict_type")
	t.Status = field.NewInt32(table, "status")
	t.CreateBy = field.NewInt32(table, "create_by")
	t.UpdateBy = field.NewInt32(table, "update_by")
	t.Remark = field.NewString(table, "remark")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tSysDictType) WithContext(ctx context.Context) *tSysDictTypeDo {
	return t.tSysDictTypeDo.WithContext(ctx)
}

func (t tSysDictType) TableName() string { return t.tSysDictTypeDo.TableName() }

func (t tSysDictType) Alias() string { return t.tSysDictTypeDo.Alias() }

func (t *tSysDictType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tSysDictType) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["dict_id"] = t.DictId
	t.fieldMap["dict_name"] = t.DictName
	t.fieldMap["dict_type"] = t.DictType
	t.fieldMap["status"] = t.Status
	t.fieldMap["create_by"] = t.CreateBy
	t.fieldMap["update_by"] = t.UpdateBy
	t.fieldMap["remark"] = t.Remark
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t tSysDictType) clone(db *gorm.DB) tSysDictType {
	t.tSysDictTypeDo.ReplaceDB(db)
	return t
}

type tSysDictTypeDo struct{ gen.DO }

func (t tSysDictTypeDo) Debug() *tSysDictTypeDo {
	return t.withDO(t.DO.Debug())
}

func (t tSysDictTypeDo) WithContext(ctx context.Context) *tSysDictTypeDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tSysDictTypeDo) Clauses(conds ...clause.Expression) *tSysDictTypeDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tSysDictTypeDo) Returning(value interface{}, columns ...string) *tSysDictTypeDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tSysDictTypeDo) Not(conds ...gen.Condition) *tSysDictTypeDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tSysDictTypeDo) Or(conds ...gen.Condition) *tSysDictTypeDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tSysDictTypeDo) Select(conds ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tSysDictTypeDo) Where(conds ...gen.Condition) *tSysDictTypeDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tSysDictTypeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tSysDictTypeDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tSysDictTypeDo) Order(conds ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tSysDictTypeDo) Distinct(cols ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tSysDictTypeDo) Omit(cols ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tSysDictTypeDo) Join(table schema.Tabler, on ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tSysDictTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tSysDictTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tSysDictTypeDo) Group(cols ...field.Expr) *tSysDictTypeDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tSysDictTypeDo) Having(conds ...gen.Condition) *tSysDictTypeDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tSysDictTypeDo) Limit(limit int) *tSysDictTypeDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tSysDictTypeDo) Offset(offset int) *tSysDictTypeDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tSysDictTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tSysDictTypeDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tSysDictTypeDo) Unscoped() *tSysDictTypeDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tSysDictTypeDo) Create(values ...*model.TSysDictType) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tSysDictTypeDo) CreateInBatches(values []*model.TSysDictType, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tSysDictTypeDo) Save(values ...*model.TSysDictType) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tSysDictTypeDo) First() (*model.TSysDictType, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysDictType), nil
	}
}

func (t tSysDictTypeDo) Take() (*model.TSysDictType, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysDictType), nil
	}
}

func (t tSysDictTypeDo) Last() (*model.TSysDictType, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysDictType), nil
	}
}

func (t tSysDictTypeDo) Find() ([]*model.TSysDictType, error) {
	result, err := t.DO.Find()
	return result.([]*model.TSysDictType), err
}

func (t tSysDictTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSysDictType, err error) {
	buf := make([]*model.TSysDictType, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tSysDictTypeDo) FindInBatches(result *[]*model.TSysDictType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tSysDictTypeDo) Attrs(attrs ...field.AssignExpr) *tSysDictTypeDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tSysDictTypeDo) Assign(attrs ...field.AssignExpr) *tSysDictTypeDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tSysDictTypeDo) Joins(field field.RelationField) *tSysDictTypeDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tSysDictTypeDo) Preload(field field.RelationField) *tSysDictTypeDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tSysDictTypeDo) FirstOrInit() (*model.TSysDictType, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysDictType), nil
	}
}

func (t tSysDictTypeDo) FirstOrCreate() (*model.TSysDictType, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysDictType), nil
	}
}

func (t tSysDictTypeDo) FindByPage(offset int, limit int) (result []*model.TSysDictType, count int64, err error) {
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

func (t tSysDictTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tSysDictTypeDo) withDO(do gen.Dao) *tSysDictTypeDo {
	t.DO = *do.(*gen.DO)
	return t
}
