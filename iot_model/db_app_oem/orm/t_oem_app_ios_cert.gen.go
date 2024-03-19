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

func newTOemAppIosCert(db *gorm.DB) tOemAppIosCert {
	_tOemAppIosCert := tOemAppIosCert{}

	_tOemAppIosCert.tOemAppIosCertDo.UseDB(db)
	_tOemAppIosCert.tOemAppIosCertDo.UseModel(&model.TOemAppIosCert{})

	tableName := _tOemAppIosCert.tOemAppIosCertDo.TableName()
	_tOemAppIosCert.ALL = field.NewField(tableName, "*")
	_tOemAppIosCert.Id = field.NewInt64(tableName, "id")
	_tOemAppIosCert.AppId = field.NewInt64(tableName, "app_id")
	_tOemAppIosCert.Version = field.NewString(tableName, "version")
	_tOemAppIosCert.DistProvision = field.NewString(tableName, "dist_provision")
	_tOemAppIosCert.DistCert = field.NewString(tableName, "dist_cert")
	_tOemAppIosCert.DistCertSecret = field.NewString(tableName, "dist_cert_secret")
	_tOemAppIosCert.DistCertOfficial = field.NewString(tableName, "dist_cert_official")

	_tOemAppIosCert.fillFieldMap()

	return _tOemAppIosCert
}

type tOemAppIosCert struct {
	tOemAppIosCertDo tOemAppIosCertDo

	ALL              field.Field
	Id               field.Int64
	AppId            field.Int64
	Version          field.String
	DistProvision    field.String
	DistCert         field.String
	DistCertSecret   field.String
	DistCertOfficial field.String

	fieldMap map[string]field.Expr
}

func (t tOemAppIosCert) Table(newTableName string) *tOemAppIosCert {
	t.tOemAppIosCertDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOemAppIosCert) As(alias string) *tOemAppIosCert {
	t.tOemAppIosCertDo.DO = *(t.tOemAppIosCertDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOemAppIosCert) updateTableName(table string) *tOemAppIosCert {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.AppId = field.NewInt64(table, "app_id")
	t.Version = field.NewString(table, "version")
	t.DistProvision = field.NewString(table, "dist_provision")
	t.DistCert = field.NewString(table, "dist_cert")
	t.DistCertSecret = field.NewString(table, "dist_cert_secret")
	t.DistCertOfficial = field.NewString(table, "dist_cert_official")

	t.fillFieldMap()

	return t
}

func (t *tOemAppIosCert) WithContext(ctx context.Context) *tOemAppIosCertDo {
	return t.tOemAppIosCertDo.WithContext(ctx)
}

func (t tOemAppIosCert) TableName() string { return t.tOemAppIosCertDo.TableName() }

func (t *tOemAppIosCert) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOemAppIosCert) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.Id
	t.fieldMap["app_id"] = t.AppId
	t.fieldMap["version"] = t.Version
	t.fieldMap["dist_provision"] = t.DistProvision
	t.fieldMap["dist_cert"] = t.DistCert
	t.fieldMap["dist_cert_secret"] = t.DistCertSecret
	t.fieldMap["dist_cert_official"] = t.DistCertOfficial
}

func (t tOemAppIosCert) clone(db *gorm.DB) tOemAppIosCert {
	t.tOemAppIosCertDo.ReplaceDB(db)
	return t
}

type tOemAppIosCertDo struct{ gen.DO }

func (t tOemAppIosCertDo) Debug() *tOemAppIosCertDo {
	return t.withDO(t.DO.Debug())
}

func (t tOemAppIosCertDo) WithContext(ctx context.Context) *tOemAppIosCertDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOemAppIosCertDo) Clauses(conds ...clause.Expression) *tOemAppIosCertDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOemAppIosCertDo) Returning(value interface{}, columns ...string) *tOemAppIosCertDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOemAppIosCertDo) Not(conds ...gen.Condition) *tOemAppIosCertDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOemAppIosCertDo) Or(conds ...gen.Condition) *tOemAppIosCertDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOemAppIosCertDo) Select(conds ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOemAppIosCertDo) Where(conds ...gen.Condition) *tOemAppIosCertDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOemAppIosCertDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOemAppIosCertDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOemAppIosCertDo) Order(conds ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOemAppIosCertDo) Distinct(cols ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOemAppIosCertDo) Omit(cols ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOemAppIosCertDo) Join(table schema.Tabler, on ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOemAppIosCertDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOemAppIosCertDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOemAppIosCertDo) Group(cols ...field.Expr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOemAppIosCertDo) Having(conds ...gen.Condition) *tOemAppIosCertDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOemAppIosCertDo) Limit(limit int) *tOemAppIosCertDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOemAppIosCertDo) Offset(offset int) *tOemAppIosCertDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOemAppIosCertDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOemAppIosCertDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOemAppIosCertDo) Unscoped() *tOemAppIosCertDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOemAppIosCertDo) Create(values ...*model.TOemAppIosCert) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOemAppIosCertDo) CreateInBatches(values []*model.TOemAppIosCert, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOemAppIosCertDo) Save(values ...*model.TOemAppIosCert) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOemAppIosCertDo) First() (*model.TOemAppIosCert, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppIosCert), nil
	}
}

func (t tOemAppIosCertDo) Take() (*model.TOemAppIosCert, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppIosCert), nil
	}
}

func (t tOemAppIosCertDo) Last() (*model.TOemAppIosCert, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppIosCert), nil
	}
}

func (t tOemAppIosCertDo) Find() ([]*model.TOemAppIosCert, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOemAppIosCert), err
}

func (t tOemAppIosCertDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOemAppIosCert, err error) {
	buf := make([]*model.TOemAppIosCert, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOemAppIosCertDo) FindInBatches(result *[]*model.TOemAppIosCert, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOemAppIosCertDo) Attrs(attrs ...field.AssignExpr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOemAppIosCertDo) Assign(attrs ...field.AssignExpr) *tOemAppIosCertDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOemAppIosCertDo) Joins(field field.RelationField) *tOemAppIosCertDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOemAppIosCertDo) Preload(field field.RelationField) *tOemAppIosCertDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOemAppIosCertDo) FirstOrInit() (*model.TOemAppIosCert, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppIosCert), nil
	}
}

func (t tOemAppIosCertDo) FirstOrCreate() (*model.TOemAppIosCert, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppIosCert), nil
	}
}

func (t tOemAppIosCertDo) FindByPage(offset int, limit int) (result []*model.TOemAppIosCert, count int64, err error) {
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

func (t tOemAppIosCertDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOemAppIosCertDo) withDO(do gen.Dao) *tOemAppIosCertDo {
	t.DO = *do.(*gen.DO)
	return t
}
