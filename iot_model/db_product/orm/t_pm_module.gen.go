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

func newTPmModule(db *gorm.DB) tPmModule {
	_tPmModule := tPmModule{}

	_tPmModule.tPmModuleDo.UseDB(db)
	_tPmModule.tPmModuleDo.UseModel(&model.TPmModule{})

	tableName := _tPmModule.tPmModuleDo.TableName()
	_tPmModule.ALL = field.NewField(tableName, "*")
	_tPmModule.Id = field.NewInt64(tableName, "id")
	_tPmModule.ModuleName = field.NewString(tableName, "module_name")
	_tPmModule.ModuleNameEn = field.NewString(tableName, "module_name_en")
	_tPmModule.FirmwareType = field.NewString(tableName, "firmware_type")
	_tPmModule.FirmwareFlag = field.NewString(tableName, "firmware_flag")
	_tPmModule.FirmwareId = field.NewInt64(tableName, "firmware_id")
	_tPmModule.Status = field.NewInt32(tableName, "status")
	_tPmModule.ImgUrl = field.NewString(tableName, "img_url")
	_tPmModule.FileUrl = field.NewString(tableName, "file_url")
	_tPmModule.Remark = field.NewString(tableName, "remark")
	_tPmModule.CreatedAt = field.NewTime(tableName, "created_at")
	_tPmModule.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tPmModule.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tPmModule.DeletedAt = field.NewField(tableName, "deleted_at")
	_tPmModule.FileName = field.NewString(tableName, "file_name")
	_tPmModule.DefaultVersion = field.NewString(tableName, "default_version")
	_tPmModule.FirmwareKey = field.NewString(tableName, "firmware_key")

	_tPmModule.fillFieldMap()

	return _tPmModule
}

type tPmModule struct {
	tPmModuleDo tPmModuleDo

	ALL            field.Field
	Id             field.Int64
	ModuleName     field.String
	ModuleNameEn   field.String
	FirmwareType   field.String
	FirmwareFlag   field.String
	FirmwareId     field.Int64
	Status         field.Int32
	ImgUrl         field.String
	FileUrl        field.String
	Remark         field.String
	CreatedAt      field.Time
	UpdatedBy      field.Int64
	UpdatedAt      field.Time
	DeletedAt      field.Field
	FileName       field.String
	DefaultVersion field.String
	FirmwareKey    field.String

	fieldMap map[string]field.Expr
}

func (t tPmModule) Table(newTableName string) *tPmModule {
	t.tPmModuleDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tPmModule) As(alias string) *tPmModule {
	t.tPmModuleDo.DO = *(t.tPmModuleDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tPmModule) updateTableName(table string) *tPmModule {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.ModuleName = field.NewString(table, "module_name")
	t.ModuleNameEn = field.NewString(table, "module_name_en")
	t.FirmwareType = field.NewString(table, "firmware_type")
	t.FirmwareFlag = field.NewString(table, "firmware_flag")
	t.FirmwareId = field.NewInt64(table, "firmware_id")
	t.Status = field.NewInt32(table, "status")
	t.ImgUrl = field.NewString(table, "img_url")
	t.FileUrl = field.NewString(table, "file_url")
	t.Remark = field.NewString(table, "remark")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.FileName = field.NewString(table, "file_name")
	t.DefaultVersion = field.NewString(table, "default_version")
	t.FirmwareKey = field.NewString(table, "firmware_key")

	t.fillFieldMap()

	return t
}

func (t *tPmModule) WithContext(ctx context.Context) *tPmModuleDo {
	return t.tPmModuleDo.WithContext(ctx)
}

func (t tPmModule) TableName() string { return t.tPmModuleDo.TableName() }

func (t tPmModule) Alias() string { return t.tPmModuleDo.Alias() }

func (t *tPmModule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tPmModule) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 17)
	t.fieldMap["id"] = t.Id
	t.fieldMap["module_name"] = t.ModuleName
	t.fieldMap["module_name_en"] = t.ModuleNameEn
	t.fieldMap["firmware_type"] = t.FirmwareType
	t.fieldMap["firmware_flag"] = t.FirmwareFlag
	t.fieldMap["firmware_id"] = t.FirmwareId
	t.fieldMap["status"] = t.Status
	t.fieldMap["img_url"] = t.ImgUrl
	t.fieldMap["file_url"] = t.FileUrl
	t.fieldMap["remark"] = t.Remark
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["file_name"] = t.FileName
	t.fieldMap["default_version"] = t.DefaultVersion
	t.fieldMap["firmware_key"] = t.FirmwareKey
}

