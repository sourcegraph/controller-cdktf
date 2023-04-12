package googlecloudrunservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudRunServiceConfig struct {
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
	// The location of the cloud run instance. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#location GoogleCloudRunService#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name must be unique within a namespace, within a Cloud Run region.
	//
	// Is required when creating resources. Name is primarily intended
	// for creation idempotence and configuration definition. Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#name GoogleCloudRunService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#autogenerate_revision_name GoogleCloudRunService#autogenerate_revision_name}.
	AutogenerateRevisionName interface{} `field:"optional" json:"autogenerateRevisionName" yaml:"autogenerateRevisionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#id GoogleCloudRunService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#metadata GoogleCloudRunService#metadata}
	Metadata *GoogleCloudRunServiceMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#project GoogleCloudRunService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#template GoogleCloudRunService#template}
	Template *GoogleCloudRunServiceTemplate `field:"optional" json:"template" yaml:"template"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#timeouts GoogleCloudRunService#timeouts}
	Timeouts *GoogleCloudRunServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#traffic GoogleCloudRunService#traffic}
	Traffic interface{} `field:"optional" json:"traffic" yaml:"traffic"`
}

