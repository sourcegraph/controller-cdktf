package googleapigeecontrolplaneaccess

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeControlPlaneAccessConfig struct {
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
	// Name of the Apigee organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_apigee_control_plane_access#name GoogleApigeeControlPlaneAccess#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Array of service accounts authorized to publish analytics data to the control plane, each specified using the following format: 'serviceAccount:service-account-name'.
	//
	// The 'service-account-name' is formatted like an email address. For example: serviceAccount@my_project_id.iam.gserviceaccount.com
	//
	// You might specify multiple service accounts, for example, if you have multiple environments and wish to assign a unique service account to each one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_apigee_control_plane_access#analytics_publisher_identities GoogleApigeeControlPlaneAccess#analytics_publisher_identities}
	AnalyticsPublisherIdentities *[]*string `field:"optional" json:"analyticsPublisherIdentities" yaml:"analyticsPublisherIdentities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_apigee_control_plane_access#id GoogleApigeeControlPlaneAccess#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Array of service accounts to grant access to control plane resources (for the Synchronizer component), each specified using the following format: 'serviceAccount:service-account-name'.
	//
	// The 'service-account-name' is formatted like an email address. For example: serviceAccount@my_project_id.iam.gserviceaccount.com
	//
	// You might specify multiple service accounts, for example, if you have multiple environments and wish to assign a unique service account to each one.
	//
	// The service accounts must have **Apigee Synchronizer Manager** role. See also [Create service accounts](https://cloud.google.com/apigee/docs/hybrid/v1.8/sa-about#create-the-service-accounts).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_apigee_control_plane_access#synchronizer_identities GoogleApigeeControlPlaneAccess#synchronizer_identities}
	SynchronizerIdentities *[]*string `field:"optional" json:"synchronizerIdentities" yaml:"synchronizerIdentities"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_apigee_control_plane_access#timeouts GoogleApigeeControlPlaneAccess#timeouts}
	Timeouts *GoogleApigeeControlPlaneAccessTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

