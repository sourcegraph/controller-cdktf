package googlefirebasehostingsite

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseHostingSiteConfig struct {
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
	// Optional. The [ID of a Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id) associated with the Hosting site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_site#app_id GoogleFirebaseHostingSite#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_site#id GoogleFirebaseHostingSite#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_site#project GoogleFirebaseHostingSite#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Required.
	//
	// Immutable. A globally unique identifier for the Hosting site. This identifier is
	// used to construct the Firebase-provisioned subdomains for the site, so it must also be a valid
	// domain name label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_site#site_id GoogleFirebaseHostingSite#site_id}
	SiteId *string `field:"optional" json:"siteId" yaml:"siteId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_firebase_hosting_site#timeouts GoogleFirebaseHostingSite#timeouts}
	Timeouts *GoogleFirebaseHostingSiteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

