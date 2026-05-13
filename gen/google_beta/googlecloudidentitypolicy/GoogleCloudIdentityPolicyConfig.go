package googlecloudidentitypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudIdentityPolicyConfig struct {
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
	// The customer that the Policy belongs to. Format: 'customers/{customer_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#customer GoogleCloudIdentityPolicy#customer}
	Customer *string `field:"required" json:"customer" yaml:"customer"`
	// policy_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#policy_query GoogleCloudIdentityPolicy#policy_query}
	PolicyQuery *GoogleCloudIdentityPolicyPolicyQuery `field:"required" json:"policyQuery" yaml:"policyQuery"`
	// setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#setting GoogleCloudIdentityPolicy#setting}
	Setting *GoogleCloudIdentityPolicySetting `field:"required" json:"setting" yaml:"setting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#id GoogleCloudIdentityPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_identity_policy#timeouts GoogleCloudIdentityPolicy#timeouts}
	Timeouts *GoogleCloudIdentityPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

