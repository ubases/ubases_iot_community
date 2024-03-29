// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_config/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTLangTranslateType(db *gorm.DB) tLangTranslateType {
	_tLangTranslateType := tLangTranslateType{}

	_tLangTranslateType.tLangTranslateTypeDo.UseDB(db)
	_tLangTranslateType.tLangTranslateTypeDo.UseModel(&model.TLangTranslateType{})

	tableName := _tLangTranslateType.tLangTranslateTypeDo.TableName()
	_tLangTranslateType.ALL = field.NewField(tableName, "*")
	_tLangTranslateType.Id = field.NewInt64(tableName, "id")
	_tLangTranslateType.Name = field.NewString(tableName, "name")
	_tLangTranslateType.Code = field.NewString(tableName, "code")
	_tLangTranslateType.Status = field.NewInt32(tableName, "status")
	_tLangTranslateType.CreatedBy = field.NewInt64(tableName, "created_by")
	_tLangTranslateType.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tLangTranslateType.CreatedAt = field.NewTime(tableName, "created_at")
	_tLangTranslateType.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tLangTranslateType.DeletedAt = field.NewField(tableName, "deleted_at")

	_tLangTranslateType.fillFieldMap()

	return _tLangTranslateType
}

type tLangTranslateType struct {
	tLangTranslateTypeDo tLangTranslateTypeDo

	ALL       field.Field
	Id        field.Int64
	Name      field.String
	Code      field.String
	Status    field.Int32
	CreatedBy field.Int64
	UpdatedBy field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (t tLangTranslateType) Table(newTableName string) *tLangTranslateType {
	t.tLangTranslateTypeDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tLangTranslateType) As(alias string) *tLangTranslateType {
	t.tLangTranslateTypeDo.DO = *(t.tLangTranslateTypeDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tLangTranslateType) updateTableName(table string) *tLangTranslateType {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.Name = field.NewString(table, "name")
	t.Code = field.NewString(table, "code")
	t.Status = field.NewInt32(table, "status")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tLangTranslateType) WithContext(ctx context.Context) *tLangTranslateTypeDo {
	return t.tLangTranslateTypeDo.WithContext(ctx)
}

func (t tLangTranslateType) TableName() string { return t.tLangTranslateTypeDo.TableName() }

func (t tLangTranslateType) Alias() string { return t.tLangTranslateTypeDo.Alias() }

func (t *tLangTranslateType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tLangTranslateType) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["id"] = t.Id
	t.fieldMap["name"] = t.Name
	t.fieldMap["code"] = t.Code
	t.fieldMap["status"] = t.Status
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t tLangTranslateType) clone(db *gorm.DB) tLangTranslateType {
	t.tLangTranslateTypeDo.ReplaceDB(db)
	return t
}

type tLangTranslateTypeDo struct{ gen.DO }

func (t tLangTranslateTypeDo) Debug() *tLangTranslateTypeDo {
	return t.withDO(t.DO.Debug())
}

func (t tLangTranslateTypeDo) WithContext(ctx context.Context) *tLangTranslateTypeDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tLangTranslateTypeDo) Clauses(conds ...clause.Expression) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tLangTranslateTypeDo) Returning(value interface{}, columns ...string) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tLangTranslateTypeDo) Not(conds ...gen.Condition) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tLangTranslateTypeDo) Or(conds ...gen.Condition) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tLangTranslateTypeDo) Select(conds ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tLangTranslateTypeDo) Where(conds ...gen.Condition) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tLangTranslateTypeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tLangTranslateTypeDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tLangTranslateTypeDo) Order(conds ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tLangTranslateTypeDo) Distinct(cols ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tLangTranslateTypeDo) Omit(cols ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tLangTranslateTypeDo) Join(table schema.Tabler, on ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tLangTranslateTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tLangTranslateTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tLangTranslateTypeDo) Group(cols ...field.Expr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tLangTranslateTypeDo) Having(conds ...gen.Condition) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tLangTranslateTypeDo) Limit(limit int) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tLangTranslateTypeDo) Offset(offset int) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tLangTranslateTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tLangTranslateTypeDo) Unscoped() *tLangTranslateTypeDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tLangTranslateTypeDo) Create(values ...*model.TLangTranslateType) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tLangTranslateTypeDo) CreateInBatches(values []*model.TLangTranslateType, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tLangTranslateTypeDo) Save(values ...*model.TLangTranslateType) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tLangTranslateTypeDo) First() (*model.TLangTranslateType, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLangTranslateType), nil
	}
}

func (t tLangTranslateTypeDo) Take() (*model.TLangTranslateType, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLangTranslateType), nil
	}
}

func (t tLangTranslateTypeDo) Last() (*model.TLangTranslateType, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLangTranslateType), nil
	}
}

func (t tLangTranslateTypeDo) Find() ([]*model.TLangTranslateType, error) {
	result, err := t.DO.Find()
	return result.([]*model.TLangTranslateType), err
}

func (t tLangTranslateTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TLangTranslateType, err error) {
	buf := make([]*model.TLangTranslateType, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tLangTranslateTypeDo) FindInBatches(result *[]*model.TLangTranslateType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tLangTranslateTypeDo) Attrs(attrs ...field.AssignExpr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tLangTranslateTypeDo) Assign(attrs ...field.AssignExpr) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tLangTranslateTypeDo) Joins(field field.RelationField) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tLangTranslateTypeDo) Preload(field field.RelationField) *tLangTranslateTypeDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tLangTranslateTypeDo) FirstOrInit() (*model.TLangTranslateType, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLangTranslateType), nil
	}
}

func (t tLangTranslateTypeDo) FirstOrCreate() (*model.TLangTranslateType, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TLangTranslateType), nil
	}
}

func (t tLangTranslateTypeDo) FindByPage(offset int, limit int) (result []*model.TLangTranslateType, count int64, err error) {
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

func (t tLangTranslateTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tLangTranslateTypeDo) withDO(do gen.Dao) *tLangTranslateTypeDo {
	t.DO = *do.(*gen.DO)
	return t
}
