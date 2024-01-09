package slo


type SloAnomalyConfigNoData struct {
	// alert_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#alert_method Slo#alert_method}
	AlertMethod interface{} `field:"required" json:"alertMethod" yaml:"alertMethod"`
}

