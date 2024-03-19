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

func newTOpmControlPanels(db *gorm.DB) tOpmControlPanels {
	_tOpmControlPanels := tOpmControlPanels{}

	_tOpmControlPanels.tOpmControlPanelsDo.UseDB(db)
	_tOpmControlPanels.tOpmControlPanelsDo.UseModel(&model.TOpmControlPanels{})

	tableName := _tOpmControlPanels.tOpmControlPanelsDo.TableName()
	_tOpmControlPanels.ALL = field.NewField(tableName, "*")
	_tOpmControlPanels.Id = field.NewInt64(tableName, "id")
	_tOpmControlPanels.Name = field.NewString(tableName, "name")
	_tOpmControlPanels.NameEn = field.NewString(tableName, "name_en")
	_tOpmControlPanels.Lang = field.NewString(tableName, "lang")
	_tOpmControlPanels.Desc = field.NewString(tableName, "desc")
	_tOpmControlPanels.Url = field.NewString(tableName, "url")
	_tOpmControlPanels.UrlName = field.NewString(tableName, "url_name")
	_tOpmControlPanels.PanelSize = field.NewInt32(tableName, "panel_size")
	_tOpmControlPanels.PreviewName = field.NewString(tableName, "preview_name")
	_tOpmControlPanels.PreviewUrl = field.NewString(tableName, "preview_url")
	_tOpmControlPanels.PreviewSize = field.NewInt32(tableName, "preview_size")
	_tOpmControlPanels.ProductTypeId = field.NewInt64(tableName, "product_type_id")
	_tOpmControlPanels.ProductId = field.NewInt64(tableName, "product_id")
	_tOpmControlPanels.CreatedBy = field.NewInt64(tableName, "created_by")
	_tOpmControlPanels.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tOpmControlPanels.CreatedAt = field.NewTime(tableName, "created_at")
	_tOpmControlPanels.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tOpmControlPanels.DeletedAt = field.NewField(tableName, "deleted_at")
	_tOpmControlPanels.TenantId = field.NewString(tableName, "tenant_id")

	_tOpmControlPanels.fillFieldMap()

	return _tOpmControlPanels
}

type tOpmControlPanels struct {
	tOpmControlPanelsDo tOpmControlPanelsDo

	ALL           field.Field
	Id            field.Int64
	Name          field.String
	NameEn        field.String
	Lang          field.String
	Desc          field.String
	Url           field.String
	UrlName       field.String
	PanelSize     field.Int32
	PreviewName   field.String
	PreviewUrl    field.String
	PreviewSize   field.Int32
	ProductTypeId field.Int64
	ProductId     field.Int64
	CreatedBy     field.Int64
	UpdatedBy     field.Int64
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	TenantId      field.String

	fieldMap map[string]field.Expr
}

func (t tOpmControlPanels) Table(newTableName string) *tOpmControlPanels {
	t.tOpmControlPanelsDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOpmControlPanels) As(alias string) *tOpmControlPanels {
	t.tOpmControlPanelsDo.DO = *(t.tOpmControlPanelsDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOpmControlPanels) updateTableName(table string) *tOpmControlPanels {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.Name = field.NewString(table, "name")
	t.NameEn = field.NewString(table, "name_en")
	t.Lang = field.NewString(table, "lang")
	t.Desc = field.NewString(table, "desc")
	t.Url = field.NewString(table, "url")
	t.UrlName = field.NewString(table, "url_name")
	t.PanelSize = field.NewInt32(table, "panel_size")
	t.PreviewName = field.NewString(table, "preview_name")
	t.PreviewUrl = field.NewString(table, "preview_url")
	t.PreviewSize = field.NewInt32(table, "preview_size")
	t.ProductTypeId = field.NewInt64(table, "product_type_id")
	t.ProductId = field.NewInt64(table, "product_id")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.TenantId = field.NewString(table, "tenant_id")

	t.fillFieldMap()

	return t
}

func (t *tOpmControlPanels) WithContext(ctx context.Context) *tOpmControlPanelsDo {
	return t.tOpmControlPanelsDo.WithContext(ctx)
}

func (t tOpmControlPanels) TableName() string { return t.tOpmControlPanelsDo.TableName() }

func (t tOpmControlPanels) Alias() string { return t.tOpmControlPanelsDo.Alias() }

func (t *tOpmControlPanels) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOpmControlPanels) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 19)
	t.fieldMap["id"] = t.Id
	t.fieldMap["name"] = t.Name
	t.fieldMap["name_en"] = t.NameEn
	t.fieldMap["lang"] = t.Lang
	t.fieldMap["desc"] = t.Desc
	t.fieldMap["url"] = t.Url
	t.fieldMap["url_name"] = t.UrlName
	t.fieldMap["panel_size"] = t.PanelSize
	t.fieldMap["preview_name"] = t.PreviewName
	t.fieldMap["preview_url"] = t.PreviewUrl
	t.fieldMap["preview_size"] = t.PreviewSize
	t.fieldMap["product_type_id"] = t.ProductTypeId
	t.fieldMap["product_id"] = t.ProductId
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["tenant_id"] = t.TenantId
}

