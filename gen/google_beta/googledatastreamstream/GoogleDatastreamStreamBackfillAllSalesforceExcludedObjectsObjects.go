package googledatastreamstream


type GoogleDatastreamStreamBackfillAllSalesforceExcludedObjectsObjects struct {
	// fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#fields GoogleDatastreamStream#fields}
	Fields interface{} `field:"optional" json:"fields" yaml:"fields"`
	// Name of object in Salesforce Org.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_datastream_stream#object_name GoogleDatastreamStream#object_name}
	ObjectName *string `field:"optional" json:"objectName" yaml:"objectName"`
}

