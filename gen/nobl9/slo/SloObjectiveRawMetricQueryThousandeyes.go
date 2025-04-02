package slo


type SloObjectiveRawMetricQueryThousandeyes struct {
	// ID of the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#test_id Slo#test_id}
	TestId *float64 `field:"required" json:"testId" yaml:"testId"`
	// Type of the test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#test_type Slo#test_type}
	TestType *string `field:"optional" json:"testType" yaml:"testType"`
}

