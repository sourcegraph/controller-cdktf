package oracledatabaseexascaledbstoragevault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabaseexascaledbstoragevault/internal"
)

type OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference interface {
	cdktf.ComplexObject
	AdditionalFlashCachePercent() *float64
	SetAdditionalFlashCachePercent(val *float64)
	AdditionalFlashCachePercentInput() *float64
	AttachedShapeAttributes() *[]*string
	AvailableShapeAttributes() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	ExascaleDbStorageDetails() OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetailsOutputReference
	ExascaleDbStorageDetailsInput() *OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails
	// Experimental.
	Fqn() *string
	InternalValue() *OracleDatabaseExascaleDbStorageVaultProperties
	SetInternalValue(val *OracleDatabaseExascaleDbStorageVaultProperties)
	Ocid() *string
	OciUri() *string
	State() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() OracleDatabaseExascaleDbStorageVaultPropertiesTimeZoneOutputReference
	TimeZoneInput() *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone
	VmClusterCount() *float64
	VmClusterIds() *[]*string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutExascaleDbStorageDetails(value *OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails)
	PutTimeZone(value *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone)
	ResetAdditionalFlashCachePercent()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference
type jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AdditionalFlashCachePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalFlashCachePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AdditionalFlashCachePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalFlashCachePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AttachedShapeAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attachedShapeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AvailableShapeAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableShapeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ExascaleDbStorageDetails() OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetailsOutputReference {
	var returns OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetailsOutputReference
	_jsii_.Get(
		j,
		"exascaleDbStorageDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ExascaleDbStorageDetailsInput() *OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails {
	var returns *OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails
	_jsii_.Get(
		j,
		"exascaleDbStorageDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InternalValue() *OracleDatabaseExascaleDbStorageVaultProperties {
	var returns *OracleDatabaseExascaleDbStorageVaultProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) OciUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TimeZone() OracleDatabaseExascaleDbStorageVaultPropertiesTimeZoneOutputReference {
	var returns OracleDatabaseExascaleDbStorageVaultPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TimeZoneInput() *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone {
	var returns *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) VmClusterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmClusterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) VmClusterIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmClusterIds",
		&returns,
	)
	return returns
}

func NewOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseExascaleDbStorageVaultPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseExascaleDbStorageVault.OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference_Override(o OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseExascaleDbStorageVault.OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference",
		[]any{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) SetAdditionalFlashCachePercent(val *float64) {
	if err := j.validateSetAdditionalFlashCachePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalFlashCachePercent",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) SetInternalValue(val *OracleDatabaseExascaleDbStorageVaultProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) PutExascaleDbStorageDetails(value *OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails) {
	if err := o.validatePutExascaleDbStorageDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putExascaleDbStorageDetails",
		[]any{value},
	)
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) PutTimeZone(value *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone) {
	if err := o.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeZone",
		[]any{value},
	)
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ResetAdditionalFlashCachePercent() {
	_jsii_.InvokeVoid(
		o,
		"resetAdditionalFlashCachePercent",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		o,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
