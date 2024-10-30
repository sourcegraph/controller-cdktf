package appaccesspolicyassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppAccessPolicyAssignmentConfig struct {
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
	// The application ID; this value is immutable and can not be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_access_policy_assignment#app_id AppAccessPolicyAssignment#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The access policy ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/okta/okta/4.11.1/docs/resources/app_access_policy_assignment#policy_id AppAccessPolicyAssignment#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}

