package schema

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchemaConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#name Schema#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The database name to alter schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#database Schema#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// When true, will also drop all the objects that are contained in the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#drop_cascade Schema#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#id Schema#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// When true, use the existing schema if it exists.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#if_not_exists Schema#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// The ROLE name who owns the schema.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#owner Schema#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/schema#policy Schema#policy}
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
}

