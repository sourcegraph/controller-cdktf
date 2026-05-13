package googlecesapp


type GoogleCesAppTimeZoneSettings struct {
	// The time zone of the app from the time zone database, e.g., America/Los_Angeles, Europe/Paris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_app#time_zone GoogleCesApp#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

