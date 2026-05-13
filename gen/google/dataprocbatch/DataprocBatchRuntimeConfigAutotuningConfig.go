package dataprocbatch


type DataprocBatchRuntimeConfigAutotuningConfig struct {
	// Optional. Scenarios for which tunings are applied. Possible values: ["AUTO", "SCALING", "BROADCAST_HASH_JOIN", "MEMORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/dataproc_batch#scenarios DataprocBatch#scenarios}
	Scenarios *[]*string `field:"optional" json:"scenarios" yaml:"scenarios"`
}

