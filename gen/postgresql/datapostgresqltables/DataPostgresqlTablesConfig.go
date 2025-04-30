package datapostgresqltables

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPostgresqlTablesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The PostgreSQL database which will be queried for table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#database DataPostgresqlTables#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#id DataPostgresqlTables#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Expression(s) which will be pattern matched against table names in the query using the PostgreSQL LIKE ALL operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#like_all_patterns DataPostgresqlTables#like_all_patterns}
	LikeAllPatterns *[]*string `field:"optional" json:"likeAllPatterns" yaml:"likeAllPatterns"`
	// Expression(s) which will be pattern matched against table names in the query using the PostgreSQL LIKE ANY operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#like_any_patterns DataPostgresqlTables#like_any_patterns}
	LikeAnyPatterns *[]*string `field:"optional" json:"likeAnyPatterns" yaml:"likeAnyPatterns"`
	// Expression(s) which will be pattern matched against table names in the query using the PostgreSQL NOT LIKE ALL operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#not_like_all_patterns DataPostgresqlTables#not_like_all_patterns}
	NotLikeAllPatterns *[]*string `field:"optional" json:"notLikeAllPatterns" yaml:"notLikeAllPatterns"`
	// Expression which will be pattern matched against table names in the query using the PostgreSQL ~ (regular expression match) operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#regex_pattern DataPostgresqlTables#regex_pattern}
	RegexPattern *string `field:"optional" json:"regexPattern" yaml:"regexPattern"`
	// The PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#schemas DataPostgresqlTables#schemas}
	Schemas *[]*string `field:"optional" json:"schemas" yaml:"schemas"`
	// The PostgreSQL table types which will be queried for table names.
	//
	// Includes all table types by default. Use 'BASE TABLE' for normal tables only
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/tables#table_types DataPostgresqlTables#table_types}
	TableTypes *[]*string `field:"optional" json:"tableTypes" yaml:"tableTypes"`
}

