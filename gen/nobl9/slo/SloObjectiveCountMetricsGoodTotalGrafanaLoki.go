package slo


type SloObjectiveCountMetricsGoodTotalGrafanaLoki struct {
	// Query for the logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#logql Slo#logql}
	Logql *string `field:"required" json:"logql" yaml:"logql"`
}

