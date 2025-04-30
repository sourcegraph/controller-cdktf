package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigBlmtConfig struct {
	// The Cloud Storage bucket name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#bucket GoogleDatastreamStream#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The bigquery connection. Format: '{project}.{location}.{name}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#connection_name GoogleDatastreamStream#connection_name}
	ConnectionName *string `field:"required" json:"connectionName" yaml:"connectionName"`
	// The file format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#file_format GoogleDatastreamStream#file_format}
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// The table format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#table_format GoogleDatastreamStream#table_format}
	TableFormat *string `field:"required" json:"tableFormat" yaml:"tableFormat"`
	// The root path inside the Cloud Storage bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#root_path GoogleDatastreamStream#root_path}
	RootPath *string `field:"optional" json:"rootPath" yaml:"rootPath"`
}

