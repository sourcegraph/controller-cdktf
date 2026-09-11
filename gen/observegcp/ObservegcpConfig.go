package observegcp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservegcpConfig struct {
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The identifier of the GCP Resource to monitor.
	//
	// The resource can be a project, folder, or organization.
	//
	// Examples: "projects/my_project-123", "folders/1234567899", "organizations/34739118321".
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The number of days to wait before deleting AbortIncompleteMultipartUpload.
	//
	// 7.
	BucketLifecycleAbortUploadDays *float64 `field:"optional" json:"bucketLifecycleAbortUploadDays" yaml:"bucketLifecycleAbortUploadDays"`
	// The number of days to wait before Delete of temporary bucket files.
	//
	// 14.
	BucketLifecycleDeleteDays *float64 `field:"optional" json:"bucketLifecycleDeleteDays" yaml:"bucketLifecycleDeleteDays"`
	// The debug level for the GCP cloud functions WARNING.
	CloudFunctionDebugLevel *string `field:"optional" json:"cloudFunctionDebugLevel" yaml:"cloudFunctionDebugLevel"`
	// Whether to enable the Cloud function that tracks GCP assets.
	//
	// true.
	EnableAssetTracking *bool `field:"optional" json:"enableAssetTracking" yaml:"enableAssetTracking"`
	// DEPRECATED: This variable has been renamed to 'enable_asset_tracking'.
	//
	// Please update your configuration to use 'enable_asset_tracking' instead.
	EnableFunction *bool `field:"optional" json:"enableFunction" yaml:"enableFunction"`
	// Whether to include all children Projects of a Folder when collecting logs true.
	FolderIncludeChildren *bool `field:"optional" json:"folderIncludeChildren" yaml:"folderIncludeChildren"`
	// Memory (in MB), available to the function.
	//
	// Default value is 512. Possible values include 128, 256, 512, 1024, etc.
	// 4096.
	FunctionAvailableMemoryMb *float64 `field:"optional" json:"functionAvailableMemoryMb" yaml:"functionAvailableMemoryMb"`
	// GCS bucket containing the Cloud Function source code observeinc.
	FunctionBucket *string `field:"optional" json:"functionBucket" yaml:"functionBucket"`
	// Whether to disable function logging.
	FunctionDisableLogging *bool `field:"optional" json:"functionDisableLogging" yaml:"functionDisableLogging"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	//
	// 100.
	FunctionMaxInstances *float64 `field:"optional" json:"functionMaxInstances" yaml:"functionMaxInstances"`
	// GCS object key of the Cloud Function source code zip file.
	//
	// Will use the latest release unless modified.
	// google-cloud-functions-latest.zip
	FunctionObject *string `field:"optional" json:"functionObject" yaml:"functionObject"`
	// A list of IAM roles to give the Cloud Function.
	//
	// roles/compute.viewer
	// roles/iam.serviceAccountViewer
	// roles/cloudscheduler.viewer
	// roles/cloudasset.viewer
	// roles/browser
	// roles/logging.viewer
	// roles/monitoring.viewer
	// roles/storage.objectCreator
	// roles/storage.objectViewer
	// roles/storage.objectAdmin
	// roles/storage.admin
	// roles/cloudfunctions.invoker
	// roles/cloudtasks.enqueuer
	// roles/cloudtasks.viewer
	// roles/cloudtasks.taskDeleter
	// roles/iam.serviceAccountUser
	FunctionRoles *[]*string `field:"optional" json:"functionRoles" yaml:"functionRoles"`
	// Cron schedule for the job 0 * * * *.
	FunctionScheduleFrequency *string `field:"optional" json:"functionScheduleFrequency" yaml:"functionScheduleFrequency"`
	// Cron schedule for the job * /5 * * * *.
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	FunctionScheduleFrequencyRestOfAssets *string `field:"optional" json:"functionScheduleFrequencyRestOfAssets" yaml:"functionScheduleFrequencyRestOfAssets"`
	// Timeout (in seconds) for the function. Default value is 300 seconds. Cannot be more than 540 seconds.
	//
	// 300.
	FunctionTimeout *float64 `field:"optional" json:"functionTimeout" yaml:"functionTimeout"`
	// The location where the Task Queue will be created.
	//
	// us-central1.
	GcpRegion *string `field:"optional" json:"gcpRegion" yaml:"gcpRegion"`
	// A map of labels to add to resources (https://cloud.google.com/resource-manager/docs/creating-managing-labels)".
	//
	// Note: Many, but not all, Google Cloud SDK resources support labels.
	//
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Log entries that match any of these exclusion filters will not be exported.
	//
	// If a log entry is matched by both logging_filter and one of logging_exclusions it will not be exported.
	//
	// Relevant docs: https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.exclusions#LogExclusion
	LoggingExclusions interface{} `field:"optional" json:"loggingExclusions" yaml:"loggingExclusions"`
	// An advanced logs filter.
	//
	// The only exported log entries are those that are
	// in the resource owning the sink and that match the filter.
	//
	// Relevant docs: https://cloud.google.com/logging/docs/view/building-queries
	LoggingFilter *string `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// The maximum number of retry attempts for a task in case of failure.
	//
	// -1.
	MaxAttempts *float64 `field:"optional" json:"maxAttempts" yaml:"maxAttempts"`
	// The maximum number of tasks that can be dispatched concurrently.
	//
	// 2.
	MaxConcurrentDispatches *float64 `field:"optional" json:"maxConcurrentDispatches" yaml:"maxConcurrentDispatches"`
	// The maximum rate at which tasks can be dispatched per second.
	//
	// 2.
	MaxDispatchesPerSecond *float64 `field:"optional" json:"maxDispatchesPerSecond" yaml:"maxDispatchesPerSecond"`
	// The time limit for retrying a task in seconds 7200s.
	MaxRetryDuration *string `field:"optional" json:"maxRetryDuration" yaml:"maxRetryDuration"`
	// The minimum amount of time to wait between retries in seconds 30s.
	MinBackoff *string `field:"optional" json:"minBackoff" yaml:"minBackoff"`
	// Module name.
	//
	// Used as a name prefix.
	// obs.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of IAM roles to give the Observe poller (through the service account key output).
	//
	// roles/monitoring.viewer
	PollerRoles *[]*string `field:"optional" json:"pollerRoles" yaml:"pollerRoles"`
	// Billing Project ID needed for asset feed.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Ack deadline for the Pub/Sub subscription (https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions) 60.
	PubsubAckDeadlineSeconds *float64 `field:"optional" json:"pubsubAckDeadlineSeconds" yaml:"pubsubAckDeadlineSeconds"`
	// Retry policy maximum backoff for the Pub/Sub subscription (https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions) 600s.
	PubsubMaximumBackoff *string `field:"optional" json:"pubsubMaximumBackoff" yaml:"pubsubMaximumBackoff"`
	// Message retention for the Pub/Sub subscription (https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions) 86400s.
	PubsubMessageRetentionDuration *string `field:"optional" json:"pubsubMessageRetentionDuration" yaml:"pubsubMessageRetentionDuration"`
	// Retry policy minimum backoff for the Pub/Sub subscription (https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.subscriptions) 10s.
	PubsubMinimumBackoff *string `field:"optional" json:"pubsubMinimumBackoff" yaml:"pubsubMinimumBackoff"`
}

