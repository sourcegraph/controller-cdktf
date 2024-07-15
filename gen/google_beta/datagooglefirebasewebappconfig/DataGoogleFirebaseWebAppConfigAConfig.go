package datagooglefirebasewebappconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleFirebaseWebAppConfigAConfig struct {
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
	// The id of the Firebase web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/data-sources/google_firebase_web_app_config#web_app_id DataGoogleFirebaseWebAppConfigA#web_app_id}
	WebAppId *string `field:"required" json:"webAppId" yaml:"webAppId"`
	// The project id of the Firebase web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/data-sources/google_firebase_web_app_config#project DataGoogleFirebaseWebAppConfigA#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

