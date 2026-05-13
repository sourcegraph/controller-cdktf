package networksecuritysecurityprofile


type NetworkSecuritySecurityProfileUrlFilteringProfileUrlFilters struct {
	// The action to take when the filter is applied. Possible values: ["ALLOW", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_security_profile#filtering_action NetworkSecuritySecurityProfile#filtering_action}
	FilteringAction *string `field:"required" json:"filteringAction" yaml:"filteringAction"`
	// The priority of the filter within the URL filtering profile.
	//
	// Must be an integer from 0 and 2147483647, inclusive. Lower integers indicate higher priorities.
	// The priority of a filter must be unique within a URL filtering profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_security_profile#priority NetworkSecuritySecurityProfile#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// A list of domain matcher strings that a domain name gets compared with to determine if the filter is applicable.
	//
	// A domain name must match with at least one of the strings in the list for a filter to be applicable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_security_profile#urls NetworkSecuritySecurityProfile#urls}
	Urls *[]*string `field:"optional" json:"urls" yaml:"urls"`
}

