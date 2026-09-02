package googleoracledatabaseexascaledbstoragevault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabaseexascaledbstoragevault/internal"
)

type GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference interface {
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
	ExascaleDbStorageDetails() GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetailsOutputReference
	ExascaleDbStorageDetailsInput() *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleOracleDatabaseExascaleDbStorageVaultProperties
	SetInternalValue(val *GoogleOracleDatabaseExascaleDbStorageVaultProperties)
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
	TimeZone() GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZoneOutputReference
	TimeZoneInput() *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZone
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
	PutExascaleDbStorageDetails(value *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails)
	PutTimeZone(value *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZone)
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

// The jsii proxy struct for GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference
type jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AdditionalFlashCachePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalFlashCachePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AdditionalFlashCachePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalFlashCachePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AttachedShapeAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attachedShapeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) AvailableShapeAttributes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableShapeAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ExascaleDbStorageDetails() GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetailsOutputReference {
	var returns GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetailsOutputReference
	_jsii_.Get(
		j,
		"exascaleDbStorageDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ExascaleDbStorageDetailsInput() *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails {
	var returns *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails
	_jsii_.Get(
		j,
		"exascaleDbStorageDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InternalValue() *GoogleOracleDatabaseExascaleDbStorageVaultProperties {
	var returns *GoogleOracleDatabaseExascaleDbStorageVaultProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) OciUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TimeZone() GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZoneOutputReference {
	var returns GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) TimeZoneInput() *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZone {
	var returns *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) VmClusterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmClusterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) VmClusterIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vmClusterIds",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseExascaleDbStorageVault.GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference_Override(g GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseExascaleDbStorageVault.GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetAdditionalFlashCachePercent(val *float64) {
	if err := j.validateSetAdditionalFlashCachePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalFlashCachePercent",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetInternalValue(val *GoogleOracleDatabaseExascaleDbStorageVaultProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) PutExascaleDbStorageDetails(value *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesExascaleDbStorageDetails) {
	if err := g.validatePutExascaleDbStorageDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExascaleDbStorageDetails",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) PutTimeZone(value *GoogleOracleDatabaseExascaleDbStorageVaultPropertiesTimeZone) {
	if err := g.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ResetAdditionalFlashCachePercent() {
	_jsii_.InvokeVoid(
		g,
		"resetAdditionalFlashCachePercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseExascaleDbStorageVaultPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

