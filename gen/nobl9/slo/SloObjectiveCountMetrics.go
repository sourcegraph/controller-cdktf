package slo

type SloObjectiveCountMetrics struct {
	// Should the metrics be incrementing or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#incremental Slo#incremental}
	Incremental any `field:"required" json:"incremental" yaml:"incremental"`
	// bad block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#bad Slo#bad}
	Bad any `field:"optional" json:"bad" yaml:"bad"`
	// good block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#good Slo#good}
	Good any `field:"optional" json:"good" yaml:"good"`
	// good_total block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#good_total Slo#good_total}
	GoodTotal any `field:"optional" json:"goodTotal" yaml:"goodTotal"`
	// total block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#total Slo#total}
	Total any `field:"optional" json:"total" yaml:"total"`
}
