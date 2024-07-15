package googlecloudtasksqueue


type GoogleCloudTasksQueueStackdriverLoggingConfig struct {
	// Specifies the fraction of operations to write to Stackdriver Logging.
	//
	// This field may contain any value between 0.0 and 1.0, inclusive. 0.0 is the
	// default and means that no operations are logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#sampling_ratio GoogleCloudTasksQueue#sampling_ratio}
	SamplingRatio *float64 `field:"required" json:"samplingRatio" yaml:"samplingRatio"`
}

