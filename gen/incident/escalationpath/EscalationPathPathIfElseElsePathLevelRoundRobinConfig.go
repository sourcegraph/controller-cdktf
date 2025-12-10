package escalationpath


type EscalationPathPathIfElseElsePathLevelRoundRobinConfig struct {
	// Whether round robin is enabled for this level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#enabled EscalationPath#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// How long should we wait before rotating to the next target in a round robin, if not set will stick with a single target per level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/incident-io/incident/5.21.1/docs/resources/escalation_path#rotate_after_seconds EscalationPath#rotate_after_seconds}
	RotateAfterSeconds *float64 `field:"optional" json:"rotateAfterSeconds" yaml:"rotateAfterSeconds"`
}

