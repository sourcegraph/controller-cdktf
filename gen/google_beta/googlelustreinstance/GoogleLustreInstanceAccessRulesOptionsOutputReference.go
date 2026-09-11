package googlelustreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlelustreinstance/internal"
)

type GoogleLustreInstanceAccessRulesOptionsOutputReference interface {
	cdktn.ComplexObject
	AccessRules() GoogleLustreInstanceAccessRulesOptionsAccessRulesList
	AccessRulesInput() interface{}
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
	DefaultSquashGid() *float64
	SetDefaultSquashGid(val *float64)
	DefaultSquashGidInput() *float64
	DefaultSquashMode() *string
	SetDefaultSquashMode(val *string)
	DefaultSquashModeInput() *string
	DefaultSquashUid() *float64
	SetDefaultSquashUid(val *float64)
	DefaultSquashUidInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleLustreInstanceAccessRulesOptions
	SetInternalValue(val *GoogleLustreInstanceAccessRulesOptions)
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
	PutAccessRules(value interface{})
	ResetAccessRules()
	ResetDefaultSquashGid()
	ResetDefaultSquashUid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleLustreInstanceAccessRulesOptionsOutputReference
type jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) AccessRules() GoogleLustreInstanceAccessRulesOptionsAccessRulesList {
	var returns GoogleLustreInstanceAccessRulesOptionsAccessRulesList
	_jsii_.Get(
		j,
		"accessRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) AccessRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) DefaultSquashGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) DefaultSquashGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) DefaultSquashMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSquashMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) DefaultSquashModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultSquashModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) DefaultSquashUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) DefaultSquashUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultSquashUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) InternalValue() *GoogleLustreInstanceAccessRulesOptions {
	var returns *GoogleLustreInstanceAccessRulesOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleLustreInstanceAccessRulesOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleLustreInstanceAccessRulesOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleLustreInstanceAccessRulesOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleLustreInstance.GoogleLustreInstanceAccessRulesOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleLustreInstanceAccessRulesOptionsOutputReference_Override(g GoogleLustreInstanceAccessRulesOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleLustreInstance.GoogleLustreInstanceAccessRulesOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetDefaultSquashGid(val *float64) {
	if err := j.validateSetDefaultSquashGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashGid",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetDefaultSquashMode(val *string) {
	if err := j.validateSetDefaultSquashModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashMode",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetDefaultSquashUid(val *float64) {
	if err := j.validateSetDefaultSquashUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSquashUid",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetInternalValue(val *GoogleLustreInstanceAccessRulesOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) PutAccessRules(value interface{}) {
	if err := g.validatePutAccessRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) ResetAccessRules() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) ResetDefaultSquashGid() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultSquashGid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) ResetDefaultSquashUid() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultSquashUid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleLustreInstanceAccessRulesOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

