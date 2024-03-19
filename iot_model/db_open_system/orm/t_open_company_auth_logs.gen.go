// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_open_system/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTOpenCompanyAuthLogs(db *gorm.DB) tOpenCompanyAuthLogs {
	_tOpenCompanyAuthLogs := tOpenCompanyAuthLogs{}

	_tOpenCompanyAuthLogs.tOpenCompanyAuthLogsDo.UseDB(db)
	_tOpenCompanyAuthLogs.tOpenCompanyAuthLogsDo.UseModel(&model.TOpenCompanyAuthLogs{})

	tableName := _tOpenCompanyAuthLogs.tOpenCompanyAuthLogsDo.TableName()
	_tOpenCompanyAuthLogs.ALL = field.NewField(tableName, "*")
	_tOpenCompanyAuthLogs.Id = field.NewInt64(tableName, "id")
	_tOpenCompanyAuthLogs.CompanyId = field.NewInt64(tableName, "company_id")
	_tOpenCompanyAuthLogs.AuthResult = field.NewString(tableName, "auth_result")
	_tOpenCompanyAuthLogs.AuthName = field.NewString(tableName, "auth_name")
	_tOpenCompanyAuthLogs.AuthDate = field.NewTime(tableName, "auth_date")
	_tOpenCompanyAuthLogs.Why = field.NewString(tableName, "why")
	_tOpenCompanyAuthLogs.CreatedBy = field.NewInt64(tableName, "created_by")
	_tOpenCompanyAuthLogs.CreatedAt = field.NewTime(tableName, "created_at")
	_tOpenCompanyAuthLogs.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tOpenCompanyAuthLogs.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tOpenCompanyAuthLogs.DeletedAt = field.NewField(tableName, "deleted_at")

	_tOpenCompanyAuthLogs.fillFieldMap()

	return _tOpenCompanyAuthLogs
}

type tOpenCompanyAuthLogs struct {
	tOpenCompanyAuthLogsDo tOpenCompanyAuthLogsDo

	ALL        field.Field
	Id         field.Int64
	CompanyId  field.Int64
	AuthResult field.String
	AuthName   field.String
	AuthDate   field.Time
	Why        field.String
	CreatedBy  field.Int64
	CreatedAt  field.Time
	UpdatedBy  field.Int64
	UpdatedAt  field.Time
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (t tOpenCompanyAuthLogs) Table(newTableName string) *tOpenCompanyAuthLogs {
	t.tOpenCompanyAuthLogsDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOpenCompanyAuthLogs) As(alias string) *tOpenCompanyAuthLogs {
	t.tOpenCompanyAuthLogsDo.DO = *(t.tOpenCompanyAuthLogsDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOpenCompanyAuthLogs) updateTableName(table string) *tOpenCompanyAuthLogs {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.CompanyId = field.NewInt64(table, "company_id")
	t.AuthResult = field.NewString(table, "auth_result")
	t.AuthName = field.NewString(table, "auth_name")
	t.AuthDate = field.NewTime(table, "auth_date")
	t.Why = field.NewString(table, "why")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tOpenCompanyAuthLogs) WithContext(ctx context.Context) *tOpenCompanyAuthLogsDo {
	return t.tOpenCompanyAuthLogsDo.WithContext(ctx)
}

func (t tOpenCompanyAuthLogs) TableName() string { return t.tOpenCompanyAuthLogsDo.TableName() }

func (t tOpenCompanyAuthLogs) Alias() string { return t.tOpenCompanyAuthLogsDo.Alias() }

func (t *tOpenCompanyAuthLogs) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOpenCompanyAuthLogs) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["id"] = t.Id
	t.fieldMap["company_id"] = t.CompanyId
	t.fieldMap["auth_result"] = t.AuthResult
	t.fieldMap["auth_name"] = t.AuthName
	t.fieldMap["auth_date"] = t.AuthDate
	t.fieldMap["why"] = t.Why
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t tOpenCompanyAuthLogs) clone(db *gorm.DB) tOpenCompanyAuthLogs {
	t.tOpenCompanyAuthLogsDo.ReplaceDB(db)
	return t
}

