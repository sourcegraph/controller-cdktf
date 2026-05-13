package googleiapwebforwardingruleserviceiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIapWebForwardingRuleServiceIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#forwarding_rule_service_name GoogleIapWebForwardingRuleServiceIamBinding#forwarding_rule_service_name}.
	ForwardingRuleServiceName *string `field:"required" json:"forwardingRuleServiceName" yaml:"forwardingRuleServiceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#members GoogleIapWebForwardingRuleServiceIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#role GoogleIapWebForwardingRuleServiceIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#condition GoogleIapWebForwardingRuleServiceIamBinding#condition}
	Condition *GoogleIapWebForwardingRuleServiceIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#id GoogleIapWebForwardingRuleServiceIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#project GoogleIapWebForwardingRuleServiceIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

