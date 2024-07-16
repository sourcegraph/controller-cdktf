package googlestoragetransferjob


type GoogleStorageTransferJobScheduleStartTimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#hours GoogleStorageTransferJob#hours}
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Minutes of hour of day. Must be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#minutes GoogleStorageTransferJob#minutes}
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#nanos GoogleStorageTransferJob#nanos}
	Nanos *float64 `field:"required" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time. Must normally be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_job#seconds GoogleStorageTransferJob#seconds}
	Seconds *float64 `field:"required" json:"seconds" yaml:"seconds"`
}

