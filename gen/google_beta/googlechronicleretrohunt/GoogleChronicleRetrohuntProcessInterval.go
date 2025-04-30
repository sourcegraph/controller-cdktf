package googlechronicleretrohunt


type GoogleChronicleRetrohuntProcessInterval struct {
	// Exclusive end of the interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_retrohunt#end_time GoogleChronicleRetrohunt#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Inclusive start of the interval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_chronicle_retrohunt#start_time GoogleChronicleRetrohunt#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

