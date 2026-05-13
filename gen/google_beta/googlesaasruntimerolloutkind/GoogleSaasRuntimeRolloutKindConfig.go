package googlesaasruntimerolloutkind

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSaasRuntimeRolloutKindConfig struct {
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
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#location GoogleSaasRuntimeRolloutKind#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID value for the new rollout kind.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#rollout_kind_id GoogleSaasRuntimeRolloutKind#rollout_kind_id}
	RolloutKindId *string `field:"required" json:"rolloutKindId" yaml:"rolloutKindId"`
	// UnitKind that this rollout kind corresponds to.
	//
	// Rollouts stemming from this
	// rollout kind will target the units of this unit kind. In other words, this
	// defines the population of target units to be upgraded by rollouts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#unit_kind GoogleSaasRuntimeRolloutKind#unit_kind}
	UnitKind *string `field:"required" json:"unitKind" yaml:"unitKind"`
	// Annotations is an unstructured key-value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// More info: https://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#annotations GoogleSaasRuntimeRolloutKind#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// error_budget block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#error_budget GoogleSaasRuntimeRolloutKind#error_budget}
	ErrorBudget *GoogleSaasRuntimeRolloutKindErrorBudget `field:"optional" json:"errorBudget" yaml:"errorBudget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#id GoogleSaasRuntimeRolloutKind#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels on the resource, which can be used for categorization. similar to Kubernetes resource labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#labels GoogleSaasRuntimeRolloutKind#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#project GoogleSaasRuntimeRolloutKind#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The strategy used for executing a Rollout. This is a required field.
	//
	// There are two supported values strategies which are used to control a rollout.
	// - "Google.Cloud.Simple.AllAtOnce"
	// - "Google.Cloud.Simple.OneLocationAtATime"
	//
	// A rollout with one of these simple strategies will rollout across
	// all locations defined in the associated UnitKind's Saas Locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#rollout_orchestration_strategy GoogleSaasRuntimeRolloutKind#rollout_orchestration_strategy}
	RolloutOrchestrationStrategy *string `field:"optional" json:"rolloutOrchestrationStrategy" yaml:"rolloutOrchestrationStrategy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#timeouts GoogleSaasRuntimeRolloutKind#timeouts}
	Timeouts *GoogleSaasRuntimeRolloutKindTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// CEL(https://github.com/google/cel-spec) formatted filter string against Unit. The filter will be applied to determine the eligible unit population. This filter can only reduce, but not expand the scope of the rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#unit_filter GoogleSaasRuntimeRolloutKind#unit_filter}
	UnitFilter *string `field:"optional" json:"unitFilter" yaml:"unitFilter"`
	// The config for updating the unit kind.
	//
	// By default, the unit kind will be
	// updated on the rollout start.
	// Possible values:
	// UPDATE_UNIT_KIND_STRATEGY_ON_START
	// UPDATE_UNIT_KIND_STRATEGY_NEVER Possible values: ["UPDATE_UNIT_KIND_STRATEGY_ON_START", "UPDATE_UNIT_KIND_STRATEGY_NEVER"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_saas_runtime_rollout_kind#update_unit_kind_strategy GoogleSaasRuntimeRolloutKind#update_unit_kind_strategy}
	UpdateUnitKindStrategy *string `field:"optional" json:"updateUnitKindStrategy" yaml:"updateUnitKindStrategy"`
}

