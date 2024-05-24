package googlemonitoringuptimecheckconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlemonitoringuptimecheckconfig/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_monitoring_uptime_check_config google_monitoring_uptime_check_config}.
type GoogleMonitoringUptimeCheckConfig interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CheckerType() *string
	SetCheckerType(val *string)
	CheckerTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentMatchers() GoogleMonitoringUptimeCheckConfigContentMatchersList
	ContentMatchersInput() interface{}
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpCheck() GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference
	HttpCheckInput() *GoogleMonitoringUptimeCheckConfigHttpCheck
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MonitoredResource() GoogleMonitoringUptimeCheckConfigMonitoredResourceOutputReference
	MonitoredResourceInput() *GoogleMonitoringUptimeCheckConfigMonitoredResource
	Name() *string
	// The tree node.
	Node() constructs.Node
	Period() *string
	SetPeriod(val *string)
	PeriodInput() *string
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
	ResourceGroup() GoogleMonitoringUptimeCheckConfigResourceGroupOutputReference
	ResourceGroupInput() *GoogleMonitoringUptimeCheckConfigResourceGroup
	SelectedRegions() *[]*string
	SetSelectedRegions(val *[]*string)
	SelectedRegionsInput() *[]*string
	SyntheticMonitor() GoogleMonitoringUptimeCheckConfigSyntheticMonitorOutputReference
	SyntheticMonitorInput() *GoogleMonitoringUptimeCheckConfigSyntheticMonitor
	TcpCheck() GoogleMonitoringUptimeCheckConfigTcpCheckOutputReference
	TcpCheckInput() *GoogleMonitoringUptimeCheckConfigTcpCheck
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	Timeouts() GoogleMonitoringUptimeCheckConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	UptimeCheckId() *string
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserLabelsInput() *map[string]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutContentMatchers(value interface{})
	PutHttpCheck(value *GoogleMonitoringUptimeCheckConfigHttpCheck)
	PutMonitoredResource(value *GoogleMonitoringUptimeCheckConfigMonitoredResource)
	PutResourceGroup(value *GoogleMonitoringUptimeCheckConfigResourceGroup)
	PutSyntheticMonitor(value *GoogleMonitoringUptimeCheckConfigSyntheticMonitor)
	PutTcpCheck(value *GoogleMonitoringUptimeCheckConfigTcpCheck)
	PutTimeouts(value *GoogleMonitoringUptimeCheckConfigTimeouts)
	ResetCheckerType()
	ResetContentMatchers()
	ResetHttpCheck()
	ResetId()
	ResetMonitoredResource()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriod()
	ResetProject()
	ResetResourceGroup()
	ResetSelectedRegions()
	ResetSyntheticMonitor()
	ResetTcpCheck()
	ResetTimeouts()
	ResetUserLabels()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleMonitoringUptimeCheckConfig
