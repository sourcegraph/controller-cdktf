package dataobservemonitorv2action

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataObserveMonitorV2ActionConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2_action#email DataObserveMonitorV2Action#email}
	Email any `field:"optional" json:"email" yaml:"email"`
	// Resource ID for this object.  One of either `id` or `name` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2_action#id DataObserveMonitorV2Action#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the monitor v2 action.  One of either `id` or `name` must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2_action#name DataObserveMonitorV2Action#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2_action#webhook DataObserveMonitorV2Action#webhook}
	Webhook any `field:"optional" json:"webhook" yaml:"webhook"`
	// OID of the workspace this object is contained in.  Must be specified if looking up by name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2_action#workspace DataObserveMonitorV2Action#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}
