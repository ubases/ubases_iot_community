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

func newTOemAppAndroidCert(db *gorm.DB) tOemAppAndroidCert {
	_tOemAppAndroidCert := tOemAppAndroidCert{}

	_tOemAppAndroidCert.tOemAppAndroidCertDo.UseDB(db)
	_tOemAppAndroidCert.tOemAppAndroidCertDo.UseModel(&model.TOemAppAndroidCert{})

	tableName := _tOemAppAndroidCert.tOemAppAndroidCertDo.TableName()
	_tOemAppAndroidCert.ALL = field.NewField(tableName, "*")
	_tOemAppAndroidCert.Id = field.NewInt64(tableName, "id")
	_tOemAppAndroidCert.AppId = field.NewInt64(tableName, "app_id")
	_tOemAppAndroidCert.Version = field.NewString(tableName, "version")
	_tOemAppAndroidCert.Resign = field.NewInt32(tableName, "resign")
	_tOemAppAndroidCert.CertSha256 = field.NewString(tableName, "cert_sha256")
	_tOemAppAndroidCert.Keypass = field.NewString(tableName, "keypass")
	_tOemAppAndroidCert.Storepass = field.NewString(tableName, "storepass")
	_tOemAppAndroidCert.AliasKeytool = field.NewString(tableName, "alias_keytool")
	_tOemAppAndroidCert.Keystore = field.NewString(tableName, "keystore")
	_tOemAppAndroidCert.KsMd5 = field.NewString(tableName, "ks_md5")
	_tOemAppAndroidCert.KsSha1 = field.NewString(tableName, "ks_sha1")
	_tOemAppAndroidCert.KsSha256 = field.NewString(tableName, "ks_sha256")
	_tOemAppAndroidCert.HwSignCert = field.NewString(tableName, "hw_sign_cert")
	_tOemAppAndroidCert.GoogleSignCert = field.NewString(tableName, "google_sign_cert")

	_tOemAppAndroidCert.fillFieldMap()

	return _tOemAppAndroidCert
}

type tOemAppAndroidCert struct {
	tOemAppAndroidCertDo tOemAppAndroidCertDo

	ALL            field.Field
	Id             field.Int64
	AppId          field.Int64
	Version        field.String
	Resign         field.Int32
	CertSha256     field.String
	Keypass        field.String
	Storepass      field.String
	AliasKeytool   field.String
	Keystore       field.String
	KsMd5          field.String
	KsSha1         field.String
	KsSha256       field.String
	HwSignCert     field.String
	GoogleSignCert field.String

	fieldMap map[string]field.Expr
}

func (t tOemAppAndroidCert) Table(newTableName string) *tOemAppAndroidCert {
	t.tOemAppAndroidCertDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOemAppAndroidCert) As(alias string) *tOemAppAndroidCert {
	t.tOemAppAndroidCertDo.DO = *(t.tOemAppAndroidCertDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOemAppAndroidCert) updateTableName(table string) *tOemAppAndroidCert {
	t.ALL = field.NewField(table, "*")
	t.Id = field.NewInt64(table, "id")
	t.AppId = field.NewInt64(table, "app_id")
	t.Version = field.NewString(table, "version")
	t.Resign = field.NewInt32(table, "resign")
	t.CertSha256 = field.NewString(table, "cert_sha256")
	t.Keypass = field.NewString(table, "keypass")
	t.Storepass = field.NewString(table, "storepass")
	t.AliasKeytool = field.NewString(table, "alias_keytool")
	t.Keystore = field.NewString(table, "keystore")
	t.KsMd5 = field.NewString(table, "ks_md5")
	t.KsSha1 = field.NewString(table, "ks_sha1")
	t.KsSha256 = field.NewString(table, "ks_sha256")
	t.HwSignCert = field.NewString(table, "hw_sign_cert")
	t.GoogleSignCert = field.NewString(table, "google_sign_cert")

	t.fillFieldMap()

	return t
}

func (t *tOemAppAndroidCert) WithContext(ctx context.Context) *tOemAppAndroidCertDo {
	return t.tOemAppAndroidCertDo.WithContext(ctx)
}

func (t tOemAppAndroidCert) TableName() string { return t.tOemAppAndroidCertDo.TableName() }

func (t *tOemAppAndroidCert) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOemAppAndroidCert) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 14)
	t.fieldMap["id"] = t.Id
	t.fieldMap["app_id"] = t.AppId
	t.fieldMap["version"] = t.Version
	t.fieldMap["resign"] = t.Resign
	t.fieldMap["cert_sha256"] = t.CertSha256
	t.fieldMap["keypass"] = t.Keypass
	t.fieldMap["storepass"] = t.Storepass
	t.fieldMap["alias_keytool"] = t.AliasKeytool
	t.fieldMap["keystore"] = t.Keystore
	t.fieldMap["ks_md5"] = t.KsMd5
	t.fieldMap["ks_sha1"] = t.KsSha1
	t.fieldMap["ks_sha256"] = t.KsSha256
	t.fieldMap["hw_sign_cert"] = t.HwSignCert
	t.fieldMap["google_sign_cert"] = t.GoogleSignCert
}

