package googledialogflowcxflow


type GoogleDialogflowCxFlowAdvancedSettings struct {
	// audio_export_gcs_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#audio_export_gcs_destination GoogleDialogflowCxFlow#audio_export_gcs_destination}
	AudioExportGcsDestination *GoogleDialogflowCxFlowAdvancedSettingsAudioExportGcsDestination `field:"optional" json:"audioExportGcsDestination" yaml:"audioExportGcsDestination"`
	// dtmf_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dialogflow_cx_flow#dtmf_settings GoogleDialogflowCxFlow#dtmf_settings}
	DtmfSettings *GoogleDialogflowCxFlowAdvancedSettingsDtmfSettings `field:"optional" json:"dtmfSettings" yaml:"dtmfSettings"`
}

