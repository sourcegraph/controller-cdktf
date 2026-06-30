package dataincidentcustomfieldoption

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataIncidentCustomFieldOptionConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ID of the custom field this option belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/custom_field_option#custom_field_id DataIncidentCustomFieldOption#custom_field_id}
	CustomFieldId *string `field:"required" json:"customFieldId" yaml:"customFieldId"`
	// Human readable name for the custom field option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/custom_field_option#value DataIncidentCustomFieldOption#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Sort key used to order the custom field options correctly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/data-sources/custom_field_option#sort_key DataIncidentCustomFieldOption#sort_key}
	SortKey *float64 `field:"optional" json:"sortKey" yaml:"sortKey"`
}
