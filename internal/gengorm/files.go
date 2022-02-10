package gengorm

import (
	"fmt"

	"github.com/complex64/protoc-gen-gorm/gormpb"
	"github.com/complex64/protoc-gen-gorm/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFile(plugin *protogen.Plugin, proto *protogen.File) (*File, error) {
	file := &File{
		plugin: plugin,
		proto:  proto,
	}
	if err := file.init(); err != nil {
		return nil, err
	}
	return file, nil
}

// File wraps the input .proto file and keeps information to generate code.
type File struct {
	plugin *protogen.Plugin
	proto  *protogen.File

	out  *protogen.GeneratedFile
	opts *gormpb.FileOptions
	msgs []*Message
}

func (f *File) init() error {
	f.initOut()
	f.initOpts()
	if err := f.initMsgs(); err != nil {
		return err
	}
	return nil
}

func (f *File) initOut() {
	// Use extension .pb.go, like `protoc-gen-go` does.
	const extension = "_gorm.pb.go"
	filename := f.proto.GeneratedFilenamePrefix + extension
	f.out = f.plugin.NewGeneratedFile(filename, f.proto.GoImportPath)
}

// initOpts reads the protoc-gen-gorm options.
// Example: option (gorm.file).model = true;
func (f *File) initOpts() {
	descOpts := f.proto.Desc.Options()
	opts, ok := proto.GetExtension(descOpts, gormpb.E_File).(*gormpb.FileOptions)
	if ok && opts != nil {
		f.opts = opts
	} else {
		f.opts = &gormpb.FileOptions{}
	}
}

func (f *File) initMsgs() error {
	return walkMessages(f.proto.Messages, []func(*protogen.Message) error{
		f.initMsg,
	})
}

func (f *File) initMsg(proto *protogen.Message) error {
	msg, err := NewMessage(f, proto)
	if err != nil {
		return err
	}
	f.msgs = append(f.msgs, msg)
	return nil
}

func (f *File) Gen() {
	f.genHeader()
	f.genPackage()
	f.genImports()
	f.genModels()
}

// genHeader generates the comment header at the top of the file.
// We warn to not edit the file and show the tool versions used to generate the file.
func (f *File) genHeader() {
	f.P("// Code generated by protoc-gen-gorm. DO NOT EDIT.")
	f.P("// versions:")

	protocGenGormVersion := version.String()
	protocVersion := "(unknown)"

	if v := f.plugin.Request.GetCompilerVersion(); v != nil {
		protocVersion = fmt.Sprintf("v%v.%v.%v", v.GetMajor(), v.GetMinor(), v.GetPatch())
		if s := v.GetSuffix(); s != "" {
			protocVersion += "-" + s
		}
	}

	f.P("// \tprotoc-gen-gorm ", protocGenGormVersion)
	f.P("// \tprotoc          ", protocVersion)

	path := f.proto.Desc.Path()
	if f.proto.Proto.GetOptions().GetDeprecated() {
		f.P("// ", path, " is a deprecated file.")
	} else {
		f.P("// source: ", path)
	}

	f.P()
}

func (f *File) genPackage() {
	f.P("package ", f.proto.GoPackageName)
	f.P()
}

func (f *File) genImports() {
	for i, imps := 0, f.proto.Desc.Imports(); i < imps.Len(); i++ {
		f.genImport(imps.Get(i))
	}
}

// genImport adds necessary import statements to the generated file.
func (f *File) genImport(imp protoreflect.FileImport) {
	if imp.IsPublic {
		panic("TODO: public imports")
	}

	impFile, ok := f.plugin.FilesByPath[imp.Path()]
	if !ok {
		return // .proto file unavailable.
	}
	if impFile.GoImportPath == f.proto.GoImportPath {
		return // Same package.
	}
	f.out.Import(impFile.GoImportPath)
	f.P()
}

func (f *File) genModels() {
	for _, msg := range f.msgs {
		msg.Gen()
	}
}

func (f *File) CRUD() bool     { return f.opts.Crud }
func (f *File) Validate() bool { return f.opts.Validate }
func (f *File) Model() bool    { return f.CRUD() || f.Validate() || f.opts.Model }

func (f *File) Annotate(symbol string, loc protogen.Location) { f.out.Annotate(symbol, loc) }
func (f *File) P(v ...interface{})                            { f.out.P(v...) }

// walkMessages invokes each function in walkFuncs for all messages, recursively.
func walkMessages(msgs []*protogen.Message, walkFuncs []func(*protogen.Message) error) error {
	for _, msg := range msgs {
		for _, f := range walkFuncs {
			if err := f(msg); err != nil {
				return err
			}
		}
		if err := walkMessages(msg.Messages, walkFuncs); err != nil {
			return err
		}
	}
	return nil
}
