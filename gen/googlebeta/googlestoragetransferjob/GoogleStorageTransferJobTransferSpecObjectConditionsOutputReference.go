package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference interface {
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
	ExcludePrefixes() *[]*string
	SetExcludePrefixes(val *[]*string)
	ExcludePrefixesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludePrefixes() *[]*string
	SetIncludePrefixes(val *[]*string)
	IncludePrefixesInput() *[]*string
	InternalValue() *GoogleStorageTransferJobTransferSpecObjectConditions
	SetInternalValue(val *GoogleStorageTransferJobTransferSpecObjectConditions)
	MaxTimeElapsedSinceLastModification() *string
	SetMaxTimeElapsedSinceLastModification(val *string)
	MaxTimeElapsedSinceLastModificationInput() *string
	MinTimeElapsedSinceLastModification() *string
	SetMinTimeElapsedSinceLastModification(val *string)
	MinTimeElapsedSinceLastModificationInput() *string
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
	ResetExcludePrefixes()
	ResetIncludePrefixes()
	ResetMaxTimeElapsedSinceLastModification()
	ResetMinTimeElapsedSinceLastModification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference
type jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ExcludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ExcludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) IncludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) IncludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) InternalValue() *GoogleStorageTransferJobTransferSpecObjectConditions {
	var returns *GoogleStorageTransferJobTransferSpecObjectConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobTransferSpecObjectConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobTransferSpecObjectConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobTransferSpecObjectConditionsOutputReference_Override(g GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetExcludePrefixes(val *[]*string) {
	if err := j.validateSetExcludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePrefixes",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetIncludePrefixes(val *[]*string) {
	if err := j.validateSetIncludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrefixes",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetInternalValue(val *GoogleStorageTransferJobTransferSpecObjectConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetMaxTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMaxTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetMinTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMinTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ResetExcludePrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludePrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ResetIncludePrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludePrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ResetMaxTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ResetMinTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		g,
		"resetMinTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecObjectConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

