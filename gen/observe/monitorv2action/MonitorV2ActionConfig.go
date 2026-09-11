package monitorv2action

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MonitorV2ActionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#name MonitorV2Action#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#type MonitorV2Action#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#workspace MonitorV2Action#workspace}.
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#description MonitorV2Action#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#email MonitorV2Action#email}
	Email *MonitorV2ActionEmail `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#id MonitorV2Action#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// webhook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/monitor_v2_action#webhook MonitorV2Action#webhook}
	Webhook *MonitorV2ActionWebhook `field:"optional" json:"webhook" yaml:"webhook"`
}

