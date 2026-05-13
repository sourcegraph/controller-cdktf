package contactcenterinsightsautolabelingrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContactCenterInsightsAutoLabelingRuleConfig struct {
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
	// Location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#location ContactCenterInsightsAutoLabelingRule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Whether the rule is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#active ContactCenterInsightsAutoLabelingRule#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// A unique ID for the new AutoLabelingRule.
	//
	// This ID will become the final
	// component of the AutoLabelingRule's resource name. If no ID is specified,
	// a server-generated ID will be used.
	//
	// This value should be 4-64 characters and must match the regular
	// expression '^[A-Za-z0-9]{4,64}$'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#auto_labeling_rule_id ContactCenterInsightsAutoLabelingRule#auto_labeling_rule_id}
	AutoLabelingRuleId *string `field:"optional" json:"autoLabelingRuleId" yaml:"autoLabelingRuleId"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#conditions ContactCenterInsightsAutoLabelingRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#description ContactCenterInsightsAutoLabelingRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display Name of the auto labeling rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#display_name ContactCenterInsightsAutoLabelingRule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#id ContactCenterInsightsAutoLabelingRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The label key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#label_key ContactCenterInsightsAutoLabelingRule#label_key}
	LabelKey *string `field:"optional" json:"labelKey" yaml:"labelKey"`
	// The type of the label key. Possible values: ["LABEL_KEY_TYPE_UNSPECIFIED", "LABEL_KEY_TYPE_CUSTOM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#label_key_type ContactCenterInsightsAutoLabelingRule#label_key_type}
	LabelKeyType *string `field:"optional" json:"labelKeyType" yaml:"labelKeyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#project ContactCenterInsightsAutoLabelingRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/contact_center_insights_auto_labeling_rule#timeouts ContactCenterInsightsAutoLabelingRule#timeouts}
	Timeouts *ContactCenterInsightsAutoLabelingRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

