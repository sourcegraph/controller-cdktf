package googlehealthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehealthcarefhirstore/internal"
)

type GoogleHealthcareFhirStoreValidationConfigOutputReference interface {
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
	DisableFhirpathValidation() interface{}
	SetDisableFhirpathValidation(val interface{})
	DisableFhirpathValidationInput() interface{}
	DisableProfileValidation() interface{}
	SetDisableProfileValidation(val interface{})
	DisableProfileValidationInput() interface{}
	DisableReferenceTypeValidation() interface{}
	SetDisableReferenceTypeValidation(val interface{})
	DisableReferenceTypeValidationInput() interface{}
	DisableRequiredFieldValidation() interface{}
	SetDisableRequiredFieldValidation(val interface{})
	DisableRequiredFieldValidationInput() interface{}
	EnabledImplementationGuides() *[]*string
	SetEnabledImplementationGuides(val *[]*string)
	EnabledImplementationGuidesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleHealthcareFhirStoreValidationConfig
	SetInternalValue(val *GoogleHealthcareFhirStoreValidationConfig)
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
	ResetDisableFhirpathValidation()
	ResetDisableProfileValidation()
	ResetDisableReferenceTypeValidation()
	ResetDisableRequiredFieldValidation()
	ResetEnabledImplementationGuides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHealthcareFhirStoreValidationConfigOutputReference
type jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableFhirpathValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFhirpathValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableFhirpathValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFhirpathValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableProfileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProfileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableProfileValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProfileValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableReferenceTypeValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferenceTypeValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableReferenceTypeValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferenceTypeValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableRequiredFieldValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRequiredFieldValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) DisableRequiredFieldValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRequiredFieldValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) EnabledImplementationGuides() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledImplementationGuides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) EnabledImplementationGuidesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledImplementationGuidesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) InternalValue() *GoogleHealthcareFhirStoreValidationConfig {
	var returns *GoogleHealthcareFhirStoreValidationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHealthcareFhirStoreValidationConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleHealthcareFhirStoreValidationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHealthcareFhirStoreValidationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStoreValidationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHealthcareFhirStoreValidationConfigOutputReference_Override(g GoogleHealthcareFhirStoreValidationConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHealthcareFhirStore.GoogleHealthcareFhirStoreValidationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetDisableFhirpathValidation(val interface{}) {
	if err := j.validateSetDisableFhirpathValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableFhirpathValidation",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetDisableProfileValidation(val interface{}) {
	if err := j.validateSetDisableProfileValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableProfileValidation",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetDisableReferenceTypeValidation(val interface{}) {
	if err := j.validateSetDisableReferenceTypeValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableReferenceTypeValidation",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetDisableRequiredFieldValidation(val interface{}) {
	if err := j.validateSetDisableRequiredFieldValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRequiredFieldValidation",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetEnabledImplementationGuides(val *[]*string) {
	if err := j.validateSetEnabledImplementationGuidesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledImplementationGuides",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetInternalValue(val *GoogleHealthcareFhirStoreValidationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ResetDisableFhirpathValidation() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableFhirpathValidation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ResetDisableProfileValidation() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableProfileValidation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ResetDisableReferenceTypeValidation() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableReferenceTypeValidation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ResetDisableRequiredFieldValidation() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableRequiredFieldValidation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ResetEnabledImplementationGuides() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabledImplementationGuides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreValidationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

