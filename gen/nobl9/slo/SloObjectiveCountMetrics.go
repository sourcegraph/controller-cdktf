package slo


type SloObjectiveCountMetrics struct {
	// Should the metrics be incrementing or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#incremental Slo#incremental}
	Incremental interface{} `field:"required" json:"incremental" yaml:"incremental"`
	// bad block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#bad Slo#bad}
	Bad interface{} `field:"optional" json:"bad" yaml:"bad"`
	// good block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#good Slo#good}
	Good interface{} `field:"optional" json:"good" yaml:"good"`
	// good_total block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#good_total Slo#good_total}
	GoodTotal interface{} `field:"optional" json:"goodTotal" yaml:"goodTotal"`
	// total block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#total Slo#total}
	Total interface{} `field:"optional" json:"total" yaml:"total"`
}

