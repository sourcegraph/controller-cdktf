package googlebigquerytable


type GoogleBigqueryTableSchemaForeignTypeInfo struct {
	// Specifies the system which defines the foreign data type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_bigquery_table#type_system GoogleBigqueryTable#type_system}
	TypeSystem *string `field:"required" json:"typeSystem" yaml:"typeSystem"`
}

