package googletpuv2vm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleTpuV2VmConfig struct {
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
	// The immutable name of the TPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#name GoogleTpuV2Vm#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Runtime version for the TPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#runtime_version GoogleTpuV2Vm#runtime_version}
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// accelerator_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#accelerator_config GoogleTpuV2Vm#accelerator_config}
	AcceleratorConfig *GoogleTpuV2VmAcceleratorConfig `field:"optional" json:"acceleratorConfig" yaml:"acceleratorConfig"`
	// TPU accelerator type for the TPU.
	//
	// 'accelerator_type' cannot be used at the same time as
	// 'accelerator_config'. If neither is specified, 'accelerator_type' defaults to 'v2-8'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#accelerator_type GoogleTpuV2Vm#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP address.
	//
	// This CIDR block must
	// be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger
	// block would be wasteful (a node can only consume one IP address). Errors will occur if the
	// CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts
	// with any subnetworks in the user's provided network, or the provided network is peered with
	// another network that is using that CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#cidr_block GoogleTpuV2Vm#cidr_block}
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// data_disks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#data_disks GoogleTpuV2Vm#data_disks}
	DataDisks interface{} `field:"optional" json:"dataDisks" yaml:"dataDisks"`
	// Text description of the TPU.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#description GoogleTpuV2Vm#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#id GoogleTpuV2Vm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user-provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#labels GoogleTpuV2Vm#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#metadata GoogleTpuV2Vm#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#network_config GoogleTpuV2Vm#network_config}
	NetworkConfig *GoogleTpuV2VmNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#project GoogleTpuV2Vm#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// scheduling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#scheduling_config GoogleTpuV2Vm#scheduling_config}
	SchedulingConfig *GoogleTpuV2VmSchedulingConfig `field:"optional" json:"schedulingConfig" yaml:"schedulingConfig"`
	// service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#service_account GoogleTpuV2Vm#service_account}
	ServiceAccount *GoogleTpuV2VmServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#shielded_instance_config GoogleTpuV2Vm#shielded_instance_config}
	ShieldedInstanceConfig *GoogleTpuV2VmShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#tags GoogleTpuV2Vm#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#timeouts GoogleTpuV2Vm#timeouts}
	Timeouts *GoogleTpuV2VmTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_tpu_v2_vm#zone GoogleTpuV2Vm#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

