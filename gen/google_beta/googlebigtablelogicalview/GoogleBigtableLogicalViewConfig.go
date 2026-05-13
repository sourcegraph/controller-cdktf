package googlebigtablelogicalview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBigtableLogicalViewConfig struct {
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
	// The unique name of the logical view in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigtable_logical_view#logical_view_id GoogleBigtableLogicalView#logical_view_id}
	LogicalViewId *string `field:"required" json:"logicalViewId" yaml:"logicalViewId"`
	// The logical view's select query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigtable_logical_view#query GoogleBigtableLogicalView#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Set to true to make the logical view protected against deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigtable_logical_view#deletion_protection GoogleBigtableLogicalView#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigtable_logical_view#id GoogleBigtableLogicalView#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the instance to create the logical view within.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigtable_logical_view#instance GoogleBigtableLogicalView#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigtable_logical_view#project GoogleBigtableLogicalView#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_bigtable_logical_view#timeouts GoogleBigtableLogicalView#timeouts}
	Timeouts *GoogleBigtableLogicalViewTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

