package googlechronicledatatable


type GoogleChronicleDataTableColumnInfo struct {
	// Column Index. 0,1,2...
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#column_index GoogleChronicleDataTable#column_index}
	ColumnIndex *float64 `field:"required" json:"columnIndex" yaml:"columnIndex"`
	// Original column name of the Data Table (present in the CSV header in case of creation of data tables using file uploads).
	//
	// It must satisfy the
	// following requirements:
	// - Starts with letter.
	// - Contains only letters, numbers and underscore.
	// - Must be unique and has length < 256
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#original_column GoogleChronicleDataTable#original_column}
	OriginalColumn *string `field:"required" json:"originalColumn" yaml:"originalColumn"`
	// Column type can be STRING, CIDR (Ex- 10.1.1.0/24), REGEX Possible values: STRING REGEX CIDR NUMBER Possible values: ["STRING", "REGEX", "CIDR", "NUMBER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#column_type GoogleChronicleDataTable#column_type}
	ColumnType *string `field:"optional" json:"columnType" yaml:"columnType"`
	// Whether to include this column in the calculation of the row ID.
	//
	// If no columns have key_column = true, all columns will be included in the
	// calculation of the row ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#key_column GoogleChronicleDataTable#key_column}
	KeyColumn interface{} `field:"optional" json:"keyColumn" yaml:"keyColumn"`
	// Entity proto field path that the column is mapped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#mapped_column_path GoogleChronicleDataTable#mapped_column_path}
	MappedColumnPath *string `field:"optional" json:"mappedColumnPath" yaml:"mappedColumnPath"`
	// Whether the column is a repeated values column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_data_table#repeated_values GoogleChronicleDataTable#repeated_values}
	RepeatedValues interface{} `field:"optional" json:"repeatedValues" yaml:"repeatedValues"`
}

