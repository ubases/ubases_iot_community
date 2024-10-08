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

func newTIotOtaVersionPublish(db *gorm.DB) tIotOtaVersionPublish {
	_tIotOtaVersionPublish := tIotOtaVersionPublish{}

	_tIotOtaVersionPublish.tIotOtaVersionPublishDo.UseDB(db)
	_tIotOtaVersionPublish.tIotOtaVersionPublishDo.UseModel(&model.TIotOtaVersionPublish{})

	tableName := _tIotOtaVersionPublish.tIotOtaVersionPublishDo.TableName()
	_tIotOtaVersionPublish.ALL = field.NewField(tableName, "*")
	_tIotOtaVersionPublish.Id = field.NewInt64(tableName, "id")
	_tIotOtaVersionPublish.VersionId = field.NewInt64(tableName, "version_id")
	_tIotOtaVersionPublish.PublishMode = field.NewInt32(tableName, "publish_mode")
	_tIotOtaVersionPublish.ScheduleTime = field.NewTime(tableName, "schedule_time")
	_tIotOtaVersionPublish.PublishTime = field.NewTime(tableName, "publish_time")
	_tIotOtaVersionPublish.Status = field.NewInt32(tableName, "status")
	_tIotOtaVersionPublish.IsGray = field.NewInt32(tableName, "is_gray")
	_tIotOtaVersionPublish.UpdateDesc = field.NewString(tableName, "update_desc")
	_tIotOtaVersionPublish.CreatedBy = field.NewInt64(tableName, "created_by")
	_tIotOtaVersionPublish.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tIotOtaVersionPublish.CreatedAt = field.NewTime(tableName, "created_at")
	_tIotOtaVersionPublish.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tIotOtaVersionPublish.DeletedAt = field.NewField(tableName, "deleted_at")

	_tIotOtaVersionPublish.fillFieldMap()

	return _tIotOtaVersionPublish
}

type tIotOtaVersionPublish struct {
	tIotOtaVersionPublishDo tIotOtaVersionPublishDo

	ALL          field.Field
	Id           field.Int64
	VersionId    field.Int64
	PublishMode  field.Int32
	ScheduleTime field.Time
	PublishTime  field.Time
	Status       field.Int32
	IsGray       field.Int32
	UpdateDesc   field.String
	CreatedBy    field.Int64
	UpdatedBy    field.Int64
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field

	fieldMap map[string]field.Expr
}

func (t tIotOtaVersionPublish) Table(newTableName string) *tIotOtaVersionPublish {
	t.tIotOtaVersionPublishDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tIotOtaVersionPublish) As(alias string) *tIotOtaVersionPublish {
	t.tIotOtaVersionPublishDo.DO = *(t.tIotOtaVersionPublishDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tIotOtaVersionPublish) updateTableName(table string) *tIotOtaVersionPublish {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.VersionId = field.NewInt64(table, "version_id")
	t.PublishMode = field.NewInt32(table, "publish_mode")
	t.ScheduleTime = field.NewTime(table, "schedule_time")
	t.PublishTime = field.NewTime(table, "publish_time")
	t.Status = field.NewInt32(table, "status")
	t.IsGray = field.NewInt32(table, "is_gray")
	t.UpdateDesc = field.NewString(table, "update_desc")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tIotOtaVersionPublish) WithContext(ctx context.Context) *tIotOtaVersionPublishDo {
	return t.tIotOtaVersionPublishDo.WithContext(ctx)
}

func (t tIotOtaVersionPublish) TableName() string { return t.tIotOtaVersionPublishDo.TableName() }

func (t tIotOtaVersionPublish) Alias() string { return t.tIotOtaVersionPublishDo.Alias() }

func (t *tIotOtaVersionPublish) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tIotOtaVersionPublish) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 13)
	t.fieldMap["id"] = t.Id
	t.fieldMap["version_id"] = t.VersionId
	t.fieldMap["publish_mode"] = t.PublishMode
	t.fieldMap["schedule_time"] = t.ScheduleTime
	t.fieldMap["publish_time"] = t.PublishTime
	t.fieldMap["status"] = t.Status
	t.fieldMap["is_gray"] = t.IsGray
	t.fieldMap["update_desc"] = t.UpdateDesc
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t tIotOtaVersionPublish) clone(db *gorm.DB) tIotOtaVersionPublish {
	t.tIotOtaVersionPublishDo.ReplaceDB(db)
	return t
}

