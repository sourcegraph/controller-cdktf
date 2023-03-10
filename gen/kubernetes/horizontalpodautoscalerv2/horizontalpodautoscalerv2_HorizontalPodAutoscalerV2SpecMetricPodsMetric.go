package horizontalpodautoscalerv2


type HorizontalPodAutoscalerV2SpecMetricPodsMetric struct {
	// name is the name of the given metric.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler_v2#name HorizontalPodAutoscalerV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// selector block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/horizontal_pod_autoscaler_v2#selector HorizontalPodAutoscalerV2#selector}
	Selector interface{} `field:"optional" json:"selector" yaml:"selector"`
}

