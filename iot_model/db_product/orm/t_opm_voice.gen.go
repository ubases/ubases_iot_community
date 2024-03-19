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

func newTOpmVoice(db *gorm.DB) tOpmVoice {
	_tOpmVoice := tOpmVoice{}

	_tOpmVoice.tOpmVoiceDo.UseDB(db)
	_tOpmVoice.tOpmVoiceDo.UseModel(&model.TOpmVoice{})

	tableName := _tOpmVoice.tOpmVoiceDo.TableName()
	_tOpmVoice.ALL = field.NewField(tableName, "*")
	_tOpmVoice.Id = field.NewInt64(tableName, "id")
	_tOpmVoice.VoiceNo = field.NewString(tableName, "voice_no")
	_tOpmVoice.VoiceName = field.NewString(tableName, "voice_name")
	_tOpmVoice.VoiceCategory = field.NewString(tableName, "voice_category")
	_tOpmVoice.VoiceDoc = field.NewString(tableName, "voice_doc")
	_tOpmVoice.VoiceLogo = field.NewString(tableName, "voice_logo")
	_tOpmVoice.VoiceEnable = field.NewInt32(tableName, "voice_enable")
	_tOpmVoice.VoiceDesc = field.NewString(tableName, "voice_desc")
	_tOpmVoice.VoiceEnName = field.NewString(tableName, "voice_en_name")

	_tOpmVoice.fillFieldMap()

	return _tOpmVoice
}

type tOpmVoice struct {
	tOpmVoiceDo tOpmVoiceDo

	ALL           field.Field
	Id            field.Int64
	VoiceNo       field.String
	VoiceName     field.String
	VoiceCategory field.String
	VoiceDoc      field.String
	VoiceLogo     field.String
	VoiceEnable   field.Int32
	VoiceDesc     field.String
	VoiceEnName   field.String

	fieldMap map[string]field.Expr
}

func (t tOpmVoice) Table(newTableName string) *tOpmVoice {
	t.tOpmVoiceDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOpmVoice) As(alias string) *tOpmVoice {
	t.tOpmVoiceDo.DO = *(t.tOpmVoiceDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOpmVoice) updateTableName(table string) *tOpmVoice {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.VoiceNo = field.NewString(table, "voice_no")
	t.VoiceName = field.NewString(table, "voice_name")
	t.VoiceCategory = field.NewString(table, "voice_category")
	t.VoiceDoc = field.NewString(table, "voice_doc")
	t.VoiceLogo = field.NewString(table, "voice_logo")
	t.VoiceEnable = field.NewInt32(table, "voice_enable")
	t.VoiceDesc = field.NewString(table, "voice_desc")
	t.VoiceEnName = field.NewString(table, "voice_en_name")

	t.fillFieldMap()

	return t
}

func (t *tOpmVoice) WithContext(ctx context.Context) *tOpmVoiceDo {
	return t.tOpmVoiceDo.WithContext(ctx)
}

func (t tOpmVoice) TableName() string { return t.tOpmVoiceDo.TableName() }

func (t tOpmVoice) Alias() string { return t.tOpmVoiceDo.Alias() }

func (t *tOpmVoice) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOpmVoice) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["id"] = t.Id
	t.fieldMap["voice_no"] = t.VoiceNo
	t.fieldMap["voice_name"] = t.VoiceName
	t.fieldMap["voice_category"] = t.VoiceCategory
	t.fieldMap["voice_doc"] = t.VoiceDoc
	t.fieldMap["voice_logo"] = t.VoiceLogo
	t.fieldMap["voice_enable"] = t.VoiceEnable
	t.fieldMap["voice_desc"] = t.VoiceDesc
	t.fieldMap["voice_en_name"] = t.VoiceEnName
}

func (t tOpmVoice) clone(db *gorm.DB) tOpmVoice {
	t.tOpmVoiceDo.ReplaceDB(db)
	return t
}

type tOpmVoiceDo struct{ gen.DO }

func (t tOpmVoiceDo) Debug() *tOpmVoiceDo {
	return t.withDO(t.DO.Debug())
}

func (t tOpmVoiceDo) WithContext(ctx context.Context) *tOpmVoiceDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOpmVoiceDo) Clauses(conds ...clause.Expression) *tOpmVoiceDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOpmVoiceDo) Returning(value interface{}, columns ...string) *tOpmVoiceDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOpmVoiceDo) Not(conds ...gen.Condition) *tOpmVoiceDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOpmVoiceDo) Or(conds ...gen.Condition) *tOpmVoiceDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOpmVoiceDo) Select(conds ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOpmVoiceDo) Where(conds ...gen.Condition) *tOpmVoiceDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOpmVoiceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOpmVoiceDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOpmVoiceDo) Order(conds ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOpmVoiceDo) Distinct(cols ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOpmVoiceDo) Omit(cols ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOpmVoiceDo) Join(table schema.Tabler, on ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOpmVoiceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOpmVoiceDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOpmVoiceDo) Group(cols ...field.Expr) *tOpmVoiceDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOpmVoiceDo) Having(conds ...gen.Condition) *tOpmVoiceDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOpmVoiceDo) Limit(limit int) *tOpmVoiceDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOpmVoiceDo) Offset(offset int) *tOpmVoiceDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOpmVoiceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOpmVoiceDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOpmVoiceDo) Unscoped() *tOpmVoiceDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOpmVoiceDo) Create(values ...*model.TOpmVoice) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOpmVoiceDo) CreateInBatches(values []*model.TOpmVoice, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOpmVoiceDo) Save(values ...*model.TOpmVoice) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOpmVoiceDo) First() (*model.TOpmVoice, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmVoice), nil
	}
}

func (t tOpmVoiceDo) Take() (*model.TOpmVoice, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmVoice), nil
	}
}

func (t tOpmVoiceDo) Last() (*model.TOpmVoice, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmVoice), nil
	}
}

func (t tOpmVoiceDo) Find() ([]*model.TOpmVoice, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOpmVoice), err
}

func (t tOpmVoiceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOpmVoice, err error) {
	buf := make([]*model.TOpmVoice, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOpmVoiceDo) FindInBatches(result *[]*model.TOpmVoice, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOpmVoiceDo) Attrs(attrs ...field.AssignExpr) *tOpmVoiceDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOpmVoiceDo) Assign(attrs ...field.AssignExpr) *tOpmVoiceDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOpmVoiceDo) Joins(field field.RelationField) *tOpmVoiceDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOpmVoiceDo) Preload(field field.RelationField) *tOpmVoiceDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOpmVoiceDo) FirstOrInit() (*model.TOpmVoice, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmVoice), nil
	}
}

func (t tOpmVoiceDo) FirstOrCreate() (*model.TOpmVoice, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOpmVoice), nil
	}
}

func (t tOpmVoiceDo) FindByPage(offset int, limit int) (result []*model.TOpmVoice, count int64, err error) {
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

func (t tOpmVoiceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOpmVoiceDo) withDO(do gen.Dao) *tOpmVoiceDo {
	t.DO = *do.(*gen.DO)
	return t
}
