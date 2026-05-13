package healthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/healthcarefhirstore/internal"
)

type HealthcareFhirStoreValidationConfigOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *HealthcareFhirStoreValidationConfig
	SetInternalValue(val *HealthcareFhirStoreValidationConfig)
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
	ResetDisableFhirpathValidation()
	ResetDisableProfileValidation()
	ResetDisableReferenceTypeValidation()
	ResetDisableRequiredFieldValidation()
	ResetEnabledImplementationGuides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthcareFhirStoreValidationConfigOutputReference
type jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableFhirpathValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFhirpathValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableFhirpathValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableFhirpathValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableProfileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProfileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableProfileValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableProfileValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableReferenceTypeValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferenceTypeValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableReferenceTypeValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableReferenceTypeValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableRequiredFieldValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRequiredFieldValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) DisableRequiredFieldValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRequiredFieldValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) EnabledImplementationGuides() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledImplementationGuides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) EnabledImplementationGuidesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledImplementationGuidesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) InternalValue() *HealthcareFhirStoreValidationConfig {
	var returns *HealthcareFhirStoreValidationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthcareFhirStoreValidationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthcareFhirStoreValidationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewHealthcareFhirStoreValidationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStoreValidationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthcareFhirStoreValidationConfigOutputReference_Override(h HealthcareFhirStoreValidationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.healthcareFhirStore.HealthcareFhirStoreValidationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetDisableFhirpathValidation(val interface{}) {
	if err := j.validateSetDisableFhirpathValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableFhirpathValidation",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetDisableProfileValidation(val interface{}) {
	if err := j.validateSetDisableProfileValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableProfileValidation",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetDisableReferenceTypeValidation(val interface{}) {
	if err := j.validateSetDisableReferenceTypeValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableReferenceTypeValidation",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetDisableRequiredFieldValidation(val interface{}) {
	if err := j.validateSetDisableRequiredFieldValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRequiredFieldValidation",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetEnabledImplementationGuides(val *[]*string) {
	if err := j.validateSetEnabledImplementationGuidesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledImplementationGuides",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetInternalValue(val *HealthcareFhirStoreValidationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ResetDisableFhirpathValidation() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableFhirpathValidation",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ResetDisableProfileValidation() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableProfileValidation",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ResetDisableReferenceTypeValidation() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableReferenceTypeValidation",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ResetDisableRequiredFieldValidation() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableRequiredFieldValidation",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ResetEnabledImplementationGuides() {
	_jsii_.InvokeVoid(
		h,
		"resetEnabledImplementationGuides",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareFhirStoreValidationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