func (t tPmModule) clone(db *gorm.DB) tPmModule {
	t.tPmModuleDo.ReplaceDB(db)
	return t
}

type tPmModuleDo struct{ gen.DO }

func (t tPmModuleDo) Debug() *tPmModuleDo {
	return t.withDO(t.DO.Debug())
}

func (t tPmModuleDo) WithContext(ctx context.Context) *tPmModuleDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tPmModuleDo) Clauses(conds ...clause.Expression) *tPmModuleDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tPmModuleDo) Returning(value interface{}, columns ...string) *tPmModuleDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tPmModuleDo) Not(conds ...gen.Condition) *tPmModuleDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tPmModuleDo) Or(conds ...gen.Condition) *tPmModuleDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tPmModuleDo) Select(conds ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tPmModuleDo) Where(conds ...gen.Condition) *tPmModuleDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tPmModuleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tPmModuleDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tPmModuleDo) Order(conds ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tPmModuleDo) Distinct(cols ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tPmModuleDo) Omit(cols ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tPmModuleDo) Join(table schema.Tabler, on ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tPmModuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tPmModuleDo) RightJoin(table schema.Tabler, on ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tPmModuleDo) Group(cols ...field.Expr) *tPmModuleDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tPmModuleDo) Having(conds ...gen.Condition) *tPmModuleDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tPmModuleDo) Limit(limit int) *tPmModuleDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tPmModuleDo) Offset(offset int) *tPmModuleDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tPmModuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tPmModuleDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tPmModuleDo) Unscoped() *tPmModuleDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tPmModuleDo) Create(values ...*model.TPmModule) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tPmModuleDo) CreateInBatches(values []*model.TPmModule, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tPmModuleDo) Save(values ...*model.TPmModule) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tPmModuleDo) First() (*model.TPmModule, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmModule), nil
	}
}

func (t tPmModuleDo) Take() (*model.TPmModule, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmModule), nil
	}
}

func (t tPmModuleDo) Last() (*model.TPmModule, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmModule), nil
	}
}

func (t tPmModuleDo) Find() ([]*model.TPmModule, error) {
	result, err := t.DO.Find()
	return result.([]*model.TPmModule), err
}

func (t tPmModuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TPmModule, err error) {
	buf := make([]*model.TPmModule, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tPmModuleDo) FindInBatches(result *[]*model.TPmModule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tPmModuleDo) Attrs(attrs ...field.AssignExpr) *tPmModuleDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tPmModuleDo) Assign(attrs ...field.AssignExpr) *tPmModuleDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tPmModuleDo) Joins(field field.RelationField) *tPmModuleDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tPmModuleDo) Preload(field field.RelationField) *tPmModuleDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tPmModuleDo) FirstOrInit() (*model.TPmModule, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmModule), nil
	}
}

func (t tPmModuleDo) FirstOrCreate() (*model.TPmModule, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TPmModule), nil
	}
}

func (t tPmModuleDo) FindByPage(offset int, limit int) (result []*model.TPmModule, count int64, err error) {
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

func (t tPmModuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tPmModuleDo) withDO(do gen.Dao) *tPmModuleDo {
	t.DO = *do.(*gen.DO)
	return t
}
