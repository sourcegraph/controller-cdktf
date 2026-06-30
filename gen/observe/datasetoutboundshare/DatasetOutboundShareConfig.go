package datasetoutboundshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasetOutboundShareConfig struct {
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
	// The OID of the dataset to be shared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#dataset DatasetOutboundShare#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#freshness_goal DatasetOutboundShare#freshness_goal}.
	FreshnessGoal *string `field:"required" json:"freshnessGoal" yaml:"freshnessGoal"`
	// A descriptive name for the dataset sharing configuration. Displayed within Observe, not used in Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#name DatasetOutboundShare#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The OID of the Observe Snowflake outbound share where the dataset will be shared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#outbound_share DatasetOutboundShare#outbound_share}
	OutboundShare *string `field:"required" json:"outboundShare" yaml:"outboundShare"`
	// The name of the schema within the shared database where the dataset view will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#schema_name DatasetOutboundShare#schema_name}
	SchemaName *string `field:"required" json:"schemaName" yaml:"schemaName"`
	// The name of the view that will be created in the shared database, within the specified schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#view_name DatasetOutboundShare#view_name}
	ViewName *string `field:"required" json:"viewName" yaml:"viewName"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#workspace DatasetOutboundShare#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// If set to true, the shared view will have change tracking enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#change_tracking DatasetOutboundShare#change_tracking}
	ChangeTracking any `field:"optional" json:"changeTracking" yaml:"changeTracking"`
	// A description of the dataset sharing configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#description DatasetOutboundShare#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#folder DatasetOutboundShare#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#id DatasetOutboundShare#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_outbound_share#timeouts DatasetOutboundShare#timeouts}
	Timeouts *DatasetOutboundShareTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
