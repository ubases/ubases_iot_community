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

func newTOpmProductAppRelation(db *gorm.DB) tOpmProductAppRelation {
	_tOpmProductAppRelation := tOpmProductAppRelation{}

	_tOpmProductAppRelation.tOpmProductAppRelationDo.UseDB(db)
	_tOpmProductAppRelation.tOpmProductAppRelationDo.UseModel(&model.TOpmProductAppRelation{})

	tableName := _tOpmProductAppRelation.tOpmProductAppRelationDo.TableName()
	_tOpmProductAppRelation.ALL = field.NewField(tableName, "*")
	_tOpmProductAppRelation.Id = field.NewInt64(tableName, "id")
	_tOpmProductAppRelation.ProductId = field.NewInt64(tableName, "product_id")
	_tOpmProductAppRelation.ProductKey = field.NewString(tableName, "product_key")
	_tOpmProductAppRelation.AppKey = field.NewString(tableName, "app_key")
	_tOpmProductAppRelation.AppName = field.NewString(tableName, "app_name")
	_tOpmProductAppRelation.CreatedBy = field.NewInt64(tableName, "created_by")
	_tOpmProductAppRelation.CreatedAt = field.NewTime(tableName, "created_at")

	_tOpmProductAppRelation.fillFieldMap()

	return _tOpmProductAppRelation
}

type tOpmProductAppRelation struct {
	tOpmProductAppRelationDo tOpmProductAppRelationDo

	ALL        field.Field
	Id         field.Int64
	ProductId  field.Int64
	ProductKey field.String
	AppKey     field.String
	AppName    field.String
	CreatedBy  field.Int64
	CreatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (t tOpmProductAppRelation) Table(newTableName string) *tOpmProductAppRelation {
	t.tOpmProductAppRelationDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOpmProductAppRelation) As(alias string) *tOpmProductAppRelation {
	t.tOpmProductAppRelationDo.DO = *(t.tOpmProductAppRelationDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOpmProductAppRelation) updateTableName(table string) *tOpmProductAppRelation {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.ProductId = field.NewInt64(table, "product_id")
	t.ProductKey = field.NewString(table, "product_key")
	t.AppKey = field.NewString(table, "app_key")
	t.AppName = field.NewString(table, "app_name")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.CreatedAt = field.NewTime(table, "created_at")

	t.fillFieldMap()

	return t
}

func (t *tOpmProductAppRelation) WithContext(ctx context.Context) *tOpmProductAppRelationDo {
	return t.tOpmProductAppRelationDo.WithContext(ctx)
}

func (t tOpmProductAppRelation) TableName() string { return t.tOpmProductAppRelationDo.TableName() }

func (t tOpmProductAppRelation) Alias() string { return t.tOpmProductAppRelationDo.Alias() }

func (t *tOpmProductAppRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOpmProductAppRelation) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.Id
	t.fieldMap["product_id"] = t.ProductId
	t.fieldMap["product_key"] = t.ProductKey
	t.fieldMap["app_key"] = t.AppKey
	t.fieldMap["app_name"] = t.AppName
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["created_at"] = t.CreatedAt
}

func (t tOpmProductAppRelation) clone(db *gorm.DB) tOpmProductAppRelation {
	t.tOpmProductAppRelationDo.ReplaceDB(db)
	return t
}

type tOpmProductAppRelationDo struct{ gen.DO }

func (t tOpmProductAppRelationDo) Debug() *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Debug())
}

func (t tOpmProductAppRelationDo) WithContext(ctx context.Context) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOpmProductAppRelationDo) Clauses(conds ...clause.Expression) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOpmProductAppRelationDo) Returning(value interface{}, columns ...string) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOpmProductAppRelationDo) Not(conds ...gen.Condition) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOpmProductAppRelationDo) Or(conds ...gen.Condition) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOpmProductAppRelationDo) Select(conds ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOpmProductAppRelationDo) Where(conds ...gen.Condition) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOpmProductAppRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOpmProductAppRelationDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOpmProductAppRelationDo) Order(conds ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOpmProductAppRelationDo) Distinct(cols ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOpmProductAppRelationDo) Omit(cols ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOpmProductAppRelationDo) Join(table schema.Tabler, on ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOpmProductAppRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOpmProductAppRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOpmProductAppRelationDo) Group(cols ...field.Expr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOpmProductAppRelationDo) Having(conds ...gen.Condition) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOpmProductAppRelationDo) Limit(limit int) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOpmProductAppRelationDo) Offset(offset int) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOpmProductAppRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOpmProductAppRelationDo) Unscoped() *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOpmProductAppRelationDo) Create(values ...*model.TOpmProductAppRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOpmProductAppRelationDo) CreateInBatches(values []*model.TOpmProductAppRelation, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOpmProductAppRelationDo) Save(values ...*model.TOpmProductAppRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOpmProductAppRelationDo) First() (*model.TOpmProductAppRelation, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductAppRelation), nil
	}
}

func (t tOpmProductAppRelationDo) Take() (*model.TOpmProductAppRelation, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductAppRelation), nil
	}
}

func (t tOpmProductAppRelationDo) Last() (*model.TOpmProductAppRelation, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductAppRelation), nil
	}
}

func (t tOpmProductAppRelationDo) Find() ([]*model.TOpmProductAppRelation, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOpmProductAppRelation), err
}

func (t tOpmProductAppRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOpmProductAppRelation, err error) {
	buf := make([]*model.TOpmProductAppRelation, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOpmProductAppRelationDo) FindInBatches(result *[]*model.TOpmProductAppRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOpmProductAppRelationDo) Attrs(attrs ...field.AssignExpr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOpmProductAppRelationDo) Assign(attrs ...field.AssignExpr) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOpmProductAppRelationDo) Joins(field field.RelationField) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOpmProductAppRelationDo) Preload(field field.RelationField) *tOpmProductAppRelationDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOpmProductAppRelationDo) FirstOrInit() (*model.TOpmProductAppRelation, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductAppRelation), nil
	}
}

func (t tOpmProductAppRelationDo) FirstOrCreate() (*model.TOpmProductAppRelation, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductAppRelation), nil
	}
}

func (t tOpmProductAppRelationDo) FindByPage(offset int, limit int) (result []*model.TOpmProductAppRelation, count int64, err error) {
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

func (t tOpmProductAppRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOpmProductAppRelationDo) withDO(do gen.Dao) *tOpmProductAppRelationDo {
	t.DO = *do.(*gen.DO)
	return t
}