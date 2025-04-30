package googledeveloperconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledeveloperconnectconnection/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_developer_connect_connection google_developer_connect_connection}.
type GoogleDeveloperConnectConnection interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BitbucketCloudConfig() GoogleDeveloperConnectConnectionBitbucketCloudConfigOutputReference
	BitbucketCloudConfigInput() *GoogleDeveloperConnectConnectionBitbucketCloudConfig
	BitbucketDataCenterConfig() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference
	BitbucketDataCenterConfigInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	CryptoKeyConfig() GoogleDeveloperConnectConnectionCryptoKeyConfigOutputReference
	CryptoKeyConfigInput() *GoogleDeveloperConnectConnectionCryptoKeyConfig
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GithubConfig() GoogleDeveloperConnectConnectionGithubConfigOutputReference
	GithubConfigInput() *GoogleDeveloperConnectConnectionGithubConfig
	GithubEnterpriseConfig() GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference
	GithubEnterpriseConfigInput() *GoogleDeveloperConnectConnectionGithubEnterpriseConfig
	GitlabConfig() GoogleDeveloperConnectConnectionGitlabConfigOutputReference
	GitlabConfigInput() *GoogleDeveloperConnectConnectionGitlabConfig
	GitlabEnterpriseConfig() GoogleDeveloperConnectConnectionGitlabEnterpriseConfigOutputReference
	GitlabEnterpriseConfigInput() *GoogleDeveloperConnectConnectionGitlabEnterpriseConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstallationState() GoogleDeveloperConnectConnectionInstallationStateList
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleDeveloperConnectConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutBitbucketCloudConfig(value *GoogleDeveloperConnectConnectionBitbucketCloudConfig)
	PutBitbucketDataCenterConfig(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig)
	PutCryptoKeyConfig(value *GoogleDeveloperConnectConnectionCryptoKeyConfig)
	PutGithubConfig(value *GoogleDeveloperConnectConnectionGithubConfig)
	PutGithubEnterpriseConfig(value *GoogleDeveloperConnectConnectionGithubEnterpriseConfig)
	PutGitlabConfig(value *GoogleDeveloperConnectConnectionGitlabConfig)
	PutGitlabEnterpriseConfig(value *GoogleDeveloperConnectConnectionGitlabEnterpriseConfig)
	PutTimeouts(value *GoogleDeveloperConnectConnectionTimeouts)
	ResetAnnotations()
	ResetBitbucketCloudConfig()
	ResetBitbucketDataCenterConfig()
	ResetCryptoKeyConfig()
	ResetDisabled()
	ResetEtag()
	ResetGithubConfig()
	ResetGithubEnterpriseConfig()
	ResetGitlabConfig()
	ResetGitlabEnterpriseConfig()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleDeveloperConnectConnection
