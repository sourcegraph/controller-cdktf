package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersResources struct {
	// Limits describes the maximum amount of compute resources allowed.
	//
	// CPU Limit details:
	// - For fractional CPU values (e.g. '0.5', '0.75', min '0.08') are also supported.
	// - CPU allocation must comply with memory limits and concurrency rules described in:
	//   https://cloud.google.com/run/docs/configuring/services/cpu
	// The values of the map is string form of the 'quantity' k8s type:
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_service#limits GoogleCloudRunService#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is
	// explicitly specified, otherwise to an implementation-defined value.
	// The values of the map is string form of the 'quantity' k8s type:
	// https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_run_service#requests GoogleCloudRunService#requests}
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

