package googlealloydbcluster


type GoogleAlloydbClusterMaintenanceUpdatePolicyMaintenanceWindowsStartTime struct {
	// Hours of day in 24 hour format. Should be from 0 to 23.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#hours GoogleAlloydbCluster#hours}
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Minutes of hour of day. Currently, only the value 0 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#minutes GoogleAlloydbCluster#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds. Currently, only the value 0 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#nanos GoogleAlloydbCluster#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time. Currently, only the value 0 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_alloydb_cluster#seconds GoogleAlloydbCluster#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

