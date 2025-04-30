package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference interface {
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
	InternalValue() *GoogleStorageTransferJobReplicationSpecObjectConditions
	SetInternalValue(val *GoogleStorageTransferJobReplicationSpecObjectConditions)
	LastModifiedBefore() *string
	SetLastModifiedBefore(val *string)
	LastModifiedBeforeInput() *string
	LastModifiedSince() *string
	SetLastModifiedSince(val *string)
	LastModifiedSinceInput() *string
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
	ResetLastModifiedBefore()
	ResetLastModifiedSince()
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

// The jsii proxy struct for GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference
type jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ExcludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ExcludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) IncludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) IncludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) InternalValue() *GoogleStorageTransferJobReplicationSpecObjectConditions {
	var returns *GoogleStorageTransferJobReplicationSpecObjectConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobReplicationSpecObjectConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference_Override(g GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageTransferJob.GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetExcludePrefixes(val *[]*string) {
	if err := j.validateSetExcludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePrefixes",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetIncludePrefixes(val *[]*string) {
	if err := j.validateSetIncludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrefixes",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetInternalValue(val *GoogleStorageTransferJobReplicationSpecObjectConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetLastModifiedBefore(val *string) {
	if err := j.validateSetLastModifiedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedBefore",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetLastModifiedSince(val *string) {
	if err := j.validateSetLastModifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedSince",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetMaxTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMaxTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetMinTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMinTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetExcludePrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludePrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetIncludePrefixes() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludePrefixes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetLastModifiedBefore() {
	_jsii_.InvokeVoid(
		g,
		"resetLastModifiedBefore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetLastModifiedSince() {
	_jsii_.InvokeVoid(
		g,
		"resetLastModifiedSince",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetMaxTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetMinTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		g,
		"resetMinTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecObjectConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

