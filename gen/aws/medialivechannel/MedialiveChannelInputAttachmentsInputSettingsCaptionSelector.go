package medialivechannel


type MedialiveChannelInputAttachmentsInputSettingsCaptionSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_channel#language_code MedialiveChannel#language_code}.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// selector_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/medialive_channel#selector_settings MedialiveChannel#selector_settings}
	SelectorSettings *MedialiveChannelInputAttachmentsInputSettingsCaptionSelectorSelectorSettings `field:"optional" json:"selectorSettings" yaml:"selectorSettings"`
}

