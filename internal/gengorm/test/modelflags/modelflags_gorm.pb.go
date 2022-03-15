// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v0.0.0
// 	protoc          (unknown)
// source: modelflags/modelflags.proto

package modelflags

import (
	context "context"
	fmt "fmt"
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	gorm "gorm.io/gorm"
)

// ModelOptionModel is the GORM model for modelflags.ModelOption.
type ModelOptionModel struct {
}

// AsProto converts a ModelOptionModel to its protobuf representation.
func (m *ModelOptionModel) AsProto() (*ModelOption, error) {
	x := new(ModelOption)
	return x, nil
}

// AsModel converts a ModelOption to its GORM model.
func (x *ModelOption) AsModel() (*ModelOptionModel, error) {
	m := new(ModelOptionModel)
	return m, nil
}

// ValidateImpliesModelModel is the GORM model for modelflags.ValidateImpliesModel.
type ValidateImpliesModelModel struct {
}

// AsProto converts a ValidateImpliesModelModel to its protobuf representation.
func (m *ValidateImpliesModelModel) AsProto() (*ValidateImpliesModel, error) {
	x := new(ValidateImpliesModel)
	return x, nil
}

// AsModel converts a ValidateImpliesModel to its GORM model.
func (x *ValidateImpliesModel) AsModel() (*ValidateImpliesModelModel, error) {
	m := new(ValidateImpliesModelModel)
	return m, nil
}

// CRUDImpliesModelModel is the GORM model for modelflags.CRUDImpliesModel.
type CRUDImpliesModelModel struct {
	Uuid string `gorm:"primaryKey"`
}

// AsProto converts a CRUDImpliesModelModel to its protobuf representation.
func (m *CRUDImpliesModelModel) AsProto() (*CRUDImpliesModel, error) {
	x := new(CRUDImpliesModel)
	x.Uuid = m.Uuid
	return x, nil
}

// AsModel converts a CRUDImpliesModel to its GORM model.
func (x *CRUDImpliesModel) AsModel() (*CRUDImpliesModelModel, error) {
	m := new(CRUDImpliesModelModel)
	m.Uuid = x.Uuid
	return m, nil
}

type CRUDImpliesModelGetOption func(tx *gorm.DB) *gorm.DB
type CRUDImpliesModelListOption func(tx *gorm.DB) *gorm.DB

type CRUDImpliesModelWithDB struct {
	x  *CRUDImpliesModel
	db *gorm.DB
}

func (x *CRUDImpliesModel) WithDB(db *gorm.DB) CRUDImpliesModelWithDB {
	return CRUDImpliesModelWithDB{x: x, db: db}
}

func (c CRUDImpliesModelWithDB) Create(ctx context.Context) (*CRUDImpliesModel, error) {
	if c.x == nil {
		return nil, nil
	}
	m, err := c.x.AsModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	if err := db.Create(m).Error; err != nil {
		return nil, err
	}
	if y, err := m.AsProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CRUDImpliesModelWithDB) Get(ctx context.Context, opts ...CRUDImpliesModelGetOption) (*CRUDImpliesModel, error) {
	if c.x == nil {
		return nil, nil
	}
	var zero string
	if c.x.Uuid == zero {
		return nil, fmt.Errorf("empty primary key")
	}
	m, err := c.x.AsModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	out := CRUDImpliesModelModel{}
	if err := db.Where(m).First(&out).Error; err != nil {
		return nil, err
	}
	if y, err := out.AsProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CRUDImpliesModelWithDB) List(ctx context.Context, opts ...CRUDImpliesModelListOption) ([]*CRUDImpliesModel, error) {
	if c.x == nil {
		return nil, nil
	}
	db := c.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	var ms []CRUDImpliesModelModel
	if err := db.Find(&ms).Error; err != nil {
		return nil, err
	}
	xs := make([]*CRUDImpliesModel, 0, len(ms))
	for _, m := range ms {
		if x, err := m.AsProto(); err != nil {
			return nil, err
		} else {
			xs = append(xs, x)
		}
	}
	return xs, nil
}

func (c CRUDImpliesModelWithDB) Update(ctx context.Context) (*CRUDImpliesModel, error) {
	if c.x == nil {
		return nil, nil
	}
	m, err := c.x.AsModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	if err := db.Save(m).Error; err != nil {
		return nil, err
	}
	return c.Get(ctx)
}

func (c CRUDImpliesModelWithDB) Patch(ctx context.Context, mask *fieldmaskpb.FieldMask) error {
	if c.x == nil {
		return nil
	}
	if mask == nil {
		_, err := c.Update(ctx)
		return err
	}
	if !mask.IsValid(c.x) {
		return fmt.Errorf("invalid field mask")
	}
	paths := mask.Paths
	if len(paths) == 0 {
		_, err := c.Update(ctx)
		return err
	}
	var zero string
	if c.x.Uuid == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.x.AsModel()
	if err != nil {
		return err
	}
	target := CRUDImpliesModelModel{Uuid: m.Uuid}
	cols := LookupCRUDImpliesModelModelColumns(paths)
	db := c.db.WithContext(ctx)
	if err := db.Model(&target).Select(cols).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (c CRUDImpliesModelWithDB) Delete(ctx context.Context) error {
	if c.x == nil {
		return nil
	}
	var zero string
	if c.x.Uuid == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.x.AsModel()
	if err != nil {
		return err
	}
	db := c.db.WithContext(ctx)
	if err := db.Where(m).Delete(&CRUDImpliesModelModel{}).Error; err != nil {
		return err
	}
	return nil
}

func WithCRUDImpliesModelGetFieldMask(mask *fieldmaskpb.FieldMask) CRUDImpliesModelGetOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupCRUDImpliesModelModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func WithCRUDImpliesModelListFieldMask(mask *fieldmaskpb.FieldMask) CRUDImpliesModelListOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupCRUDImpliesModelModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func LookupCRUDImpliesModelModelColumn(path string) string {
	switch path {
	case "uuid":
		return "Uuid"
	}
	panic(path)
}

func LookupCRUDImpliesModelModelColumns(paths []string) (cols []string) {
	for _, p := range paths {
		cols = append(cols, LookupCRUDImpliesModelModelColumn(p))
	}
	return
}
