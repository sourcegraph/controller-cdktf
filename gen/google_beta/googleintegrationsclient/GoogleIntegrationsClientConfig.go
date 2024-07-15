package googleintegrationsclient

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIntegrationsClientConfig struct {
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
	// Location in which client needs to be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#location GoogleIntegrationsClient#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// cloud_kms_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#cloud_kms_config GoogleIntegrationsClient#cloud_kms_config}
	CloudKmsConfig *GoogleIntegrationsClientCloudKmsConfig `field:"optional" json:"cloudKmsConfig" yaml:"cloudKmsConfig"`
	// Indicates if sample integrations should be created along with provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#create_sample_integrations GoogleIntegrationsClient#create_sample_integrations}
	CreateSampleIntegrations interface{} `field:"optional" json:"createSampleIntegrations" yaml:"createSampleIntegrations"`
	// Indicates if sample workflow should be created along with provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#create_sample_workflows GoogleIntegrationsClient#create_sample_workflows}
	CreateSampleWorkflows interface{} `field:"optional" json:"createSampleWorkflows" yaml:"createSampleWorkflows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#id GoogleIntegrationsClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#project GoogleIntegrationsClient#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Indicates provision with GMEK or CMEK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#provision_gmek GoogleIntegrationsClient#provision_gmek}
	ProvisionGmek interface{} `field:"optional" json:"provisionGmek" yaml:"provisionGmek"`
	// User input run-as service account, if empty, will bring up a new default service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#run_as_service_account GoogleIntegrationsClient#run_as_service_account}
	RunAsServiceAccount *string `field:"optional" json:"runAsServiceAccount" yaml:"runAsServiceAccount"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_client#timeouts GoogleIntegrationsClient#timeouts}
	Timeouts *GoogleIntegrationsClientTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

