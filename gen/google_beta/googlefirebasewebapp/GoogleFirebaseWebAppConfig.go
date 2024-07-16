package googlefirebasewebapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseWebAppConfig struct {
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
	// The user-assigned display name of the App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_web_app#display_name GoogleFirebaseWebApp#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the WebApp.
	//
	// If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the WebApp.
	// This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_web_app#api_key_id GoogleFirebaseWebApp#api_key_id}
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// Set to 'ABANDON' to allow the WebApp to be untracked from terraform state rather than deleted upon 'terraform destroy'.
	//
	// This is useful becaue the WebApp may be
	// serving traffic. Set to 'DELETE' to delete the WebApp. Default to 'DELETE'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_web_app#deletion_policy GoogleFirebaseWebApp#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_web_app#id GoogleFirebaseWebApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_web_app#project GoogleFirebaseWebApp#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_web_app#timeouts GoogleFirebaseWebApp#timeouts}
	Timeouts *GoogleFirebaseWebAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

