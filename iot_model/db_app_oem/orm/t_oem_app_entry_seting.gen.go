// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_app_oem/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTOemAppEntrySeting(db *gorm.DB) tOemAppEntrySeting {
	_tOemAppEntrySeting := tOemAppEntrySeting{}

	_tOemAppEntrySeting.tOemAppEntrySetingDo.UseDB(db)
	_tOemAppEntrySeting.tOemAppEntrySetingDo.UseModel(&model.TOemAppEntrySeting{})

	tableName := _tOemAppEntrySeting.tOemAppEntrySetingDo.TableName()
	_tOemAppEntrySeting.ALL = field.NewField(tableName, "*")
	_tOemAppEntrySeting.Id = field.NewInt64(tableName, "id")
	_tOemAppEntrySeting.DirId = field.NewInt64(tableName, "dir_id")
	_tOemAppEntrySeting.Sort = field.NewInt32(tableName, "sort")
	_tOemAppEntrySeting.IsEnable = field.NewInt32(tableName, "is_enable")
	_tOemAppEntrySeting.IsNormal = field.NewInt32(tableName, "is_normal")

	_tOemAppEntrySeting.fillFieldMap()

	return _tOemAppEntrySeting
}

type tOemAppEntrySeting struct {
	tOemAppEntrySetingDo tOemAppEntrySetingDo

	ALL      field.Field
	Id       field.Int64
	DirId    field.Int64
	Sort     field.Int32
	IsEnable field.Int32
	IsNormal field.Int32

	fieldMap map[string]field.Expr
}

func (t tOemAppEntrySeting) Table(newTableName string) *tOemAppEntrySeting {
	t.tOemAppEntrySetingDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOemAppEntrySeting) As(alias string) *tOemAppEntrySeting {
	t.tOemAppEntrySetingDo.DO = *(t.tOemAppEntrySetingDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOemAppEntrySeting) updateTableName(table string) *tOemAppEntrySeting {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.DirId = field.NewInt64(table, "dir_id")
	t.Sort = field.NewInt32(table, "sort")
	t.IsEnable = field.NewInt32(table, "is_enable")
	t.IsNormal = field.NewInt32(table, "is_normal")

	t.fillFieldMap()

	return t
}

func (t *tOemAppEntrySeting) WithContext(ctx context.Context) *tOemAppEntrySetingDo {
	return t.tOemAppEntrySetingDo.WithContext(ctx)
}

func (t tOemAppEntrySeting) TableName() string { return t.tOemAppEntrySetingDo.TableName() }

func (t *tOemAppEntrySeting) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOemAppEntrySeting) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["id"] = t.Id
	t.fieldMap["dir_id"] = t.DirId
	t.fieldMap["sort"] = t.Sort
	t.fieldMap["is_enable"] = t.IsEnable
	t.fieldMap["is_normal"] = t.IsNormal
}

func (t tOemAppEntrySeting) clone(db *gorm.DB) tOemAppEntrySeting {
	t.tOemAppEntrySetingDo.ReplaceDB(db)
	return t
}

type tOemAppEntrySetingDo struct{ gen.DO }

func (t tOemAppEntrySetingDo) Debug() *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Debug())
}

func (t tOemAppEntrySetingDo) WithContext(ctx context.Context) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOemAppEntrySetingDo) Clauses(conds ...clause.Expression) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOemAppEntrySetingDo) Returning(value interface{}, columns ...string) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOemAppEntrySetingDo) Not(conds ...gen.Condition) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOemAppEntrySetingDo) Or(conds ...gen.Condition) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOemAppEntrySetingDo) Select(conds ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOemAppEntrySetingDo) Where(conds ...gen.Condition) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOemAppEntrySetingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOemAppEntrySetingDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOemAppEntrySetingDo) Order(conds ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOemAppEntrySetingDo) Distinct(cols ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOemAppEntrySetingDo) Omit(cols ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOemAppEntrySetingDo) Join(table schema.Tabler, on ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOemAppEntrySetingDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOemAppEntrySetingDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOemAppEntrySetingDo) Group(cols ...field.Expr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOemAppEntrySetingDo) Having(conds ...gen.Condition) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOemAppEntrySetingDo) Limit(limit int) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOemAppEntrySetingDo) Offset(offset int) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOemAppEntrySetingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOemAppEntrySetingDo) Unscoped() *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOemAppEntrySetingDo) Create(values ...*model.TOemAppEntrySeting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOemAppEntrySetingDo) CreateInBatches(values []*model.TOemAppEntrySeting, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOemAppEntrySetingDo) Save(values ...*model.TOemAppEntrySeting) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOemAppEntrySetingDo) First() (*model.TOemAppEntrySeting, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppEntrySeting), nil
	}
}

func (t tOemAppEntrySetingDo) Take() (*model.TOemAppEntrySeting, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppEntrySeting), nil
	}
}

func (t tOemAppEntrySetingDo) Last() (*model.TOemAppEntrySeting, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppEntrySeting), nil
	}
}

func (t tOemAppEntrySetingDo) Find() ([]*model.TOemAppEntrySeting, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOemAppEntrySeting), err
}

func (t tOemAppEntrySetingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOemAppEntrySeting, err error) {
	buf := make([]*model.TOemAppEntrySeting, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOemAppEntrySetingDo) FindInBatches(result *[]*model.TOemAppEntrySeting, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOemAppEntrySetingDo) Attrs(attrs ...field.AssignExpr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOemAppEntrySetingDo) Assign(attrs ...field.AssignExpr) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOemAppEntrySetingDo) Joins(field field.RelationField) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOemAppEntrySetingDo) Preload(field field.RelationField) *tOemAppEntrySetingDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOemAppEntrySetingDo) FirstOrInit() (*model.TOemAppEntrySeting, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppEntrySeting), nil
	}
}

func (t tOemAppEntrySetingDo) FirstOrCreate() (*model.TOemAppEntrySeting, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppEntrySeting), nil
	}
}

func (t tOemAppEntrySetingDo) FindByPage(offset int, limit int) (result []*model.TOemAppEntrySeting, count int64, err error) {
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

func (t tOemAppEntrySetingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOemAppEntrySetingDo) withDO(do gen.Dao) *tOemAppEntrySetingDo {
	t.DO = *do.(*gen.DO)
	return t
}
