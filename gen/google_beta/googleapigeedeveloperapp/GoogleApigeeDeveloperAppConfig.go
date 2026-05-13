package googleapigeedeveloperapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeDeveloperAppConfig struct {
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
	// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to developer apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#callback_url GoogleApigeeDeveloperApp#callback_url}
	CallbackUrl *string `field:"required" json:"callbackUrl" yaml:"callbackUrl"`
	// Email address of the developer.
	//
	// This value is used to uniquely identify the developer in Apigee hybrid.
	// Note that the email address has to be in lowercase only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#developer_email GoogleApigeeDeveloperApp#developer_email}
	DeveloperEmail *string `field:"required" json:"developerEmail" yaml:"developerEmail"`
	// Name of the developer app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#name GoogleApigeeDeveloperApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Apigee Organization associated with the Apigee instance, in the format 'organizations/{{org_name}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#org_id GoogleApigeeDeveloperApp#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// List of API products associated with the developer app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#api_products GoogleApigeeDeveloperApp#api_products}
	ApiProducts *[]*string `field:"optional" json:"apiProducts" yaml:"apiProducts"`
	// Developer app family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#app_family GoogleApigeeDeveloperApp#app_family}
	AppFamily *string `field:"optional" json:"appFamily" yaml:"appFamily"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#attributes GoogleApigeeDeveloperApp#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#id GoogleApigeeDeveloperApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Expiration time, in milliseconds, for the consumer key that is generated for the developer app.
	//
	// If not set or left to the default value of -1,
	// the API key never expires. The expiration time can't be updated after it is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#key_expires_in GoogleApigeeDeveloperApp#key_expires_in}
	KeyExpiresIn *string `field:"optional" json:"keyExpiresIn" yaml:"keyExpiresIn"`
	// Scopes to apply to the developer app.
	//
	// The specified scopes must already exist for the API product that
	// you associate with the developer app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#scopes GoogleApigeeDeveloperApp#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Status of the credential. Valid values include approved or revoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#status GoogleApigeeDeveloperApp#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_apigee_developer_app#timeouts GoogleApigeeDeveloperApp#timeouts}
	Timeouts *GoogleApigeeDeveloperAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

