package linkdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinkDefinitionConfig struct {
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
	// Description of the associated relationship.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_definition#associated_description LinkDefinition#associated_description}
	AssociatedDescription *string `field:"required" json:"associatedDescription" yaml:"associatedDescription"`
	// API name of the associated link.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_definition#associated_name LinkDefinition#associated_name}
	AssociatedName *string `field:"required" json:"associatedName" yaml:"associatedName"`
	// Display name of the associated link.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_definition#associated_title LinkDefinition#associated_title}
	AssociatedTitle *string `field:"required" json:"associatedTitle" yaml:"associatedTitle"`
	// Description of the primary relationship.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_definition#primary_description LinkDefinition#primary_description}
	PrimaryDescription *string `field:"required" json:"primaryDescription" yaml:"primaryDescription"`
	// API name of the primary link.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_definition#primary_name LinkDefinition#primary_name}
	PrimaryName *string `field:"required" json:"primaryName" yaml:"primaryName"`
	// Display name of the primary link.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_definition#primary_title LinkDefinition#primary_title}
	PrimaryTitle *string `field:"required" json:"primaryTitle" yaml:"primaryTitle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/link_definition#id LinkDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

