package googledataprocbatch


type GoogleDataprocBatchRuntimeConfigAutotuningConfig struct {
	// Optional. Scenarios for which tunings are applied. Possible values: ["SCALING", "BROADCAST_HASH_JOIN", "MEMORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_dataproc_batch#scenarios GoogleDataprocBatch#scenarios}
	Scenarios *[]*string `field:"optional" json:"scenarios" yaml:"scenarios"`
}

