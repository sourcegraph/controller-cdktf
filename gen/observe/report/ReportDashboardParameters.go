package report


type ReportDashboardParameters struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#key Report#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/report#value Report#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

