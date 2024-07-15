package googledataflowflextemplatejob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataflowFlexTemplateJobConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#container_spec_gcs_path GoogleDataflowFlexTemplateJob#container_spec_gcs_path}.
	ContainerSpecGcsPath *string `field:"required" json:"containerSpecGcsPath" yaml:"containerSpecGcsPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#name GoogleDataflowFlexTemplateJob#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#additional_experiments GoogleDataflowFlexTemplateJob#additional_experiments}
	AdditionalExperiments *[]*string `field:"optional" json:"additionalExperiments" yaml:"additionalExperiments"`
	// The algorithm to use for autoscaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#autoscaling_algorithm GoogleDataflowFlexTemplateJob#autoscaling_algorithm}
	AutoscalingAlgorithm *string `field:"optional" json:"autoscalingAlgorithm" yaml:"autoscalingAlgorithm"`
	// Indicates if the job should use the streaming engine feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#enable_streaming_engine GoogleDataflowFlexTemplateJob#enable_streaming_engine}
	EnableStreamingEngine interface{} `field:"optional" json:"enableStreamingEngine" yaml:"enableStreamingEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#id GoogleDataflowFlexTemplateJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#ip_configuration GoogleDataflowFlexTemplateJob#ip_configuration}
	IpConfiguration *string `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#kms_key_name GoogleDataflowFlexTemplateJob#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// User labels to be specified for the job.
	//
	// Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#labels GoogleDataflowFlexTemplateJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The machine type to use for launching the job. The default is n1-standard-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#launcher_machine_type GoogleDataflowFlexTemplateJob#launcher_machine_type}
	LauncherMachineType *string `field:"optional" json:"launcherMachineType" yaml:"launcherMachineType"`
	// The machine type to use for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#machine_type GoogleDataflowFlexTemplateJob#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The maximum number of Google Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#max_workers GoogleDataflowFlexTemplateJob#max_workers}
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#network GoogleDataflowFlexTemplateJob#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The initial number of Google Compute Engine instances for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#num_workers GoogleDataflowFlexTemplateJob#num_workers}
	NumWorkers *float64 `field:"optional" json:"numWorkers" yaml:"numWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#on_delete GoogleDataflowFlexTemplateJob#on_delete}.
	OnDelete *string `field:"optional" json:"onDelete" yaml:"onDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#parameters GoogleDataflowFlexTemplateJob#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#project GoogleDataflowFlexTemplateJob#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region in which the created job should run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#region GoogleDataflowFlexTemplateJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docker registry location of container image to use for the 'worker harness.
	//
	// Default is the container for the version of the SDK. Note this field is only valid for portable pipelines.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#sdk_container_image GoogleDataflowFlexTemplateJob#sdk_container_image}
	SdkContainerImage *string `field:"optional" json:"sdkContainerImage" yaml:"sdkContainerImage"`
	// The Service Account email used to create the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#service_account_email GoogleDataflowFlexTemplateJob#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from terraform state and moving on.
	//
	// WARNING: this will lead to job name conflicts if you do not ensure that the job names are different, e.g. by embedding a release ID or by using a random_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#skip_wait_on_job_termination GoogleDataflowFlexTemplateJob#skip_wait_on_job_termination}
	SkipWaitOnJobTermination interface{} `field:"optional" json:"skipWaitOnJobTermination" yaml:"skipWaitOnJobTermination"`
	// The Cloud Storage path to use for staging files. Must be a valid Cloud Storage URL, beginning with gs://.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#staging_location GoogleDataflowFlexTemplateJob#staging_location}
	StagingLocation *string `field:"optional" json:"stagingLocation" yaml:"stagingLocation"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#subnetwork GoogleDataflowFlexTemplateJob#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#temp_location GoogleDataflowFlexTemplateJob#temp_location}
	TempLocation *string `field:"optional" json:"tempLocation" yaml:"tempLocation"`
	// Only applicable when updating a pipeline.
	//
	// Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataflow_flex_template_job#transform_name_mapping GoogleDataflowFlexTemplateJob#transform_name_mapping}
	TransformNameMapping *map[string]*string `field:"optional" json:"transformNameMapping" yaml:"transformNameMapping"`
}

