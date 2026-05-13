package googlefirebaseailogicconfig


type GoogleFirebaseAiLogicConfigTelemetryConfig struct {
	// The current monitoring mode used for this project. Possible values: NONE ALL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#mode GoogleFirebaseAiLogicConfig#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The percentage of requests to be sampled, expressed as a fraction in the range (0,1].
	//
	// Note that the actual sampling rate may be lower than
	// the specified value if the system is overloaded. Default is 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_firebase_ai_logic_config#sampling_rate GoogleFirebaseAiLogicConfig#sampling_rate}
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

