package cesapp


type CesAppLoggingSettingsRedactionConfig struct {
	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name to instruct on how to de-identify content. Format: 'projects/{project}/locations/{location}/deidentifyTemplates/{deidentify_template}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#deidentify_template CesApp#deidentify_template}
	DeidentifyTemplate *string `field:"optional" json:"deidentifyTemplate" yaml:"deidentifyTemplate"`
	// If true, redaction will be applied in various logging scenarios, including conversation history, Cloud Logging and audio recording.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#enable_redaction CesApp#enable_redaction}
	EnableRedaction interface{} `field:"optional" json:"enableRedaction" yaml:"enableRedaction"`
	// [DLP](https://cloud.google.com/dlp/docs) inspect template name to configure detection of sensitive data types. Format: 'projects/{project}/locations/{location}/inspectTemplates/{inspect_template}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#inspect_template CesApp#inspect_template}
	InspectTemplate *string `field:"optional" json:"inspectTemplate" yaml:"inspectTemplate"`
}

