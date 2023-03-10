package configconformancepack


type ConfigConformancePackInputParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack#parameter_name ConfigConformancePack#parameter_name}.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/config_conformance_pack#parameter_value ConfigConformancePack#parameter_value}.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