func (t tOpmControlPanels) clone(db *gorm.DB) tOpmControlPanels {
	t.tOpmControlPanelsDo.ReplaceDB(db)
	return t
}

type tOpmControlPanelsDo struct{ gen.DO }

func (t tOpmControlPanelsDo) Debug() *tOpmControlPanelsDo {
	return t.withDO(t.DO.Debug())
}

func (t tOpmControlPanelsDo) WithContext(ctx context.Context) *tOpmControlPanelsDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOpmControlPanelsDo) Clauses(conds ...clause.Expression) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOpmControlPanelsDo) Returning(value interface{}, columns ...string) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOpmControlPanelsDo) Not(conds ...gen.Condition) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOpmControlPanelsDo) Or(conds ...gen.Condition) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOpmControlPanelsDo) Select(conds ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOpmControlPanelsDo) Where(conds ...gen.Condition) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOpmControlPanelsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOpmControlPanelsDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOpmControlPanelsDo) Order(conds ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOpmControlPanelsDo) Distinct(cols ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOpmControlPanelsDo) Omit(cols ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOpmControlPanelsDo) Join(table schema.Tabler, on ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOpmControlPanelsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOpmControlPanelsDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOpmControlPanelsDo) Group(cols ...field.Expr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOpmControlPanelsDo) Having(conds ...gen.Condition) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOpmControlPanelsDo) Limit(limit int) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOpmControlPanelsDo) Offset(offset int) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOpmControlPanelsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOpmControlPanelsDo) Unscoped() *tOpmControlPanelsDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOpmControlPanelsDo) Create(values ...*model.TOpmControlPanels) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOpmControlPanelsDo) CreateInBatches(values []*model.TOpmControlPanels, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOpmControlPanelsDo) Save(values ...*model.TOpmControlPanels) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOpmControlPanelsDo) First() (*model.TOpmControlPanels, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmControlPanels), nil
	}
}

func (t tOpmControlPanelsDo) Take() (*model.TOpmControlPanels, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmControlPanels), nil
	}
}

func (t tOpmControlPanelsDo) Last() (*model.TOpmControlPanels, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmControlPanels), nil
	}
}

func (t tOpmControlPanelsDo) Find() ([]*model.TOpmControlPanels, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOpmControlPanels), err
}

func (t tOpmControlPanelsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOpmControlPanels, err error) {
	buf := make([]*model.TOpmControlPanels, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOpmControlPanelsDo) FindInBatches(result *[]*model.TOpmControlPanels, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOpmControlPanelsDo) Attrs(attrs ...field.AssignExpr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOpmControlPanelsDo) Assign(attrs ...field.AssignExpr) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOpmControlPanelsDo) Joins(field field.RelationField) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOpmControlPanelsDo) Preload(field field.RelationField) *tOpmControlPanelsDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOpmControlPanelsDo) FirstOrInit() (*model.TOpmControlPanels, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmControlPanels), nil
	}
}

func (t tOpmControlPanelsDo) FirstOrCreate() (*model.TOpmControlPanels, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmControlPanels), nil
	}
}

func (t tOpmControlPanelsDo) FindByPage(offset int, limit int) (result []*model.TOpmControlPanels, count int64, err error) {
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

func (t tOpmControlPanelsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOpmControlPanelsDo) withDO(do gen.Dao) *tOpmControlPanelsDo {
	t.DO = *do.(*gen.DO)
	return t
}
