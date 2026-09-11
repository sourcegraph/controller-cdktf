package monitoraction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorActionConfig struct {
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
	// Monitor action name. Must be unique within workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#name MonitorAction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#workspace MonitorAction#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// A brief description of the monitor action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#description MonitorAction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#email MonitorAction#email}
	Email *MonitorActionEmail `field:"optional" json:"email" yaml:"email"`
	// Icon to be displayed for this object. Icons are sourced from the [fluency-filled](https://icons8.com/icons/fluency-systems-filled) icon set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#icon_url MonitorAction#icon_url}
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#id MonitorAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enables a final update when a monitor action notification is closed (no longer triggered).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#notify_on_close MonitorAction#notify_on_close}
	NotifyOnClose interface{} `field:"optional" json:"notifyOnClose" yaml:"notifyOnClose"`
	// Limits 10 alerts to the defined time period.
	//
	// For email actions the minimum
	// is 10m. For webhook actions the minimum is 1s.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#rate_limit MonitorAction#rate_limit}
	RateLimit *string `field:"optional" json:"rateLimit" yaml:"rateLimit"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_action#webhook MonitorAction#webhook}
	Webhook *MonitorActionWebhook `field:"optional" json:"webhook" yaml:"webhook"`
}

