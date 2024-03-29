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

func newTSysAuthRule(db *gorm.DB) tSysAuthRule {
	_tSysAuthRule := tSysAuthRule{}

	_tSysAuthRule.tSysAuthRuleDo.UseDB(db)
	_tSysAuthRule.tSysAuthRuleDo.UseModel(&model.TSysAuthRule{})

	tableName := _tSysAuthRule.tSysAuthRuleDo.TableName()
	_tSysAuthRule.ALL = field.NewField(tableName, "*")
	_tSysAuthRule.Id = field.NewInt64(tableName, "id")
	_tSysAuthRule.Pid = field.NewInt64(tableName, "pid")
	_tSysAuthRule.Name = field.NewString(tableName, "name")
	_tSysAuthRule.Title = field.NewString(tableName, "title")
	_tSysAuthRule.Icon = field.NewString(tableName, "icon")
	_tSysAuthRule.Condition = field.NewString(tableName, "condition")
	_tSysAuthRule.Remark = field.NewString(tableName, "remark")
	_tSysAuthRule.MenuType = field.NewInt32(tableName, "menu_type")
	_tSysAuthRule.Weigh = field.NewInt32(tableName, "weigh")
	_tSysAuthRule.Status = field.NewInt32(tableName, "status")
	_tSysAuthRule.AlwaysShow = field.NewInt32(tableName, "always_show")
	_tSysAuthRule.Path = field.NewString(tableName, "path")
	_tSysAuthRule.JumpPath = field.NewString(tableName, "jump_path")
	_tSysAuthRule.Component = field.NewString(tableName, "component")
	_tSysAuthRule.IsFrame = field.NewInt32(tableName, "is_frame")
	_tSysAuthRule.ModuleType = field.NewString(tableName, "module_type")
	_tSysAuthRule.ModelId = field.NewInt64(tableName, "model_id")
	_tSysAuthRule.IsCache = field.NewInt32(tableName, "is_cache")
	_tSysAuthRule.IsHideChildMenu = field.NewInt32(tableName, "is_hide_child_menu")
	_tSysAuthRule.CreatedAt = field.NewTime(tableName, "created_at")
	_tSysAuthRule.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tSysAuthRule.DeletedAt = field.NewField(tableName, "deleted_at")

	_tSysAuthRule.fillFieldMap()

	return _tSysAuthRule
}

type tSysAuthRule struct {
	tSysAuthRuleDo tSysAuthRuleDo

	ALL             field.Field
	Id              field.Int64
	Pid             field.Int64
	Name            field.String
	Title           field.String
	Icon            field.String
	Condition       field.String
	Remark          field.String
	MenuType        field.Int32
	Weigh           field.Int32
	Status          field.Int32
	AlwaysShow      field.Int32
	Path            field.String
	JumpPath        field.String
	Component       field.String
	IsFrame         field.Int32
	ModuleType      field.String
	ModelId         field.Int64
	IsCache         field.Int32
	IsHideChildMenu field.Int32
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (t tSysAuthRule) Table(newTableName string) *tSysAuthRule {
	t.tSysAuthRuleDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tSysAuthRule) As(alias string) *tSysAuthRule {
	t.tSysAuthRuleDo.DO = *(t.tSysAuthRuleDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tSysAuthRule) updateTableName(table string) *tSysAuthRule {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.Pid = field.NewInt64(table, "pid")
	t.Name = field.NewString(table, "name")
	t.Title = field.NewString(table, "title")
	t.Icon = field.NewString(table, "icon")
	t.Condition = field.NewString(table, "condition")
	t.Remark = field.NewString(table, "remark")
	t.MenuType = field.NewInt32(table, "menu_type")
	t.Weigh = field.NewInt32(table, "weigh")
	t.Status = field.NewInt32(table, "status")
	t.AlwaysShow = field.NewInt32(table, "always_show")
	t.Path = field.NewString(table, "path")
	t.JumpPath = field.NewString(table, "jump_path")
	t.Component = field.NewString(table, "component")
	t.IsFrame = field.NewInt32(table, "is_frame")
	t.ModuleType = field.NewString(table, "module_type")
	t.ModelId = field.NewInt64(table, "model_id")
	t.IsCache = field.NewInt32(table, "is_cache")
	t.IsHideChildMenu = field.NewInt32(table, "is_hide_child_menu")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tSysAuthRule) WithContext(ctx context.Context) *tSysAuthRuleDo {
	return t.tSysAuthRuleDo.WithContext(ctx)
}

func (t tSysAuthRule) TableName() string { return t.tSysAuthRuleDo.TableName() }

func (t tSysAuthRule) Alias() string { return t.tSysAuthRuleDo.Alias() }

func (t *tSysAuthRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tSysAuthRule) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 22)
	t.fieldMap["id"] = t.Id
	t.fieldMap["pid"] = t.Pid
	t.fieldMap["name"] = t.Name
	t.fieldMap["title"] = t.Title
	t.fieldMap["icon"] = t.Icon
	t.fieldMap["condition"] = t.Condition
	t.fieldMap["remark"] = t.Remark
	t.fieldMap["menu_type"] = t.MenuType
	t.fieldMap["weigh"] = t.Weigh
	t.fieldMap["status"] = t.Status
	t.fieldMap["always_show"] = t.AlwaysShow
	t.fieldMap["path"] = t.Path
	t.fieldMap["jump_path"] = t.JumpPath
	t.fieldMap["component"] = t.Component
	t.fieldMap["is_frame"] = t.IsFrame
	t.fieldMap["module_type"] = t.ModuleType
	t.fieldMap["model_id"] = t.ModelId
	t.fieldMap["is_cache"] = t.IsCache
	t.fieldMap["is_hide_child_menu"] = t.IsHideChildMenu
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t tSysAuthRule) clone(db *gorm.DB) tSysAuthRule {
	t.tSysAuthRuleDo.ReplaceDB(db)
	return t
}

