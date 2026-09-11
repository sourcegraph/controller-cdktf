package googlediscoveryenginedataconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlediscoveryenginedataconnector/internal"
)

type GoogleDiscoveryEngineDataConnectorActionConfigOutputReference interface {
	cdktn.ComplexObject
	ActionParams() *map[string]*string
	SetActionParams(val *map[string]*string)
	ActionParamsInput() *map[string]*string
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
	CreateBapConnection() interface{}
	SetCreateBapConnection(val interface{})
	CreateBapConnectionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDiscoveryEngineDataConnectorActionConfig
	SetInternalValue(val *GoogleDiscoveryEngineDataConnectorActionConfig)
	IsActionConfigured() cdktn.IResolvable
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
	ResetActionParams()
	ResetCreateBapConnection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDiscoveryEngineDataConnectorActionConfigOutputReference
type jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ActionParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"actionParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ActionParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"actionParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) CreateBapConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createBapConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) CreateBapConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createBapConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) InternalValue() *GoogleDiscoveryEngineDataConnectorActionConfig {
	var returns *GoogleDiscoveryEngineDataConnectorActionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) IsActionConfigured() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isActionConfigured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDiscoveryEngineDataConnectorActionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDiscoveryEngineDataConnectorActionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDiscoveryEngineDataConnectorActionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnectorActionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDiscoveryEngineDataConnectorActionConfigOutputReference_Override(g GoogleDiscoveryEngineDataConnectorActionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDiscoveryEngineDataConnector.GoogleDiscoveryEngineDataConnectorActionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference)SetActionParams(val *map[string]*string) {
	if err := j.validateSetActionParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionParams",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference)SetCreateBapConnection(val interface{}) {
	if err := j.validateSetCreateBapConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createBapConnection",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference)SetInternalValue(val *GoogleDiscoveryEngineDataConnectorActionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ResetActionParams() {
	_jsii_.InvokeVoid(
		g,
		"resetActionParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ResetCreateBapConnection() {
	_jsii_.InvokeVoid(
		g,
		"resetCreateBapConnection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDiscoveryEngineDataConnectorActionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

