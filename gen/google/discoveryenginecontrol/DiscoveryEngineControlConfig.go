package discoveryenginecontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineControlConfig struct {
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
	// The unique id of the control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#control_id DiscoveryEngineControl#control_id}
	ControlId *string `field:"required" json:"controlId" yaml:"controlId"`
	// The display name of the control.
	//
	// This field must be a UTF-8 encoded
	// string with a length limit of 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#display_name DiscoveryEngineControl#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The engine to add the control to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#engine_id DiscoveryEngineControl#engine_id}
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#location DiscoveryEngineControl#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The solution type that the control belongs to. Possible values: ["SOLUTION_TYPE_RECOMMENDATION", "SOLUTION_TYPE_SEARCH", "SOLUTION_TYPE_CHAT", "SOLUTION_TYPE_GENERATIVE_CHAT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#solution_type DiscoveryEngineControl#solution_type}
	SolutionType *string `field:"required" json:"solutionType" yaml:"solutionType"`
	// boost_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#boost_action DiscoveryEngineControl#boost_action}
	BoostAction *DiscoveryEngineControlBoostAction `field:"optional" json:"boostAction" yaml:"boostAction"`
	// The collection ID. Currently only accepts "default_collection".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#collection_id DiscoveryEngineControl#collection_id}
	CollectionId *string `field:"optional" json:"collectionId" yaml:"collectionId"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#conditions DiscoveryEngineControl#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// filter_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#filter_action DiscoveryEngineControl#filter_action}
	FilterAction *DiscoveryEngineControlFilterAction `field:"optional" json:"filterAction" yaml:"filterAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#id DiscoveryEngineControl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#project DiscoveryEngineControl#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// promote_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#promote_action DiscoveryEngineControl#promote_action}
	PromoteAction *DiscoveryEngineControlPromoteAction `field:"optional" json:"promoteAction" yaml:"promoteAction"`
	// redirect_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#redirect_action DiscoveryEngineControl#redirect_action}
	RedirectAction *DiscoveryEngineControlRedirectAction `field:"optional" json:"redirectAction" yaml:"redirectAction"`
	// synonyms_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#synonyms_action DiscoveryEngineControl#synonyms_action}
	SynonymsAction *DiscoveryEngineControlSynonymsAction `field:"optional" json:"synonymsAction" yaml:"synonymsAction"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#timeouts DiscoveryEngineControl#timeouts}
	Timeouts *DiscoveryEngineControlTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The use cases that the control is used for. Possible values: ["SEARCH_USE_CASE_SEARCH", "SEARCH_USE_CASE_BROWSE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#use_cases DiscoveryEngineControl#use_cases}
	UseCases *[]*string `field:"optional" json:"useCases" yaml:"useCases"`
}

