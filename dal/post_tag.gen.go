// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"github.com/go-sonic/sonic/model/entity"
)

func newPostTag(db *gorm.DB, opts ...gen.DOOption) postTag {
	_postTag := postTag{}

	_postTag.postTagDo.UseDB(db, opts...)
	_postTag.postTagDo.UseModel(&entity.PostTag{})

	tableName := _postTag.postTagDo.TableName()
	_postTag.ALL = field.NewAsterisk(tableName)
	_postTag.ID = field.NewInt32(tableName, "id")
	_postTag.CreateTime = field.NewTime(tableName, "create_time")
	_postTag.UpdateTime = field.NewTime(tableName, "update_time")
	_postTag.PostID = field.NewInt32(tableName, "post_id")
	_postTag.TagID = field.NewInt32(tableName, "tag_id")

	_postTag.fillFieldMap()

	return _postTag
}

type postTag struct {
	postTagDo postTagDo

	ALL        field.Asterisk
	ID         field.Int32
	CreateTime field.Time
	UpdateTime field.Time
	PostID     field.Int32
	TagID      field.Int32

	fieldMap map[string]field.Expr
}

func (p postTag) Table(newTableName string) *postTag {
	p.postTagDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p postTag) As(alias string) *postTag {
	p.postTagDo.DO = *(p.postTagDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *postTag) updateTableName(table string) *postTag {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.CreateTime = field.NewTime(table, "create_time")
	p.UpdateTime = field.NewTime(table, "update_time")
	p.PostID = field.NewInt32(table, "post_id")
	p.TagID = field.NewInt32(table, "tag_id")

	p.fillFieldMap()

	return p
}

func (p *postTag) WithContext(ctx context.Context) *postTagDo { return p.postTagDo.WithContext(ctx) }

func (p postTag) TableName() string { return p.postTagDo.TableName() }

func (p postTag) Alias() string { return p.postTagDo.Alias() }

func (p *postTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *postTag) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["id"] = p.ID
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["update_time"] = p.UpdateTime
	p.fieldMap["post_id"] = p.PostID
	p.fieldMap["tag_id"] = p.TagID
}

func (p postTag) clone(db *gorm.DB) postTag {
	p.postTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p postTag) replaceDB(db *gorm.DB) postTag {
	p.postTagDo.ReplaceDB(db)
	return p
}

type postTagDo struct{ gen.DO }

func (p postTagDo) Debug() *postTagDo {
	return p.withDO(p.DO.Debug())
}

func (p postTagDo) WithContext(ctx context.Context) *postTagDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postTagDo) ReadDB() *postTagDo {
	return p.Clauses(dbresolver.Read)
}

func (p postTagDo) WriteDB() *postTagDo {
	return p.Clauses(dbresolver.Write)
}

func (p postTagDo) Session(config *gorm.Session) *postTagDo {
	return p.withDO(p.DO.Session(config))
}

func (p postTagDo) Clauses(conds ...clause.Expression) *postTagDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postTagDo) Returning(value interface{}, columns ...string) *postTagDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postTagDo) Not(conds ...gen.Condition) *postTagDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postTagDo) Or(conds ...gen.Condition) *postTagDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postTagDo) Select(conds ...field.Expr) *postTagDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postTagDo) Where(conds ...gen.Condition) *postTagDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postTagDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *postTagDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p postTagDo) Order(conds ...field.Expr) *postTagDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postTagDo) Distinct(cols ...field.Expr) *postTagDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postTagDo) Omit(cols ...field.Expr) *postTagDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postTagDo) Join(table schema.Tabler, on ...field.Expr) *postTagDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) *postTagDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postTagDo) RightJoin(table schema.Tabler, on ...field.Expr) *postTagDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postTagDo) Group(cols ...field.Expr) *postTagDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postTagDo) Having(conds ...gen.Condition) *postTagDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postTagDo) Limit(limit int) *postTagDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postTagDo) Offset(offset int) *postTagDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *postTagDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postTagDo) Unscoped() *postTagDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postTagDo) Create(values ...*entity.PostTag) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postTagDo) CreateInBatches(values []*entity.PostTag, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postTagDo) Save(values ...*entity.PostTag) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postTagDo) First() (*entity.PostTag, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PostTag), nil
	}
}

func (p postTagDo) Take() (*entity.PostTag, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PostTag), nil
	}
}

func (p postTagDo) Last() (*entity.PostTag, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PostTag), nil
	}
}

func (p postTagDo) Find() ([]*entity.PostTag, error) {
	result, err := p.DO.Find()
	return result.([]*entity.PostTag), err
}

func (p postTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.PostTag, err error) {
	buf := make([]*entity.PostTag, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postTagDo) FindInBatches(result *[]*entity.PostTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postTagDo) Attrs(attrs ...field.AssignExpr) *postTagDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postTagDo) Assign(attrs ...field.AssignExpr) *postTagDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postTagDo) Joins(fields ...field.RelationField) *postTagDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postTagDo) Preload(fields ...field.RelationField) *postTagDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postTagDo) FirstOrInit() (*entity.PostTag, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PostTag), nil
	}
}

func (p postTagDo) FirstOrCreate() (*entity.PostTag, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PostTag), nil
	}
}

func (p postTagDo) FindByPage(offset int, limit int) (result []*entity.PostTag, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postTagDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postTagDo) Delete(models ...*entity.PostTag) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postTagDo) withDO(do gen.Dao) *postTagDo {
	p.DO = *do.(*gen.DO)
	return p
}
