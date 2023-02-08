// executors
package executors

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExecutorsOptions struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// A label tag to add to all the executors;
	//
	// can be used for filtering out the right instances in stackdriver monitoring.
	InstanceTag *string `field:"required" json:"instanceTag" yaml:"instanceTag"`
	// The value for environment by which to filter the custom metrics.
	MetricsEnvironmentLabel *string `field:"required" json:"metricsEnvironmentLabel" yaml:"metricsEnvironmentLabel"`
	// undefined.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The queue from which the executor should dequeue jobs.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// The shared password used to authenticate requests to the internal executor proxy.
	SourcegraphExecutorProxyPassword *string `field:"required" json:"sourcegraphExecutorProxyPassword" yaml:"sourcegraphExecutorProxyPassword"`
	// The externally accessible URL of the target Sourcegraph instance.
	SourcegraphExternalUrl *string `field:"required" json:"sourcegraphExternalUrl" yaml:"sourcegraphExternalUrl"`
	// undefined.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// zone.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// If false, no public IP will be associated with the executors.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Executor node disk size in GB.
	BootDiskSize *float64 `field:"optional" json:"bootDiskSize" yaml:"bootDiskSize"`
	// A URL to a docker registry mirror to use (falling back to docker.io).
	DockerRegistryMirror *string `field:"optional" json:"dockerRegistryMirror" yaml:"dockerRegistryMirror"`
	// A URL to a docker registry mirror node exporter to scrape (optional).
	DockerRegistryMirrorNodeExporterUrl *string `field:"optional" json:"dockerRegistryMirrorNodeExporterUrl" yaml:"dockerRegistryMirrorNodeExporterUrl"`
	// The amount of disk space to give to each firecracker VM.
	FirecrackerDiskSpace *string `field:"optional" json:"firecrackerDiskSpace" yaml:"firecrackerDiskSpace"`
	// The amount of memory to give to each firecracker VM.
	FirecrackerMemory *string `field:"optional" json:"firecrackerMemory" yaml:"firecrackerMemory"`
	// The number of CPUs to give to each firecracker VM.
	FirecrackerNumCpus *float64 `field:"optional" json:"firecrackerNumCpus" yaml:"firecrackerNumCpus"`
	// DEPRECATED.
	//
	// This is not used anymore.
	HttpAccessCidrRanges *[]*string `field:"optional" json:"httpAccessCidrRanges" yaml:"httpAccessCidrRanges"`
	// The amount of memory to allocate to each virtual machine or container.
	JobMemory *string `field:"optional" json:"jobMemory" yaml:"jobMemory"`
	// The number of CPUs to allocate to each virtual machine or container.
	JobNumCpus *float64 `field:"optional" json:"jobNumCpus" yaml:"jobNumCpus"`
	// The amount of jobs a single instance should have in queue.
	//
	// Used for autoscaling.
	JobsPerInstanceScaling *float64 `field:"optional" json:"jobsPerInstanceScaling" yaml:"jobsPerInstanceScaling"`
	// Executor node machine disk image to use for creating the boot volume.
	//
	// Leave empty to use latest compatible with the Sourcegraph version.
	MachineImage *string `field:"optional" json:"machineImage" yaml:"machineImage"`
	// Executor node machine type.
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The maximum time that can be spent by the worker dequeueing records to be handled.
	MaxActiveTime *string `field:"optional" json:"maxActiveTime" yaml:"maxActiveTime"`
	// The number of jobs to run concurrently per executor instance.
	MaximumNumJobs *float64 `field:"optional" json:"maximumNumJobs" yaml:"maximumNumJobs"`
	// The maximum wall time that can be spent on a single job.
	MaximumRuntimePerJob *string `field:"optional" json:"maximumRuntimePerJob" yaml:"maximumRuntimePerJob"`
	// The maximum number of executor instances to run in the autoscaling group.
	MaxReplicas *float64 `field:"optional" json:"maxReplicas" yaml:"maxReplicas"`
	// The minimum number of executor instances to run in the autoscaling group.
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// The maximum number of jobs that will be dequeued by the worker.
	NumTotalJobs *float64 `field:"optional" json:"numTotalJobs" yaml:"numTotalJobs"`
	// Whether to use preemptible machines instead of standard machines;
	//
	// usually way cheaper but might be terminated at any time.
	PreemptibleMachines *bool `field:"optional" json:"preemptibleMachines" yaml:"preemptibleMachines"`
	// An optional prefix to add to all resources created.
	ResourcePrefix *string `field:"optional" json:"resourcePrefix" yaml:"resourcePrefix"`
	// Whether to isolate commands in virtual machines.
	UseFirecracker *bool `field:"optional" json:"useFirecracker" yaml:"useFirecracker"`
}

