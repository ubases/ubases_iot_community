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

func newTOpmNetworkGuideStep(db *gorm.DB) tOpmNetworkGuideStep {
	_tOpmNetworkGuideStep := tOpmNetworkGuideStep{}

	_tOpmNetworkGuideStep.tOpmNetworkGuideStepDo.UseDB(db)
	_tOpmNetworkGuideStep.tOpmNetworkGuideStepDo.UseModel(&model.TOpmNetworkGuideStep{})

	tableName := _tOpmNetworkGuideStep.tOpmNetworkGuideStepDo.TableName()
	_tOpmNetworkGuideStep.ALL = field.NewField(tableName, "*")
	_tOpmNetworkGuideStep.Id = field.NewInt64(tableName, "id")
	_tOpmNetworkGuideStep.NetworkGuideId = field.NewInt64(tableName, "network_guide_id")
	_tOpmNetworkGuideStep.ProductId = field.NewInt64(tableName, "product_id")
	_tOpmNetworkGuideStep.Instruction = field.NewString(tableName, "instruction")
	_tOpmNetworkGuideStep.InstructionEn = field.NewString(tableName, "instruction_en")
	_tOpmNetworkGuideStep.ImageUrl = field.NewString(tableName, "image_url")
	_tOpmNetworkGuideStep.VideoUrl = field.NewString(tableName, "video_url")
	_tOpmNetworkGuideStep.Sort = field.NewInt32(tableName, "sort")
	_tOpmNetworkGuideStep.CreatedBy = field.NewInt64(tableName, "created_by")
	_tOpmNetworkGuideStep.CreatedAt = field.NewTime(tableName, "created_at")
	_tOpmNetworkGuideStep.UpdatedBy = field.NewInt64(tableName, "updated_by")
	_tOpmNetworkGuideStep.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tOpmNetworkGuideStep.TenantId = field.NewString(tableName, "tenant_id")

	_tOpmNetworkGuideStep.fillFieldMap()

	return _tOpmNetworkGuideStep
}

type tOpmNetworkGuideStep struct {
	tOpmNetworkGuideStepDo tOpmNetworkGuideStepDo

	ALL            field.Field
	Id             field.Int64
	NetworkGuideId field.Int64
	ProductId      field.Int64
	Instruction    field.String
	InstructionEn  field.String
	ImageUrl       field.String
	VideoUrl       field.String
	Sort           field.Int32
	CreatedBy      field.Int64
	CreatedAt      field.Time
	UpdatedBy      field.Int64
	UpdatedAt      field.Time
	TenantId       field.String

	fieldMap map[string]field.Expr
}

func (t tOpmNetworkGuideStep) Table(newTableName string) *tOpmNetworkGuideStep {
	t.tOpmNetworkGuideStepDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOpmNetworkGuideStep) As(alias string) *tOpmNetworkGuideStep {
	t.tOpmNetworkGuideStepDo.DO = *(t.tOpmNetworkGuideStepDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOpmNetworkGuideStep) updateTableName(table string) *tOpmNetworkGuideStep {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.NetworkGuideId = field.NewInt64(table, "network_guide_id")
	t.ProductId = field.NewInt64(table, "product_id")
	t.Instruction = field.NewString(table, "instruction")
	t.InstructionEn = field.NewString(table, "instruction_en")
	t.ImageUrl = field.NewString(table, "image_url")
	t.VideoUrl = field.NewString(table, "video_url")
	t.Sort = field.NewInt32(table, "sort")
	t.CreatedBy = field.NewInt64(table, "created_by")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedBy = field.NewInt64(table, "updated_by")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.TenantId = field.NewString(table, "tenant_id")

	t.fillFieldMap()

	return t
}

func (t *tOpmNetworkGuideStep) WithContext(ctx context.Context) *tOpmNetworkGuideStepDo {
	return t.tOpmNetworkGuideStepDo.WithContext(ctx)
}

func (t tOpmNetworkGuideStep) TableName() string { return t.tOpmNetworkGuideStepDo.TableName() }

func (t tOpmNetworkGuideStep) Alias() string { return t.tOpmNetworkGuideStepDo.Alias() }

func (t *tOpmNetworkGuideStep) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOpmNetworkGuideStep) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 13)
	t.fieldMap["id"] = t.Id
	t.fieldMap["network_guide_id"] = t.NetworkGuideId
	t.fieldMap["product_id"] = t.ProductId
	t.fieldMap["instruction"] = t.Instruction
	t.fieldMap["instruction_en"] = t.InstructionEn
	t.fieldMap["image_url"] = t.ImageUrl
	t.fieldMap["video_url"] = t.VideoUrl
	t.fieldMap["sort"] = t.Sort
	t.fieldMap["created_by"] = t.CreatedBy
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_by"] = t.UpdatedBy
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["tenant_id"] = t.TenantId
}

