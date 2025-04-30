package googleapigeedeveloper

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeDeveloperConfig struct {
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
	// Email address of the developer.
	//
	// This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#email GoogleApigeeDeveloper#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// First name of the developer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#first_name GoogleApigeeDeveloper#first_name}
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// Last name of the developer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#last_name GoogleApigeeDeveloper#last_name}
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#org_id GoogleApigeeDeveloper#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// User name of the developer. Not used by Apigee hybrid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#user_name GoogleApigeeDeveloper#user_name}
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#attributes GoogleApigeeDeveloper#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#id GoogleApigeeDeveloper#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_apigee_developer#timeouts GoogleApigeeDeveloper#timeouts}
	Timeouts *GoogleApigeeDeveloperTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

