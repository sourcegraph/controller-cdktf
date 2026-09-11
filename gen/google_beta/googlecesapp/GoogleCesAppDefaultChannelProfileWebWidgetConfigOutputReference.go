package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

type GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference interface {
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
	InternalValue() *GoogleCesAppDefaultChannelProfileWebWidgetConfig
	SetInternalValue(val *GoogleCesAppDefaultChannelProfileWebWidgetConfig)
	Modality() *string
	SetModality(val *string)
	ModalityInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Theme() *string
	SetTheme(val *string)
	ThemeInput() *string
	WebWidgetTitle() *string
	SetWebWidgetTitle(val *string)
	WebWidgetTitleInput() *string
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
	ResetModality()
	ResetTheme()
	ResetWebWidgetTitle()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference
type jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) InternalValue() *GoogleCesAppDefaultChannelProfileWebWidgetConfig {
	var returns *GoogleCesAppDefaultChannelProfileWebWidgetConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) Modality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ModalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) Theme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"theme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"themeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) WebWidgetTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webWidgetTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) WebWidgetTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webWidgetTitleInput",
		&returns,
	)
	return returns
}


func NewGoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference_Override(g GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetInternalValue(val *GoogleCesAppDefaultChannelProfileWebWidgetConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetModality(val *string) {
	if err := j.validateSetModalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modality",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetTheme(val *string) {
	if err := j.validateSetThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"theme",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference)SetWebWidgetTitle(val *string) {
	if err := j.validateSetWebWidgetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webWidgetTitle",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ResetModality() {
	_jsii_.InvokeVoid(
		g,
		"resetModality",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ResetTheme() {
	_jsii_.InvokeVoid(
		g,
		"resetTheme",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ResetWebWidgetTitle() {
	_jsii_.InvokeVoid(
		g,
		"resetWebWidgetTitle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppDefaultChannelProfileWebWidgetConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

