package googlecomputenodegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeNodeGroupConfig struct {
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
	// The URL of the node template to which this node group belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#node_template GoogleComputeNodeGroup#node_template}
	NodeTemplate *string `field:"required" json:"nodeTemplate" yaml:"nodeTemplate"`
	// autoscaling_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#autoscaling_policy GoogleComputeNodeGroup#autoscaling_policy}
	AutoscalingPolicy *GoogleComputeNodeGroupAutoscalingPolicy `field:"optional" json:"autoscalingPolicy" yaml:"autoscalingPolicy"`
	// An optional textual description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#description GoogleComputeNodeGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#id GoogleComputeNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The initial number of nodes in the node group.
	//
	// One of 'initial_size' or 'autoscaling_policy' must be configured on resource creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#initial_size GoogleComputeNodeGroup#initial_size}
	InitialSize *float64 `field:"optional" json:"initialSize" yaml:"initialSize"`
	// Specifies the frequency of planned maintenance events.
	//
	// Set to one of the following:
	//   - AS_NEEDED: Hosts are eligible to receive infrastructure and hypervisor updates as they become available.
	//   - RECURRENT: Hosts receive planned infrastructure and hypervisor updates on a periodic basis, but not more frequently than every 28 days. This minimizes the number of planned maintenance operations on individual hosts and reduces the frequency of disruptions, both live migrations and terminations, on individual VMs. Possible values: ["AS_NEEDED", "RECURRENT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#maintenance_interval GoogleComputeNodeGroup#maintenance_interval}
	MaintenanceInterval *string `field:"optional" json:"maintenanceInterval" yaml:"maintenanceInterval"`
	// Specifies how to handle instances when a node in the group undergoes maintenance.
	//
	// Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#maintenance_policy GoogleComputeNodeGroup#maintenance_policy}
	MaintenancePolicy *string `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#maintenance_window GoogleComputeNodeGroup#maintenance_window}
	MaintenanceWindow *GoogleComputeNodeGroupMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#name GoogleComputeNodeGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#project GoogleComputeNodeGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// share_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#share_settings GoogleComputeNodeGroup#share_settings}
	ShareSettings *GoogleComputeNodeGroupShareSettings `field:"optional" json:"shareSettings" yaml:"shareSettings"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#timeouts GoogleComputeNodeGroup#timeouts}
	Timeouts *GoogleComputeNodeGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zone where this node group is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_group#zone GoogleComputeNodeGroup#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