func (t tOpmNetworkGuideStep) clone(db *gorm.DB) tOpmNetworkGuideStep {
	t.tOpmNetworkGuideStepDo.ReplaceDB(db)
	return t
}

type tOpmNetworkGuideStepDo struct{ gen.DO }

func (t tOpmNetworkGuideStepDo) Debug() *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Debug())
}

func (t tOpmNetworkGuideStepDo) WithContext(ctx context.Context) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOpmNetworkGuideStepDo) Clauses(conds ...clause.Expression) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOpmNetworkGuideStepDo) Returning(value interface{}, columns ...string) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOpmNetworkGuideStepDo) Not(conds ...gen.Condition) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOpmNetworkGuideStepDo) Or(conds ...gen.Condition) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOpmNetworkGuideStepDo) Select(conds ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOpmNetworkGuideStepDo) Where(conds ...gen.Condition) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOpmNetworkGuideStepDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOpmNetworkGuideStepDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOpmNetworkGuideStepDo) Order(conds ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOpmNetworkGuideStepDo) Distinct(cols ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOpmNetworkGuideStepDo) Omit(cols ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOpmNetworkGuideStepDo) Join(table schema.Tabler, on ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOpmNetworkGuideStepDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOpmNetworkGuideStepDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOpmNetworkGuideStepDo) Group(cols ...field.Expr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOpmNetworkGuideStepDo) Having(conds ...gen.Condition) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOpmNetworkGuideStepDo) Limit(limit int) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOpmNetworkGuideStepDo) Offset(offset int) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOpmNetworkGuideStepDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOpmNetworkGuideStepDo) Unscoped() *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOpmNetworkGuideStepDo) Create(values ...*model.TOpmNetworkGuideStep) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOpmNetworkGuideStepDo) CreateInBatches(values []*model.TOpmNetworkGuideStep, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOpmNetworkGuideStepDo) Save(values ...*model.TOpmNetworkGuideStep) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOpmNetworkGuideStepDo) First() (*model.TOpmNetworkGuideStep, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmNetworkGuideStep), nil
	}
}

func (t tOpmNetworkGuideStepDo) Take() (*model.TOpmNetworkGuideStep, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmNetworkGuideStep), nil
	}
}

func (t tOpmNetworkGuideStepDo) Last() (*model.TOpmNetworkGuideStep, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmNetworkGuideStep), nil
	}
}

func (t tOpmNetworkGuideStepDo) Find() ([]*model.TOpmNetworkGuideStep, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOpmNetworkGuideStep), err
}

func (t tOpmNetworkGuideStepDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOpmNetworkGuideStep, err error) {
	buf := make([]*model.TOpmNetworkGuideStep, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOpmNetworkGuideStepDo) FindInBatches(result *[]*model.TOpmNetworkGuideStep, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOpmNetworkGuideStepDo) Attrs(attrs ...field.AssignExpr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOpmNetworkGuideStepDo) Assign(attrs ...field.AssignExpr) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOpmNetworkGuideStepDo) Joins(field field.RelationField) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOpmNetworkGuideStepDo) Preload(field field.RelationField) *tOpmNetworkGuideStepDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOpmNetworkGuideStepDo) FirstOrInit() (*model.TOpmNetworkGuideStep, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmNetworkGuideStep), nil
	}
}

func (t tOpmNetworkGuideStepDo) FirstOrCreate() (*model.TOpmNetworkGuideStep, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmNetworkGuideStep), nil
	}
}

func (t tOpmNetworkGuideStepDo) FindByPage(offset int, limit int) (result []*model.TOpmNetworkGuideStep, count int64, err error) {
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

func (t tOpmNetworkGuideStepDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOpmNetworkGuideStepDo) withDO(do gen.Dao) *tOpmNetworkGuideStepDo {
	t.DO = *do.(*gen.DO)
	return t
}
