package googlecomputeorganizationsecuritypolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeOrganizationSecurityPolicyRuleConfig struct {
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
	// The Action to perform when the client connection triggers the rule. Can currently be either "allow", "deny" or "goto_next".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#action GoogleComputeOrganizationSecurityPolicyRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#match GoogleComputeOrganizationSecurityPolicyRule#match}
	Match *GoogleComputeOrganizationSecurityPolicyRuleMatch `field:"required" json:"match" yaml:"match"`
	// The ID of the OrganizationSecurityPolicy this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#policy_id GoogleComputeOrganizationSecurityPolicyRule#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// An integer indicating the priority of a rule in the list.
	//
	// The priority must be a value
	// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	// highest priority and 2147483647 is the lowest prority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#priority GoogleComputeOrganizationSecurityPolicyRule#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// A description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#description GoogleComputeOrganizationSecurityPolicyRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The direction in which this rule applies. If unspecified an INGRESS rule is created. Possible values: ["INGRESS", "EGRESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#direction GoogleComputeOrganizationSecurityPolicyRule#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Denotes whether to enable logging for a particular rule.
	//
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#enable_logging GoogleComputeOrganizationSecurityPolicyRule#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#id GoogleComputeOrganizationSecurityPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If set to true, the specified action is not enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#preview GoogleComputeOrganizationSecurityPolicyRule#preview}
	Preview interface{} `field:"optional" json:"preview" yaml:"preview"`
	// A list of network resource URLs to which this rule applies.
	//
	// This field allows you to control which network's VMs get
	// this rule. If this field is left blank, all VMs
	// within the organization will receive the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#target_resources GoogleComputeOrganizationSecurityPolicyRule#target_resources}
	TargetResources *[]*string `field:"optional" json:"targetResources" yaml:"targetResources"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#target_service_accounts GoogleComputeOrganizationSecurityPolicyRule#target_service_accounts}
	TargetServiceAccounts *[]*string `field:"optional" json:"targetServiceAccounts" yaml:"targetServiceAccounts"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_organization_security_policy_rule#timeouts GoogleComputeOrganizationSecurityPolicyRule#timeouts}
	Timeouts *GoogleComputeOrganizationSecurityPolicyRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

