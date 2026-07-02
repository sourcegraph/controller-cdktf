package alertattribute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertAttributeConfig struct {
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
	// Whether this attribute is an array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_attribute#array AlertAttribute#array}
	Array any `field:"required" json:"array" yaml:"array"`
	// Unique name of this attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_attribute#name AlertAttribute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Engine resource name for this attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_attribute#type AlertAttribute#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether this attribute is required. If this field is not set, the existing setting will be preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/alert_attribute#required AlertAttribute#required}
	Required any `field:"optional" json:"required" yaml:"required"`
}
