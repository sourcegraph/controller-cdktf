package googleintegrationsclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleintegrationsclient/internal"
)

type GoogleIntegrationsClientCloudKmsConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleIntegrationsClientCloudKmsConfig
	SetInternalValue(val *GoogleIntegrationsClientCloudKmsConfig)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	KeyVersion() *string
	SetKeyVersion(val *string)
	KeyVersionInput() *string
	KmsLocation() *string
	SetKmsLocation(val *string)
	KmsLocationInput() *string
	KmsProjectId() *string
	SetKmsProjectId(val *string)
	KmsProjectIdInput() *string
	KmsRing() *string
	SetKmsRing(val *string)
	KmsRingInput() *string
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
	ResetKeyVersion()
	ResetKmsProjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleIntegrationsClientCloudKmsConfigOutputReference
type jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) InternalValue() *GoogleIntegrationsClientCloudKmsConfig {
	var returns *GoogleIntegrationsClientCloudKmsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KeyVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KeyVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KmsLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KmsLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KmsProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KmsProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KmsRing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsRing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) KmsRingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsRingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleIntegrationsClientCloudKmsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleIntegrationsClientCloudKmsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleIntegrationsClientCloudKmsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationsClient.GoogleIntegrationsClientCloudKmsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleIntegrationsClientCloudKmsConfigOutputReference_Override(g GoogleIntegrationsClientCloudKmsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleIntegrationsClient.GoogleIntegrationsClientCloudKmsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetInternalValue(val *GoogleIntegrationsClientCloudKmsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetKeyVersion(val *string) {
	if err := j.validateSetKeyVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetKmsLocation(val *string) {
	if err := j.validateSetKmsLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsLocation",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetKmsProjectId(val *string) {
	if err := j.validateSetKmsProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsProjectId",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetKmsRing(val *string) {
	if err := j.validateSetKmsRingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsRing",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) ResetKeyVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetKeyVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) ResetKmsProjectId() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsProjectId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleIntegrationsClientCloudKmsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

