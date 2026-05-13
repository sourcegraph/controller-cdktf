package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentCloudControlMetadata struct {
	// cloud_control_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#cloud_control_details CloudSecurityComplianceFrameworkDeployment#cloud_control_details}
	CloudControlDetails *CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetails `field:"required" json:"cloudControlDetails" yaml:"cloudControlDetails"`
	// Enforcement mode for the framework deployment. Possible values: PREVENTIVE DETECTIVE AUDIT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#enforcement_mode CloudSecurityComplianceFrameworkDeployment#enforcement_mode}
	EnforcementMode *string `field:"required" json:"enforcementMode" yaml:"enforcementMode"`
}

