package ssoadminpermissionsboundaryattachment


type SsoadminPermissionsBoundaryAttachmentPermissionsBoundary struct {
	// customer_managed_policy_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssoadmin_permissions_boundary_attachment#customer_managed_policy_reference SsoadminPermissionsBoundaryAttachment#customer_managed_policy_reference}
	CustomerManagedPolicyReference *SsoadminPermissionsBoundaryAttachmentPermissionsBoundaryCustomerManagedPolicyReference `field:"optional" json:"customerManagedPolicyReference" yaml:"customerManagedPolicyReference"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ssoadmin_permissions_boundary_attachment#managed_policy_arn SsoadminPermissionsBoundaryAttachment#managed_policy_arn}.
	ManagedPolicyArn *string `field:"optional" json:"managedPolicyArn" yaml:"managedPolicyArn"`
}

