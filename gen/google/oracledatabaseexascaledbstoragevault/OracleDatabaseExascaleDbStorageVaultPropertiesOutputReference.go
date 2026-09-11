package oracledatabaseexascaledbstoragevault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabaseexascaledbstoragevault/internal"
)

type OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference interface {
	cdktn.ComplexObject
	AdditionalFlashCachePercent() *float64
	SetAdditionalFlashCachePercent(val *float64)
	AdditionalFlashCachePercentInput() *float64
	AttachedShapeAttributes() *[]*string
	AvailableShapeAttributes() *[]*string
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
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeZone() OracleDatabaseExascaleDbStorageVaultPropertiesTimeZoneOutputReference
	TimeZoneInput() *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone
	VmClusterCount() *float64
	VmClusterIds() *[]*string
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
	PutExascaleDbStorageDetails(value *OracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails)
	PutTimeZone(value *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone)
	ResetAdditionalFlashCachePercent()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference
type jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
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

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
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


func NewOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseExascaleDbStorageVaultPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseExascaleDbStorageVault.OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference_Override(o OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.oracleDatabaseExascaleDbStorageVault.OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetAdditionalFlashCachePercent(val *float64) {
	if err := j.validateSetAdditionalFlashCachePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalFlashCachePercent",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetInternalValue(val *OracleDatabaseExascaleDbStorageVaultProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
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

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
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
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
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
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) PutTimeZone(value *OracleDatabaseExascaleDbStorageVaultPropertiesTimeZone) {
	if err := o.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeZone",
		[]interface{}{value},
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

func (o *jsiiProxy_OracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
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

