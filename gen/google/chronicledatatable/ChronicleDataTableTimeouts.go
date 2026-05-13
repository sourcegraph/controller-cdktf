package chronicledatatable


type ChronicleDataTableTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/chronicle_data_table#create ChronicleDataTable#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/chronicle_data_table#delete ChronicleDataTable#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/chronicle_data_table#update ChronicleDataTable#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

