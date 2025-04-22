package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesFromNotSourcesResources struct {
	// iam_service_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#iam_service_account GoogleNetworkSecurityAuthzPolicy#iam_service_account}
	IamServiceAccount *GoogleNetworkSecurityAuthzPolicyHttpRulesFromNotSourcesResourcesIamServiceAccount `field:"optional" json:"iamServiceAccount" yaml:"iamServiceAccount"`
	// tag_value_id_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#tag_value_id_set GoogleNetworkSecurityAuthzPolicy#tag_value_id_set}
	TagValueIdSet *GoogleNetworkSecurityAuthzPolicyHttpRulesFromNotSourcesResourcesTagValueIdSet `field:"optional" json:"tagValueIdSet" yaml:"tagValueIdSet"`
}

