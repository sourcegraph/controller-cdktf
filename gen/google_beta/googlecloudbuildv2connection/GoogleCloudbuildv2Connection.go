package googlecloudbuildv2connection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudbuildv2connection/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloudbuildv2_connection google_cloudbuildv2_connection}.
type GoogleCloudbuildv2Connection interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BitbucketCloudConfig() GoogleCloudbuildv2ConnectionBitbucketCloudConfigOutputReference
	BitbucketCloudConfigInput() *GoogleCloudbuildv2ConnectionBitbucketCloudConfig
	BitbucketDataCenterConfig() GoogleCloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference
	BitbucketDataCenterConfigInput() *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EffectiveAnnotations() cdktf.StringMap
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GithubConfig() GoogleCloudbuildv2ConnectionGithubConfigOutputReference
	GithubConfigInput() *GoogleCloudbuildv2ConnectionGithubConfig
	GithubEnterpriseConfig() GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference
	GithubEnterpriseConfigInput() *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig
	GitlabConfig() GoogleCloudbuildv2ConnectionGitlabConfigOutputReference
	GitlabConfigInput() *GoogleCloudbuildv2ConnectionGitlabConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstallationState() GoogleCloudbuildv2ConnectionInstallationStateList
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudbuildv2ConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutBitbucketCloudConfig(value *GoogleCloudbuildv2ConnectionBitbucketCloudConfig)
	PutBitbucketDataCenterConfig(value *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfig)
	PutGithubConfig(value *GoogleCloudbuildv2ConnectionGithubConfig)
	PutGithubEnterpriseConfig(value *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig)
	PutGitlabConfig(value *GoogleCloudbuildv2ConnectionGitlabConfig)
	PutTimeouts(value *GoogleCloudbuildv2ConnectionTimeouts)
	ResetAnnotations()
	ResetBitbucketCloudConfig()
	ResetBitbucketDataCenterConfig()
	ResetDisabled()
	ResetGithubConfig()
	ResetGithubEnterpriseConfig()
	ResetGitlabConfig()
	ResetId()
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

// The jsii proxy struct for GoogleCloudbuildv2Connection
type jsiiProxy_GoogleCloudbuildv2Connection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) BitbucketCloudConfig() GoogleCloudbuildv2ConnectionBitbucketCloudConfigOutputReference {
	var returns GoogleCloudbuildv2ConnectionBitbucketCloudConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketCloudConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) BitbucketCloudConfigInput() *GoogleCloudbuildv2ConnectionBitbucketCloudConfig {
	var returns *GoogleCloudbuildv2ConnectionBitbucketCloudConfig
	_jsii_.Get(
		j,
		"bitbucketCloudConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) BitbucketDataCenterConfig() GoogleCloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference {
	var returns GoogleCloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) BitbucketDataCenterConfigInput() *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfig {
	var returns *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfig
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) GithubConfig() GoogleCloudbuildv2ConnectionGithubConfigOutputReference {
	var returns GoogleCloudbuildv2ConnectionGithubConfigOutputReference
	_jsii_.Get(
		j,
		"githubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) GithubConfigInput() *GoogleCloudbuildv2ConnectionGithubConfig {
	var returns *GoogleCloudbuildv2ConnectionGithubConfig
	_jsii_.Get(
		j,
		"githubConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) GithubEnterpriseConfig() GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference {
	var returns GoogleCloudbuildv2ConnectionGithubEnterpriseConfigOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) GithubEnterpriseConfigInput() *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig {
	var returns *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig
	_jsii_.Get(
		j,
		"githubEnterpriseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) GitlabConfig() GoogleCloudbuildv2ConnectionGitlabConfigOutputReference {
	var returns GoogleCloudbuildv2ConnectionGitlabConfigOutputReference
	_jsii_.Get(
		j,
		"gitlabConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) GitlabConfigInput() *GoogleCloudbuildv2ConnectionGitlabConfig {
	var returns *GoogleCloudbuildv2ConnectionGitlabConfig
	_jsii_.Get(
		j,
		"gitlabConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) InstallationState() GoogleCloudbuildv2ConnectionInstallationStateList {
	var returns GoogleCloudbuildv2ConnectionInstallationStateList
	_jsii_.Get(
		j,
		"installationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) Timeouts() GoogleCloudbuildv2ConnectionTimeoutsOutputReference {
	var returns GoogleCloudbuildv2ConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloudbuildv2_connection google_cloudbuildv2_connection} Resource.
func NewGoogleCloudbuildv2Connection(scope constructs.Construct, id *string, config *GoogleCloudbuildv2ConnectionConfig) GoogleCloudbuildv2Connection {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildv2ConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildv2Connection{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2Connection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_cloudbuildv2_connection google_cloudbuildv2_connection} Resource.
func NewGoogleCloudbuildv2Connection_Override(g GoogleCloudbuildv2Connection, scope constructs.Construct, id *string, config *GoogleCloudbuildv2ConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2Connection",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildv2Connection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleCloudbuildv2Connection resource upon running "cdktf plan <stack-name>".
func GoogleCloudbuildv2Connection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCloudbuildv2Connection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2Connection",
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
func GoogleCloudbuildv2Connection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudbuildv2Connection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2Connection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudbuildv2Connection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudbuildv2Connection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2Connection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudbuildv2Connection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudbuildv2Connection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2Connection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCloudbuildv2Connection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleCloudbuildv2Connection.GoogleCloudbuildv2Connection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudbuildv2Connection) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) PutBitbucketCloudConfig(value *GoogleCloudbuildv2ConnectionBitbucketCloudConfig) {
	if err := g.validatePutBitbucketCloudConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBitbucketCloudConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) PutBitbucketDataCenterConfig(value *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfig) {
	if err := g.validatePutBitbucketDataCenterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBitbucketDataCenterConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) PutGithubConfig(value *GoogleCloudbuildv2ConnectionGithubConfig) {
	if err := g.validatePutGithubConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGithubConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) PutGithubEnterpriseConfig(value *GoogleCloudbuildv2ConnectionGithubEnterpriseConfig) {
	if err := g.validatePutGithubEnterpriseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGithubEnterpriseConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) PutGitlabConfig(value *GoogleCloudbuildv2ConnectionGitlabConfig) {
	if err := g.validatePutGitlabConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGitlabConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) PutTimeouts(value *GoogleCloudbuildv2ConnectionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetBitbucketCloudConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBitbucketCloudConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetBitbucketDataCenterConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBitbucketDataCenterConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetGithubConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGithubConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetGithubEnterpriseConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGithubEnterpriseConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetGitlabConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGitlabConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildv2Connection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

