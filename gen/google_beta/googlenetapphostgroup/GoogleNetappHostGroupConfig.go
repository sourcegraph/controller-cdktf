package googlenetapphostgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetappHostGroupConfig struct {
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
	// The list of hosts associated with the host group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#hosts GoogleNetappHostGroup#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// Location (region) of the Host Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#location GoogleNetappHostGroup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the Host Group. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#name GoogleNetappHostGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The OS type of the host group.
	//
	// It indicates the type of operating system
	// used by all of the hosts in the HostGroup. All hosts in a HostGroup must be
	// of the same OS type. This can be set only when creating a HostGroup. Possible values: ["LINUX", "WINDOWS", "ESXI"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#os_type GoogleNetappHostGroup#os_type}
	OsType *string `field:"required" json:"osType" yaml:"osType"`
	// Type of the host group. Possible values: ["ISCSI_INITIATOR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#type GoogleNetappHostGroup#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#description GoogleNetappHostGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#id GoogleNetappHostGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#labels GoogleNetappHostGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#project GoogleNetappHostGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_host_group#timeouts GoogleNetappHostGroup#timeouts}
	Timeouts *GoogleNetappHostGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

