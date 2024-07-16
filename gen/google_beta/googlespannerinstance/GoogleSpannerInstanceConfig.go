package googlespannerinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSpannerInstanceConfig struct {
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
	// The name of the instance's configuration (similar but not quite the same as a region) which defines the geographic placement and replication of your databases in this instance.
	//
	// It determines where your data
	// is stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.
	// In order to obtain a valid list please consult the
	// [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#config GoogleSpannerInstance#config}
	Config *string `field:"required" json:"config" yaml:"config"`
	// The descriptive name for this instance as it appears in UIs.
	//
	// Must be
	// unique per project and between 4 and 30 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#display_name GoogleSpannerInstance#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// autoscaling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#autoscaling_config GoogleSpannerInstance#autoscaling_config}
	AutoscalingConfig *GoogleSpannerInstanceAutoscalingConfig `field:"optional" json:"autoscalingConfig" yaml:"autoscalingConfig"`
	// When deleting a spanner instance, this boolean option will delete all backups of this instance.
	//
	// This must be set to true if you created a backup manually in the console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#force_destroy GoogleSpannerInstance#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#id GoogleSpannerInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#labels GoogleSpannerInstance#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// A unique identifier for the instance, which cannot be changed after the instance is created.
	//
	// The name must be between 6 and 30 characters
	// in length.
	//
	//
	// If not provided, a random string starting with 'tf-' will be selected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#name GoogleSpannerInstance#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of nodes allocated to this instance. Exactly one of either node_count or processing_units must be present in terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#num_nodes GoogleSpannerInstance#num_nodes}
	NumNodes *float64 `field:"optional" json:"numNodes" yaml:"numNodes"`
	// The number of processing units allocated to this instance. Exactly one of processing_units or node_count must be present in terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#processing_units GoogleSpannerInstance#processing_units}
	ProcessingUnits *float64 `field:"optional" json:"processingUnits" yaml:"processingUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#project GoogleSpannerInstance#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_spanner_instance#timeouts GoogleSpannerInstance#timeouts}
	Timeouts *GoogleSpannerInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

