package slo


type SloObjectiveRawMetricQuerySplunkObservability struct {
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#program Slo#program}
	Program *string `field:"required" json:"program" yaml:"program"`
}

