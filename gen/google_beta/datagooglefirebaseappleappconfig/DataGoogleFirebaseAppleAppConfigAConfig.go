package datagooglefirebaseappleappconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleFirebaseAppleAppConfigAConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The id of the Firebase iOS App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/data-sources/google_firebase_apple_app_config#app_id DataGoogleFirebaseAppleAppConfigA#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The project id of the Firebase iOS App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/data-sources/google_firebase_apple_app_config#project DataGoogleFirebaseAppleAppConfigA#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

