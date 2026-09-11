package googleaccesscontextmanageraccesslevel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleaccesscontextmanageraccesslevel/internal"
)

type GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumVersion() *string
	SetMinimumVersion(val *string)
	MinimumVersionInput() *string
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	RequireVerifiedChromeOs() interface{}
	SetRequireVerifiedChromeOs(val interface{})
	RequireVerifiedChromeOsInput() interface{}
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
	ResetMinimumVersion()
	ResetRequireVerifiedChromeOs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference
type jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) MinimumVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) MinimumVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) RequireVerifiedChromeOs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireVerifiedChromeOs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) RequireVerifiedChromeOsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireVerifiedChromeOsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAccessContextManagerAccessLevel.GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference_Override(g GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleAccessContextManagerAccessLevel.GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetMinimumVersion(val *string) {
	if err := j.validateSetMinimumVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetRequireVerifiedChromeOs(val interface{}) {
	if err := j.validateSetRequireVerifiedChromeOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireVerifiedChromeOs",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) ResetMinimumVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetMinimumVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) ResetRequireVerifiedChromeOs() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireVerifiedChromeOs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelBasicConditionsDevicePolicyOsConstraintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

