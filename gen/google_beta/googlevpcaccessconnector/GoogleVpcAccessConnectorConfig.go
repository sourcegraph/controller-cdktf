package googlevpcaccessconnector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVpcAccessConnectorConfig struct {
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
	// The name of the resource (Max 25 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#name GoogleVpcAccessConnector#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#id GoogleVpcAccessConnector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The range of internal addresses that follows RFC 4632 notation. Example: '10.132.0.0/28'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#ip_cidr_range GoogleVpcAccessConnector#ip_cidr_range}
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Machine type of VM Instance underlying connector. Default is e2-micro.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#machine_type GoogleVpcAccessConnector#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Maximum value of instances in autoscaling group underlying the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#max_instances GoogleVpcAccessConnector#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// Maximum throughput of the connector in Mbps, must be greater than 'min_throughput'. Default is 300.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#max_throughput GoogleVpcAccessConnector#max_throughput}
	MaxThroughput *float64 `field:"optional" json:"maxThroughput" yaml:"maxThroughput"`
	// Minimum value of instances in autoscaling group underlying the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#min_instances GoogleVpcAccessConnector#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// Minimum throughput of the connector in Mbps. Default and min is 200.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#min_throughput GoogleVpcAccessConnector#min_throughput}
	MinThroughput *float64 `field:"optional" json:"minThroughput" yaml:"minThroughput"`
	// Name or self_link of the VPC network. Required if 'ip_cidr_range' is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#network GoogleVpcAccessConnector#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#project GoogleVpcAccessConnector#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the VPC Access connector resides. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#region GoogleVpcAccessConnector#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// subnet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#subnet GoogleVpcAccessConnector#subnet}
	Subnet *GoogleVpcAccessConnectorSubnet `field:"optional" json:"subnet" yaml:"subnet"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_vpc_access_connector#timeouts GoogleVpcAccessConnector#timeouts}
	Timeouts *GoogleVpcAccessConnectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

