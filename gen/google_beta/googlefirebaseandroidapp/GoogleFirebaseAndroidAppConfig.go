package googlefirebaseandroidapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAndroidAppConfig struct {
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
	// The user-assigned display name of the AndroidApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#display_name GoogleFirebaseAndroidApp#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The canonical package name of the Android app as would appear in the Google Play Developer Console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#package_name GoogleFirebaseAndroidApp#package_name}
	PackageName *string `field:"required" json:"packageName" yaml:"packageName"`
	// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp.
	//
	// If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp.
	// This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#api_key_id GoogleFirebaseAndroidApp#api_key_id}
	ApiKeyId *string `field:"optional" json:"apiKeyId" yaml:"apiKeyId"`
	// (Optional) Set to 'ABANDON' to allow the AndroidApp to be untracked from terraform state rather than deleted upon 'terraform destroy'.
	//
	// This is useful because the AndroidApp may be
	// serving traffic. Set to 'DELETE' to delete the AndroidApp. Defaults to 'DELETE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#deletion_policy GoogleFirebaseAndroidApp#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#id GoogleFirebaseAndroidApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#project GoogleFirebaseAndroidApp#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The SHA1 certificate hashes for the AndroidApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#sha1_hashes GoogleFirebaseAndroidApp#sha1_hashes}
	Sha1Hashes *[]*string `field:"optional" json:"sha1Hashes" yaml:"sha1Hashes"`
	// The SHA256 certificate hashes for the AndroidApp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#sha256_hashes GoogleFirebaseAndroidApp#sha256_hashes}
	Sha256Hashes *[]*string `field:"optional" json:"sha256Hashes" yaml:"sha256Hashes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_android_app#timeouts GoogleFirebaseAndroidApp#timeouts}
	Timeouts *GoogleFirebaseAndroidAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

