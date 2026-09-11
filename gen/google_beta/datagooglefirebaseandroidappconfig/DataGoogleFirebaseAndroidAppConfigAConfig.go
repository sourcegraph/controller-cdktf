package datagooglefirebaseandroidappconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataGoogleFirebaseAndroidAppConfigAConfig struct {
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
	// The id of the Firebase Android App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/data-sources/google_firebase_android_app_config#app_id DataGoogleFirebaseAndroidAppConfigA#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The project id of the Firebase Android App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.32.0/docs/data-sources/google_firebase_android_app_config#project DataGoogleFirebaseAndroidAppConfigA#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

