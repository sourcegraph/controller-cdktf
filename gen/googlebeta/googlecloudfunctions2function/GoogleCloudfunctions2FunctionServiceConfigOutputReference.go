package googlecloudfunctions2function

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecloudfunctions2function/internal"
)

type GoogleCloudfunctions2FunctionServiceConfigOutputReference interface {
	cdktf.ComplexObject
	AllTrafficOnLatestRevision() interface{}
	SetAllTrafficOnLatestRevision(val interface{})
	AllTrafficOnLatestRevisionInput() interface{}
	AvailableMemory() *string
	SetAvailableMemory(val *string)
	AvailableMemoryInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	GcfUri() *string
	IngressSettings() *string
	SetIngressSettings(val *string)
	IngressSettingsInput() *string
	InternalValue() *GoogleCloudfunctions2FunctionServiceConfig
	SetInternalValue(val *GoogleCloudfunctions2FunctionServiceConfig)
	MaxInstanceCount() *float64
	SetMaxInstanceCount(val *float64)
	MaxInstanceCountInput() *float64
	MinInstanceCount() *float64
	SetMinInstanceCount(val *float64)
	MinInstanceCountInput() *float64
	SecretEnvironmentVariables() GoogleCloudfunctions2FunctionServiceConfigSecretEnvironmentVariablesList
	SecretEnvironmentVariablesInput() interface{}
	SecretVolumes() GoogleCloudfunctions2FunctionServiceConfigSecretVolumesList
	SecretVolumesInput() interface{}
	Service() *string
	SetService(val *string)
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	ServiceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
	Uri() *string
	VpcConnector() *string
	SetVpcConnector(val *string)
	VpcConnectorEgressSettings() *string
	SetVpcConnectorEgressSettings(val *string)
	VpcConnectorEgressSettingsInput() *string
	VpcConnectorInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutSecretEnvironmentVariables(value interface{})
	PutSecretVolumes(value interface{})
	ResetAllTrafficOnLatestRevision()
	ResetAvailableMemory()
	ResetEnvironmentVariables()
	ResetIngressSettings()
	ResetMaxInstanceCount()
	ResetMinInstanceCount()
	ResetSecretEnvironmentVariables()
	ResetSecretVolumes()
	ResetService()
	ResetServiceAccountEmail()
	ResetTimeoutSeconds()
	ResetVpcConnector()
	ResetVpcConnectorEgressSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudfunctions2FunctionServiceConfigOutputReference
type jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) AllTrafficOnLatestRevision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allTrafficOnLatestRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) AllTrafficOnLatestRevisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allTrafficOnLatestRevisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) AvailableMemory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) AvailableMemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GcfUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcfUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) IngressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) IngressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) InternalValue() *GoogleCloudfunctions2FunctionServiceConfig {
	var returns *GoogleCloudfunctions2FunctionServiceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) MaxInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) MaxInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) MinInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) MinInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) SecretEnvironmentVariables() GoogleCloudfunctions2FunctionServiceConfigSecretEnvironmentVariablesList {
	var returns GoogleCloudfunctions2FunctionServiceConfigSecretEnvironmentVariablesList
	_jsii_.Get(
		j,
		"secretEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) SecretEnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) SecretVolumes() GoogleCloudfunctions2FunctionServiceConfigSecretVolumesList {
	var returns GoogleCloudfunctions2FunctionServiceConfigSecretVolumesList
	_jsii_.Get(
		j,
		"secretVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) SecretVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) VpcConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) VpcConnectorEgressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) VpcConnectorEgressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) VpcConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorInput",
		&returns,
	)
	return returns
}


func NewGoogleCloudfunctions2FunctionServiceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudfunctions2FunctionServiceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudfunctions2FunctionServiceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionServiceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudfunctions2FunctionServiceConfigOutputReference_Override(g GoogleCloudfunctions2FunctionServiceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleCloudfunctions2Function.GoogleCloudfunctions2FunctionServiceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetAllTrafficOnLatestRevision(val interface{}) {
	if err := j.validateSetAllTrafficOnLatestRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allTrafficOnLatestRevision",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetAvailableMemory(val *string) {
	if err := j.validateSetAvailableMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableMemory",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetIngressSettings(val *string) {
	if err := j.validateSetIngressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressSettings",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetInternalValue(val *GoogleCloudfunctions2FunctionServiceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetMaxInstanceCount(val *float64) {
	if err := j.validateSetMaxInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceCount",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetMinInstanceCount(val *float64) {
	if err := j.validateSetMinInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstanceCount",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetVpcConnector(val *string) {
	if err := j.validateSetVpcConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnector",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference)SetVpcConnectorEgressSettings(val *string) {
	if err := j.validateSetVpcConnectorEgressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnectorEgressSettings",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) PutSecretEnvironmentVariables(value interface{}) {
	if err := g.validatePutSecretEnvironmentVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretEnvironmentVariables",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) PutSecretVolumes(value interface{}) {
	if err := g.validatePutSecretVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecretVolumes",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetAllTrafficOnLatestRevision() {
	_jsii_.InvokeVoid(
		g,
		"resetAllTrafficOnLatestRevision",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetAvailableMemory() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailableMemory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetIngressSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetMaxInstanceCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxInstanceCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetMinInstanceCount() {
	_jsii_.InvokeVoid(
		g,
		"resetMinInstanceCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetSecretEnvironmentVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretEnvironmentVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetSecretVolumes() {
	_jsii_.InvokeVoid(
		g,
		"resetSecretVolumes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		g,
		"resetService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetVpcConnector() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcConnector",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ResetVpcConnectorEgressSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcConnectorEgressSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudfunctions2FunctionServiceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

