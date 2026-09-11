package worksheet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorksheetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Worksheet name. Must be unique within workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/worksheet#name Worksheet#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Worksheet definition in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/worksheet#queries Worksheet#queries}
	Queries *string `field:"required" json:"queries" yaml:"queries"`
	// OID of workspace worksheet is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/worksheet#workspace Worksheet#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// Icon image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/worksheet#icon_url Worksheet#icon_url}
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/worksheet#id Worksheet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

