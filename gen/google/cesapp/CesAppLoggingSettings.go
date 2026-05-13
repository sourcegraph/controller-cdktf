package cesapp


type CesAppLoggingSettings struct {
	// audio_recording_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#audio_recording_config CesApp#audio_recording_config}
	AudioRecordingConfig *CesAppLoggingSettingsAudioRecordingConfig `field:"optional" json:"audioRecordingConfig" yaml:"audioRecordingConfig"`
	// bigquery_export_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#bigquery_export_settings CesApp#bigquery_export_settings}
	BigqueryExportSettings *CesAppLoggingSettingsBigqueryExportSettings `field:"optional" json:"bigqueryExportSettings" yaml:"bigqueryExportSettings"`
	// cloud_logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#cloud_logging_settings CesApp#cloud_logging_settings}
	CloudLoggingSettings *CesAppLoggingSettingsCloudLoggingSettings `field:"optional" json:"cloudLoggingSettings" yaml:"cloudLoggingSettings"`
	// conversation_logging_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#conversation_logging_settings CesApp#conversation_logging_settings}
	ConversationLoggingSettings *CesAppLoggingSettingsConversationLoggingSettings `field:"optional" json:"conversationLoggingSettings" yaml:"conversationLoggingSettings"`
	// redaction_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_app#redaction_config CesApp#redaction_config}
	RedactionConfig *CesAppLoggingSettingsRedactionConfig `field:"optional" json:"redactionConfig" yaml:"redactionConfig"`
}

