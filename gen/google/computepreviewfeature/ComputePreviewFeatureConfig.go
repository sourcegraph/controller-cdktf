package computepreviewfeature

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputePreviewFeatureConfig struct {
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
	// The activation status of the preview feature. Possible values: ["ENABLED", "ACTIVATION_STATE_UNSPECIFIED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_preview_feature#activation_status ComputePreviewFeature#activation_status}
	ActivationStatus *string `field:"required" json:"activationStatus" yaml:"activationStatus"`
	// The name of the preview feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_preview_feature#name ComputePreviewFeature#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_preview_feature#id ComputePreviewFeature#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_preview_feature#project ComputePreviewFeature#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rollout_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_preview_feature#rollout_operation ComputePreviewFeature#rollout_operation}
	RolloutOperation *ComputePreviewFeatureRolloutOperation `field:"optional" json:"rolloutOperation" yaml:"rolloutOperation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/compute_preview_feature#timeouts ComputePreviewFeature#timeouts}
	Timeouts *ComputePreviewFeatureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

