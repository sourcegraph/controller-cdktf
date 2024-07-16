package googlefirebaseappleapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAppleAppConfig struct {
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
	// The canonical bundle ID of the Apple app as it would appear in the Apple AppStore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#bundle_id GoogleFirebaseAppleApp#bundle_id}
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// The user-assigned display name of the App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#display_name GoogleFirebaseAppleApp#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AppleApp.
	//
	// If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AppleApp.
	// This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#api_key_id GoogleFirebaseAppleApp#api_key_id}
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// The automatically generated Apple ID assigned to the Apple app by Apple in the Apple App Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#app_store_id GoogleFirebaseAppleApp#app_store_id}
	AppStoreId *string `field:"optional" json:"appStoreId" yaml:"appStoreId"`
	// (Optional) Set to 'ABANDON' to allow the Apple to be untracked from terraform state rather than deleted upon 'terraform destroy'.
	//
	// This is useful because the Apple may be
	// serving traffic. Set to 'DELETE' to delete the Apple. Defaults to 'DELETE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#deletion_policy GoogleFirebaseAppleApp#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#id GoogleFirebaseAppleApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#project GoogleFirebaseAppleApp#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Apple Developer Team ID associated with the App in the App Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#team_id GoogleFirebaseAppleApp#team_id}
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_apple_app#timeouts GoogleFirebaseAppleApp#timeouts}
	Timeouts *GoogleFirebaseAppleAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

