package functionresource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionResourceConfig struct {
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
	// Body of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#body FunctionResource#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// Name of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#name FunctionResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// arg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#arg FunctionResource#arg}
	Arg interface{} `field:"optional" json:"arg" yaml:"arg"`
	// The database where the function is located. If not specified, the provider default database is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#database FunctionResource#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Automatically drop objects that depend on the function (such as operators or triggers), and in turn all objects that depend on those objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#drop_cascade FunctionResource#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#id FunctionResource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Language of theof the function. One of: internal, sql, c, plpgsql.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#language FunctionResource#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
	// If the function can be executed in parallel for a single query execution. One of: UNSAFE, RESTRICTED, SAFE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#parallel FunctionResource#parallel}
	Parallel *string `field:"optional" json:"parallel" yaml:"parallel"`
	// Function return type. If not specified, it will be calculated based on the output arguments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#returns FunctionResource#returns}
	Returns *string `field:"optional" json:"returns" yaml:"returns"`
	// Schema where the function is located. If not specified, the provider default schema is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#schema FunctionResource#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// If the function should execute with the permissions of the function owner instead of the permissions of the caller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#security_definer FunctionResource#security_definer}
	SecurityDefiner interface{} `field:"optional" json:"securityDefiner" yaml:"securityDefiner"`
	// If the function should always return NULL if any of it's inputs is NULL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#strict FunctionResource#strict}
	Strict interface{} `field:"optional" json:"strict" yaml:"strict"`
	// Volatility of the function. One of: VOLATILE, STABLE, IMMUTABLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cyrilgdn/postgresql/1.25.0/docs/resources/function#volatility FunctionResource#volatility}
	Volatility *string `field:"optional" json:"volatility" yaml:"volatility"`
}

