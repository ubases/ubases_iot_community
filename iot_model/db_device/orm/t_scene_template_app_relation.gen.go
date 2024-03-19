// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_device/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTSceneTemplateAppRelation(db *gorm.DB) tSceneTemplateAppRelation {
	_tSceneTemplateAppRelation := tSceneTemplateAppRelation{}

	_tSceneTemplateAppRelation.tSceneTemplateAppRelationDo.UseDB(db)
	_tSceneTemplateAppRelation.tSceneTemplateAppRelationDo.UseModel(&model.TSceneTemplateAppRelation{})

	tableName := _tSceneTemplateAppRelation.tSceneTemplateAppRelationDo.TableName()
	_tSceneTemplateAppRelation.ALL = field.NewField(tableName, "*")
	_tSceneTemplateAppRelation.Id = field.NewInt64(tableName, "id")
	_tSceneTemplateAppRelation.SceneTemplateId = field.NewInt64(tableName, "scene_template_id")
	_tSceneTemplateAppRelation.AppId = field.NewInt64(tableName, "app_id")
	_tSceneTemplateAppRelation.AppKey = field.NewString(tableName, "app_key")
	_tSceneTemplateAppRelation.AppName = field.NewString(tableName, "app_name")
	_tSceneTemplateAppRelation.TenantId = field.NewString(tableName, "tenant_id")
	_tSceneTemplateAppRelation.CreatedBy = field.NewInt64(tableName, "created_by")
	_tSceneTemplateAppRelation.CreatedAt = field.NewTime(tableName, "created_at")

	_tSceneTemplateAppRelation.fillFieldMap()

	return _tSceneTemplateAppRelation
}

type tSceneTemplateAppRelation struct {
	tSceneTemplateAppRelationDo tSceneTemplateAppRelationDo

	ALL             field.Field
	Id              field.Int64
	SceneTemplateId field.Int64
	AppId           field.Int64
	AppKey          field.String
	AppName         field.String
	TenantId        field.String
	CreatedBy       field.Int64
	CreatedAt       field.Time

	fieldMap map[string]field.Expr
}

func (t tSceneTemplateAppRelation) Table(newTableName string) *tSceneTemplateAppRelation {
	t.tSceneTemplateAppRelationDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tSceneTemplateAppRelation) As(alias string) *tSceneTemplateAppRelation {
	t.tSceneTemplateAppRelationDo.DO = *(t.tSceneTemplateAppRelationDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tSceneTemplateAppRelation) updateTableName(table string) *tSceneTemplateAppRelation {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.SceneTemplateId = field.NewInt64(table, "scene_template_id")
	t.AppId = field.NewInt64(table, "app_id")
	t.AppKey = field.NewString(table, "app_key")
	t.AppName = field.NewString(table, "app_name")
	t.TenantId = field.NewString(table, "tenant_id")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.CreatedAt = field.NewTime(table, "created_at")

	t.fillFieldMap()

	return t
}

func (t *tSceneTemplateAppRelation) WithContext(ctx context.Context) *tSceneTemplateAppRelationDo {
	return t.tSceneTemplateAppRelationDo.WithContext(ctx)
}

func (t tSceneTemplateAppRelation) TableName() string {
	return t.tSceneTemplateAppRelationDo.TableName()
}

func (t tSceneTemplateAppRelation) Alias() string { return t.tSceneTemplateAppRelationDo.Alias() }

func (t *tSceneTemplateAppRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tSceneTemplateAppRelation) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["id"] = t.Id
	t.fieldMap["scene_template_id"] = t.SceneTemplateId
	t.fieldMap["app_id"] = t.AppId
	t.fieldMap["app_key"] = t.AppKey
	t.fieldMap["app_name"] = t.AppName
	t.fieldMap["tenant_id"] = t.TenantId
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["created_at"] = t.CreatedAt
}

