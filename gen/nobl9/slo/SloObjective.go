package slo


type SloObjective struct {
	// The numeric target for your objective.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#target Slo#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// composite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#composite Slo#composite}
	Composite *SloObjectiveComposite `field:"optional" json:"composite" yaml:"composite"`
	// count_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#count_metrics Slo#count_metrics}
	CountMetrics interface{} `field:"optional" json:"countMetrics" yaml:"countMetrics"`
	// Name to be displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#display_name Slo#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Objective's name. This field is computed if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#name Slo#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// For threshold metrics, the logical operator applied to the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#op Slo#op}
	Op *string `field:"optional" json:"op" yaml:"op"`
	// Is objective marked as primary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#primary Slo#primary}
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// raw_metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#raw_metric Slo#raw_metric}
	RawMetric interface{} `field:"optional" json:"rawMetric" yaml:"rawMetric"`
	// Designated value for slice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#time_slice_target Slo#time_slice_target}
	TimeSliceTarget *float64 `field:"optional" json:"timeSliceTarget" yaml:"timeSliceTarget"`
	// Required for threshold and ratio metrics.
	//
	// Optional for composite SLOs. For threshold metrics, the threshold value. For ratio metrics, this must be a unique value per objective (for legacy reasons). For composite SLOs, it should be omitted. If, for composite SLO, it was set previously to a non-zero value, then it must remain unchanged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#value Slo#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

