package podsecuritypolicy


type PodSecurityPolicySpecSeLinux struct {
	// rule is the strategy that will dictate the allowable labels that may be set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy#rule PodSecurityPolicy#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// se_linux_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy#se_linux_options PodSecurityPolicy#se_linux_options}
	SeLinuxOptions interface{} `field:"optional" json:"seLinuxOptions" yaml:"seLinuxOptions"`
}

