package googlefirebaseapphostingdefaultdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAppHostingDefaultDomainConfig struct {
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
	// The ID of the Backend that this Domain is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_default_domain#backend GoogleFirebaseAppHostingDefaultDomain#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Id of the domain. For default domain, it should be {{backend}}--{{project_id}}.{{location}}.hosted.app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_default_domain#domain_id GoogleFirebaseAppHostingDefaultDomain#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The location of the Backend that this Domain is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_default_domain#location GoogleFirebaseAppHostingDefaultDomain#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Whether the domain is disabled. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_default_domain#disabled GoogleFirebaseAppHostingDefaultDomain#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_default_domain#id GoogleFirebaseAppHostingDefaultDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_default_domain#project GoogleFirebaseAppHostingDefaultDomain#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_app_hosting_default_domain#timeouts GoogleFirebaseAppHostingDefaultDomain#timeouts}
	Timeouts *GoogleFirebaseAppHostingDefaultDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

