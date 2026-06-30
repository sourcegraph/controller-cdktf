package osconfigv2policyorchestratorforfolder

type OsConfigV2PolicyOrchestratorForFolderOrchestratedResourceOsPolicyAssignmentV1PayloadInstanceFilter struct {
	// Target all VMs in the project. If true, no other criteria is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#all OsConfigV2PolicyOrchestratorForFolder#all}
	All any `field:"optional" json:"all" yaml:"all"`
	// exclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#exclusion_labels OsConfigV2PolicyOrchestratorForFolder#exclusion_labels}
	ExclusionLabels any `field:"optional" json:"exclusionLabels" yaml:"exclusionLabels"`
	// inclusion_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#inclusion_labels OsConfigV2PolicyOrchestratorForFolder#inclusion_labels}
	InclusionLabels any `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
	// inventories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/os_config_v2_policy_orchestrator_for_folder#inventories OsConfigV2PolicyOrchestratorForFolder#inventories}
	Inventories any `field:"optional" json:"inventories" yaml:"inventories"`
}
