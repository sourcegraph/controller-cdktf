package cronjob


type CronJobSpecJobTemplateSpecTemplateSpecInitContainerStartupProbe struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#exec CronJob#exec}
	Exec *CronJobSpecJobTemplateSpecTemplateSpecInitContainerStartupProbeExec `field:"optional" json:"exec" yaml:"exec"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#failure_threshold CronJob#failure_threshold}
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#http_get CronJob#http_get}
	HttpGet *CronJobSpecJobTemplateSpecTemplateSpecInitContainerStartupProbeHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#initial_delay_seconds CronJob#initial_delay_seconds}
	InitialDelaySeconds *float64 `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#period_seconds CronJob#period_seconds}
	PeriodSeconds *float64 `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#success_threshold CronJob#success_threshold}
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#tcp_socket CronJob#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
	// Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/cron_job#timeout_seconds CronJob#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