func (t tOemAppAndroidCert) clone(db *gorm.DB) tOemAppAndroidCert {
	t.tOemAppAndroidCertDo.ReplaceDB(db)
	return t
}

type tOemAppAndroidCertDo struct{ gen.DO }

func (t tOemAppAndroidCertDo) Debug() *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Debug())
}

func (t tOemAppAndroidCertDo) WithContext(ctx context.Context) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOemAppAndroidCertDo) Clauses(conds ...clause.Expression) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOemAppAndroidCertDo) Returning(value interface{}, columns ...string) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOemAppAndroidCertDo) Not(conds ...gen.Condition) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOemAppAndroidCertDo) Or(conds ...gen.Condition) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOemAppAndroidCertDo) Select(conds ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOemAppAndroidCertDo) Where(conds ...gen.Condition) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOemAppAndroidCertDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tOemAppAndroidCertDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOemAppAndroidCertDo) Order(conds ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOemAppAndroidCertDo) Distinct(cols ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOemAppAndroidCertDo) Omit(cols ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOemAppAndroidCertDo) Join(table schema.Tabler, on ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOemAppAndroidCertDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOemAppAndroidCertDo) RightJoin(table schema.Tabler, on ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOemAppAndroidCertDo) Group(cols ...field.Expr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOemAppAndroidCertDo) Having(conds ...gen.Condition) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOemAppAndroidCertDo) Limit(limit int) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOemAppAndroidCertDo) Offset(offset int) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOemAppAndroidCertDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOemAppAndroidCertDo) Unscoped() *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOemAppAndroidCertDo) Create(values ...*model.TOemAppAndroidCert) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOemAppAndroidCertDo) CreateInBatches(values []*model.TOemAppAndroidCert, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOemAppAndroidCertDo) Save(values ...*model.TOemAppAndroidCert) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOemAppAndroidCertDo) First() (*model.TOemAppAndroidCert, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppAndroidCert), nil
	}
}

func (t tOemAppAndroidCertDo) Take() (*model.TOemAppAndroidCert, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppAndroidCert), nil
	}
}

func (t tOemAppAndroidCertDo) Last() (*model.TOemAppAndroidCert, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppAndroidCert), nil
	}
}

func (t tOemAppAndroidCertDo) Find() ([]*model.TOemAppAndroidCert, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOemAppAndroidCert), err
}

func (t tOemAppAndroidCertDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOemAppAndroidCert, err error) {
	buf := make([]*model.TOemAppAndroidCert, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOemAppAndroidCertDo) FindInBatches(result *[]*model.TOemAppAndroidCert, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOemAppAndroidCertDo) Attrs(attrs ...field.AssignExpr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOemAppAndroidCertDo) Assign(attrs ...field.AssignExpr) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOemAppAndroidCertDo) Joins(field field.RelationField) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Joins(field))
}

func (t tOemAppAndroidCertDo) Preload(field field.RelationField) *tOemAppAndroidCertDo {
	return t.withDO(t.DO.Preload(field))
}

func (t tOemAppAndroidCertDo) FirstOrInit() (*model.TOemAppAndroidCert, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppAndroidCert), nil
	}
}

func (t tOemAppAndroidCertDo) FirstOrCreate() (*model.TOemAppAndroidCert, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOemAppAndroidCert), nil
	}
}

func (t tOemAppAndroidCertDo) FindByPage(offset int, limit int) (result []*model.TOemAppAndroidCert, count int64, err error) {
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

func (t tOemAppAndroidCertDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t *tOemAppAndroidCertDo) withDO(do gen.Dao) *tOemAppAndroidCertDo {
	t.DO = *do.(*gen.DO)
	return t
}
