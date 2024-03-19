// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package orm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"cloud_platform/iot_model/db_app_oem/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newTOemAppDocDir(db *gorm.DB) tOemAppDocDir {
	_tOemAppDocDir := tOemAppDocDir{}

	_tOemAppDocDir.tOemAppDocDirDo.UseDB(db)
	_tOemAppDocDir.tOemAppDocDirDo.UseModel(&model.TOemAppDocDir{})

	tableName := _tOemAppDocDir.tOemAppDocDirDo.TableName()
	_tOemAppDocDir.ALL = field.NewField(tableName, "*")
	_tOemAppDocDir.Id = field.NewInt64(tableName, "id")
	_tOemAppDocDir.DocId = field.NewInt64(tableName, "doc_id")
	_tOemAppDocDir.ParentId = field.NewInt64(tableName, "parent_id")
	_tOemAppDocDir.DirName = field.NewString(tableName, "dir_name")
	_tOemAppDocDir.DirImg = field.NewString(tableName, "dir_img")
	_tOemAppDocDir.Sort = field.NewInt32(tableName, "sort")

	_tOemAppDocDir.fillFieldMap()

	return _tOemAppDocDir
}

type tOemAppDocDir struct {
	tOemAppDocDirDo tOemAppDocDirDo

	ALL      field.Field
	Id       field.Int64
	DocId    field.Int64
	ParentId field.Int64
	DirName  field.String
	DirImg   field.String
	Sort     field.Int32

	fieldMap map[string]field.Expr
}

func (t tOemAppDocDir) Table(newTableName string) *tOemAppDocDir {
	t.tOemAppDocDirDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOemAppDocDir) As(alias string) *tOemAppDocDir {
	t.tOemAppDocDirDo.DO = *(t.tOemAppDocDirDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOemAppDocDir) updateTableName(table string) *tOemAppDocDir {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.DocId = field.NewInt64(table, "doc_id")
	t.ParentId = field.NewInt64(table, "parent_id")
	t.DirName = field.NewString(table, "dir_name")
	t.DirImg = field.NewString(table, "dir_img")
	t.Sort = field.NewInt32(table, "sort")

	t.fillFieldMap()

	return t
}

func (t *tOemAppDocDir) WithContext(ctx context.Context) *tOemAppDocDirDo {
	return t.tOemAppDocDirDo.WithContext(ctx)
}

func (t tOemAppDocDir) TableName() string { return t.tOemAppDocDirDo.TableName() }

func (t *tOemAppDocDir) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOemAppDocDir) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["id"] = t.Id
	t.fieldMap["doc_id"] = t.DocId
	t.fieldMap["parent_id"] = t.ParentId
	t.fieldMap["dir_name"] = t.DirName
	t.fieldMap["dir_img"] = t.DirImg
	t.fieldMap["sort"] = t.Sort
}

func (t tOemAppDocDir) clone(db *gorm.DB) tOemAppDocDir {
	t.tOemAppDocDirDo.ReplaceDB(db)
	return t
}

type tOemAppDocDirDo struct{ gen.DO }

func (t tOemAppDocDirDo) Debug() *tOemAppDocDirDo {
	return t.withDO(t.DO.Debug())
}

func (t tOemAppDocDirDo) WithContext(ctx context.Context) *tOemAppDocDirDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOemAppDocDirDo) Clauses(conds ...clause.Expression) *tOemAppDocDirDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOemAppDocDirDo) Returning(value interface{}, columns ...string) *tOemAppDocDirDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOemAppDocDirDo) Not(conds ...gen.Condition) *tOemAppDocDirDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOemAppDocDirDo) Or(conds ...gen.Condition) *tOemAppDocDirDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOemAppDocDirDo) Select(conds ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOemAppDocDirDo) Where(conds ...gen.Condition) *tOemAppDocDirDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOemAppDocDirDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOemAppDocDirDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOemAppDocDirDo) Order(conds ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOemAppDocDirDo) Distinct(cols ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOemAppDocDirDo) Omit(cols ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOemAppDocDirDo) Join(table schema.Tabler, on ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOemAppDocDirDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOemAppDocDirDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOemAppDocDirDo) Group(cols ...field.Expr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOemAppDocDirDo) Having(conds ...gen.Condition) *tOemAppDocDirDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOemAppDocDirDo) Limit(limit int) *tOemAppDocDirDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOemAppDocDirDo) Offset(offset int) *tOemAppDocDirDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOemAppDocDirDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOemAppDocDirDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOemAppDocDirDo) Unscoped() *tOemAppDocDirDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOemAppDocDirDo) Create(values ...*model.TOemAppDocDir) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOemAppDocDirDo) CreateInBatches(values []*model.TOemAppDocDir, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOemAppDocDirDo) Save(values ...*model.TOemAppDocDir) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOemAppDocDirDo) First() (*model.TOemAppDocDir, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocDir), nil
	}
}

func (t tOemAppDocDirDo) Take() (*model.TOemAppDocDir, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocDir), nil
	}
}

func (t tOemAppDocDirDo) Last() (*model.TOemAppDocDir, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocDir), nil
	}
}

func (t tOemAppDocDirDo) Find() ([]*model.TOemAppDocDir, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOemAppDocDir), err
}

func (t tOemAppDocDirDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOemAppDocDir, err error) {
	buf := make([]*model.TOemAppDocDir, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOemAppDocDirDo) FindInBatches(result *[]*model.TOemAppDocDir, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOemAppDocDirDo) Attrs(attrs ...field.AssignExpr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOemAppDocDirDo) Assign(attrs ...field.AssignExpr) *tOemAppDocDirDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOemAppDocDirDo) Joins(field field.RelationField) *tOemAppDocDirDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOemAppDocDirDo) Preload(field field.RelationField) *tOemAppDocDirDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOemAppDocDirDo) FirstOrInit() (*model.TOemAppDocDir, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocDir), nil
	}
}

func (t tOemAppDocDirDo) FirstOrCreate() (*model.TOemAppDocDir, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppDocDir), nil
	}
}

func (t tOemAppDocDirDo) FindByPage(offset int, limit int) (result []*model.TOemAppDocDir, count int64, err error) {
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

func (t tOemAppDocDirDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOemAppDocDirDo) withDO(do gen.Dao) *tOemAppDocDirDo {
	t.DO = *do.(*gen.DO)
	return t
}
