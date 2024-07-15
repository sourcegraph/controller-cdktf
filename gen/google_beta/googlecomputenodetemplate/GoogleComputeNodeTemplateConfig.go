package googlecomputenodetemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeNodeTemplateConfig struct {
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
	// CPU overcommit. Default value: "NONE" Possible values: ["ENABLED", "NONE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#cpu_overcommit_type GoogleComputeNodeTemplate#cpu_overcommit_type}
	CpuOvercommitType *string `field:"optional" json:"cpuOvercommitType" yaml:"cpuOvercommitType"`
	// An optional textual description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#description GoogleComputeNodeTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#id GoogleComputeNodeTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#name GoogleComputeNodeTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#node_affinity_labels GoogleComputeNodeTemplate#node_affinity_labels}
	NodeAffinityLabels *map[string]*string `field:"optional" json:"nodeAffinityLabels" yaml:"nodeAffinityLabels"`
	// Node type to use for nodes group that are created from this template.
	//
	// Only one of nodeTypeFlexibility and nodeType can be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#node_type GoogleComputeNodeTemplate#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// node_type_flexibility block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#node_type_flexibility GoogleComputeNodeTemplate#node_type_flexibility}
	NodeTypeFlexibility *GoogleComputeNodeTemplateNodeTypeFlexibility `field:"optional" json:"nodeTypeFlexibility" yaml:"nodeTypeFlexibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#project GoogleComputeNodeTemplate#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where nodes using the node template will be created. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#region GoogleComputeNodeTemplate#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// server_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#server_binding GoogleComputeNodeTemplate#server_binding}
	ServerBinding *GoogleComputeNodeTemplateServerBinding `field:"optional" json:"serverBinding" yaml:"serverBinding"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_node_template#timeouts GoogleComputeNodeTemplate#timeouts}
	Timeouts *GoogleComputeNodeTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

