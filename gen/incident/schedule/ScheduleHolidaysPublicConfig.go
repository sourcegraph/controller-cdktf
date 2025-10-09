package schedule


type ScheduleHolidaysPublicConfig struct {
	// ISO 3166-1 alpha-2 country codes for the countries that this schedule is configured to view holidays for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.17.1/docs/resources/schedule#country_codes Schedule#country_codes}
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
}

