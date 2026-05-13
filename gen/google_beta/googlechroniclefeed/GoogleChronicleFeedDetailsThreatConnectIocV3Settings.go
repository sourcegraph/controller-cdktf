package googlechroniclefeed


type GoogleChronicleFeedDetailsThreatConnectIocV3Settings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsThreatConnectIocV3SettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#fields GoogleChronicleFeed#fields}
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Owners.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#owners GoogleChronicleFeed#owners}
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
	// Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#schedule GoogleChronicleFeed#schedule}
	Schedule *float64 `field:"optional" json:"schedule" yaml:"schedule"`
	// ThreatConnect Query Language filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#tql_query GoogleChronicleFeed#tql_query}
	TqlQuery *string `field:"optional" json:"tqlQuery" yaml:"tqlQuery"`
}

