module github.com/complex64/protoc-gen-gorm

go 1.17

replace github.com/complex64/protoc-gen-gorm/gormpb => ./gormpb

require (
	github.com/google/go-cmp v0.5.9
	github.com/stretchr/testify v1.8.1
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.24.3
)

require (
	github.com/complex64/protoc-gen-gorm/gormpb v0.0.0-20230118151903-c16dc98323c5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
