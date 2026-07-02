package sourcedataset

type SourceDatasetField struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#name SourceDataset#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#sql_type SourceDataset#sql_type}.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#type SourceDataset#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#is_const SourceDataset#is_const}.
	IsConst any `field:"optional" json:"isConst" yaml:"isConst"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#is_enum SourceDataset#is_enum}.
	IsEnum any `field:"optional" json:"isEnum" yaml:"isEnum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#is_hidden SourceDataset#is_hidden}.
	IsHidden any `field:"optional" json:"isHidden" yaml:"isHidden"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#is_metric SourceDataset#is_metric}.
	IsMetric any `field:"optional" json:"isMetric" yaml:"isMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/source_dataset#is_searchable SourceDataset#is_searchable}.
	IsSearchable any `field:"optional" json:"isSearchable" yaml:"isSearchable"`
}
