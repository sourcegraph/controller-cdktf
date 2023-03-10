package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/provider/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta google-beta}.
type GoogleBetaProvider interface {
	cdktf.TerraformProvider
	AccessApprovalCustomEndpoint() *string
	SetAccessApprovalCustomEndpoint(val *string)
	AccessApprovalCustomEndpointInput() *string
	AccessContextManagerCustomEndpoint() *string
	SetAccessContextManagerCustomEndpoint(val *string)
	AccessContextManagerCustomEndpointInput() *string
	AccessToken() *string
	SetAccessToken(val *string)
	AccessTokenInput() *string
	ActiveDirectoryCustomEndpoint() *string
	SetActiveDirectoryCustomEndpoint(val *string)
	ActiveDirectoryCustomEndpointInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiGatewayCustomEndpoint() *string
	SetApiGatewayCustomEndpoint(val *string)
	ApiGatewayCustomEndpointInput() *string
	ApigeeCustomEndpoint() *string
	SetApigeeCustomEndpoint(val *string)
	ApigeeCustomEndpointInput() *string
	ApikeysCustomEndpoint() *string
	SetApikeysCustomEndpoint(val *string)
	ApikeysCustomEndpointInput() *string
	AppEngineCustomEndpoint() *string
	SetAppEngineCustomEndpoint(val *string)
	AppEngineCustomEndpointInput() *string
	ArtifactRegistryCustomEndpoint() *string
	SetArtifactRegistryCustomEndpoint(val *string)
	ArtifactRegistryCustomEndpointInput() *string
	AssuredWorkloadsCustomEndpoint() *string
	SetAssuredWorkloadsCustomEndpoint(val *string)
	AssuredWorkloadsCustomEndpointInput() *string
	Batching() *GoogleBetaProviderBatching
	SetBatching(val *GoogleBetaProviderBatching)
	BatchingInput() *GoogleBetaProviderBatching
	BigqueryAnalyticsHubCustomEndpoint() *string
	SetBigqueryAnalyticsHubCustomEndpoint(val *string)
	BigqueryAnalyticsHubCustomEndpointInput() *string
	BigqueryConnectionCustomEndpoint() *string
	SetBigqueryConnectionCustomEndpoint(val *string)
	BigqueryConnectionCustomEndpointInput() *string
	BigQueryCustomEndpoint() *string
	SetBigQueryCustomEndpoint(val *string)
	BigQueryCustomEndpointInput() *string
	BigqueryDataTransferCustomEndpoint() *string
	SetBigqueryDataTransferCustomEndpoint(val *string)
	BigqueryDataTransferCustomEndpointInput() *string
	BigqueryReservationCustomEndpoint() *string
	SetBigqueryReservationCustomEndpoint(val *string)
	BigqueryReservationCustomEndpointInput() *string
	BigtableCustomEndpoint() *string
	SetBigtableCustomEndpoint(val *string)
	BigtableCustomEndpointInput() *string
	BillingCustomEndpoint() *string
	SetBillingCustomEndpoint(val *string)
	BillingCustomEndpointInput() *string
	BillingProject() *string
	SetBillingProject(val *string)
	BillingProjectInput() *string
	BinaryAuthorizationCustomEndpoint() *string
	SetBinaryAuthorizationCustomEndpoint(val *string)
	BinaryAuthorizationCustomEndpointInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateManagerCustomEndpoint() *string
	SetCertificateManagerCustomEndpoint(val *string)
	CertificateManagerCustomEndpointInput() *string
	CloudAssetCustomEndpoint() *string
	SetCloudAssetCustomEndpoint(val *string)
	CloudAssetCustomEndpointInput() *string
	CloudBillingCustomEndpoint() *string
	SetCloudBillingCustomEndpoint(val *string)
	CloudBillingCustomEndpointInput() *string
	CloudBuildCustomEndpoint() *string
	SetCloudBuildCustomEndpoint(val *string)
	CloudBuildCustomEndpointInput() *string
	CloudBuildWorkerPoolCustomEndpoint() *string
	SetCloudBuildWorkerPoolCustomEndpoint(val *string)
	CloudBuildWorkerPoolCustomEndpointInput() *string
	ClouddeployCustomEndpoint() *string
	SetClouddeployCustomEndpoint(val *string)
	ClouddeployCustomEndpointInput() *string
	Cloudfunctions2CustomEndpoint() *string
	SetCloudfunctions2CustomEndpoint(val *string)
	Cloudfunctions2CustomEndpointInput() *string
	CloudFunctionsCustomEndpoint() *string
	SetCloudFunctionsCustomEndpoint(val *string)
	CloudFunctionsCustomEndpointInput() *string
	CloudIdentityCustomEndpoint() *string
	SetCloudIdentityCustomEndpoint(val *string)
	CloudIdentityCustomEndpointInput() *string
	CloudIotCustomEndpoint() *string
	SetCloudIotCustomEndpoint(val *string)
	CloudIotCustomEndpointInput() *string
	CloudResourceManagerCustomEndpoint() *string
	SetCloudResourceManagerCustomEndpoint(val *string)
	CloudResourceManagerCustomEndpointInput() *string
	CloudRunCustomEndpoint() *string
	SetCloudRunCustomEndpoint(val *string)
	CloudRunCustomEndpointInput() *string
	CloudSchedulerCustomEndpoint() *string
	SetCloudSchedulerCustomEndpoint(val *string)
	CloudSchedulerCustomEndpointInput() *string
	CloudTasksCustomEndpoint() *string
	SetCloudTasksCustomEndpoint(val *string)
	CloudTasksCustomEndpointInput() *string
	ComposerCustomEndpoint() *string
	SetComposerCustomEndpoint(val *string)
	ComposerCustomEndpointInput() *string
	ComputeCustomEndpoint() *string
	SetComputeCustomEndpoint(val *string)
	ComputeCustomEndpointInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerAnalysisCustomEndpoint() *string
	SetContainerAnalysisCustomEndpoint(val *string)
	ContainerAnalysisCustomEndpointInput() *string
	ContainerAwsCustomEndpoint() *string
	SetContainerAwsCustomEndpoint(val *string)
	ContainerAwsCustomEndpointInput() *string
	ContainerAzureCustomEndpoint() *string
	SetContainerAzureCustomEndpoint(val *string)
	ContainerAzureCustomEndpointInput() *string
	ContainerCustomEndpoint() *string
	SetContainerCustomEndpoint(val *string)
	ContainerCustomEndpointInput() *string
	Credentials() *string
	SetCredentials(val *string)
	CredentialsInput() *string
	DataCatalogCustomEndpoint() *string
	SetDataCatalogCustomEndpoint(val *string)
	DataCatalogCustomEndpointInput() *string
	DataflowCustomEndpoint() *string
	SetDataflowCustomEndpoint(val *string)
	DataflowCustomEndpointInput() *string
	DataFusionCustomEndpoint() *string
	SetDataFusionCustomEndpoint(val *string)
	DataFusionCustomEndpointInput() *string
	DataLossPreventionCustomEndpoint() *string
	SetDataLossPreventionCustomEndpoint(val *string)
	DataLossPreventionCustomEndpointInput() *string
	DataplexCustomEndpoint() *string
	SetDataplexCustomEndpoint(val *string)
	DataplexCustomEndpointInput() *string
	DataprocCustomEndpoint() *string
	SetDataprocCustomEndpoint(val *string)
	DataprocCustomEndpointInput() *string
	DataprocMetastoreCustomEndpoint() *string
	SetDataprocMetastoreCustomEndpoint(val *string)
	DataprocMetastoreCustomEndpointInput() *string
	DatastoreCustomEndpoint() *string
	SetDatastoreCustomEndpoint(val *string)
	DatastoreCustomEndpointInput() *string
	DatastreamCustomEndpoint() *string
	SetDatastreamCustomEndpoint(val *string)
	DatastreamCustomEndpointInput() *string
	DeploymentManagerCustomEndpoint() *string
	SetDeploymentManagerCustomEndpoint(val *string)
	DeploymentManagerCustomEndpointInput() *string
	DialogflowCustomEndpoint() *string
	SetDialogflowCustomEndpoint(val *string)
	DialogflowCustomEndpointInput() *string
	DialogflowCxCustomEndpoint() *string
	SetDialogflowCxCustomEndpoint(val *string)
	DialogflowCxCustomEndpointInput() *string
	DnsCustomEndpoint() *string
	SetDnsCustomEndpoint(val *string)
	DnsCustomEndpointInput() *string
	DocumentAiCustomEndpoint() *string
	SetDocumentAiCustomEndpoint(val *string)
	DocumentAiCustomEndpointInput() *string
	EssentialContactsCustomEndpoint() *string
	SetEssentialContactsCustomEndpoint(val *string)
	EssentialContactsCustomEndpointInput() *string
	EventarcCustomEndpoint() *string
	SetEventarcCustomEndpoint(val *string)
	EventarcCustomEndpointInput() *string
	FilestoreCustomEndpoint() *string
	SetFilestoreCustomEndpoint(val *string)
	FilestoreCustomEndpointInput() *string
	FirebaseCustomEndpoint() *string
	SetFirebaseCustomEndpoint(val *string)
	FirebaseCustomEndpointInput() *string
	FirebaserulesCustomEndpoint() *string
	SetFirebaserulesCustomEndpoint(val *string)
	FirebaserulesCustomEndpointInput() *string
	FirestoreCustomEndpoint() *string
	SetFirestoreCustomEndpoint(val *string)
	FirestoreCustomEndpointInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GameServicesCustomEndpoint() *string
	SetGameServicesCustomEndpoint(val *string)
	GameServicesCustomEndpointInput() *string
	GkeHubCustomEndpoint() *string
	SetGkeHubCustomEndpoint(val *string)
	GkeHubCustomEndpointInput() *string
	GkehubFeatureCustomEndpoint() *string
	SetGkehubFeatureCustomEndpoint(val *string)
	GkehubFeatureCustomEndpointInput() *string
	HealthcareCustomEndpoint() *string
	SetHealthcareCustomEndpoint(val *string)
	HealthcareCustomEndpointInput() *string
	Iam2CustomEndpoint() *string
	SetIam2CustomEndpoint(val *string)
	Iam2CustomEndpointInput() *string
	IamBetaCustomEndpoint() *string
	SetIamBetaCustomEndpoint(val *string)
	IamBetaCustomEndpointInput() *string
	IamCredentialsCustomEndpoint() *string
	SetIamCredentialsCustomEndpoint(val *string)
	IamCredentialsCustomEndpointInput() *string
	IamCustomEndpoint() *string
	SetIamCustomEndpoint(val *string)
	IamCustomEndpointInput() *string
	IapCustomEndpoint() *string
	SetIapCustomEndpoint(val *string)
	IapCustomEndpointInput() *string
	IdentityPlatformCustomEndpoint() *string
	SetIdentityPlatformCustomEndpoint(val *string)
	IdentityPlatformCustomEndpointInput() *string
	ImpersonateServiceAccount() *string
	SetImpersonateServiceAccount(val *string)
	ImpersonateServiceAccountDelegates() *[]*string
	SetImpersonateServiceAccountDelegates(val *[]*string)
	ImpersonateServiceAccountDelegatesInput() *[]*string
	ImpersonateServiceAccountInput() *string
	KmsCustomEndpoint() *string
	SetKmsCustomEndpoint(val *string)
	KmsCustomEndpointInput() *string
	LoggingCustomEndpoint() *string
	SetLoggingCustomEndpoint(val *string)
	LoggingCustomEndpointInput() *string
	MemcacheCustomEndpoint() *string
	SetMemcacheCustomEndpoint(val *string)
	MemcacheCustomEndpointInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MlEngineCustomEndpoint() *string
	SetMlEngineCustomEndpoint(val *string)
	MlEngineCustomEndpointInput() *string
	MonitoringCustomEndpoint() *string
	SetMonitoringCustomEndpoint(val *string)
	MonitoringCustomEndpointInput() *string
	NetworkConnectivityCustomEndpoint() *string
	SetNetworkConnectivityCustomEndpoint(val *string)
	NetworkConnectivityCustomEndpointInput() *string
	NetworkManagementCustomEndpoint() *string
	SetNetworkManagementCustomEndpoint(val *string)
	NetworkManagementCustomEndpointInput() *string
	NetworkServicesCustomEndpoint() *string
	SetNetworkServicesCustomEndpoint(val *string)
	NetworkServicesCustomEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	NotebooksCustomEndpoint() *string
	SetNotebooksCustomEndpoint(val *string)
	NotebooksCustomEndpointInput() *string
	OrgPolicyCustomEndpoint() *string
	SetOrgPolicyCustomEndpoint(val *string)
	OrgPolicyCustomEndpointInput() *string
	OsConfigCustomEndpoint() *string
	SetOsConfigCustomEndpoint(val *string)
	OsConfigCustomEndpointInput() *string
	OsLoginCustomEndpoint() *string
	SetOsLoginCustomEndpoint(val *string)
	OsLoginCustomEndpointInput() *string
	PrivatecaCustomEndpoint() *string
	SetPrivatecaCustomEndpoint(val *string)
	PrivatecaCustomEndpointInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	PubsubCustomEndpoint() *string
	SetPubsubCustomEndpoint(val *string)
	PubsubCustomEndpointInput() *string
	PubsubLiteCustomEndpoint() *string
	SetPubsubLiteCustomEndpoint(val *string)
	PubsubLiteCustomEndpointInput() *string
	// Experimental.
	RawOverrides() interface{}
	RecaptchaEnterpriseCustomEndpoint() *string
	SetRecaptchaEnterpriseCustomEndpoint(val *string)
	RecaptchaEnterpriseCustomEndpointInput() *string
	RedisCustomEndpoint() *string
	SetRedisCustomEndpoint(val *string)
	RedisCustomEndpointInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RequestReason() *string
	SetRequestReason(val *string)
	RequestReasonInput() *string
	RequestTimeout() *string
	SetRequestTimeout(val *string)
	RequestTimeoutInput() *string
	ResourceManagerCustomEndpoint() *string
	SetResourceManagerCustomEndpoint(val *string)
	ResourceManagerCustomEndpointInput() *string
	ResourceManagerV3CustomEndpoint() *string
	SetResourceManagerV3CustomEndpoint(val *string)
	ResourceManagerV3CustomEndpointInput() *string
	RuntimeconfigCustomEndpoint() *string
	SetRuntimeconfigCustomEndpoint(val *string)
	RuntimeConfigCustomEndpoint() *string
	SetRuntimeConfigCustomEndpoint(val *string)
	RuntimeconfigCustomEndpointInput() *string
	RuntimeConfigCustomEndpointInput() *string
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	SecretManagerCustomEndpoint() *string
	SetSecretManagerCustomEndpoint(val *string)
	SecretManagerCustomEndpointInput() *string
	SecurityCenterCustomEndpoint() *string
	SetSecurityCenterCustomEndpoint(val *string)
	SecurityCenterCustomEndpointInput() *string
	SecurityScannerCustomEndpoint() *string
	SetSecurityScannerCustomEndpoint(val *string)
	SecurityScannerCustomEndpointInput() *string
	ServiceDirectoryCustomEndpoint() *string
	SetServiceDirectoryCustomEndpoint(val *string)
	ServiceDirectoryCustomEndpointInput() *string
	ServiceManagementCustomEndpoint() *string
	SetServiceManagementCustomEndpoint(val *string)
	ServiceManagementCustomEndpointInput() *string
	ServiceNetworkingCustomEndpoint() *string
	SetServiceNetworkingCustomEndpoint(val *string)
	ServiceNetworkingCustomEndpointInput() *string
	ServiceUsageCustomEndpoint() *string
	SetServiceUsageCustomEndpoint(val *string)
	ServiceUsageCustomEndpointInput() *string
	SourceRepoCustomEndpoint() *string
	SetSourceRepoCustomEndpoint(val *string)
	SourceRepoCustomEndpointInput() *string
	SpannerCustomEndpoint() *string
	SetSpannerCustomEndpoint(val *string)
	SpannerCustomEndpointInput() *string
	SqlCustomEndpoint() *string
	SetSqlCustomEndpoint(val *string)
	SqlCustomEndpointInput() *string
	StorageCustomEndpoint() *string
	SetStorageCustomEndpoint(val *string)
	StorageCustomEndpointInput() *string
	StorageTransferCustomEndpoint() *string
	SetStorageTransferCustomEndpoint(val *string)
	StorageTransferCustomEndpointInput() *string
	TagsCustomEndpoint() *string
	SetTagsCustomEndpoint(val *string)
	TagsCustomEndpointInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TpuCustomEndpoint() *string
	SetTpuCustomEndpoint(val *string)
	TpuCustomEndpointInput() *string
	UserProjectOverride() interface{}
	SetUserProjectOverride(val interface{})
	UserProjectOverrideInput() interface{}
	VertexAiCustomEndpoint() *string
	SetVertexAiCustomEndpoint(val *string)
	VertexAiCustomEndpointInput() *string
	VpcAccessCustomEndpoint() *string
	SetVpcAccessCustomEndpoint(val *string)
	VpcAccessCustomEndpointInput() *string
	WorkflowsCustomEndpoint() *string
	SetWorkflowsCustomEndpoint(val *string)
	WorkflowsCustomEndpointInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAccessApprovalCustomEndpoint()
	ResetAccessContextManagerCustomEndpoint()
	ResetAccessToken()
	ResetActiveDirectoryCustomEndpoint()
	ResetAlias()
	ResetApiGatewayCustomEndpoint()
	ResetApigeeCustomEndpoint()
	ResetApikeysCustomEndpoint()
	ResetAppEngineCustomEndpoint()
	ResetArtifactRegistryCustomEndpoint()
	ResetAssuredWorkloadsCustomEndpoint()
	ResetBatching()
	ResetBigqueryAnalyticsHubCustomEndpoint()
	ResetBigqueryConnectionCustomEndpoint()
	ResetBigQueryCustomEndpoint()
	ResetBigqueryDataTransferCustomEndpoint()
	ResetBigqueryReservationCustomEndpoint()
	ResetBigtableCustomEndpoint()
	ResetBillingCustomEndpoint()
	ResetBillingProject()
	ResetBinaryAuthorizationCustomEndpoint()
	ResetCertificateManagerCustomEndpoint()
	ResetCloudAssetCustomEndpoint()
	ResetCloudBillingCustomEndpoint()
	ResetCloudBuildCustomEndpoint()
	ResetCloudBuildWorkerPoolCustomEndpoint()
	ResetClouddeployCustomEndpoint()
	ResetCloudfunctions2CustomEndpoint()
	ResetCloudFunctionsCustomEndpoint()
	ResetCloudIdentityCustomEndpoint()
	ResetCloudIotCustomEndpoint()
	ResetCloudResourceManagerCustomEndpoint()
	ResetCloudRunCustomEndpoint()
	ResetCloudSchedulerCustomEndpoint()
	ResetCloudTasksCustomEndpoint()
	ResetComposerCustomEndpoint()
	ResetComputeCustomEndpoint()
	ResetContainerAnalysisCustomEndpoint()
	ResetContainerAwsCustomEndpoint()
	ResetContainerAzureCustomEndpoint()
	ResetContainerCustomEndpoint()
	ResetCredentials()
	ResetDataCatalogCustomEndpoint()
	ResetDataflowCustomEndpoint()
	ResetDataFusionCustomEndpoint()
	ResetDataLossPreventionCustomEndpoint()
	ResetDataplexCustomEndpoint()
	ResetDataprocCustomEndpoint()
	ResetDataprocMetastoreCustomEndpoint()
	ResetDatastoreCustomEndpoint()
	ResetDatastreamCustomEndpoint()
	ResetDeploymentManagerCustomEndpoint()
	ResetDialogflowCustomEndpoint()
	ResetDialogflowCxCustomEndpoint()
	ResetDnsCustomEndpoint()
	ResetDocumentAiCustomEndpoint()
	ResetEssentialContactsCustomEndpoint()
	ResetEventarcCustomEndpoint()
	ResetFilestoreCustomEndpoint()
	ResetFirebaseCustomEndpoint()
	ResetFirebaserulesCustomEndpoint()
	ResetFirestoreCustomEndpoint()
	ResetGameServicesCustomEndpoint()
	ResetGkeHubCustomEndpoint()
	ResetGkehubFeatureCustomEndpoint()
	ResetHealthcareCustomEndpoint()
	ResetIam2CustomEndpoint()
	ResetIamBetaCustomEndpoint()
	ResetIamCredentialsCustomEndpoint()
	ResetIamCustomEndpoint()
	ResetIapCustomEndpoint()
	ResetIdentityPlatformCustomEndpoint()
	ResetImpersonateServiceAccount()
	ResetImpersonateServiceAccountDelegates()
	ResetKmsCustomEndpoint()
	ResetLoggingCustomEndpoint()
	ResetMemcacheCustomEndpoint()
	ResetMlEngineCustomEndpoint()
	ResetMonitoringCustomEndpoint()
	ResetNetworkConnectivityCustomEndpoint()
	ResetNetworkManagementCustomEndpoint()
	ResetNetworkServicesCustomEndpoint()
	ResetNotebooksCustomEndpoint()
	ResetOrgPolicyCustomEndpoint()
	ResetOsConfigCustomEndpoint()
	ResetOsLoginCustomEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivatecaCustomEndpoint()
	ResetProject()
	ResetPubsubCustomEndpoint()
	ResetPubsubLiteCustomEndpoint()
	ResetRecaptchaEnterpriseCustomEndpoint()
	ResetRedisCustomEndpoint()
	ResetRegion()
	ResetRequestReason()
	ResetRequestTimeout()
	ResetResourceManagerCustomEndpoint()
	ResetResourceManagerV3CustomEndpoint()
	ResetRuntimeconfigCustomEndpoint()
	ResetRuntimeConfigCustomEndpoint()
	ResetScopes()
	ResetSecretManagerCustomEndpoint()
	ResetSecurityCenterCustomEndpoint()
	ResetSecurityScannerCustomEndpoint()
	ResetServiceDirectoryCustomEndpoint()
	ResetServiceManagementCustomEndpoint()
	ResetServiceNetworkingCustomEndpoint()
	ResetServiceUsageCustomEndpoint()
	ResetSourceRepoCustomEndpoint()
	ResetSpannerCustomEndpoint()
	ResetSqlCustomEndpoint()
	ResetStorageCustomEndpoint()
	ResetStorageTransferCustomEndpoint()
	ResetTagsCustomEndpoint()
	ResetTpuCustomEndpoint()
	ResetUserProjectOverride()
	ResetVertexAiCustomEndpoint()
	ResetVpcAccessCustomEndpoint()
	ResetWorkflowsCustomEndpoint()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleBetaProvider
type jsiiProxy_GoogleBetaProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_GoogleBetaProvider) AccessApprovalCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessApprovalCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AccessApprovalCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessApprovalCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AccessContextManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessContextManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AccessContextManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessContextManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ActiveDirectoryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ActiveDirectoryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ApiGatewayCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGatewayCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ApiGatewayCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGatewayCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ApigeeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ApigeeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apigeeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ApikeysCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apikeysCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ApikeysCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apikeysCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AppEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AppEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ArtifactRegistryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ArtifactRegistryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactRegistryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AssuredWorkloadsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assuredWorkloadsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) AssuredWorkloadsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assuredWorkloadsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Batching() *GoogleBetaProviderBatching {
	var returns *GoogleBetaProviderBatching
	_jsii_.Get(
		j,
		"batching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BatchingInput() *GoogleBetaProviderBatching {
	var returns *GoogleBetaProviderBatching
	_jsii_.Get(
		j,
		"batchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryAnalyticsHubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryAnalyticsHubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryAnalyticsHubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryAnalyticsHubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryConnectionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryConnectionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryConnectionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryConnectionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigQueryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigQueryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigQueryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigQueryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryDataTransferCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDataTransferCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryDataTransferCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryDataTransferCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryReservationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryReservationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigqueryReservationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigqueryReservationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigtableCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigtableCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BigtableCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bigtableCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BillingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BillingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BillingProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BillingProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BinaryAuthorizationCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAuthorizationCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) BinaryAuthorizationCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binaryAuthorizationCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CertificateManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CertificateManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudAssetCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAssetCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudAssetCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudAssetCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudBillingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBillingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudBillingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBillingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudBuildCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudBuildCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudBuildWorkerPoolCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildWorkerPoolCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudBuildWorkerPoolCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuildWorkerPoolCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ClouddeployCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddeployCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ClouddeployCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clouddeployCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Cloudfunctions2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfunctions2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Cloudfunctions2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfunctions2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudFunctionsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudFunctionsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudIdentityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdentityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudIdentityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIdentityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudIotCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIotCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudIotCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudIotCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudResourceManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudResourceManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudResourceManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudResourceManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudRunCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudRunCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRunCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudSchedulerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSchedulerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudSchedulerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudSchedulerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudTasksCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTasksCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CloudTasksCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudTasksCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ComposerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ComposerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ComputeCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ComputeCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerAnalysisCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAnalysisCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerAnalysisCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAnalysisCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerAwsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAwsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerAwsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAwsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerAzureCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAzureCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerAzureCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAzureCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ContainerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataCatalogCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCatalogCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataCatalogCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCatalogCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataflowCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataflowCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataflowCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataflowCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataFusionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFusionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataFusionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFusionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataLossPreventionCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLossPreventionCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataLossPreventionCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLossPreventionCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataplexCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataplexCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataplexCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataplexCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataprocCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataprocCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataprocMetastoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocMetastoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DataprocMetastoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocMetastoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DatastoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DatastoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DatastreamCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastreamCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DatastreamCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datastreamCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DeploymentManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DeploymentManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DialogflowCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DialogflowCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DialogflowCxCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCxCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DialogflowCxCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowCxCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DnsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DnsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DocumentAiCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) DocumentAiCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentAiCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) EssentialContactsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"essentialContactsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) EssentialContactsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"essentialContactsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) EventarcCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) EventarcCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FilestoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filestoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FilestoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filestoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FirebaseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FirebaseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FirebaserulesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaserulesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FirebaserulesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firebaserulesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FirestoreCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FirestoreCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firestoreCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) GameServicesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServicesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) GameServicesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gameServicesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) GkeHubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) GkeHubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeHubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) GkehubFeatureCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkehubFeatureCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) GkehubFeatureCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkehubFeatureCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) HealthcareCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcareCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) HealthcareCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthcareCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Iam2CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam2CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Iam2CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iam2CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IamBetaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamBetaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IamBetaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamBetaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IamCredentialsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCredentialsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IamCredentialsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCredentialsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IamCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IamCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IapCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iapCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IapCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iapCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IdentityPlatformCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPlatformCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) IdentityPlatformCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPlatformCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ImpersonateServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ImpersonateServiceAccountDelegates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"impersonateServiceAccountDelegates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ImpersonateServiceAccountDelegatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"impersonateServiceAccountDelegatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ImpersonateServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impersonateServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) KmsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) KmsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) LoggingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) LoggingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) MemcacheCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memcacheCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) MemcacheCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memcacheCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) MlEngineCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlEngineCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) MlEngineCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mlEngineCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) MonitoringCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) MonitoringCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NetworkConnectivityCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NetworkConnectivityCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkConnectivityCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NetworkManagementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkManagementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NetworkManagementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkManagementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NetworkServicesCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkServicesCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NetworkServicesCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkServicesCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NotebooksCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebooksCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) NotebooksCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notebooksCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) OrgPolicyCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgPolicyCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) OrgPolicyCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgPolicyCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) OsConfigCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) OsConfigCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osConfigCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) OsLoginCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osLoginCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) OsLoginCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osLoginCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) PrivatecaCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privatecaCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) PrivatecaCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privatecaCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) PubsubCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) PubsubCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) PubsubLiteCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubLiteCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) PubsubLiteCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubLiteCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RecaptchaEnterpriseCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaEnterpriseCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RecaptchaEnterpriseCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recaptchaEnterpriseCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RedisCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RedisCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redisCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RequestReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RequestReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RequestTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RequestTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ResourceManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ResourceManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ResourceManagerV3CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerV3CustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ResourceManagerV3CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceManagerV3CustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RuntimeconfigCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeconfigCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RuntimeConfigCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeConfigCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RuntimeconfigCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeconfigCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) RuntimeConfigCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeConfigCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SecretManagerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SecretManagerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretManagerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SecurityCenterCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SecurityCenterCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityCenterCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SecurityScannerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityScannerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SecurityScannerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityScannerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceDirectoryCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceDirectoryCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectoryCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceManagementCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceManagementCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceManagementCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceNetworkingCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkingCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceNetworkingCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNetworkingCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceUsageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ServiceUsageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUsageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SourceRepoCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRepoCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SourceRepoCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRepoCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SpannerCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spannerCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SpannerCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spannerCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SqlCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) SqlCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) StorageCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) StorageCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) StorageTransferCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTransferCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) StorageTransferCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTransferCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) TagsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) TagsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) TpuCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) TpuCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) UserProjectOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProjectOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) UserProjectOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userProjectOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) VertexAiCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexAiCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) VertexAiCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vertexAiCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) VpcAccessCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcAccessCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) VpcAccessCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcAccessCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) WorkflowsCustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsCustomEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) WorkflowsCustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsCustomEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBetaProvider) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta google-beta} Resource.
func NewGoogleBetaProvider(scope constructs.Construct, id *string, config *GoogleBetaProviderConfig) GoogleBetaProvider {
	_init_.Initialize()

	if err := validateNewGoogleBetaProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBetaProvider{}

	_jsii_.Create(
		"google-beta.provider.GoogleBetaProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta google-beta} Resource.
func NewGoogleBetaProvider_Override(g GoogleBetaProvider, scope constructs.Construct, id *string, config *GoogleBetaProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.provider.GoogleBetaProvider",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetAccessApprovalCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"accessApprovalCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetAccessContextManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"accessContextManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetAccessToken(val *string) {
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetActiveDirectoryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetApiGatewayCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apiGatewayCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetApigeeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apigeeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetApikeysCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"apikeysCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetAppEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"appEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetArtifactRegistryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"artifactRegistryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetAssuredWorkloadsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"assuredWorkloadsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBatching(val *GoogleBetaProviderBatching) {
	if err := j.validateSetBatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batching",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBigqueryAnalyticsHubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryAnalyticsHubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBigqueryConnectionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryConnectionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBigQueryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigQueryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBigqueryDataTransferCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryDataTransferCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBigqueryReservationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigqueryReservationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBigtableCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"bigtableCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBillingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"billingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBillingProject(val *string) {
	_jsii_.Set(
		j,
		"billingProject",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetBinaryAuthorizationCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"binaryAuthorizationCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCertificateManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"certificateManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudAssetCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudAssetCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudBillingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBillingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudBuildCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBuildCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudBuildWorkerPoolCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudBuildWorkerPoolCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetClouddeployCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"clouddeployCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudfunctions2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudfunctions2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudFunctionsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudFunctionsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudIdentityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudIdentityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudIotCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudIotCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudResourceManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudResourceManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudRunCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudRunCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudSchedulerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudSchedulerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCloudTasksCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"cloudTasksCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetComposerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"composerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetComputeCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"computeCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetContainerAnalysisCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAnalysisCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetContainerAwsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAwsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetContainerAzureCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerAzureCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetContainerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"containerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetCredentials(val *string) {
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDataCatalogCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataCatalogCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDataflowCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataflowCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDataFusionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataFusionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDataLossPreventionCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataLossPreventionCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDataplexCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataplexCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDataprocCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataprocCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDataprocMetastoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dataprocMetastoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDatastoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"datastoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDatastreamCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"datastreamCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDeploymentManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"deploymentManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDialogflowCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dialogflowCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDialogflowCxCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dialogflowCxCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDnsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"dnsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetDocumentAiCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"documentAiCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetEssentialContactsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"essentialContactsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetEventarcCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"eventarcCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetFilestoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"filestoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetFirebaseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetFirebaserulesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firebaserulesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetFirestoreCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"firestoreCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetGameServicesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gameServicesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetGkeHubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkeHubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetGkehubFeatureCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gkehubFeatureCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetHealthcareCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"healthcareCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetIam2CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iam2CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetIamBetaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamBetaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetIamCredentialsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamCredentialsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetIamCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iamCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetIapCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"iapCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetIdentityPlatformCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"identityPlatformCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetImpersonateServiceAccount(val *string) {
	_jsii_.Set(
		j,
		"impersonateServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetImpersonateServiceAccountDelegates(val *[]*string) {
	_jsii_.Set(
		j,
		"impersonateServiceAccountDelegates",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetKmsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"kmsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetLoggingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"loggingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetMemcacheCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"memcacheCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetMlEngineCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"mlEngineCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetMonitoringCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"monitoringCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetNetworkConnectivityCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkConnectivityCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetNetworkManagementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkManagementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetNetworkServicesCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"networkServicesCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetNotebooksCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"notebooksCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetOrgPolicyCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"orgPolicyCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetOsConfigCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"osConfigCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetOsLoginCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"osLoginCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetPrivatecaCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"privatecaCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetProject(val *string) {
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetPubsubCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"pubsubCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetPubsubLiteCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"pubsubLiteCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetRecaptchaEnterpriseCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"recaptchaEnterpriseCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetRedisCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"redisCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetRequestReason(val *string) {
	_jsii_.Set(
		j,
		"requestReason",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetRequestTimeout(val *string) {
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetResourceManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourceManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetResourceManagerV3CustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"resourceManagerV3CustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetRuntimeconfigCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"runtimeconfigCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetRuntimeConfigCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"runtimeConfigCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetSecretManagerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"secretManagerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetSecurityCenterCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securityCenterCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetSecurityScannerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"securityScannerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetServiceDirectoryCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceDirectoryCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetServiceManagementCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceManagementCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetServiceNetworkingCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceNetworkingCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetServiceUsageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"serviceUsageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetSourceRepoCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"sourceRepoCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetSpannerCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"spannerCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetSqlCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"sqlCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetStorageCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetStorageTransferCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"storageTransferCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetTagsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tagsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetTpuCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tpuCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetUserProjectOverride(val interface{}) {
	if err := j.validateSetUserProjectOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userProjectOverride",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetVertexAiCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vertexAiCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetVpcAccessCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"vpcAccessCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetWorkflowsCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"workflowsCustomEndpoint",
		val,
	)
}

func (j *jsiiProxy_GoogleBetaProvider)SetZone(val *string) {
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleBetaProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBetaProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.provider.GoogleBetaProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBetaProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.provider.GoogleBetaProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBetaProvider) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBetaProvider) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetAccessApprovalCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessApprovalCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetAccessContextManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessContextManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetAccessToken() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetActiveDirectoryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetActiveDirectoryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		g,
		"resetAlias",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetApiGatewayCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApiGatewayCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetApigeeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApigeeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetApikeysCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetApikeysCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetAppEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAppEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetArtifactRegistryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifactRegistryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetAssuredWorkloadsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetAssuredWorkloadsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBatching() {
	_jsii_.InvokeVoid(
		g,
		"resetBatching",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBigqueryAnalyticsHubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryAnalyticsHubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBigqueryConnectionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryConnectionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBigQueryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigQueryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBigqueryDataTransferCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryDataTransferCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBigqueryReservationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigqueryReservationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBigtableCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBigtableCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBillingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBillingProject() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetBinaryAuthorizationCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryAuthorizationCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCertificateManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCertificateManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudAssetCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudAssetCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudBillingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBillingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudBuildCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBuildCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudBuildWorkerPoolCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudBuildWorkerPoolCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetClouddeployCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetClouddeployCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudfunctions2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudfunctions2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudFunctionsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudFunctionsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudIdentityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudIdentityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudIotCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudIotCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudResourceManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudResourceManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudRunCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRunCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudSchedulerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudSchedulerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCloudTasksCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudTasksCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetComposerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetComposerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetComputeCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetComputeCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetContainerAnalysisCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAnalysisCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetContainerAwsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAwsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetContainerAzureCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerAzureCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetContainerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetContainerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDataCatalogCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataCatalogCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDataflowCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataflowCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDataFusionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataFusionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDataLossPreventionCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataLossPreventionCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDataplexCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataplexCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDataprocCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDataprocMetastoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDataprocMetastoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDatastoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDatastreamCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDatastreamCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDeploymentManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDialogflowCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDialogflowCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDialogflowCxCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDialogflowCxCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDnsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDnsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetDocumentAiCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentAiCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetEssentialContactsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEssentialContactsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetEventarcCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetEventarcCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetFilestoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFilestoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetFirebaseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetFirebaserulesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirebaserulesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetFirestoreCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetFirestoreCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetGameServicesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGameServicesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetGkeHubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkeHubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetGkehubFeatureCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetGkehubFeatureCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetHealthcareCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthcareCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetIam2CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIam2CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetIamBetaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamBetaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetIamCredentialsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamCredentialsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetIamCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIamCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetIapCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIapCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetIdentityPlatformCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentityPlatformCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetImpersonateServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetImpersonateServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetImpersonateServiceAccountDelegates() {
	_jsii_.InvokeVoid(
		g,
		"resetImpersonateServiceAccountDelegates",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetKmsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetLoggingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetMemcacheCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMemcacheCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetMlEngineCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMlEngineCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetMonitoringCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoringCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetNetworkConnectivityCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConnectivityCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetNetworkManagementCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkManagementCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetNetworkServicesCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkServicesCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetNotebooksCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetNotebooksCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetOrgPolicyCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOrgPolicyCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetOsConfigCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOsConfigCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetOsLoginCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetOsLoginCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetPrivatecaCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPrivatecaCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetPubsubCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetPubsubLiteCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubLiteCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetRecaptchaEnterpriseCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRecaptchaEnterpriseCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetRedisCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRedisCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetRequestReason() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestReason",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetResourceManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetResourceManagerV3CustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerV3CustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetRuntimeconfigCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeconfigCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetRuntimeConfigCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetRuntimeConfigCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetScopes() {
	_jsii_.InvokeVoid(
		g,
		"resetScopes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetSecretManagerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretManagerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetSecurityCenterCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityCenterCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetSecurityScannerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityScannerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetServiceDirectoryCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceDirectoryCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetServiceManagementCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceManagementCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetServiceNetworkingCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceNetworkingCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetServiceUsageCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceUsageCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetSourceRepoCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceRepoCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetSpannerCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSpannerCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetSqlCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetStorageCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetStorageTransferCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageTransferCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetTagsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTagsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetTpuCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetTpuCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetUserProjectOverride() {
	_jsii_.InvokeVoid(
		g,
		"resetUserProjectOverride",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetVertexAiCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVertexAiCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetVpcAccessCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccessCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetWorkflowsCustomEndpoint() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkflowsCustomEndpoint",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) ResetZone() {
	_jsii_.InvokeVoid(
		g,
		"resetZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBetaProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBetaProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

