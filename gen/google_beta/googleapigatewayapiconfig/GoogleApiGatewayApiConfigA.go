package googleapigatewayapiconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigatewayapiconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config google_api_gateway_api_config}.
type GoogleApiGatewayApiConfigA interface {
	cdktf.TerraformResource
	Api() *string
	SetApi(val *string)
	ApiConfigId() *string
	SetApiConfigId(val *string)
	ApiConfigIdInput() *string
	ApiConfigIdPrefix() *string
	SetApiConfigIdPrefix(val *string)
	ApiConfigIdPrefixInput() *string
	ApiInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayConfig() GoogleApiGatewayApiConfigGatewayConfigOutputReference
	GatewayConfigInput() *GoogleApiGatewayApiConfigGatewayConfig
	GrpcServices() GoogleApiGatewayApiConfigGrpcServicesList
	GrpcServicesInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManagedServiceConfigs() GoogleApiGatewayApiConfigManagedServiceConfigsList
	ManagedServiceConfigsInput() interface{}
	Name() *string
	// The tree node.
	Node() constructs.Node
	OpenapiDocuments() GoogleApiGatewayApiConfigOpenapiDocumentsList
	OpenapiDocumentsInput() interface{}
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
	ServiceConfigId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleApiGatewayApiConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutGatewayConfig(value *GoogleApiGatewayApiConfigGatewayConfig)
	PutGrpcServices(value interface{})
	PutManagedServiceConfigs(value interface{})
	PutOpenapiDocuments(value interface{})
	PutTimeouts(value *GoogleApiGatewayApiConfigTimeouts)
	ResetApiConfigId()
	ResetApiConfigIdPrefix()
	ResetDisplayName()
	ResetGatewayConfig()
	ResetGrpcServices()
	ResetId()
	ResetLabels()
	ResetManagedServiceConfigs()
	ResetOpenapiDocuments()
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

// The jsii proxy struct for GoogleApiGatewayApiConfigA
type jsiiProxy_GoogleApiGatewayApiConfigA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Api() *string {
	var returns *string
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ApiConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ApiConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ApiConfigIdPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiConfigIdPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ApiConfigIdPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiConfigIdPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ApiInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) GatewayConfig() GoogleApiGatewayApiConfigGatewayConfigOutputReference {
	var returns GoogleApiGatewayApiConfigGatewayConfigOutputReference
	_jsii_.Get(
		j,
		"gatewayConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) GatewayConfigInput() *GoogleApiGatewayApiConfigGatewayConfig {
	var returns *GoogleApiGatewayApiConfigGatewayConfig
	_jsii_.Get(
		j,
		"gatewayConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) GrpcServices() GoogleApiGatewayApiConfigGrpcServicesList {
	var returns GoogleApiGatewayApiConfigGrpcServicesList
	_jsii_.Get(
		j,
		"grpcServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) GrpcServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grpcServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ManagedServiceConfigs() GoogleApiGatewayApiConfigManagedServiceConfigsList {
	var returns GoogleApiGatewayApiConfigManagedServiceConfigsList
	_jsii_.Get(
		j,
		"managedServiceConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ManagedServiceConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedServiceConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) OpenapiDocuments() GoogleApiGatewayApiConfigOpenapiDocumentsList {
	var returns GoogleApiGatewayApiConfigOpenapiDocumentsList
	_jsii_.Get(
		j,
		"openapiDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) OpenapiDocumentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openapiDocumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) ServiceConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) Timeouts() GoogleApiGatewayApiConfigTimeoutsOutputReference {
	var returns GoogleApiGatewayApiConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config google_api_gateway_api_config} Resource.
func NewGoogleApiGatewayApiConfigA(scope constructs.Construct, id *string, config *GoogleApiGatewayApiConfigAConfig) GoogleApiGatewayApiConfigA {
	_init_.Initialize()

	if err := validateNewGoogleApiGatewayApiConfigAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApiGatewayApiConfigA{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApiGatewayApiConfig.GoogleApiGatewayApiConfigA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_api_gateway_api_config google_api_gateway_api_config} Resource.
func NewGoogleApiGatewayApiConfigA_Override(g GoogleApiGatewayApiConfigA, scope constructs.Construct, id *string, config *GoogleApiGatewayApiConfigAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApiGatewayApiConfig.GoogleApiGatewayApiConfigA",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetApi(val *string) {
	if err := j.validateSetApiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetApiConfigId(val *string) {
	if err := j.validateSetApiConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiConfigId",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetApiConfigIdPrefix(val *string) {
	if err := j.validateSetApiConfigIdPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiConfigIdPrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleApiGatewayApiConfigA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleApiGatewayApiConfigA resource upon running "cdktf plan <stack-name>".
func GoogleApiGatewayApiConfigA_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleApiGatewayApiConfigA_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApiGatewayApiConfig.GoogleApiGatewayApiConfigA",
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
func GoogleApiGatewayApiConfigA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApiGatewayApiConfigA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApiGatewayApiConfig.GoogleApiGatewayApiConfigA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApiGatewayApiConfigA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApiGatewayApiConfigA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApiGatewayApiConfig.GoogleApiGatewayApiConfigA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleApiGatewayApiConfigA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleApiGatewayApiConfigA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleApiGatewayApiConfig.GoogleApiGatewayApiConfigA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleApiGatewayApiConfigA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleApiGatewayApiConfig.GoogleApiGatewayApiConfigA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) PutGatewayConfig(value *GoogleApiGatewayApiConfigGatewayConfig) {
	if err := g.validatePutGatewayConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGatewayConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) PutGrpcServices(value interface{}) {
	if err := g.validatePutGrpcServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGrpcServices",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) PutManagedServiceConfigs(value interface{}) {
	if err := g.validatePutManagedServiceConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManagedServiceConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) PutOpenapiDocuments(value interface{}) {
	if err := g.validatePutOpenapiDocumentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOpenapiDocuments",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) PutTimeouts(value *GoogleApiGatewayApiConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetApiConfigId() {
	_jsii_.InvokeVoid(
		g,
		"resetApiConfigId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetApiConfigIdPrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetApiConfigIdPrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetGatewayConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGatewayConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetGrpcServices() {
	_jsii_.InvokeVoid(
		g,
		"resetGrpcServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetManagedServiceConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetManagedServiceConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetOpenapiDocuments() {
	_jsii_.InvokeVoid(
		g,
		"resetOpenapiDocuments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApiGatewayApiConfigA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

