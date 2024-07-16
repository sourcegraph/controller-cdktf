package googlevertexaideploymentresourcepool


type GoogleVertexAiDeploymentResourcePoolDedicatedResourcesAutoscalingMetricSpecs struct {
	// The resource metric name. Supported metrics: For Online Prediction: * 'aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle' * 'aiplatform.googleapis.com/prediction/online/cpu/utilization'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_deployment_resource_pool#metric_name GoogleVertexAiDeploymentResourcePool#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The target resource utilization in percentage (1% - 100%) for the given metric;
	//
	// once the real usage deviates from the target by a certain percentage, the machine replicas change. The default value is 60 (representing 60%) if not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_deployment_resource_pool#target GoogleVertexAiDeploymentResourcePool#target}
	Target *float64 `field:"optional" json:"target" yaml:"target"`
}