type tSysAuthRuleDo struct{ gen.DO }

func (t tSysAuthRuleDo) Debug() *tSysAuthRuleDo {
	return t.withDO(t.DO.Debug())
}

func (t tSysAuthRuleDo) WithContext(ctx context.Context) *tSysAuthRuleDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tSysAuthRuleDo) Clauses(conds ...clause.Expression) *tSysAuthRuleDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tSysAuthRuleDo) Returning(value interface{}, columns ...string) *tSysAuthRuleDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tSysAuthRuleDo) Not(conds ...gen.Condition) *tSysAuthRuleDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tSysAuthRuleDo) Or(conds ...gen.Condition) *tSysAuthRuleDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tSysAuthRuleDo) Select(conds ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tSysAuthRuleDo) Where(conds ...gen.Condition) *tSysAuthRuleDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tSysAuthRuleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tSysAuthRuleDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tSysAuthRuleDo) Order(conds ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tSysAuthRuleDo) Distinct(cols ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tSysAuthRuleDo) Omit(cols ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tSysAuthRuleDo) Join(table schema.Tabler, on ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tSysAuthRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tSysAuthRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tSysAuthRuleDo) Group(cols ...field.Expr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tSysAuthRuleDo) Having(conds ...gen.Condition) *tSysAuthRuleDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tSysAuthRuleDo) Limit(limit int) *tSysAuthRuleDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tSysAuthRuleDo) Offset(offset int) *tSysAuthRuleDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tSysAuthRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tSysAuthRuleDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tSysAuthRuleDo) Unscoped() *tSysAuthRuleDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tSysAuthRuleDo) Create(values ...*model.TSysAuthRule) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tSysAuthRuleDo) CreateInBatches(values []*model.TSysAuthRule, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tSysAuthRuleDo) Save(values ...*model.TSysAuthRule) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tSysAuthRuleDo) First() (*model.TSysAuthRule, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAuthRule), nil
	}
}

func (t tSysAuthRuleDo) Take() (*model.TSysAuthRule, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAuthRule), nil
	}
}

func (t tSysAuthRuleDo) Last() (*model.TSysAuthRule, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAuthRule), nil
	}
}

func (t tSysAuthRuleDo) Find() ([]*model.TSysAuthRule, error) {
	result, err := t.DO.Find()
	return result.([]*model.TSysAuthRule), err
}

func (t tSysAuthRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TSysAuthRule, err error) {
	buf := make([]*model.TSysAuthRule, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tSysAuthRuleDo) FindInBatches(result *[]*model.TSysAuthRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tSysAuthRuleDo) Attrs(attrs ...field.AssignExpr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tSysAuthRuleDo) Assign(attrs ...field.AssignExpr) *tSysAuthRuleDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tSysAuthRuleDo) Joins(field field.RelationField) *tSysAuthRuleDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tSysAuthRuleDo) Preload(field field.RelationField) *tSysAuthRuleDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tSysAuthRuleDo) FirstOrInit() (*model.TSysAuthRule, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAuthRule), nil
	}
}

func (t tSysAuthRuleDo) FirstOrCreate() (*model.TSysAuthRule, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TSysAuthRule), nil
	}
}

func (t tSysAuthRuleDo) FindByPage(offset int, limit int) (result []*model.TSysAuthRule, count int64, err error) {
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

func (t tSysAuthRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tSysAuthRuleDo) withDO(do gen.Dao) *tSysAuthRuleDo {
	t.DO = *do.(*gen.DO)
	return t
}
