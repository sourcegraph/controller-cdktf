package googleappenginestandardappversion


type GoogleAppEngineStandardAppVersionVpcAccessConnector struct {
	// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#name GoogleAppEngineStandardAppVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The egress setting for the connector, controlling what traffic is diverted through it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_standard_app_version#egress_setting GoogleAppEngineStandardAppVersion#egress_setting}
	EgressSetting *string `field:"optional" json:"egressSetting" yaml:"egressSetting"`
}

