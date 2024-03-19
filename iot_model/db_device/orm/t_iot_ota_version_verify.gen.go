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

func newTIotOtaVersionVerify(db *gorm.DB) tIotOtaVersionVerify {
	_tIotOtaVersionVerify := tIotOtaVersionVerify{}

	_tIotOtaVersionVerify.tIotOtaVersionVerifyDo.UseDB(db)
	_tIotOtaVersionVerify.tIotOtaVersionVerifyDo.UseModel(&model.TIotOtaVersionVerify{})

	tableName := _tIotOtaVersionVerify.tIotOtaVersionVerifyDo.TableName()
	_tIotOtaVersionVerify.ALL = field.NewField(tableName, "*")
	_tIotOtaVersionVerify.Id = field.NewInt64(tableName, "id")
	_tIotOtaVersionVerify.VersionId = field.NewInt64(tableName, "version_id")
	_tIotOtaVersionVerify.DeviceVersion = field.NewString(tableName, "device_version")
	_tIotOtaVersionVerify.Did = field.NewString(tableName, "did")
	_tIotOtaVersionVerify.DeviceId = field.NewInt64(tableName, "device_id")
	_tIotOtaVersionVerify.Status = field.NewInt32(tableName, "status")
	_tIotOtaVersionVerify.DeviceLog = field.NewString(tableName, "device_log")
	_tIotOtaVersionVerify.CreatedBy = field.NewInt64(tableName, "created_by")
	_tIotOtaVersionVerify.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tIotOtaVersionVerify.CreatedAt = field.NewTime(tableName, "created_at")
	_tIotOtaVersionVerify.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tIotOtaVersionVerify.DeletedAt = field.NewField(tableName, "deleted_at")

	_tIotOtaVersionVerify.fillFieldMap()

	return _tIotOtaVersionVerify
}

type tIotOtaVersionVerify struct {
	tIotOtaVersionVerifyDo tIotOtaVersionVerifyDo

	ALL           field.Field
	Id            field.Int64
	VersionId     field.Int64
	DeviceVersion field.String
	Did           field.String
	DeviceId      field.Int64
	Status        field.Int32
	DeviceLog     field.String
	CreatedBy     field.Int64
	UpdatedBy     field.Int64
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field

	fieldMap map[string]field.Expr
}

func (t tIotOtaVersionVerify) Table(newTableName string) *tIotOtaVersionVerify {
	t.tIotOtaVersionVerifyDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tIotOtaVersionVerify) As(alias string) *tIotOtaVersionVerify {
	t.tIotOtaVersionVerifyDo.DO = *(t.tIotOtaVersionVerifyDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tIotOtaVersionVerify) updateTableName(table string) *tIotOtaVersionVerify {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.VersionId = field.NewInt64(table, "version_id")
	t.DeviceVersion = field.NewString(table, "device_version")
	t.Did = field.NewString(table, "did")
	t.DeviceId = field.NewInt64(table, "device_id")
	t.Status = field.NewInt32(table, "status")
	t.DeviceLog = field.NewString(table, "device_log")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tIotOtaVersionVerify) WithContext(ctx context.Context) *tIotOtaVersionVerifyDo {
	return t.tIotOtaVersionVerifyDo.WithContext(ctx)
}

func (t tIotOtaVersionVerify) TableName() string { return t.tIotOtaVersionVerifyDo.TableName() }

func (t tIotOtaVersionVerify) Alias() string { return t.tIotOtaVersionVerifyDo.Alias() }

func (t *tIotOtaVersionVerify) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tIotOtaVersionVerify) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.Id
	t.fieldMap["version_id"] = t.VersionId
	t.fieldMap["device_version"] = t.DeviceVersion
	t.fieldMap["did"] = t.Did
	t.fieldMap["device_id"] = t.DeviceId
	t.fieldMap["status"] = t.Status
	t.fieldMap["device_log"] = t.DeviceLog
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t tIotOtaVersionVerify) clone(db *gorm.DB) tIotOtaVersionVerify {
	t.tIotOtaVersionVerifyDo.ReplaceDB(db)
	return t
}

