package googledialogflowcxflow


type GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestination struct {
	// The Google Cloud Storage URI for the exported objects.
	//
	// Whether a full object name, or just a prefix, its usage depends on the Dialogflow operation.
	// Format: gs://bucket/object-name-or-prefix
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#uri GoogleDialogflowCxFlow#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

