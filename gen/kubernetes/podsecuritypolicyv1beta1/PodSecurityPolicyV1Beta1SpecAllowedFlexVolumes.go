package podsecuritypolicyv1beta1


type PodSecurityPolicyV1Beta1SpecAllowedFlexVolumes struct {
	// driver is the name of the Flexvolume driver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.15.0/docs/resources/pod_security_policy_v1beta1#driver PodSecurityPolicyV1Beta1#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
}