func (t tSceneTemplateAppRelation) clone(db *gorm.DB) tSceneTemplateAppRelation {
	t.tSceneTemplateAppRelationDo.ReplaceDB(db)
	return t
}

type tSceneTemplateAppRelationDo struct{ gen.DO }

func (t tSceneTemplateAppRelationDo) Debug() *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Debug())
}

func (t tSceneTemplateAppRelationDo) WithContext(ctx context.Context) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tSceneTemplateAppRelationDo) Clauses(conds ...clause.Expression) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tSceneTemplateAppRelationDo) Returning(value interface{}, columns ...string) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tSceneTemplateAppRelationDo) Not(conds ...gen.Condition) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tSceneTemplateAppRelationDo) Or(conds ...gen.Condition) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tSceneTemplateAppRelationDo) Select(conds ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tSceneTemplateAppRelationDo) Where(conds ...gen.Condition) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tSceneTemplateAppRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tSceneTemplateAppRelationDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tSceneTemplateAppRelationDo) Order(conds ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tSceneTemplateAppRelationDo) Distinct(cols ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tSceneTemplateAppRelationDo) Omit(cols ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tSceneTemplateAppRelationDo) Join(table schema.Tabler, on ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tSceneTemplateAppRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tSceneTemplateAppRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tSceneTemplateAppRelationDo) Group(cols ...field.Expr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tSceneTemplateAppRelationDo) Having(conds ...gen.Condition) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tSceneTemplateAppRelationDo) Limit(limit int) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tSceneTemplateAppRelationDo) Offset(offset int) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tSceneTemplateAppRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tSceneTemplateAppRelationDo) Unscoped() *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tSceneTemplateAppRelationDo) Create(values ...*model.TSceneTemplateAppRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tSceneTemplateAppRelationDo) CreateInBatches(values []*model.TSceneTemplateAppRelation, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tSceneTemplateAppRelationDo) Save(values ...*model.TSceneTemplateAppRelation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tSceneTemplateAppRelationDo) First() (*model.TSceneTemplateAppRelation, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSceneTemplateAppRelation), nil
	}
}

func (t tSceneTemplateAppRelationDo) Take() (*model.TSceneTemplateAppRelation, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSceneTemplateAppRelation), nil
	}
}

func (t tSceneTemplateAppRelationDo) Last() (*model.TSceneTemplateAppRelation, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSceneTemplateAppRelation), nil
	}
}

func (t tSceneTemplateAppRelationDo) Find() ([]*model.TSceneTemplateAppRelation, error) {
	result, err := t.DO.Find()
	return result.([]*model.TSceneTemplateAppRelation), err
}

func (t tSceneTemplateAppRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSceneTemplateAppRelation, err error) {
	buf := make([]*model.TSceneTemplateAppRelation, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tSceneTemplateAppRelationDo) FindInBatches(result *[]*model.TSceneTemplateAppRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tSceneTemplateAppRelationDo) Attrs(attrs ...field.AssignExpr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tSceneTemplateAppRelationDo) Assign(attrs ...field.AssignExpr) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tSceneTemplateAppRelationDo) Joins(field field.RelationField) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tSceneTemplateAppRelationDo) Preload(field field.RelationField) *tSceneTemplateAppRelationDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tSceneTemplateAppRelationDo) FirstOrInit() (*model.TSceneTemplateAppRelation, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSceneTemplateAppRelation), nil
	}
}

func (t tSceneTemplateAppRelationDo) FirstOrCreate() (*model.TSceneTemplateAppRelation, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSceneTemplateAppRelation), nil
	}
}

func (t tSceneTemplateAppRelationDo) FindByPage(offset int, limit int) (result []*model.TSceneTemplateAppRelation, count int64, err error) {
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

func (t tSceneTemplateAppRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tSceneTemplateAppRelationDo) withDO(do gen.Dao) *tSceneTemplateAppRelationDo {
	t.DO = *do.(*gen.DO)
	return t
}
