package googlecesapp


type GoogleCesAppLoggingSettings struct {
	// audio_recording_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#audio_recording_config GoogleCesApp#audio_recording_config}
	AudioRecordingConfig *GoogleCesAppLoggingSettingsAudioRecordingConfig `field:"optional" json:"audioRecordingConfig" yaml:"audioRecordingConfig"`
	// bigquery_export_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#bigquery_export_settings GoogleCesApp#bigquery_export_settings}
	BigqueryExportSettings *GoogleCesAppLoggingSettingsBigqueryExportSettings `field:"optional" json:"bigqueryExportSettings" yaml:"bigqueryExportSettings"`
	// cloud_logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#cloud_logging_settings GoogleCesApp#cloud_logging_settings}
	CloudLoggingSettings *GoogleCesAppLoggingSettingsCloudLoggingSettings `field:"optional" json:"cloudLoggingSettings" yaml:"cloudLoggingSettings"`
	// conversation_logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#conversation_logging_settings GoogleCesApp#conversation_logging_settings}
	ConversationLoggingSettings *GoogleCesAppLoggingSettingsConversationLoggingSettings `field:"optional" json:"conversationLoggingSettings" yaml:"conversationLoggingSettings"`
	// redaction_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#redaction_config GoogleCesApp#redaction_config}
	RedactionConfig *GoogleCesAppLoggingSettingsRedactionConfig `field:"optional" json:"redactionConfig" yaml:"redactionConfig"`
}

