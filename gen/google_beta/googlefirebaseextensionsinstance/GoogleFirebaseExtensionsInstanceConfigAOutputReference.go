package googlefirebaseextensionsinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlefirebaseextensionsinstance/internal"
)

type GoogleFirebaseExtensionsInstanceConfigAOutputReference interface {
	cdktf.ComplexObject
	AllowedEventTypes() *[]*string
	SetAllowedEventTypes(val *[]*string)
	AllowedEventTypesInput() *[]*string
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
	CreateTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EventarcChannel() *string
	SetEventarcChannel(val *string)
	EventarcChannelInput() *string
	ExtensionRef() *string
	SetExtensionRef(val *string)
	ExtensionRefInput() *string
	ExtensionVersion() *string
	SetExtensionVersion(val *string)
	ExtensionVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleFirebaseExtensionsInstanceConfigA
	SetInternalValue(val *GoogleFirebaseExtensionsInstanceConfigA)
	Name() *string
	Params() *map[string]*string
	SetParams(val *map[string]*string)
	ParamsInput() *map[string]*string
	PopulatedPostinstallContent() *string
	SystemParams() *map[string]*string
	SetSystemParams(val *map[string]*string)
	SystemParamsInput() *map[string]*string
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
	ResetAllowedEventTypes()
	ResetEventarcChannel()
	ResetExtensionVersion()
	ResetSystemParams()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleFirebaseExtensionsInstanceConfigAOutputReference
type jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) AllowedEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) AllowedEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) EventarcChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) EventarcChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventarcChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ExtensionRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ExtensionRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ExtensionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ExtensionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) InternalValue() *GoogleFirebaseExtensionsInstanceConfigA {
	var returns *GoogleFirebaseExtensionsInstanceConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) Params() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) PopulatedPostinstallContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"populatedPostinstallContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) SystemParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"systemParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) SystemParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"systemParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleFirebaseExtensionsInstanceConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleFirebaseExtensionsInstanceConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleFirebaseExtensionsInstanceConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleFirebaseExtensionsInstance.GoogleFirebaseExtensionsInstanceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleFirebaseExtensionsInstanceConfigAOutputReference_Override(g GoogleFirebaseExtensionsInstanceConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleFirebaseExtensionsInstance.GoogleFirebaseExtensionsInstanceConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetAllowedEventTypes(val *[]*string) {
	if err := j.validateSetAllowedEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEventTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetEventarcChannel(val *string) {
	if err := j.validateSetEventarcChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventarcChannel",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetExtensionRef(val *string) {
	if err := j.validateSetExtensionRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionRef",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetExtensionVersion(val *string) {
	if err := j.validateSetExtensionVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetInternalValue(val *GoogleFirebaseExtensionsInstanceConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetParams(val *map[string]*string) {
	if err := j.validateSetParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"params",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetSystemParams(val *map[string]*string) {
	if err := j.validateSetSystemParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemParams",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ResetAllowedEventTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedEventTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ResetEventarcChannel() {
	_jsii_.InvokeVoid(
		g,
		"resetEventarcChannel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ResetExtensionVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetExtensionVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ResetSystemParams() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemParams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleFirebaseExtensionsInstanceConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

