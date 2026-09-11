package dataobservemonitoraction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataObserveMonitorActionConfig struct {
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
	// Resource ID for this object. One of `id` or `name` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_action#id DataObserveMonitorAction#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Monitor name.
	//
	// Must be unique within workspace.
	// One of `name` or `id` must be set. If `name` is provided, `workspace` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_action#name DataObserveMonitorAction#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_action#workspace DataObserveMonitorAction#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

