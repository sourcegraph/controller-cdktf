package googledatalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googledatalosspreventiondiscoveryconfig/internal"
)

type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference interface {
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
	InspectTemplateModifiedCadence() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadenceOutputReference
	InspectTemplateModifiedCadenceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadence
	InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence
	SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence)
	RefreshFrequency() *string
	SetRefreshFrequency(val *string)
	RefreshFrequencyInput() *string
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
	PutInspectTemplateModifiedCadence(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadence)
	ResetInspectTemplateModifiedCadence()
	ResetRefreshFrequency()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference
type jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) InspectTemplateModifiedCadence() GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadenceOutputReference {
	var returns GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadenceOutputReference
	_jsii_.Get(
		j,
		"inspectTemplateModifiedCadence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) InspectTemplateModifiedCadenceInput() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadence {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadence
	_jsii_.Get(
		j,
		"inspectTemplateModifiedCadenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) InternalValue() *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence {
	var returns *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) RefreshFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) RefreshFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference_Override(g GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleDataLossPreventionDiscoveryConfig.GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference)SetInternalValue(val *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference)SetRefreshFrequency(val *string) {
	if err := j.validateSetRefreshFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshFrequency",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) PutInspectTemplateModifiedCadence(value *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadence) {
	if err := g.validatePutInspectTemplateModifiedCadenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInspectTemplateModifiedCadence",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) ResetInspectTemplateModifiedCadence() {
	_jsii_.InvokeVoid(
		g,
		"resetInspectTemplateModifiedCadence",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) ResetRefreshFrequency() {
	_jsii_.InvokeVoid(
		g,
		"resetRefreshFrequency",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

