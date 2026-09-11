package dataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The inputs map binds dataset OIDs to labels which can be referenced within stage pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#inputs Dataset#inputs}
	Inputs *map[string]*string `field:"required" json:"inputs" yaml:"inputs"`
	// Dataset name. Must be unique within workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#name Dataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#stage Dataset#stage}
	Stage interface{} `field:"required" json:"stage" yaml:"stage"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#workspace Dataset#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// Disables periodic materialization of the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#acceleration_disabled Dataset#acceleration_disabled}
	AccelerationDisabled interface{} `field:"optional" json:"accelerationDisabled" yaml:"accelerationDisabled"`
	// Source of disabled materialization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#acceleration_disabled_source Dataset#acceleration_disabled_source}
	AccelerationDisabledSource *string `field:"optional" json:"accelerationDisabledSource" yaml:"accelerationDisabledSource"`
	// JSON representation of state used for dataset formatting in the UI.
	//
	// Not intended to be configured by hand, please use export functionality.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#data_table_view_state Dataset#data_table_view_state}
	DataTableViewState *string `field:"optional" json:"dataTableViewState" yaml:"dataTableViewState"`
	// Dataset description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#description Dataset#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Target freshness for results.
	//
	// Tighten the freshness to increase the
	// frequency with which queries are run, which incurs higher transform costs.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#freshness Dataset#freshness}
	Freshness *string `field:"optional" json:"freshness" yaml:"freshness"`
	// Icon to be displayed for this object. Icons are sourced from the [fluency-filled](https://icons8.com/icons/fluency-systems-filled) icon set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#icon_url Dataset#icon_url}
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#id Dataset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum on-demand materialization length for the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#on_demand_materialization_length Dataset#on_demand_materialization_length}
	OnDemandMaterializationLength *string `field:"optional" json:"onDemandMaterializationLength" yaml:"onDemandMaterializationLength"`
	// Path cost incurred by this dataset when computing graph link.
	//
	// Increasing
	// this value will reduce the preference for using this dataset when computing
	// paths between two datasets.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#path_cost Dataset#path_cost}
	PathCost *float64 `field:"optional" json:"pathCost" yaml:"pathCost"`
	// Specifies rematerialization mode when updating a dataset.
	//
	// Options include
	// "rematerialize" (default), "skip_rematerialization", and "must_skip_rematerialization".
	// "skip_rematerialization" will skip rematerialization if certain conditions are met, will rematerialize otherwise.
	// "must_skip_rematerialization" will never rematerialize, update will fail if skipping rematerialization is not possible.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#rematerialization_mode Dataset#rematerialization_mode}
	RematerializationMode *string `field:"optional" json:"rematerializationMode" yaml:"rematerializationMode"`
	// The ID of the storage integration associated with this dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset#storage_integration Dataset#storage_integration}
	StorageIntegration *string `field:"optional" json:"storageIntegration" yaml:"storageIntegration"`
}

