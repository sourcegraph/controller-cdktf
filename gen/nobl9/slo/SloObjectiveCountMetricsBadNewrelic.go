package slo


type SloObjectiveCountMetricsBadNewrelic struct {
	// Query for the metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#nrql Slo#nrql}
	Nrql *string `field:"required" json:"nrql" yaml:"nrql"`
}

