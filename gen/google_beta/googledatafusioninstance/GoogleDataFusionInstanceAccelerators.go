package googledatafusioninstance


type GoogleDataFusionInstanceAccelerators struct {
	// The type of an accelator for a CDF instance. Possible values: ["CDC", "HEALTHCARE", "CCAI_INSIGHTS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#accelerator_type GoogleDataFusionInstance#accelerator_type}
	AcceleratorType *string `field:"required" json:"acceleratorType" yaml:"acceleratorType"`
	// The type of an accelator for a CDF instance. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_fusion_instance#state GoogleDataFusionInstance#state}
	State *string `field:"required" json:"state" yaml:"state"`
}

