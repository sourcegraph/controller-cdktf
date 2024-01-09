package slo


type SloAnomalyConfig struct {
	// no_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#no_data Slo#no_data}
	NoData *SloAnomalyConfigNoData `field:"required" json:"noData" yaml:"noData"`
}

