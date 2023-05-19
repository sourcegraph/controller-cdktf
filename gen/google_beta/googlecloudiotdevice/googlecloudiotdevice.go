package googlecloudiotdevice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecloudiotdevice/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_device google_cloudiot_device}.
type GoogleCloudiotDevice interface {
	cdktf.TerraformResource
	Blocked() interface{}
	SetBlocked(val interface{})
	BlockedInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Config() GoogleCloudiotDeviceConfigAList
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
	Credentials() GoogleCloudiotDeviceCredentialsList
	CredentialsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayConfig() GoogleCloudiotDeviceGatewayConfigOutputReference
	GatewayConfigInput() *GoogleCloudiotDeviceGatewayConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastConfigAckTime() *string
	LastConfigSendTime() *string
	LastErrorStatus() GoogleCloudiotDeviceLastErrorStatusList
	LastErrorTime() *string
	LastEventTime() *string
	LastHeartbeatTime() *string
	LastStateTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NumId() *string
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
	Registry() *string
	SetRegistry(val *string)
	RegistryInput() *string
	State() GoogleCloudiotDeviceStateList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudiotDeviceTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutCredentials(value interface{})
	PutGatewayConfig(value *GoogleCloudiotDeviceGatewayConfig)
	PutTimeouts(value *GoogleCloudiotDeviceTimeouts)
	ResetBlocked()
	ResetCredentials()
	ResetGatewayConfig()
	ResetId()
	ResetLogLevel()
	ResetMetadata()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleCloudiotDevice
type jsiiProxy_GoogleCloudiotDevice struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudiotDevice) Blocked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) BlockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Config() GoogleCloudiotDeviceConfigAList {
	var returns GoogleCloudiotDeviceConfigAList
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Credentials() GoogleCloudiotDeviceCredentialsList {
	var returns GoogleCloudiotDeviceCredentialsList
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) CredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) GatewayConfig() GoogleCloudiotDeviceGatewayConfigOutputReference {
	var returns GoogleCloudiotDeviceGatewayConfigOutputReference
	_jsii_.Get(
		j,
		"gatewayConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) GatewayConfigInput() *GoogleCloudiotDeviceGatewayConfig {
	var returns *GoogleCloudiotDeviceGatewayConfig
	_jsii_.Get(
		j,
		"gatewayConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LastConfigAckTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastConfigAckTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LastConfigSendTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastConfigSendTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LastErrorStatus() GoogleCloudiotDeviceLastErrorStatusList {
	var returns GoogleCloudiotDeviceLastErrorStatusList
	_jsii_.Get(
		j,
		"lastErrorStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LastErrorTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastErrorTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LastEventTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastEventTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LastHeartbeatTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastHeartbeatTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LastStateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastStateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) NumId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"numId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Registry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) RegistryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) State() GoogleCloudiotDeviceStateList {
	var returns GoogleCloudiotDeviceStateList
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) Timeouts() GoogleCloudiotDeviceTimeoutsOutputReference {
	var returns GoogleCloudiotDeviceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudiotDevice) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_device google_cloudiot_device} Resource.
func NewGoogleCloudiotDevice(scope constructs.Construct, id *string, config *GoogleCloudiotDeviceConfig) GoogleCloudiotDevice {
	_init_.Initialize()

	if err := validateNewGoogleCloudiotDeviceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudiotDevice{}

	_jsii_.Create(
		"google-beta.googleCloudiotDevice.GoogleCloudiotDevice",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_cloudiot_device google_cloudiot_device} Resource.
func NewGoogleCloudiotDevice_Override(g GoogleCloudiotDevice, scope constructs.Construct, id *string, config *GoogleCloudiotDeviceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleCloudiotDevice.GoogleCloudiotDevice",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetBlocked(val interface{}) {
	if err := j.validateSetBlockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blocked",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudiotDevice)SetRegistry(val *string) {
	if err := j.validateSetRegistryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registry",
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
func GoogleCloudiotDevice_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudiotDevice_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleCloudiotDevice.GoogleCloudiotDevice",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudiotDevice_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudiotDevice_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleCloudiotDevice.GoogleCloudiotDevice",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudiotDevice_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudiotDevice_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"google-beta.googleCloudiotDevice.GoogleCloudiotDevice",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCloudiotDevice_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"google-beta.googleCloudiotDevice.GoogleCloudiotDevice",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudiotDevice) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudiotDevice) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudiotDevice) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudiotDevice) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) PutCredentials(value interface{}) {
	if err := g.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCredentials",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) PutGatewayConfig(value *GoogleCloudiotDeviceGatewayConfig) {
	if err := g.validatePutGatewayConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGatewayConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) PutTimeouts(value *GoogleCloudiotDeviceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetBlocked() {
	_jsii_.InvokeVoid(
		g,
		"resetBlocked",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetGatewayConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGatewayConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetLogLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudiotDevice) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudiotDevice) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudiotDevice) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudiotDevice) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

