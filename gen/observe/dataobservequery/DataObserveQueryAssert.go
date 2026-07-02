package dataobservequery

type DataObserveQueryAssert struct {
	// Filename containing expected query output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#golden_file DataObserveQuery#golden_file}
	GoldenFile *string `field:"required" json:"goldenFile" yaml:"goldenFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#update DataObserveQuery#update}.
	Update any `field:"optional" json:"update" yaml:"update"`
}
