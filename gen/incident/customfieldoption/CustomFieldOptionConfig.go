package customfieldoption

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CustomFieldOptionConfig struct {
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
	// ID of the custom field this option belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field_option#custom_field_id CustomFieldOption#custom_field_id}
	CustomFieldId *string `field:"required" json:"customFieldId" yaml:"customFieldId"`
	// Human readable name for the custom field option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field_option#value CustomFieldOption#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Sort key used to order the custom field options correctly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/custom_field_option#sort_key CustomFieldOption#sort_key}
	SortKey *float64 `field:"optional" json:"sortKey" yaml:"sortKey"`
}

