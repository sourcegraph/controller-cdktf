package googleclouddeploytarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleClouddeployTargetConfig struct {
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
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#location GoogleClouddeployTarget#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the `Target`. Format is `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#name GoogleClouddeployTarget#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional.
	//
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#annotations GoogleClouddeployTarget#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// anthos_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#anthos_cluster GoogleClouddeployTarget#anthos_cluster}
	AnthosCluster *GoogleClouddeployTargetAnthosCluster `field:"optional" json:"anthosCluster" yaml:"anthosCluster"`
	// custom_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#custom_target GoogleClouddeployTarget#custom_target}
	CustomTarget *GoogleClouddeployTargetCustomTarget `field:"optional" json:"customTarget" yaml:"customTarget"`
	// Optional. The deploy parameters to use for this target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#deploy_parameters GoogleClouddeployTarget#deploy_parameters}
	DeployParameters *map[string]*string `field:"optional" json:"deployParameters" yaml:"deployParameters"`
	// Optional. Description of the `Target`. Max length is 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#description GoogleClouddeployTarget#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// execution_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#execution_configs GoogleClouddeployTarget#execution_configs}
	ExecutionConfigs interface{} `field:"optional" json:"executionConfigs" yaml:"executionConfigs"`
	// gke block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#gke GoogleClouddeployTarget#gke}
	Gke *GoogleClouddeployTargetGke `field:"optional" json:"gke" yaml:"gke"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#id GoogleClouddeployTarget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional.
	//
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effective_labels` for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#labels GoogleClouddeployTarget#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// multi_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#multi_target GoogleClouddeployTarget#multi_target}
	MultiTarget *GoogleClouddeployTargetMultiTarget `field:"optional" json:"multiTarget" yaml:"multiTarget"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#project GoogleClouddeployTarget#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Optional. Whether or not the `Target` requires approval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#require_approval GoogleClouddeployTarget#require_approval}
	RequireApproval interface{} `field:"optional" json:"requireApproval" yaml:"requireApproval"`
	// run block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#run GoogleClouddeployTarget#run}
	Run *GoogleClouddeployTargetRun `field:"optional" json:"run" yaml:"run"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_clouddeploy_target#timeouts GoogleClouddeployTarget#timeouts}
	Timeouts *GoogleClouddeployTargetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

