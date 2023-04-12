package datagooglefirebasewebappconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleFirebaseWebAppConfigAConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_firebase_web_app_config#web_app_id DataGoogleFirebaseWebAppConfigA#web_app_id}
	WebAppId *string `field:"required" json:"webAppId" yaml:"webAppId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_firebase_web_app_config#id DataGoogleFirebaseWebAppConfigA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project id of the Firebase web App.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_firebase_web_app_config#project DataGoogleFirebaseWebAppConfigA#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

