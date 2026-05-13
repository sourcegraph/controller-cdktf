package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig struct {
	// folder_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#folder_creation_config CloudSecurityComplianceFrameworkDeployment#folder_creation_config}
	FolderCreationConfig *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfig `field:"optional" json:"folderCreationConfig" yaml:"folderCreationConfig"`
	// project_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#project_creation_config CloudSecurityComplianceFrameworkDeployment#project_creation_config}
	ProjectCreationConfig *CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigProjectCreationConfig `field:"optional" json:"projectCreationConfig" yaml:"projectCreationConfig"`
}

