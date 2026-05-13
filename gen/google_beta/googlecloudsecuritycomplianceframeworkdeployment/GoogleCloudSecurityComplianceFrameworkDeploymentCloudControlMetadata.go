package googlecloudsecuritycomplianceframeworkdeployment


type GoogleCloudSecurityComplianceFrameworkDeploymentCloudControlMetadata struct {
	// cloud_control_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework_deployment#cloud_control_details GoogleCloudSecurityComplianceFrameworkDeployment#cloud_control_details}
	CloudControlDetails *GoogleCloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetails `field:"required" json:"cloudControlDetails" yaml:"cloudControlDetails"`
	// Enforcement mode for the framework deployment. Possible values: PREVENTIVE DETECTIVE AUDIT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework_deployment#enforcement_mode GoogleCloudSecurityComplianceFrameworkDeployment#enforcement_mode}
	EnforcementMode *string `field:"required" json:"enforcementMode" yaml:"enforcementMode"`
}

