package deploymentv1


type DeploymentV1SpecTemplateSpecContainerReadinessProbeHttpGet struct {
	// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#host DeploymentV1#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#http_header DeploymentV1#http_header}
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// Path to access on the HTTP server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#path DeploymentV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Name or number of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#port DeploymentV1#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Scheme to use for connecting to the host.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment_v1#scheme DeploymentV1#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}
