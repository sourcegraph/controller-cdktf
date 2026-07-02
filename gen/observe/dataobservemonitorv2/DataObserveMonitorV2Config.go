package dataobservemonitorv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataObserveMonitorV2Config struct {
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
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#actions DataObserveMonitorV2#actions}
	Actions any `field:"optional" json:"actions" yaml:"actions"`
	// groupings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#groupings DataObserveMonitorV2#groupings}
	Groupings any `field:"optional" json:"groupings" yaml:"groupings"`
	// Resource ID for this object. One of `name` or `id` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#id DataObserveMonitorV2#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Monitor name. One of `name` or `id` must be set. If `name` is provided, `workspace` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#name DataObserveMonitorV2#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// no_data_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#no_data_rules DataObserveMonitorV2#no_data_rules}
	NoDataRules any `field:"optional" json:"noDataRules" yaml:"noDataRules"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#rules DataObserveMonitorV2#rules}
	Rules any `field:"optional" json:"rules" yaml:"rules"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#scheduling DataObserveMonitorV2#scheduling}
	Scheduling any `field:"optional" json:"scheduling" yaml:"scheduling"`
	// stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#stage DataObserveMonitorV2#stage}
	Stage any `field:"optional" json:"stage" yaml:"stage"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/monitor_v2#workspace DataObserveMonitorV2#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}
