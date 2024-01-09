package slo


type SloObjectiveCountMetrics struct {
	// Should the metrics be incrementing or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#incremental Slo#incremental}
	Incremental interface{} `field:"required" json:"incremental" yaml:"incremental"`
	// total block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#total Slo#total}
	Total interface{} `field:"required" json:"total" yaml:"total"`
	// bad block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#bad Slo#bad}
	Bad interface{} `field:"optional" json:"bad" yaml:"bad"`
	// good block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.22.0/docs/resources/slo#good Slo#good}
	Good interface{} `field:"optional" json:"good" yaml:"good"`
}

