package vertexaiendpointwithmodelgardendeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/vertexaiendpointwithmodelgardendeployment/internal"
)

type VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference interface {
	cdktn.ComplexObject
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
	EnablePrivateServiceConnect() interface{}
	SetEnablePrivateServiceConnect(val interface{})
	EnablePrivateServiceConnectInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfig
	SetInternalValue(val *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfig)
	ProjectAllowlist() *[]*string
	SetProjectAllowlist(val *[]*string)
	ProjectAllowlistInput() *[]*string
	PscAutomationConfigs() VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigsOutputReference
	PscAutomationConfigsInput() *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs
	ServiceAttachment() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutPscAutomationConfigs(value *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs)
	ResetProjectAllowlist()
	ResetPscAutomationConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference
type jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) EnablePrivateServiceConnect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateServiceConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) EnablePrivateServiceConnectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateServiceConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) InternalValue() *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfig {
	var returns *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ProjectAllowlist() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectAllowlist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ProjectAllowlistInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projectAllowlistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) PscAutomationConfigs() VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigsOutputReference {
	var returns VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigsOutputReference
	_jsii_.Get(
		j,
		"pscAutomationConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) PscAutomationConfigsInput() *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs {
	var returns *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs
	_jsii_.Get(
		j,
		"pscAutomationConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ServiceAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference_Override(v VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.vertexAiEndpointWithModelGardenDeployment.VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference)SetEnablePrivateServiceConnect(val interface{}) {
	if err := j.validateSetEnablePrivateServiceConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateServiceConnect",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference)SetInternalValue(val *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference)SetProjectAllowlist(val *[]*string) {
	if err := j.validateSetProjectAllowlistParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectAllowlist",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) PutPscAutomationConfigs(value *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs) {
	if err := v.validatePutPscAutomationConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPscAutomationConfigs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ResetProjectAllowlist() {
	_jsii_.InvokeVoid(
		v,
		"resetProjectAllowlist",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ResetPscAutomationConfigs() {
	_jsii_.InvokeVoid(
		v,
		"resetPscAutomationConfigs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

