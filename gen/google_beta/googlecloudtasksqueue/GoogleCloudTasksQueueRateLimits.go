package googlecloudtasksqueue


type GoogleCloudTasksQueueRateLimits struct {
	// The maximum number of concurrent tasks that Cloud Tasks allows to be dispatched for this queue.
	//
	// After this threshold has been
	// reached, Cloud Tasks stops dispatching tasks until the number of
	// concurrent requests decreases.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#max_concurrent_dispatches GoogleCloudTasksQueue#max_concurrent_dispatches}
	MaxConcurrentDispatches *float64 `field:"optional" json:"maxConcurrentDispatches" yaml:"maxConcurrentDispatches"`
	// The maximum rate at which tasks are dispatched from this queue.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_tasks_queue#max_dispatches_per_second GoogleCloudTasksQueue#max_dispatches_per_second}
	MaxDispatchesPerSecond *float64 `field:"optional" json:"maxDispatchesPerSecond" yaml:"maxDispatchesPerSecond"`
}

