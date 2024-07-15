package googlebiglaketable


type GoogleBiglakeTableHiveOptionsStorageDescriptor struct {
	// The fully qualified Java class name of the input format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_table#input_format GoogleBiglakeTable#input_format}
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Cloud Storage folder URI where the table data is stored, starting with "gs://".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_table#location_uri GoogleBiglakeTable#location_uri}
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// The fully qualified Java class name of the output format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_biglake_table#output_format GoogleBiglakeTable#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

