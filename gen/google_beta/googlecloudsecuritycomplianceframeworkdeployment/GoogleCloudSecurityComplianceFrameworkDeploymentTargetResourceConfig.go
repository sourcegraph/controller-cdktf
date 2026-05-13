package googlecloudsecuritycomplianceframeworkdeployment


type GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfig struct {
	// CRM node in format organizations/{organization}, folders/{folder}, or projects/{project}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework_deployment#existing_target_resource GoogleCloudSecurityComplianceFrameworkDeployment#existing_target_resource}
	ExistingTargetResource *string `field:"optional" json:"existingTargetResource" yaml:"existingTargetResource"`
	// target_resource_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework_deployment#target_resource_creation_config GoogleCloudSecurityComplianceFrameworkDeployment#target_resource_creation_config}
	TargetResourceCreationConfig *GoogleCloudSecurityComplianceFrameworkDeploymentTargetResourceConfigTargetResourceCreationConfig `field:"optional" json:"targetResourceCreationConfig" yaml:"targetResourceCreationConfig"`
}

