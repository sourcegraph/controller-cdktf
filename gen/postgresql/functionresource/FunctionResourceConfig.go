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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#body FunctionResource#body}
	Body *string `field:"required" json:"body" yaml:"body"`
	// Name of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#name FunctionResource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// arg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#arg FunctionResource#arg}
	Arg interface{} `field:"optional" json:"arg" yaml:"arg"`
	// The database where the function is located. If not specified, the provider default database is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#database FunctionResource#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Automatically drop objects that depend on the function (such as operators or triggers), and in turn all objects that depend on those objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#drop_cascade FunctionResource#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#id FunctionResource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Function return type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#returns FunctionResource#returns}
	Returns *string `field:"optional" json:"returns" yaml:"returns"`
	// Schema where the function is located. If not specified, the provider default schema is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/sourcegraph/postgresql/1.18.0/docs/resources/function#schema FunctionResource#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

