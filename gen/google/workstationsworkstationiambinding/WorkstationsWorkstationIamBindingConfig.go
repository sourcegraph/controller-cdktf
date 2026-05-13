package workstationsworkstationiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkstationsWorkstationIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#members WorkstationsWorkstationIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#role WorkstationsWorkstationIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#workstation_cluster_id WorkstationsWorkstationIamBinding#workstation_cluster_id}.
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#workstation_config_id WorkstationsWorkstationIamBinding#workstation_config_id}.
	WorkstationConfigId *string `field:"required" json:"workstationConfigId" yaml:"workstationConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#workstation_id WorkstationsWorkstationIamBinding#workstation_id}.
	WorkstationId *string `field:"required" json:"workstationId" yaml:"workstationId"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#condition WorkstationsWorkstationIamBinding#condition}
	Condition *WorkstationsWorkstationIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#id WorkstationsWorkstationIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#location WorkstationsWorkstationIamBinding#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/workstations_workstation_iam_binding#project WorkstationsWorkstationIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

