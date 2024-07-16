package googleedgenetworksubnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleEdgenetworkSubnetConfig struct {
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
	// The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#location GoogleEdgenetworkSubnet#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the network to which this router belongs. Must be of the form: 'projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#network GoogleEdgenetworkSubnet#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// A unique ID that identifies this subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#subnet_id GoogleEdgenetworkSubnet#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The name of the target Distributed Cloud Edge zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#zone GoogleEdgenetworkSubnet#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#description GoogleEdgenetworkSubnet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#id GoogleEdgenetworkSubnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ranges of ipv4 addresses that are owned by this subnetwork, in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#ipv4_cidr GoogleEdgenetworkSubnet#ipv4_cidr}
	Ipv4Cidr *[]*string `field:"optional" json:"ipv4Cidr" yaml:"ipv4Cidr"`
	// The ranges of ipv6 addresses that are owned by this subnetwork, in CIDR format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#ipv6_cidr GoogleEdgenetworkSubnet#ipv6_cidr}
	Ipv6Cidr *[]*string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
	// Labels associated with this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#labels GoogleEdgenetworkSubnet#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#project GoogleEdgenetworkSubnet#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#timeouts GoogleEdgenetworkSubnet#timeouts}
	Timeouts *GoogleEdgenetworkSubnetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// VLAN ID for this subnetwork. If not specified, one is assigned automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_edgenetwork_subnet#vlan_id GoogleEdgenetworkSubnet#vlan_id}
	VlanId *float64 `field:"optional" json:"vlanId" yaml:"vlanId"`
}

