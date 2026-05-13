package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationIntRange struct {
	// Maximum allowed value for the numeric parameter (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#max CloudSecurityComplianceCloudControl#max}
	Max *string `field:"required" json:"max" yaml:"max"`
	// Minimum allowed value for the numeric parameter (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#min CloudSecurityComplianceCloudControl#min}
	Min *string `field:"required" json:"min" yaml:"min"`
}

