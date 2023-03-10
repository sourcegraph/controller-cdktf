package googlecomputeperinstanceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputePerInstanceConfigConfig struct {
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
	// The instance group manager this instance config is part of.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#instance_group_manager GoogleComputePerInstanceConfig#instance_group_manager}
	InstanceGroupManager *string `field:"required" json:"instanceGroupManager" yaml:"instanceGroupManager"`
	// The name for this per-instance config and its corresponding instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#name GoogleComputePerInstanceConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#id GoogleComputePerInstanceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#minimal_action GoogleComputePerInstanceConfig#minimal_action}.
	MinimalAction *string `field:"optional" json:"minimalAction" yaml:"minimalAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#most_disruptive_allowed_action GoogleComputePerInstanceConfig#most_disruptive_allowed_action}.
	MostDisruptiveAllowedAction *string `field:"optional" json:"mostDisruptiveAllowedAction" yaml:"mostDisruptiveAllowedAction"`
	// preserved_state block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#preserved_state GoogleComputePerInstanceConfig#preserved_state}
	PreservedState *GoogleComputePerInstanceConfigPreservedState `field:"optional" json:"preservedState" yaml:"preservedState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#project GoogleComputePerInstanceConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#remove_instance_state_on_destroy GoogleComputePerInstanceConfig#remove_instance_state_on_destroy}.
	RemoveInstanceStateOnDestroy interface{} `field:"optional" json:"removeInstanceStateOnDestroy" yaml:"removeInstanceStateOnDestroy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#timeouts GoogleComputePerInstanceConfig#timeouts}
	Timeouts *GoogleComputePerInstanceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zone where the containing instance group manager is located.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_per_instance_config#zone GoogleComputePerInstanceConfig#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

