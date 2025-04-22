package googleosconfigv2policyorchestrator


type GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecValidate struct {
	// Required. The script interpreter to use. Possible values: INTERPRETER_UNSPECIFIED NONE SHELL POWERSHELL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#interpreter GoogleOsConfigV2PolicyOrchestrator#interpreter}
	Interpreter *string `field:"required" json:"interpreter" yaml:"interpreter"`
	// Optional arguments to pass to the source during execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#args GoogleOsConfigV2PolicyOrchestrator#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#file GoogleOsConfigV2PolicyOrchestrator#file}
	File *GoogleOsConfigV2PolicyOrchestratorOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecValidateFile `field:"optional" json:"file" yaml:"file"`
	// Only recorded for enforce Exec.
	//
	// Path to an output file (that is created by this Exec) whose
	// content will be recorded in OSPolicyResourceCompliance after a
	// successful run. Absence or failure to read this file will result in
	// this ExecResource being non-compliant. Output file size is limited to
	// 500K bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#output_file_path GoogleOsConfigV2PolicyOrchestrator#output_file_path}
	OutputFilePath *string `field:"optional" json:"outputFilePath" yaml:"outputFilePath"`
	// An inline script. The size of the script is limited to 32KiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_os_config_v2_policy_orchestrator#script GoogleOsConfigV2PolicyOrchestrator#script}
	Script *string `field:"optional" json:"script" yaml:"script"`
}

