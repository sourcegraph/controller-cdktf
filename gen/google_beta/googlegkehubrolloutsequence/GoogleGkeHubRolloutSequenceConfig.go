package googlegkehubrolloutsequence

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGkeHubRolloutSequenceConfig struct {
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
	// The user-provided identifier of the RolloutSequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#rollout_sequence_id GoogleGkeHubRolloutSequence#rollout_sequence_id}
	RolloutSequenceId *string `field:"required" json:"rolloutSequenceId" yaml:"rolloutSequenceId"`
	// stages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#stages GoogleGkeHubRolloutSequence#stages}
	Stages interface{} `field:"required" json:"stages" yaml:"stages"`
	// Human readable display name of the Rollout Sequence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#display_name GoogleGkeHubRolloutSequence#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#id GoogleGkeHubRolloutSequence#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ignored_clusters_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#ignored_clusters_selector GoogleGkeHubRolloutSequence#ignored_clusters_selector}
	IgnoredClustersSelector *GoogleGkeHubRolloutSequenceIgnoredClustersSelector `field:"optional" json:"ignoredClustersSelector" yaml:"ignoredClustersSelector"`
	// Labels for this Rollout Sequence.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#labels GoogleGkeHubRolloutSequence#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#project GoogleGkeHubRolloutSequence#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_gke_hub_rollout_sequence#timeouts GoogleGkeHubRolloutSequence#timeouts}
	Timeouts *GoogleGkeHubRolloutSequenceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

