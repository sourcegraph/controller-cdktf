package googleappengineapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleAppEngineApplicationConfig struct {
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
	// The location to serve the app from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#location_id GoogleAppEngineApplication#location_id}
	LocationId *string `field:"required" json:"locationId" yaml:"locationId"`
	// The domain to authenticate users with when using App Engine's User API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#auth_domain GoogleAppEngineApplication#auth_domain}
	AuthDomain *string `field:"optional" json:"authDomain" yaml:"authDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#database_type GoogleAppEngineApplication#database_type}.
	DatabaseType *string `field:"optional" json:"databaseType" yaml:"databaseType"`
	// feature_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#feature_settings GoogleAppEngineApplication#feature_settings}
	FeatureSettings *GoogleAppEngineApplicationFeatureSettings `field:"optional" json:"featureSettings" yaml:"featureSettings"`
	// iap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#iap GoogleAppEngineApplication#iap}
	Iap *GoogleAppEngineApplicationIap `field:"optional" json:"iap" yaml:"iap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#id GoogleAppEngineApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The project ID to create the application under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#project GoogleAppEngineApplication#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The serving status of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#serving_status GoogleAppEngineApplication#serving_status}
	ServingStatus *string `field:"optional" json:"servingStatus" yaml:"servingStatus"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#timeouts GoogleAppEngineApplication#timeouts}
	Timeouts *GoogleAppEngineApplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

