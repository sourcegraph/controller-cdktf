package googledialogflowenvironment


type GoogleDialogflowEnvironmentTextToSpeechSettings struct {
	// Indicates whether text to speech is enabled.
	//
	// Even when this field is false, other settings in this proto are still retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#enable_text_to_speech GoogleDialogflowEnvironment#enable_text_to_speech}
	EnableTextToSpeech interface{} `field:"optional" json:"enableTextToSpeech" yaml:"enableTextToSpeech"`
	// Audio encoding of the synthesized audio content. Possible values: ["OUTPUT_AUDIO_ENCODING_UNSPECIFIED", "OUTPUT_AUDIO_ENCODING_LINEAR_16", "OUTPUT_AUDIO_ENCODING_MP3", "OUTPUT_AUDIO_ENCODING_MP3_64_KBPS", "OUTPUT_AUDIO_ENCODING_OGG_OPUS", "OUTPUT_AUDIO_ENCODING_MULAW", "OUTPUT_AUDIO_ENCODING_ALAW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#output_audio_encoding GoogleDialogflowEnvironment#output_audio_encoding}
	OutputAudioEncoding *string `field:"optional" json:"outputAudioEncoding" yaml:"outputAudioEncoding"`
	// The synthesis sample rate (in hertz) for this audio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#sample_rate_hertz GoogleDialogflowEnvironment#sample_rate_hertz}
	SampleRateHertz *float64 `field:"optional" json:"sampleRateHertz" yaml:"sampleRateHertz"`
	// synthesize_speech_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_dialogflow_environment#synthesize_speech_configs GoogleDialogflowEnvironment#synthesize_speech_configs}
	SynthesizeSpeechConfigs interface{} `field:"optional" json:"synthesizeSpeechConfigs" yaml:"synthesizeSpeechConfigs"`
}

