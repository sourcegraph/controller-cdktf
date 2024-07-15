package googlefirebaseappcheckappattestconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAppCheckAppAttestConfigConfig struct {
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
	// The ID of an [Apple App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_app_check_app_attest_config#app_id GoogleFirebaseAppCheckAppAttestConfig#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_app_check_app_attest_config#id GoogleFirebaseAppCheckAppAttestConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_app_check_app_attest_config#project GoogleFirebaseAppCheckAppAttestConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_app_check_app_attest_config#timeouts GoogleFirebaseAppCheckAppAttestConfig#timeouts}
	Timeouts *GoogleFirebaseAppCheckAppAttestConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid.
	//
	// If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_app_check_app_attest_config#token_ttl GoogleFirebaseAppCheckAppAttestConfig#token_ttl}
	TokenTtl *string `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
}

