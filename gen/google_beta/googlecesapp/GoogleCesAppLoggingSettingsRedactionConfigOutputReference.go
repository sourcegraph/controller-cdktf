package googlecesapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecesapp/internal"
)

type GoogleCesAppLoggingSettingsRedactionConfigOutputReference interface {
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
	DeidentifyTemplate() *string
	SetDeidentifyTemplate(val *string)
	DeidentifyTemplateInput() *string
	EnableRedaction() interface{}
	SetEnableRedaction(val interface{})
	EnableRedactionInput() interface{}
	// Experimental.
	Fqn() *string
	InspectTemplate() *string
	SetInspectTemplate(val *string)
	InspectTemplateInput() *string
	InternalValue() *GoogleCesAppLoggingSettingsRedactionConfig
	SetInternalValue(val *GoogleCesAppLoggingSettingsRedactionConfig)
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
	ResetDeidentifyTemplate()
	ResetEnableRedaction()
	ResetInspectTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCesAppLoggingSettingsRedactionConfigOutputReference
type jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) DeidentifyTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) DeidentifyTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deidentifyTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) EnableRedaction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRedaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) EnableRedactionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRedactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) InspectTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) InspectTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) InternalValue() *GoogleCesAppLoggingSettingsRedactionConfig {
	var returns *GoogleCesAppLoggingSettingsRedactionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleCesAppLoggingSettingsRedactionConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleCesAppLoggingSettingsRedactionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCesAppLoggingSettingsRedactionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppLoggingSettingsRedactionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCesAppLoggingSettingsRedactionConfigOutputReference_Override(g GoogleCesAppLoggingSettingsRedactionConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleCesApp.GoogleCesAppLoggingSettingsRedactionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetDeidentifyTemplate(val *string) {
	if err := j.validateSetDeidentifyTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deidentifyTemplate",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetEnableRedaction(val interface{}) {
	if err := j.validateSetEnableRedactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRedaction",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetInspectTemplate(val *string) {
	if err := j.validateSetInspectTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectTemplate",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetInternalValue(val *GoogleCesAppLoggingSettingsRedactionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) ResetDeidentifyTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetDeidentifyTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) ResetEnableRedaction() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableRedaction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) ResetInspectTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetInspectTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleCesAppLoggingSettingsRedactionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

