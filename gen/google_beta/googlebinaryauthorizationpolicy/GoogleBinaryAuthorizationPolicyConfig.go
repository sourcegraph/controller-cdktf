package googlebinaryauthorizationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleBinaryAuthorizationPolicyConfig struct {
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
	// default_admission_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#default_admission_rule GoogleBinaryAuthorizationPolicy#default_admission_rule}
	DefaultAdmissionRule *GoogleBinaryAuthorizationPolicyDefaultAdmissionRule `field:"required" json:"defaultAdmissionRule" yaml:"defaultAdmissionRule"`
	// admission_whitelist_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#admission_whitelist_patterns GoogleBinaryAuthorizationPolicy#admission_whitelist_patterns}
	AdmissionWhitelistPatterns interface{} `field:"optional" json:"admissionWhitelistPatterns" yaml:"admissionWhitelistPatterns"`
	// cluster_admission_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#cluster_admission_rules GoogleBinaryAuthorizationPolicy#cluster_admission_rules}
	ClusterAdmissionRules interface{} `field:"optional" json:"clusterAdmissionRules" yaml:"clusterAdmissionRules"`
	// A descriptive comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#description GoogleBinaryAuthorizationPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Controls the evaluation of a Google-maintained global admission policy for common system-level images.
	//
	// Images not covered by the global
	// policy will be subject to the project admission policy. Possible values: ["ENABLE", "DISABLE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#global_policy_evaluation_mode GoogleBinaryAuthorizationPolicy#global_policy_evaluation_mode}
	GlobalPolicyEvaluationMode *string `field:"optional" json:"globalPolicyEvaluationMode" yaml:"globalPolicyEvaluationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#id GoogleBinaryAuthorizationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#project GoogleBinaryAuthorizationPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_binary_authorization_policy#timeouts GoogleBinaryAuthorizationPolicy#timeouts}
	Timeouts *GoogleBinaryAuthorizationPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

