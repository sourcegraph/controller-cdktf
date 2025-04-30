package datapostgresqlsequences

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPostgresqlSequencesConfig struct {
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
	// The PostgreSQL database which will be queried for sequence names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences#database DataPostgresqlSequences#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences#id DataPostgresqlSequences#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Expression(s) which will be pattern matched against sequence names in the query using the PostgreSQL LIKE ALL operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences#like_all_patterns DataPostgresqlSequences#like_all_patterns}
	LikeAllPatterns *[]*string `field:"optional" json:"likeAllPatterns" yaml:"likeAllPatterns"`
	// Expression(s) which will be pattern matched against sequence names in the query using the PostgreSQL LIKE ANY operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences#like_any_patterns DataPostgresqlSequences#like_any_patterns}
	LikeAnyPatterns *[]*string `field:"optional" json:"likeAnyPatterns" yaml:"likeAnyPatterns"`
	// Expression(s) which will be pattern matched against sequence names in the query using the PostgreSQL NOT LIKE ALL operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences#not_like_all_patterns DataPostgresqlSequences#not_like_all_patterns}
	NotLikeAllPatterns *[]*string `field:"optional" json:"notLikeAllPatterns" yaml:"notLikeAllPatterns"`
	// Expression which will be pattern matched against sequence names in the query using the PostgreSQL ~ (regular expression match) operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences#regex_pattern DataPostgresqlSequences#regex_pattern}
	RegexPattern *string `field:"optional" json:"regexPattern" yaml:"regexPattern"`
	// The PostgreSQL schema(s) which will be queried for sequence names. Queries all schemas in the database by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/data-sources/sequences#schemas DataPostgresqlSequences#schemas}
	Schemas *[]*string `field:"optional" json:"schemas" yaml:"schemas"`
}

