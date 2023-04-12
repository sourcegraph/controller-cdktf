package googledataplexzone


type GoogleDataplexZoneResourceSpec struct {
	// Required.
	//
	// Immutable. The location type of the resources that are allowed to be attached to the assets within this zone. Possible values: LOCATION_TYPE_UNSPECIFIED, SINGLE_REGION, MULTI_REGION
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataplex_zone#location_type GoogleDataplexZone#location_type}
	LocationType *string `field:"required" json:"locationType" yaml:"locationType"`
}

