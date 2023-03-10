package computesharedvpcserviceproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSharedVpcServiceProjectConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The ID of a host project to associate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_shared_vpc_service_project#host_project ComputeSharedVpcServiceProject#host_project}
	HostProject *string `field:"required" json:"hostProject" yaml:"hostProject"`
	// The ID of the project that will serve as a Shared VPC service project.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_shared_vpc_service_project#service_project ComputeSharedVpcServiceProject#service_project}
	ServiceProject *string `field:"required" json:"serviceProject" yaml:"serviceProject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_shared_vpc_service_project#id ComputeSharedVpcServiceProject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_shared_vpc_service_project#timeouts ComputeSharedVpcServiceProject#timeouts}
	Timeouts *ComputeSharedVpcServiceProjectTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

