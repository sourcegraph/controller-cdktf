package googleworkstationsworkstationiampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleWorkstationsWorkstationIamPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_policy#policy_data GoogleWorkstationsWorkstationIamPolicy#policy_data}.
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_policy#workstation_cluster_id GoogleWorkstationsWorkstationIamPolicy#workstation_cluster_id}.
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_policy#workstation_config_id GoogleWorkstationsWorkstationIamPolicy#workstation_config_id}.
	WorkstationConfigId *string `field:"required" json:"workstationConfigId" yaml:"workstationConfigId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_policy#workstation_id GoogleWorkstationsWorkstationIamPolicy#workstation_id}.
	WorkstationId *string `field:"required" json:"workstationId" yaml:"workstationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_policy#id GoogleWorkstationsWorkstationIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_policy#location GoogleWorkstationsWorkstationIamPolicy#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workstations_workstation_iam_policy#project GoogleWorkstationsWorkstationIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

