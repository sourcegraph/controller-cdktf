package chronicledatatablerow


type ChronicleDataTableRowTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/chronicle_data_table_row#create ChronicleDataTableRow#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/chronicle_data_table_row#delete ChronicleDataTableRow#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/chronicle_data_table_row#update ChronicleDataTableRow#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

