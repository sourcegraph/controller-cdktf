package monitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorConfig struct {
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
	// The inputs map binds dataset OIDs to labels which can be referenced within stage pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#inputs Monitor#inputs}
	Inputs *map[string]*string `field:"required" json:"inputs" yaml:"inputs"`
	// Monitor name. Must be unique within workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#name Monitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#rule Monitor#rule}
	Rule *MonitorRule `field:"required" json:"rule" yaml:"rule"`
	// stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#stage Monitor#stage}
	Stage interface{} `field:"required" json:"stage" yaml:"stage"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#workspace Monitor#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// A long-form comment describing the content of the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#comment Monitor#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Monitor definition in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#definition Monitor#definition}
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// A brief description of the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#description Monitor#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Set to `true` to disable monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#disabled Monitor#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Target freshness for results.
	//
	// Tighten the freshness to increase the
	// frequency with which queries are run, which incurs higher transform costs.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#freshness Monitor#freshness}
	Freshness *string `field:"optional" json:"freshness" yaml:"freshness"`
	// Icon to be displayed for this object. Icons are sourced from the [fluency-filled](https://icons8.com/icons/fluency-systems-filled) icon set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#icon_url Monitor#icon_url}
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#id Monitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This resource is template for creating new monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#is_template Monitor#is_template}
	IsTemplate interface{} `field:"optional" json:"isTemplate" yaml:"isTemplate"`
	// notification_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor#notification_spec Monitor#notification_spec}
	NotificationSpec *MonitorNotificationSpec `field:"optional" json:"notificationSpec" yaml:"notificationSpec"`
}

