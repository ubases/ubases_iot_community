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

func newTPmProductPanelRelation(db *gorm.DB) tPmProductPanelRelation {
	_tPmProductPanelRelation := tPmProductPanelRelation{}

	_tPmProductPanelRelation.tPmProductPanelRelationDo.UseDB(db)
	_tPmProductPanelRelation.tPmProductPanelRelationDo.UseModel(&model.TPmProductPanelRelation{})

	tableName := _tPmProductPanelRelation.tPmProductPanelRelationDo.TableName()
	_tPmProductPanelRelation.ALL = field.NewField(tableName, "*")
	_tPmProductPanelRelation.Id = field.NewInt64(tableName, "id")
	_tPmProductPanelRelation.ProductId = field.NewInt64(tableName, "product_id")
	_tPmProductPanelRelation.ControlPanelId = field.NewInt64(tableName, "control_panel_id")

	_tPmProductPanelRelation.fillFieldMap()

	return _tPmProductPanelRelation
}

type tPmProductPanelRelation struct {
	tPmProductPanelRelationDo tPmProductPanelRelationDo

	ALL            field.Field
	Id             field.Int64
	ProductId      field.Int64
	ControlPanelId field.Int64

	fieldMap map[string]field.Expr
}

func (t tPmProductPanelRelation) Table(newTableName string) *tPmProductPanelRelation {
	t.tPmProductPanelRelationDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tPmProductPanelRelation) As(alias string) *tPmProductPanelRelation {
	t.tPmProductPanelRelationDo.DO = *(t.tPmProductPanelRelationDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tPmProductPanelRelation) updateTableName(table string) *tPmProductPanelRelation {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.ProductId = field.NewInt64(table, "product_id")
	t.ControlPanelId = field.NewInt64(table, "control_panel_id")

	t.fillFieldMap()

	return t
}

func (t *tPmProductPanelRelation) WithContext(ctx context.Context) *tPmProductPanelRelationDo {
	return t.tPmProductPanelRelationDo.WithContext(ctx)
}

func (t tPmProductPanelRelation) TableName() string { return t.tPmProductPanelRelationDo.TableName() }

func (t tPmProductPanelRelation) Alias() string { return t.tPmProductPanelRelationDo.Alias() }

func (t *tPmProductPanelRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tPmProductPanelRelation) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["id"] = t.Id
	t.fieldMap["product_id"] = t.ProductId
	t.fieldMap["control_panel_id"] = t.ControlPanelId
}

func (t tPmProductPanelRelation) clone(db *gorm.DB) tPmProductPanelRelation {
	t.tPmProductPanelRelationDo.ReplaceDB(db)
	return t
}

type tPmProductPanelRelationDo struct{ gen.DO }

func (t tPmProductPanelRelationDo) Debug() *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Debug())
}

func (t tPmProductPanelRelationDo) WithContext(ctx context.Context) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tPmProductPanelRelationDo) Clauses(conds ...clause.Expression) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tPmProductPanelRelationDo) Returning(value interface{}, columns ...string) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tPmProductPanelRelationDo) Not(conds ...gen.Condition) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tPmProductPanelRelationDo) Or(conds ...gen.Condition) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tPmProductPanelRelationDo) Select(conds ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tPmProductPanelRelationDo) Where(conds ...gen.Condition) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tPmProductPanelRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tPmProductPanelRelationDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tPmProductPanelRelationDo) Order(conds ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tPmProductPanelRelationDo) Distinct(cols ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tPmProductPanelRelationDo) Omit(cols ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tPmProductPanelRelationDo) Join(table schema.Tabler, on ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tPmProductPanelRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tPmProductPanelRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tPmProductPanelRelationDo) Group(cols ...field.Expr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tPmProductPanelRelationDo) Having(conds ...gen.Condition) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tPmProductPanelRelationDo) Limit(limit int) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tPmProductPanelRelationDo) Offset(offset int) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tPmProductPanelRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tPmProductPanelRelationDo) Unscoped() *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tPmProductPanelRelationDo) Create(values ...*model.TPmProductPanelRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tPmProductPanelRelationDo) CreateInBatches(values []*model.TPmProductPanelRelation, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tPmProductPanelRelationDo) Save(values ...*model.TPmProductPanelRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tPmProductPanelRelationDo) First() (*model.TPmProductPanelRelation, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmProductPanelRelation), nil
	}
}

func (t tPmProductPanelRelationDo) Take() (*model.TPmProductPanelRelation, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmProductPanelRelation), nil
	}
}

func (t tPmProductPanelRelationDo) Last() (*model.TPmProductPanelRelation, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmProductPanelRelation), nil
	}
}

func (t tPmProductPanelRelationDo) Find() ([]*model.TPmProductPanelRelation, error) {
	result, err := t.DO.Find()
	return result.([]*model.TPmProductPanelRelation), err
}

func (t tPmProductPanelRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPmProductPanelRelation, err error) {
	buf := make([]*model.TPmProductPanelRelation, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tPmProductPanelRelationDo) FindInBatches(result *[]*model.TPmProductPanelRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tPmProductPanelRelationDo) Attrs(attrs ...field.AssignExpr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tPmProductPanelRelationDo) Assign(attrs ...field.AssignExpr) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tPmProductPanelRelationDo) Joins(field field.RelationField) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tPmProductPanelRelationDo) Preload(field field.RelationField) *tPmProductPanelRelationDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tPmProductPanelRelationDo) FirstOrInit() (*model.TPmProductPanelRelation, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmProductPanelRelation), nil
	}
}

func (t tPmProductPanelRelationDo) FirstOrCreate() (*model.TPmProductPanelRelation, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmProductPanelRelation), nil
	}
}

func (t tPmProductPanelRelationDo) FindByPage(offset int, limit int) (result []*model.TPmProductPanelRelation, count int64, err error) {
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

func (t tPmProductPanelRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tPmProductPanelRelationDo) withDO(do gen.Dao) *tPmProductPanelRelationDo {
	t.DO = *do.(*gen.DO)
	return t
}
