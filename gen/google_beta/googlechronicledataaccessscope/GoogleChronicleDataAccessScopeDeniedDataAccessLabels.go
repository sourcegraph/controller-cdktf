package googlechronicledataaccessscope


type GoogleChronicleDataAccessScopeDeniedDataAccessLabels struct {
	// The asset namespace configured in the forwarder of the customer's events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_chronicle_data_access_scope#asset_namespace GoogleChronicleDataAccessScope#asset_namespace}
	AssetNamespace *string `field:"optional" json:"assetNamespace" yaml:"assetNamespace"`
	// The name of the data access label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_chronicle_data_access_scope#data_access_label GoogleChronicleDataAccessScope#data_access_label}
	DataAccessLabel *string `field:"optional" json:"dataAccessLabel" yaml:"dataAccessLabel"`
	// ingestion_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_chronicle_data_access_scope#ingestion_label GoogleChronicleDataAccessScope#ingestion_label}
	IngestionLabel *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel `field:"optional" json:"ingestionLabel" yaml:"ingestionLabel"`
	// The name of the log type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_chronicle_data_access_scope#log_type GoogleChronicleDataAccessScope#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