type jsiiProxy_GoogleMonitoringUptimeCheckConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) CheckerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) CheckerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checkerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ContentMatchers() GoogleMonitoringUptimeCheckConfigContentMatchersList {
	var returns GoogleMonitoringUptimeCheckConfigContentMatchersList
	_jsii_.Get(
		j,
		"contentMatchers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ContentMatchersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentMatchersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) HttpCheck() GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigHttpCheckOutputReference
	_jsii_.Get(
		j,
		"httpCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) HttpCheckInput() *GoogleMonitoringUptimeCheckConfigHttpCheck {
	var returns *GoogleMonitoringUptimeCheckConfigHttpCheck
	_jsii_.Get(
		j,
		"httpCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MonitoredResource() GoogleMonitoringUptimeCheckConfigMonitoredResourceOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigMonitoredResourceOutputReference
	_jsii_.Get(
		j,
		"monitoredResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) MonitoredResourceInput() *GoogleMonitoringUptimeCheckConfigMonitoredResource {
	var returns *GoogleMonitoringUptimeCheckConfigMonitoredResource
	_jsii_.Get(
		j,
		"monitoredResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResourceGroup() GoogleMonitoringUptimeCheckConfigResourceGroupOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigResourceGroupOutputReference
	_jsii_.Get(
		j,
		"resourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResourceGroupInput() *GoogleMonitoringUptimeCheckConfigResourceGroup {
	var returns *GoogleMonitoringUptimeCheckConfigResourceGroup
	_jsii_.Get(
		j,
		"resourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SelectedRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectedRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SelectedRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectedRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SyntheticMonitor() GoogleMonitoringUptimeCheckConfigSyntheticMonitorOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigSyntheticMonitorOutputReference
	_jsii_.Get(
		j,
		"syntheticMonitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SyntheticMonitorInput() *GoogleMonitoringUptimeCheckConfigSyntheticMonitor {
	var returns *GoogleMonitoringUptimeCheckConfigSyntheticMonitor
	_jsii_.Get(
		j,
		"syntheticMonitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TcpCheck() GoogleMonitoringUptimeCheckConfigTcpCheckOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigTcpCheckOutputReference
	_jsii_.Get(
		j,
		"tcpCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TcpCheckInput() *GoogleMonitoringUptimeCheckConfigTcpCheck {
	var returns *GoogleMonitoringUptimeCheckConfigTcpCheck
	_jsii_.Get(
		j,
		"tcpCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) Timeouts() GoogleMonitoringUptimeCheckConfigTimeoutsOutputReference {
	var returns GoogleMonitoringUptimeCheckConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) UptimeCheckId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uptimeCheckId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig) UserLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabelsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_monitoring_uptime_check_config google_monitoring_uptime_check_config} Resource.
func NewGoogleMonitoringUptimeCheckConfig(scope constructs.Construct, id *string, config *GoogleMonitoringUptimeCheckConfigConfig) GoogleMonitoringUptimeCheckConfig {
	_init_.Initialize()

	if err := validateNewGoogleMonitoringUptimeCheckConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleMonitoringUptimeCheckConfig{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_monitoring_uptime_check_config google_monitoring_uptime_check_config} Resource.
func NewGoogleMonitoringUptimeCheckConfig_Override(g GoogleMonitoringUptimeCheckConfig, scope constructs.Construct, id *string, config *GoogleMonitoringUptimeCheckConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetCheckerType(val *string) {
	if err := j.validateSetCheckerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkerType",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetSelectedRegions(val *[]*string) {
	if err := j.validateSetSelectedRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectedRegions",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GoogleMonitoringUptimeCheckConfig)SetUserLabels(val *map[string]*string) {
	if err := j.validateSetUserLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userLabels",
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
func GoogleMonitoringUptimeCheckConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringUptimeCheckConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMonitoringUptimeCheckConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringUptimeCheckConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleMonitoringUptimeCheckConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleMonitoringUptimeCheckConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleMonitoringUptimeCheckConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleMonitoringUptimeCheckConfig.GoogleMonitoringUptimeCheckConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutContentMatchers(value interface{}) {
	if err := g.validatePutContentMatchersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putContentMatchers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutHttpCheck(value *GoogleMonitoringUptimeCheckConfigHttpCheck) {
	if err := g.validatePutHttpCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutMonitoredResource(value *GoogleMonitoringUptimeCheckConfigMonitoredResource) {
	if err := g.validatePutMonitoredResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonitoredResource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutResourceGroup(value *GoogleMonitoringUptimeCheckConfigResourceGroup) {
	if err := g.validatePutResourceGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putResourceGroup",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutSyntheticMonitor(value *GoogleMonitoringUptimeCheckConfigSyntheticMonitor) {
	if err := g.validatePutSyntheticMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSyntheticMonitor",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutTcpCheck(value *GoogleMonitoringUptimeCheckConfigTcpCheck) {
	if err := g.validatePutTcpCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTcpCheck",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) PutTimeouts(value *GoogleMonitoringUptimeCheckConfigTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetCheckerType() {
	_jsii_.InvokeVoid(
		g,
		"resetCheckerType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetContentMatchers() {
	_jsii_.InvokeVoid(
		g,
		"resetContentMatchers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetHttpCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetMonitoredResource() {
	_jsii_.InvokeVoid(
		g,
		"resetMonitoredResource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetResourceGroup() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceGroup",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetSelectedRegions() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedRegions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetSyntheticMonitor() {
	_jsii_.InvokeVoid(
		g,
		"resetSyntheticMonitor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetTcpCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetTcpCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ResetUserLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetUserLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleMonitoringUptimeCheckConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