type tIotOtaVersionPublishDo struct{ gen.DO }

func (t tIotOtaVersionPublishDo) Debug() *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Debug())
}

func (t tIotOtaVersionPublishDo) WithContext(ctx context.Context) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tIotOtaVersionPublishDo) Clauses(conds ...clause.Expression) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tIotOtaVersionPublishDo) Returning(value interface{}, columns ...string) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tIotOtaVersionPublishDo) Not(conds ...gen.Condition) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tIotOtaVersionPublishDo) Or(conds ...gen.Condition) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tIotOtaVersionPublishDo) Select(conds ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tIotOtaVersionPublishDo) Where(conds ...gen.Condition) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tIotOtaVersionPublishDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tIotOtaVersionPublishDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tIotOtaVersionPublishDo) Order(conds ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tIotOtaVersionPublishDo) Distinct(cols ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tIotOtaVersionPublishDo) Omit(cols ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tIotOtaVersionPublishDo) Join(table schema.Tabler, on ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tIotOtaVersionPublishDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tIotOtaVersionPublishDo) RightJoin(table schema.Tabler, on ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tIotOtaVersionPublishDo) Group(cols ...field.Expr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tIotOtaVersionPublishDo) Having(conds ...gen.Condition) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tIotOtaVersionPublishDo) Limit(limit int) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tIotOtaVersionPublishDo) Offset(offset int) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tIotOtaVersionPublishDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tIotOtaVersionPublishDo) Unscoped() *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tIotOtaVersionPublishDo) Create(values ...*model.TIotOtaVersionPublish) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tIotOtaVersionPublishDo) CreateInBatches(values []*model.TIotOtaVersionPublish, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tIotOtaVersionPublishDo) Save(values ...*model.TIotOtaVersionPublish) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tIotOtaVersionPublishDo) First() (*model.TIotOtaVersionPublish, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionPublish), nil
	}
}

func (t tIotOtaVersionPublishDo) Take() (*model.TIotOtaVersionPublish, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionPublish), nil
	}
}

func (t tIotOtaVersionPublishDo) Last() (*model.TIotOtaVersionPublish, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionPublish), nil
	}
}

func (t tIotOtaVersionPublishDo) Find() ([]*model.TIotOtaVersionPublish, error) {
	result, err := t.DO.Find()
	return result.([]*model.TIotOtaVersionPublish), err
}

func (t tIotOtaVersionPublishDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TIotOtaVersionPublish, err error) {
	buf := make([]*model.TIotOtaVersionPublish, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tIotOtaVersionPublishDo) FindInBatches(result *[]*model.TIotOtaVersionPublish, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tIotOtaVersionPublishDo) Attrs(attrs ...field.AssignExpr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tIotOtaVersionPublishDo) Assign(attrs ...field.AssignExpr) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tIotOtaVersionPublishDo) Joins(field field.RelationField) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tIotOtaVersionPublishDo) Preload(field field.RelationField) *tIotOtaVersionPublishDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tIotOtaVersionPublishDo) FirstOrInit() (*model.TIotOtaVersionPublish, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionPublish), nil
	}
}

func (t tIotOtaVersionPublishDo) FirstOrCreate() (*model.TIotOtaVersionPublish, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TIotOtaVersionPublish), nil
	}
}

func (t tIotOtaVersionPublishDo) FindByPage(offset int, limit int) (result []*model.TIotOtaVersionPublish, count int64, err error) {
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

func (t tIotOtaVersionPublishDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tIotOtaVersionPublishDo) withDO(do gen.Dao) *tIotOtaVersionPublishDo {
	t.DO = *do.(*gen.DO)
	return t
}
