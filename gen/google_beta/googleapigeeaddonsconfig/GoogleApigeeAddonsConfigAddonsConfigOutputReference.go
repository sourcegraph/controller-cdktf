package googleapigeeaddonsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleapigeeaddonsconfig/internal"
)

type GoogleApigeeAddonsConfigAddonsConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedApiOpsConfig() GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigOutputReference
	AdvancedApiOpsConfigInput() *GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig
	ApiSecurityConfig() GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfigOutputReference
	ApiSecurityConfigInput() *GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfig
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
	ConnectorsPlatformConfig() GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigOutputReference
	ConnectorsPlatformConfigInput() *GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IntegrationConfig() GoogleApigeeAddonsConfigAddonsConfigIntegrationConfigOutputReference
	IntegrationConfigInput() *GoogleApigeeAddonsConfigAddonsConfigIntegrationConfig
	InternalValue() *GoogleApigeeAddonsConfigAddonsConfig
	SetInternalValue(val *GoogleApigeeAddonsConfigAddonsConfig)
	MonetizationConfig() GoogleApigeeAddonsConfigAddonsConfigMonetizationConfigOutputReference
	MonetizationConfigInput() *GoogleApigeeAddonsConfigAddonsConfigMonetizationConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutAdvancedApiOpsConfig(value *GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig)
	PutApiSecurityConfig(value *GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfig)
	PutConnectorsPlatformConfig(value *GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig)
	PutIntegrationConfig(value *GoogleApigeeAddonsConfigAddonsConfigIntegrationConfig)
	PutMonetizationConfig(value *GoogleApigeeAddonsConfigAddonsConfigMonetizationConfig)
	ResetAdvancedApiOpsConfig()
	ResetApiSecurityConfig()
	ResetConnectorsPlatformConfig()
	ResetIntegrationConfig()
	ResetMonetizationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleApigeeAddonsConfigAddonsConfigOutputReference
type jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) AdvancedApiOpsConfig() GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigOutputReference {
	var returns GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigOutputReference
	_jsii_.Get(
		j,
		"advancedApiOpsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) AdvancedApiOpsConfigInput() *GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig {
	var returns *GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig
	_jsii_.Get(
		j,
		"advancedApiOpsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ApiSecurityConfig() GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfigOutputReference {
	var returns GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"apiSecurityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ApiSecurityConfigInput() *GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfig {
	var returns *GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfig
	_jsii_.Get(
		j,
		"apiSecurityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ConnectorsPlatformConfig() GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigOutputReference {
	var returns GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigOutputReference
	_jsii_.Get(
		j,
		"connectorsPlatformConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ConnectorsPlatformConfigInput() *GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig {
	var returns *GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig
	_jsii_.Get(
		j,
		"connectorsPlatformConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) IntegrationConfig() GoogleApigeeAddonsConfigAddonsConfigIntegrationConfigOutputReference {
	var returns GoogleApigeeAddonsConfigAddonsConfigIntegrationConfigOutputReference
	_jsii_.Get(
		j,
		"integrationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) IntegrationConfigInput() *GoogleApigeeAddonsConfigAddonsConfigIntegrationConfig {
	var returns *GoogleApigeeAddonsConfigAddonsConfigIntegrationConfig
	_jsii_.Get(
		j,
		"integrationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) InternalValue() *GoogleApigeeAddonsConfigAddonsConfig {
	var returns *GoogleApigeeAddonsConfigAddonsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) MonetizationConfig() GoogleApigeeAddonsConfigAddonsConfigMonetizationConfigOutputReference {
	var returns GoogleApigeeAddonsConfigAddonsConfigMonetizationConfigOutputReference
	_jsii_.Get(
		j,
		"monetizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) MonetizationConfigInput() *GoogleApigeeAddonsConfigAddonsConfigMonetizationConfig {
	var returns *GoogleApigeeAddonsConfigAddonsConfigMonetizationConfig
	_jsii_.Get(
		j,
		"monetizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleApigeeAddonsConfigAddonsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleApigeeAddonsConfigAddonsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleApigeeAddonsConfigAddonsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeAddonsConfig.GoogleApigeeAddonsConfigAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleApigeeAddonsConfigAddonsConfigOutputReference_Override(g GoogleApigeeAddonsConfigAddonsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleApigeeAddonsConfig.GoogleApigeeAddonsConfigAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference)SetInternalValue(val *GoogleApigeeAddonsConfigAddonsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) PutAdvancedApiOpsConfig(value *GoogleApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig) {
	if err := g.validatePutAdvancedApiOpsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdvancedApiOpsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) PutApiSecurityConfig(value *GoogleApigeeAddonsConfigAddonsConfigApiSecurityConfig) {
	if err := g.validatePutApiSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiSecurityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) PutConnectorsPlatformConfig(value *GoogleApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig) {
	if err := g.validatePutConnectorsPlatformConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConnectorsPlatformConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) PutIntegrationConfig(value *GoogleApigeeAddonsConfigAddonsConfigIntegrationConfig) {
	if err := g.validatePutIntegrationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIntegrationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) PutMonetizationConfig(value *GoogleApigeeAddonsConfigAddonsConfigMonetizationConfig) {
	if err := g.validatePutMonetizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMonetizationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ResetAdvancedApiOpsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedApiOpsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ResetApiSecurityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetApiSecurityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ResetConnectorsPlatformConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectorsPlatformConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ResetIntegrationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIntegrationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ResetMonetizationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMonetizationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleApigeeAddonsConfigAddonsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

