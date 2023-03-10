package emrserverlessapplication


type EmrserverlessApplicationInitialCapacityInitialCapacityConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#worker_count EmrserverlessApplication#worker_count}.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
	// worker_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emrserverless_application#worker_configuration EmrserverlessApplication#worker_configuration}
	WorkerConfiguration *EmrserverlessApplicationInitialCapacityInitialCapacityConfigWorkerConfiguration `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

