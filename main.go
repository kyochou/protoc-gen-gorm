package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
	"path/filepath"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm"
	"github.com/complex64/protoc-gen-gorm/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	docURL = "https://github.com/complex64/protoc-gen-gorm"
)

var (
	showVersion = flag.Bool("version", false, "")
	showHelp    = flag.Bool("help", false, "")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}
	if *showHelp {
		fmt.Fprintf(os.Stdout, "Please see %s for usage information.\n", docURL)
		os.Exit(0)
	}

	var (
		flags flag.FlagSet
	)
	protogen.Options{
		// Calls `flags.Set(param, value)` for each from `--gorm_out=<param1>=<value1>,...`.
		ParamFunc: flags.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range plugin.Files {
			if f.Generate {
				if err := gengorm.GenerateFile(flags, plugin, f); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
