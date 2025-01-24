package googlesccmanagementprojectsecurityhealthanalyticscustommodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlesccmanagementprojectsecurityhealthanalyticscustommodule/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_management_project_security_health_analytics_custom_module google_scc_management_project_security_health_analytics_custom_module}.
type GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule interface {
	cdktf.TerraformResource
	AncestorModule() *string
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
	CustomConfig() GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference
	CustomConfigInput() *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnablementState() *string
	SetEnablementState(val *string)
	EnablementStateInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastEditor() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsOutputReference
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
	PutCustomConfig(value *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig)
	PutTimeouts(value *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleTimeouts)
	ResetCustomConfig()
	ResetDisplayName()
	ResetEnablementState()
	ResetId()
	ResetLocation()
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

// The jsii proxy struct for GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule
type jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) AncestorModule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ancestorModule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) CustomConfig() GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference {
	var returns GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference
	_jsii_.Get(
		j,
		"customConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) CustomConfigInput() *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig {
	var returns *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig
	_jsii_.Get(
		j,
		"customConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) EnablementState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablementState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) EnablementStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablementStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) LastEditor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastEditor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) Timeouts() GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsOutputReference {
	var returns GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_management_project_security_health_analytics_custom_module google_scc_management_project_security_health_analytics_custom_module} Resource.
func NewGoogleSccManagementProjectSecurityHealthAnalyticsCustomModule(scope constructs.Construct, id *string, config *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleConfig) GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule {
	_init_.Initialize()

	if err := validateNewGoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSccManagementProjectSecurityHealthAnalyticsCustomModule.GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_scc_management_project_security_health_analytics_custom_module google_scc_management_project_security_health_analytics_custom_module} Resource.
func NewGoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_Override(g GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule, scope constructs.Construct, id *string, config *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSccManagementProjectSecurityHealthAnalyticsCustomModule.GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetEnablementState(val *string) {
	if err := j.validateSetEnablementStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablementState",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule resource upon running "cdktf plan <stack-name>".
func GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSccManagementProjectSecurityHealthAnalyticsCustomModule.GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule",
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
func GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSccManagementProjectSecurityHealthAnalyticsCustomModule.GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSccManagementProjectSecurityHealthAnalyticsCustomModule.GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleSccManagementProjectSecurityHealthAnalyticsCustomModule.GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleSccManagementProjectSecurityHealthAnalyticsCustomModule.GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) PutCustomConfig(value *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig) {
	if err := g.validatePutCustomConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) PutTimeouts(value *GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetCustomConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetEnablementState() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablementState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

