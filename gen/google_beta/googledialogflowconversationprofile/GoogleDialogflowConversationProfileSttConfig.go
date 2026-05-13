package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileSttConfig struct {
	// Audio encoding of the audio content to process. Possible values: ["AUDIO_ENCODING_UNSPECIFIED", "AUDIO_ENCODING_LINEAR_16", "AUDIO_ENCODING_FLAC", "AUDIO_ENCODING_MULAW", "AUDIO_ENCODING_AMR", "AUDIO_ENCODING_AMR_WB", "AUDIO_ENCODING_OGG_OPUS", "AUDIOENCODING_SPEEX_WITH_HEADER_BYTE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#audio_encoding GoogleDialogflowConversationProfile#audio_encoding}
	AudioEncoding *string `field:"optional" json:"audioEncoding" yaml:"audioEncoding"`
	// If true, Dialogflow returns SpeechWordInfo in StreamingRecognitionResult with information about the recognized speech words.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#enable_word_info GoogleDialogflowConversationProfile#enable_word_info}
	EnableWordInfo interface{} `field:"optional" json:"enableWordInfo" yaml:"enableWordInfo"`
	// The language of the supplied audio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#language_code GoogleDialogflowConversationProfile#language_code}
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Which Speech model to select. Leave this field unspecified to use Agent Speech settings for model selection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#model GoogleDialogflowConversationProfile#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// Sample rate (in Hertz) of the audio content sent in the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#sample_rate_hertz GoogleDialogflowConversationProfile#sample_rate_hertz}
	SampleRateHertz *float64 `field:"optional" json:"sampleRateHertz" yaml:"sampleRateHertz"`
	// The speech model used in speech to text. Possible values: ["SPEECH_MODEL_VARIANT_UNSPECIFIED", "USE_BEST_AVAILABLE", "USE_STANDARD", "USE_ENHANCED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#speech_model_variant GoogleDialogflowConversationProfile#speech_model_variant}
	SpeechModelVariant *string `field:"optional" json:"speechModelVariant" yaml:"speechModelVariant"`
	// Use timeout based endpointing, interpreting endpointer sensitivy as seconds of timeout value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_conversation_profile#use_timeout_based_endpointing GoogleDialogflowConversationProfile#use_timeout_based_endpointing}
	UseTimeoutBasedEndpointing interface{} `field:"optional" json:"useTimeoutBasedEndpointing" yaml:"useTimeoutBasedEndpointing"`
}