type tIotOtaVersionVerifyDo struct{ gen.DO }

func (t tIotOtaVersionVerifyDo) Debug() *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Debug())
}

func (t tIotOtaVersionVerifyDo) WithContext(ctx context.Context) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tIotOtaVersionVerifyDo) Clauses(conds ...clause.Expression) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tIotOtaVersionVerifyDo) Returning(value interface{}, columns ...string) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tIotOtaVersionVerifyDo) Not(conds ...gen.Condition) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tIotOtaVersionVerifyDo) Or(conds ...gen.Condition) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tIotOtaVersionVerifyDo) Select(conds ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tIotOtaVersionVerifyDo) Where(conds ...gen.Condition) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tIotOtaVersionVerifyDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tIotOtaVersionVerifyDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tIotOtaVersionVerifyDo) Order(conds ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tIotOtaVersionVerifyDo) Distinct(cols ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tIotOtaVersionVerifyDo) Omit(cols ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tIotOtaVersionVerifyDo) Join(table schema.Tabler, on ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tIotOtaVersionVerifyDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tIotOtaVersionVerifyDo) RightJoin(table schema.Tabler, on ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tIotOtaVersionVerifyDo) Group(cols ...field.Expr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tIotOtaVersionVerifyDo) Having(conds ...gen.Condition) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tIotOtaVersionVerifyDo) Limit(limit int) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tIotOtaVersionVerifyDo) Offset(offset int) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tIotOtaVersionVerifyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tIotOtaVersionVerifyDo) Unscoped() *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tIotOtaVersionVerifyDo) Create(values ...*model.TIotOtaVersionVerify) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tIotOtaVersionVerifyDo) CreateInBatches(values []*model.TIotOtaVersionVerify, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tIotOtaVersionVerifyDo) Save(values ...*model.TIotOtaVersionVerify) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tIotOtaVersionVerifyDo) First() (*model.TIotOtaVersionVerify, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionVerify), nil
	}
}

func (t tIotOtaVersionVerifyDo) Take() (*model.TIotOtaVersionVerify, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionVerify), nil
	}
}

func (t tIotOtaVersionVerifyDo) Last() (*model.TIotOtaVersionVerify, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionVerify), nil
	}
}

func (t tIotOtaVersionVerifyDo) Find() ([]*model.TIotOtaVersionVerify, error) {
	result, err := t.DO.Find()
	return result.([]*model.TIotOtaVersionVerify), err
}

func (t tIotOtaVersionVerifyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TIotOtaVersionVerify, err error) {
	buf := make([]*model.TIotOtaVersionVerify, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tIotOtaVersionVerifyDo) FindInBatches(result *[]*model.TIotOtaVersionVerify, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tIotOtaVersionVerifyDo) Attrs(attrs ...field.AssignExpr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tIotOtaVersionVerifyDo) Assign(attrs ...field.AssignExpr) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tIotOtaVersionVerifyDo) Joins(field field.RelationField) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tIotOtaVersionVerifyDo) Preload(field field.RelationField) *tIotOtaVersionVerifyDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tIotOtaVersionVerifyDo) FirstOrInit() (*model.TIotOtaVersionVerify, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionVerify), nil
	}
}

func (t tIotOtaVersionVerifyDo) FirstOrCreate() (*model.TIotOtaVersionVerify, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionVerify), nil
	}
}

func (t tIotOtaVersionVerifyDo) FindByPage(offset int, limit int) (result []*model.TIotOtaVersionVerify, count int64, err error) {
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

func (t tIotOtaVersionVerifyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tIotOtaVersionVerifyDo) withDO(do gen.Dao) *tIotOtaVersionVerifyDo {
	t.DO = *do.(*gen.DO)
	return t
}
