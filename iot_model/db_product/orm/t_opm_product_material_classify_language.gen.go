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

func newTOpmProductMaterialClassifyLanguage(db *gorm.DB) tOpmProductMaterialClassifyLanguage {
	_tOpmProductMaterialClassifyLanguage := tOpmProductMaterialClassifyLanguage{}

	_tOpmProductMaterialClassifyLanguage.tOpmProductMaterialClassifyLanguageDo.UseDB(db)
	_tOpmProductMaterialClassifyLanguage.tOpmProductMaterialClassifyLanguageDo.UseModel(&model.TOpmProductMaterialClassifyLanguage{})

	tableName := _tOpmProductMaterialClassifyLanguage.tOpmProductMaterialClassifyLanguageDo.TableName()
	_tOpmProductMaterialClassifyLanguage.ALL = field.NewField(tableName, "*")
	_tOpmProductMaterialClassifyLanguage.Id = field.NewInt64(tableName, "id")
	_tOpmProductMaterialClassifyLanguage.ClassifyId = field.NewInt64(tableName, "classify_id")
	_tOpmProductMaterialClassifyLanguage.Lang = field.NewString(tableName, "lang")
	_tOpmProductMaterialClassifyLanguage.ClassifyName = field.NewString(tableName, "classify_name")

	_tOpmProductMaterialClassifyLanguage.fillFieldMap()

	return _tOpmProductMaterialClassifyLanguage
}

type tOpmProductMaterialClassifyLanguage struct {
	tOpmProductMaterialClassifyLanguageDo tOpmProductMaterialClassifyLanguageDo

	ALL          field.Field
	Id           field.Int64
	ClassifyId   field.Int64
	Lang         field.String
	ClassifyName field.String

	fieldMap map[string]field.Expr
}

func (t tOpmProductMaterialClassifyLanguage) Table(newTableName string) *tOpmProductMaterialClassifyLanguage {
	t.tOpmProductMaterialClassifyLanguageDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOpmProductMaterialClassifyLanguage) As(alias string) *tOpmProductMaterialClassifyLanguage {
	t.tOpmProductMaterialClassifyLanguageDo.DO = *(t.tOpmProductMaterialClassifyLanguageDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOpmProductMaterialClassifyLanguage) updateTableName(table string) *tOpmProductMaterialClassifyLanguage {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.ClassifyId = field.NewInt64(table, "classify_id")
	t.Lang = field.NewString(table, "lang")
	t.ClassifyName = field.NewString(table, "classify_name")

	t.fillFieldMap()

	return t
}

func (t *tOpmProductMaterialClassifyLanguage) WithContext(ctx context.Context) *tOpmProductMaterialClassifyLanguageDo {
	return t.tOpmProductMaterialClassifyLanguageDo.WithContext(ctx)
}

func (t tOpmProductMaterialClassifyLanguage) TableName() string {
	return t.tOpmProductMaterialClassifyLanguageDo.TableName()
}

func (t tOpmProductMaterialClassifyLanguage) Alias() string {
	return t.tOpmProductMaterialClassifyLanguageDo.Alias()
}

func (t *tOpmProductMaterialClassifyLanguage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOpmProductMaterialClassifyLanguage) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["id"] = t.Id
	t.fieldMap["classify_id"] = t.ClassifyId
	t.fieldMap["lang"] = t.Lang
	t.fieldMap["classify_name"] = t.ClassifyName
}

func (t tOpmProductMaterialClassifyLanguage) clone(db *gorm.DB) tOpmProductMaterialClassifyLanguage {
	t.tOpmProductMaterialClassifyLanguageDo.ReplaceDB(db)
	return t
}

type tOpmProductMaterialClassifyLanguageDo struct{ gen.DO }

func (t tOpmProductMaterialClassifyLanguageDo) Debug() *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Debug())
}

func (t tOpmProductMaterialClassifyLanguageDo) WithContext(ctx context.Context) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOpmProductMaterialClassifyLanguageDo) Clauses(conds ...clause.Expression) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Returning(value interface{}, columns ...string) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Not(conds ...gen.Condition) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Or(conds ...gen.Condition) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Select(conds ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Where(conds ...gen.Condition) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOpmProductMaterialClassifyLanguageDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOpmProductMaterialClassifyLanguageDo) Order(conds ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Distinct(cols ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Omit(cols ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Join(table schema.Tabler, on ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOpmProductMaterialClassifyLanguageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOpmProductMaterialClassifyLanguageDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Group(cols ...field.Expr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Having(conds ...gen.Condition) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Limit(limit int) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOpmProductMaterialClassifyLanguageDo) Offset(offset int) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOpmProductMaterialClassifyLanguageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Unscoped() *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOpmProductMaterialClassifyLanguageDo) Create(values ...*model.TOpmProductMaterialClassifyLanguage) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOpmProductMaterialClassifyLanguageDo) CreateInBatches(values []*model.TOpmProductMaterialClassifyLanguage, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOpmProductMaterialClassifyLanguageDo) Save(values ...*model.TOpmProductMaterialClassifyLanguage) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOpmProductMaterialClassifyLanguageDo) First() (*model.TOpmProductMaterialClassifyLanguage, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductMaterialClassifyLanguage), nil
	}
}

func (t tOpmProductMaterialClassifyLanguageDo) Take() (*model.TOpmProductMaterialClassifyLanguage, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductMaterialClassifyLanguage), nil
	}
}

func (t tOpmProductMaterialClassifyLanguageDo) Last() (*model.TOpmProductMaterialClassifyLanguage, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductMaterialClassifyLanguage), nil
	}
}

func (t tOpmProductMaterialClassifyLanguageDo) Find() ([]*model.TOpmProductMaterialClassifyLanguage, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOpmProductMaterialClassifyLanguage), err
}

func (t tOpmProductMaterialClassifyLanguageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOpmProductMaterialClassifyLanguage, err error) {
	buf := make([]*model.TOpmProductMaterialClassifyLanguage, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOpmProductMaterialClassifyLanguageDo) FindInBatches(result *[]*model.TOpmProductMaterialClassifyLanguage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOpmProductMaterialClassifyLanguageDo) Attrs(attrs ...field.AssignExpr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Assign(attrs ...field.AssignExpr) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOpmProductMaterialClassifyLanguageDo) Joins(field field.RelationField) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOpmProductMaterialClassifyLanguageDo) Preload(field field.RelationField) *tOpmProductMaterialClassifyLanguageDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOpmProductMaterialClassifyLanguageDo) FirstOrInit() (*model.TOpmProductMaterialClassifyLanguage, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductMaterialClassifyLanguage), nil
	}
}

func (t tOpmProductMaterialClassifyLanguageDo) FirstOrCreate() (*model.TOpmProductMaterialClassifyLanguage, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmProductMaterialClassifyLanguage), nil
	}
}

func (t tOpmProductMaterialClassifyLanguageDo) FindByPage(offset int, limit int) (result []*model.TOpmProductMaterialClassifyLanguage, count int64, err error) {
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

func (t tOpmProductMaterialClassifyLanguageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOpmProductMaterialClassifyLanguageDo) withDO(do gen.Dao) *tOpmProductMaterialClassifyLanguageDo {
	t.DO = *do.(*gen.DO)
	return t
}
