package slo


type SloObjectiveRawMetricQueryAzureMonitorWorkspace struct {
	// Resource group of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#resource_group Slo#resource_group}
	ResourceGroup *string `field:"required" json:"resourceGroup" yaml:"resourceGroup"`
	// Subscription ID of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#subscription_id Slo#subscription_id}
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// ID of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/slo#workspace_id Slo#workspace_id}
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

