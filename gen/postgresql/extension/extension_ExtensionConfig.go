package extension

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExtensionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#name Extension#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When true, will also create any extensions that this extension depends on that are not already installed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#create_cascade Extension#create_cascade}
	CreateCascade interface{} `field:"optional" json:"createCascade" yaml:"createCascade"`
	// Sets the database to add the extension to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#database Extension#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// When true, will also drop all the objects that depend on the extension, and in turn all objects that depend on those objects.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#drop_cascade Extension#drop_cascade}
	DropCascade interface{} `field:"optional" json:"dropCascade" yaml:"dropCascade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#id Extension#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Sets the schema of an extension.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#schema Extension#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Sets the version number of the extension.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/postgresql/r/extension#version Extension#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

