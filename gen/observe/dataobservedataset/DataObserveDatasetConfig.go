package dataobservedataset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataObserveDatasetConfig struct {
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
	// correlation_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/dataset#correlation_tag DataObserveDataset#correlation_tag}
	CorrelationTag interface{} `field:"optional" json:"correlationTag" yaml:"correlationTag"`
	// Resource ID for this object. One of `name` or `id` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/dataset#id DataObserveDataset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Dataset name.
	//
	// Must be unique within workspace.
	// One of `name` or `id` must be set. If `name` is provided, `workspace` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/dataset#name DataObserveDataset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/dataset#stage DataObserveDataset#stage}
	Stage interface{} `field:"optional" json:"stage" yaml:"stage"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/dataset#workspace DataObserveDataset#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

