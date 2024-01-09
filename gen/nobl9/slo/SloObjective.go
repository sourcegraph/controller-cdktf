package slo


type SloObjective struct {
	// Name to be displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#display_name Slo#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Designated value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#target Slo#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// Value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#value Slo#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// count_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#count_metrics Slo#count_metrics}
	CountMetrics interface{} `field:"optional" json:"countMetrics" yaml:"countMetrics"`
	// Objective's name. This field is computed if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#name Slo#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Type of logical operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#op Slo#op}
	Op *string `field:"optional" json:"op" yaml:"op"`
	// raw_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#raw_metric Slo#raw_metric}
	RawMetric interface{} `field:"optional" json:"rawMetric" yaml:"rawMetric"`
	// Designated value for slice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#time_slice_target Slo#time_slice_target}
	TimeSliceTarget *float64 `field:"optional" json:"timeSliceTarget" yaml:"timeSliceTarget"`
}

