package googlenetappvolume


type GoogleNetappVolumeLargeCapacityConfig struct {
	// The number of internal constituents (e.g., FlexVols) for this large volume. The minimum number of constituents is 2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_netapp_volume#constituent_count GoogleNetappVolume#constituent_count}
	ConstituentCount *float64 `field:"optional" json:"constituentCount" yaml:"constituentCount"`
}