type jsiiProxy_GoogleDeveloperConnectConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) BitbucketCloudConfig() GoogleDeveloperConnectConnectionBitbucketCloudConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionBitbucketCloudConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketCloudConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) BitbucketCloudConfigInput() *GoogleDeveloperConnectConnectionBitbucketCloudConfig {
	var returns *GoogleDeveloperConnectConnectionBitbucketCloudConfig
	_jsii_.Get(
		j,
		"bitbucketCloudConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) BitbucketDataCenterConfig() GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionBitbucketDataCenterConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) BitbucketDataCenterConfigInput() *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig {
	var returns *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) CryptoKeyConfig() GoogleDeveloperConnectConnectionCryptoKeyConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionCryptoKeyConfigOutputReference
	_jsii_.Get(
		j,
		"cryptoKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) CryptoKeyConfigInput() *GoogleDeveloperConnectConnectionCryptoKeyConfig {
	var returns *GoogleDeveloperConnectConnectionCryptoKeyConfig
	_jsii_.Get(
		j,
		"cryptoKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GithubConfig() GoogleDeveloperConnectConnectionGithubConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionGithubConfigOutputReference
	_jsii_.Get(
		j,
		"githubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GithubConfigInput() *GoogleDeveloperConnectConnectionGithubConfig {
	var returns *GoogleDeveloperConnectConnectionGithubConfig
	_jsii_.Get(
		j,
		"githubConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GithubEnterpriseConfig() GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionGithubEnterpriseConfigOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GithubEnterpriseConfigInput() *GoogleDeveloperConnectConnectionGithubEnterpriseConfig {
	var returns *GoogleDeveloperConnectConnectionGithubEnterpriseConfig
	_jsii_.Get(
		j,
		"githubEnterpriseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GitlabConfig() GoogleDeveloperConnectConnectionGitlabConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionGitlabConfigOutputReference
	_jsii_.Get(
		j,
		"gitlabConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GitlabConfigInput() *GoogleDeveloperConnectConnectionGitlabConfig {
	var returns *GoogleDeveloperConnectConnectionGitlabConfig
	_jsii_.Get(
		j,
		"gitlabConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GitlabEnterpriseConfig() GoogleDeveloperConnectConnectionGitlabEnterpriseConfigOutputReference {
	var returns GoogleDeveloperConnectConnectionGitlabEnterpriseConfigOutputReference
	_jsii_.Get(
		j,
		"gitlabEnterpriseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) GitlabEnterpriseConfigInput() *GoogleDeveloperConnectConnectionGitlabEnterpriseConfig {
	var returns *GoogleDeveloperConnectConnectionGitlabEnterpriseConfig
	_jsii_.Get(
		j,
		"gitlabEnterpriseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) InstallationState() GoogleDeveloperConnectConnectionInstallationStateList {
	var returns GoogleDeveloperConnectConnectionInstallationStateList
	_jsii_.Get(
		j,
		"installationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Timeouts() GoogleDeveloperConnectConnectionTimeoutsOutputReference {
	var returns GoogleDeveloperConnectConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_developer_connect_connection google_developer_connect_connection} Resource.
func NewGoogleDeveloperConnectConnection(scope constructs.Construct, id *string, config *GoogleDeveloperConnectConnectionConfig) GoogleDeveloperConnectConnection {
	_init_.Initialize()

	if err := validateNewGoogleDeveloperConnectConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDeveloperConnectConnection{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_developer_connect_connection google_developer_connect_connection} Resource.
func NewGoogleDeveloperConnectConnection_Override(g GoogleDeveloperConnectConnection, scope constructs.Construct, id *string, config *GoogleDeveloperConnectConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnection",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleDeveloperConnectConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleDeveloperConnectConnection resource upon running "cdktf plan <stack-name>".
func GoogleDeveloperConnectConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleDeveloperConnectConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnection",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func GoogleDeveloperConnectConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDeveloperConnectConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDeveloperConnectConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDeveloperConnectConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleDeveloperConnectConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleDeveloperConnectConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleDeveloperConnectConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleDeveloperConnectConnection.GoogleDeveloperConnectConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutBitbucketCloudConfig(value *GoogleDeveloperConnectConnectionBitbucketCloudConfig) {
	if err := g.validatePutBitbucketCloudConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBitbucketCloudConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutBitbucketDataCenterConfig(value *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig) {
	if err := g.validatePutBitbucketDataCenterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBitbucketDataCenterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutCryptoKeyConfig(value *GoogleDeveloperConnectConnectionCryptoKeyConfig) {
	if err := g.validatePutCryptoKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCryptoKeyConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutGithubConfig(value *GoogleDeveloperConnectConnectionGithubConfig) {
	if err := g.validatePutGithubConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGithubConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutGithubEnterpriseConfig(value *GoogleDeveloperConnectConnectionGithubEnterpriseConfig) {
	if err := g.validatePutGithubEnterpriseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGithubEnterpriseConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutGitlabConfig(value *GoogleDeveloperConnectConnectionGitlabConfig) {
	if err := g.validatePutGitlabConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGitlabConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutGitlabEnterpriseConfig(value *GoogleDeveloperConnectConnectionGitlabEnterpriseConfig) {
	if err := g.validatePutGitlabEnterpriseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGitlabEnterpriseConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) PutTimeouts(value *GoogleDeveloperConnectConnectionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetBitbucketCloudConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBitbucketCloudConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetBitbucketDataCenterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBitbucketDataCenterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetCryptoKeyConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCryptoKeyConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetEtag() {
	_jsii_.InvokeVoid(
		g,
		"resetEtag",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetGithubConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGithubConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetGithubEnterpriseConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGithubEnterpriseConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetGitlabConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGitlabConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetGitlabEnterpriseConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGitlabEnterpriseConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDeveloperConnectConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