type tOpenCompanyAuthLogsDo struct{ gen.DO }

func (t tOpenCompanyAuthLogsDo) Debug() *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Debug())
}

func (t tOpenCompanyAuthLogsDo) WithContext(ctx context.Context) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOpenCompanyAuthLogsDo) Clauses(conds ...clause.Expression) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOpenCompanyAuthLogsDo) Returning(value interface{}, columns ...string) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOpenCompanyAuthLogsDo) Not(conds ...gen.Condition) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOpenCompanyAuthLogsDo) Or(conds ...gen.Condition) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOpenCompanyAuthLogsDo) Select(conds ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOpenCompanyAuthLogsDo) Where(conds ...gen.Condition) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOpenCompanyAuthLogsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOpenCompanyAuthLogsDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOpenCompanyAuthLogsDo) Order(conds ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOpenCompanyAuthLogsDo) Distinct(cols ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOpenCompanyAuthLogsDo) Omit(cols ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOpenCompanyAuthLogsDo) Join(table schema.Tabler, on ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOpenCompanyAuthLogsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOpenCompanyAuthLogsDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOpenCompanyAuthLogsDo) Group(cols ...field.Expr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOpenCompanyAuthLogsDo) Having(conds ...gen.Condition) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOpenCompanyAuthLogsDo) Limit(limit int) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOpenCompanyAuthLogsDo) Offset(offset int) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOpenCompanyAuthLogsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOpenCompanyAuthLogsDo) Unscoped() *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOpenCompanyAuthLogsDo) Create(values ...*model.TOpenCompanyAuthLogs) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOpenCompanyAuthLogsDo) CreateInBatches(values []*model.TOpenCompanyAuthLogs, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOpenCompanyAuthLogsDo) Save(values ...*model.TOpenCompanyAuthLogs) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOpenCompanyAuthLogsDo) First() (*model.TOpenCompanyAuthLogs, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpenCompanyAuthLogs), nil
	}
}

func (t tOpenCompanyAuthLogsDo) Take() (*model.TOpenCompanyAuthLogs, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpenCompanyAuthLogs), nil
	}
}

func (t tOpenCompanyAuthLogsDo) Last() (*model.TOpenCompanyAuthLogs, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpenCompanyAuthLogs), nil
	}
}

func (t tOpenCompanyAuthLogsDo) Find() ([]*model.TOpenCompanyAuthLogs, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOpenCompanyAuthLogs), err
}

func (t tOpenCompanyAuthLogsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOpenCompanyAuthLogs, err error) {
	buf := make([]*model.TOpenCompanyAuthLogs, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOpenCompanyAuthLogsDo) FindInBatches(result *[]*model.TOpenCompanyAuthLogs, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOpenCompanyAuthLogsDo) Attrs(attrs ...field.AssignExpr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOpenCompanyAuthLogsDo) Assign(attrs ...field.AssignExpr) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOpenCompanyAuthLogsDo) Joins(field field.RelationField) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOpenCompanyAuthLogsDo) Preload(field field.RelationField) *tOpenCompanyAuthLogsDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOpenCompanyAuthLogsDo) FirstOrInit() (*model.TOpenCompanyAuthLogs, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpenCompanyAuthLogs), nil
	}
}

func (t tOpenCompanyAuthLogsDo) FirstOrCreate() (*model.TOpenCompanyAuthLogs, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpenCompanyAuthLogs), nil
	}
}

func (t tOpenCompanyAuthLogsDo) FindByPage(offset int, limit int) (result []*model.TOpenCompanyAuthLogs, count int64, err error) {
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

func (t tOpenCompanyAuthLogsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOpenCompanyAuthLogsDo) withDO(do gen.Dao) *tOpenCompanyAuthLogsDo {
	t.DO = *do.(*gen.DO)
	return t
}
