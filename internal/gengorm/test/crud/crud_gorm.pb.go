// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm v0.0.0
// 	protoc          (unknown)
// source: crud/crud.proto

package crud

import (
	context "context"
	fmt "fmt"
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
	gengorm "github.com/complex64/protoc-gen-gorm/pkg/gengorm"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	gorm "gorm.io/gorm"
)

// CrudModel is the GORM model for crud.Crud.
type CrudModel struct {
	Uuid        string `gorm:"primaryKey"`
	StringField string
}

// AsProto converts a CrudModel to its protobuf representation.
func (m *CrudModel) AsProto() (*Crud, error) {
	x := new(Crud)
	x.Uuid = m.Uuid
	x.StringField = m.StringField
	return x, nil
}

// AsModel converts a Crud to its GORM model.
func (x *Crud) AsModel() (*CrudModel, error) {
	m := new(CrudModel)
	m.Uuid = x.Uuid
	m.StringField = x.StringField
	return m, nil
}

type CrudWithDB struct {
	x  *Crud
	db *gorm.DB
}

func (x *Crud) WithDB(db *gorm.DB) CrudWithDB {
	return CrudWithDB{x: x, db: db}
}

func (c CrudWithDB) Create(ctx context.Context, opts ...gengorm.CreateOption) (*Crud, error) {
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

func (c CrudWithDB) Get(ctx context.Context, opts ...gengorm.GetOption) (*Crud, error) {
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
	n := CrudModel{}
	db := c.db.WithContext(ctx)
	if err := db.Where(m).First(&n).Error; err != nil {
		return nil, err
	}
	if y, err := n.AsProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CrudWithDB) List(ctx context.Context, opts ...gengorm.ListOption) ([]*Crud, error) {
	if c.x == nil {
		return nil, nil
	}
	var ms []CrudModel
	db := c.db.WithContext(ctx)
	if err := db.Find(&ms).Error; err != nil {
		return nil, err
	}
	xs := make([]*Crud, 0, len(ms))
	for _, m := range ms {
		if x, err := m.AsProto(); err != nil {
			return nil, err
		} else {
			xs = append(xs, x)
		}
	}
	return xs, nil
}

func (c CrudWithDB) Update(ctx context.Context, opts ...gengorm.UpdateOption) (*Crud, error) {
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
	if y, err := m.AsProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CrudWithDB) Patch(ctx context.Context, mask *fieldmaskpb.FieldMask, opts ...gengorm.PatchOption) (*Crud, error) {
	if c.x == nil {
		return nil, nil
	}
	if mask == nil {
		return c.Update(ctx)
	}
	if !mask.IsValid(c.x) {
		return nil, fmt.Errorf("invalid field mask")
	}
	paths := mask.Paths
	if len(paths) == 0 {
		return c.Update(ctx)
	}
	var zero string
	if c.x.Uuid == zero {
		return nil, fmt.Errorf("empty primary key")
	}
	var cols []string
	for _, path := range paths {
		switch path {
		case "uuid":
			cols = append(cols, "Uuid")
		case "string_field":
			cols = append(cols, "StringField")
		}
	}
	m, err := c.x.AsModel()
	if err != nil {
		return nil, err
	}
	target := CrudModel{Uuid: m.Uuid}
	db := c.db.WithContext(ctx)
	if err := db.Model(&target).Select(cols).Updates(m).Error; err != nil {
		return nil, err
	}
	if y, err := m.AsProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CrudWithDB) Delete(ctx context.Context, opts ...gengorm.DeleteOption) error {
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
	if err := db.Where(m).Delete(&CrudModel{}).Error; err != nil {
		return err
	}
	return nil
}