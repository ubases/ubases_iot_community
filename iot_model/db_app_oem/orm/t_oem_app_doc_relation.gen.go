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

func newTOemAppDocRelation(db *gorm.DB) tOemAppDocRelation {
	_tOemAppDocRelation := tOemAppDocRelation{}

	_tOemAppDocRelation.tOemAppDocRelationDo.UseDB(db)
	_tOemAppDocRelation.tOemAppDocRelationDo.UseModel(&model.TOemAppDocRelation{})

	tableName := _tOemAppDocRelation.tOemAppDocRelationDo.TableName()
	_tOemAppDocRelation.ALL = field.NewField(tableName, "*")
	_tOemAppDocRelation.Id = field.NewInt64(tableName, "id")
	_tOemAppDocRelation.DocId = field.NewInt64(tableName, "doc_id")
	_tOemAppDocRelation.AppId = field.NewInt64(tableName, "app_id")
	_tOemAppDocRelation.AppKey = field.NewString(tableName, "app_key")

	_tOemAppDocRelation.fillFieldMap()

	return _tOemAppDocRelation
}

type tOemAppDocRelation struct {
	tOemAppDocRelationDo tOemAppDocRelationDo

	ALL    field.Field
	Id     field.Int64
	DocId  field.Int64
	AppId  field.Int64
	AppKey field.String

	fieldMap map[string]field.Expr
}

func (t tOemAppDocRelation) Table(newTableName string) *tOemAppDocRelation {
	t.tOemAppDocRelationDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOemAppDocRelation) As(alias string) *tOemAppDocRelation {
	t.tOemAppDocRelationDo.DO = *(t.tOemAppDocRelationDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOemAppDocRelation) updateTableName(table string) *tOemAppDocRelation {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.DocId = field.NewInt64(table, "doc_id")
	t.AppId = field.NewInt64(table, "app_id")
	t.AppKey = field.NewString(table, "app_key")

	t.fillFieldMap()

	return t
}

func (t *tOemAppDocRelation) WithContext(ctx context.Context) *tOemAppDocRelationDo {
	return t.tOemAppDocRelationDo.WithContext(ctx)
}

func (t tOemAppDocRelation) TableName() string { return t.tOemAppDocRelationDo.TableName() }

func (t *tOemAppDocRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOemAppDocRelation) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["id"] = t.Id
	t.fieldMap["doc_id"] = t.DocId
	t.fieldMap["app_id"] = t.AppId
	t.fieldMap["app_key"] = t.AppKey
}

func (t tOemAppDocRelation) clone(db *gorm.DB) tOemAppDocRelation {
	t.tOemAppDocRelationDo.ReplaceDB(db)
	return t
}

type tOemAppDocRelationDo struct{ gen.DO }

func (t tOemAppDocRelationDo) Debug() *tOemAppDocRelationDo {
	return t.withDO(t.DO.Debug())
}

func (t tOemAppDocRelationDo) WithContext(ctx context.Context) *tOemAppDocRelationDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOemAppDocRelationDo) Clauses(conds ...clause.Expression) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOemAppDocRelationDo) Returning(value interface{}, columns ...string) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOemAppDocRelationDo) Not(conds ...gen.Condition) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOemAppDocRelationDo) Or(conds ...gen.Condition) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOemAppDocRelationDo) Select(conds ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOemAppDocRelationDo) Where(conds ...gen.Condition) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOemAppDocRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOemAppDocRelationDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOemAppDocRelationDo) Order(conds ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOemAppDocRelationDo) Distinct(cols ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOemAppDocRelationDo) Omit(cols ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOemAppDocRelationDo) Join(table schema.Tabler, on ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOemAppDocRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOemAppDocRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOemAppDocRelationDo) Group(cols ...field.Expr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOemAppDocRelationDo) Having(conds ...gen.Condition) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOemAppDocRelationDo) Limit(limit int) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOemAppDocRelationDo) Offset(offset int) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOemAppDocRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOemAppDocRelationDo) Unscoped() *tOemAppDocRelationDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOemAppDocRelationDo) Create(values ...*model.TOemAppDocRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOemAppDocRelationDo) CreateInBatches(values []*model.TOemAppDocRelation, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOemAppDocRelationDo) Save(values ...*model.TOemAppDocRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOemAppDocRelationDo) First() (*model.TOemAppDocRelation, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocRelation), nil
	}
}

func (t tOemAppDocRelationDo) Take() (*model.TOemAppDocRelation, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocRelation), nil
	}
}

func (t tOemAppDocRelationDo) Last() (*model.TOemAppDocRelation, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocRelation), nil
	}
}

func (t tOemAppDocRelationDo) Find() ([]*model.TOemAppDocRelation, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOemAppDocRelation), err
}

func (t tOemAppDocRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOemAppDocRelation, err error) {
	buf := make([]*model.TOemAppDocRelation, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOemAppDocRelationDo) FindInBatches(result *[]*model.TOemAppDocRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOemAppDocRelationDo) Attrs(attrs ...field.AssignExpr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOemAppDocRelationDo) Assign(attrs ...field.AssignExpr) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOemAppDocRelationDo) Joins(field field.RelationField) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOemAppDocRelationDo) Preload(field field.RelationField) *tOemAppDocRelationDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOemAppDocRelationDo) FirstOrInit() (*model.TOemAppDocRelation, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocRelation), nil
	}
}

func (t tOemAppDocRelationDo) FirstOrCreate() (*model.TOemAppDocRelation, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocRelation), nil
	}
}

func (t tOemAppDocRelationDo) FindByPage(offset int, limit int) (result []*model.TOemAppDocRelation, count int64, err error) {
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

func (t tOemAppDocRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOemAppDocRelationDo) withDO(do gen.Dao) *tOemAppDocRelationDo {
	t.DO = *do.(*gen.DO)
	return t
}
