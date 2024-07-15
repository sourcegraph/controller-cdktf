package googlecomputesecuritypolicy


type GoogleComputeSecurityPolicyAdaptiveProtectionConfigAutoDeployConfig struct {
	// Rules are only automatically deployed for alerts on potential attacks with confidence scores greater than this threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#confidence_threshold GoogleComputeSecurityPolicy#confidence_threshold}
	ConfidenceThreshold *float64 `field:"optional" json:"confidenceThreshold" yaml:"confidenceThreshold"`
	// Google Cloud Armor stops applying the action in the automatically deployed rule to an identified attacker after this duration.
	//
	// The rule continues to operate against new requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#expiration_sec GoogleComputeSecurityPolicy#expiration_sec}
	ExpirationSec *float64 `field:"optional" json:"expirationSec" yaml:"expirationSec"`
	// Rules are only automatically deployed when the estimated impact to baseline traffic from the suggested mitigation is below this threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#impacted_baseline_threshold GoogleComputeSecurityPolicy#impacted_baseline_threshold}
	ImpactedBaselineThreshold *float64 `field:"optional" json:"impactedBaselineThreshold" yaml:"impactedBaselineThreshold"`
	// Identifies new attackers only when the load to the backend service that is under attack exceeds this threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_security_policy#load_threshold GoogleComputeSecurityPolicy#load_threshold}
	LoadThreshold *float64 `field:"optional" json:"loadThreshold" yaml:"loadThreshold"`
}

