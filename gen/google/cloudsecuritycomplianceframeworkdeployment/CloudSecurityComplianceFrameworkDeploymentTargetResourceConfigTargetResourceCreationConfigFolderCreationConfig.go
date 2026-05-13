package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfigFolderCreationConfig struct {
	// Display name of the folder to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#folder_display_name CloudSecurityComplianceFrameworkDeployment#folder_display_name}
	FolderDisplayName *string `field:"required" json:"folderDisplayName" yaml:"folderDisplayName"`
	// The parent of the folder to be created. It can be an organizations/{org} or folders/{folder}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#parent CloudSecurityComplianceFrameworkDeployment#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
}

