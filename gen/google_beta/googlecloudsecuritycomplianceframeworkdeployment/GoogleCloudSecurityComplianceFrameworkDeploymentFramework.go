package googlecloudsecuritycomplianceframeworkdeployment


type GoogleCloudSecurityComplianceFrameworkDeploymentFramework struct {
	// In the format: organizations/{org}/locations/{location}/frameworks/{framework}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework_deployment#framework GoogleCloudSecurityComplianceFrameworkDeployment#framework}
	Framework *string `field:"required" json:"framework" yaml:"framework"`
	// Major revision id of the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework_deployment#major_revision_id GoogleCloudSecurityComplianceFrameworkDeployment#major_revision_id}
	MajorRevisionId *string `field:"required" json:"majorRevisionId" yaml:"majorRevisionId"`
}

