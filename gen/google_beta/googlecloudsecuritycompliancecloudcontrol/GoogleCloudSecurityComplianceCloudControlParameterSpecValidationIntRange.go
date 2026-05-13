package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRange struct {
	// Maximum allowed value for the numeric parameter (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#max GoogleCloudSecurityComplianceCloudControl#max}
	Max *string `field:"required" json:"max" yaml:"max"`
	// Minimum allowed value for the numeric parameter (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#min GoogleCloudSecurityComplianceCloudControl#min}
	Min *string `field:"required" json:"min" yaml:"min"`
}

