package horizontalpodautoscaler

type HorizontalPodAutoscalerSpecBehavior struct {
	// scale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/horizontal_pod_autoscaler#scale_down HorizontalPodAutoscaler#scale_down}
	ScaleDown any `field:"optional" json:"scaleDown" yaml:"scaleDown"`
	// scale_up block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/horizontal_pod_autoscaler#scale_up HorizontalPodAutoscaler#scale_up}
	ScaleUp any `field:"optional" json:"scaleUp" yaml:"scaleUp"`
}
