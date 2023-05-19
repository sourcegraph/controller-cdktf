package googlednsrecordset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlednsrecordset/internal"
)

type GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference interface {
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
	HealthCheckedTargets() GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference
	HealthCheckedTargetsInput() *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Rrdatas() *[]*string
	SetRrdatas(val *[]*string)
	RrdatasInput() *[]*string
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
	PutHealthCheckedTargets(value *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets)
	ResetHealthCheckedTargets()
	ResetRrdatas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference
type jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) HealthCheckedTargets() GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference {
	var returns GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargetsOutputReference
	_jsii_.Get(
		j,
		"healthCheckedTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) HealthCheckedTargetsInput() *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets {
	var returns *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets
	_jsii_.Get(
		j,
		"healthCheckedTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Rrdatas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rrdatas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) RrdatasInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rrdatasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference{}

	_jsii_.Create(
		"google-beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference_Override(g GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleDnsRecordSet.GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetRrdatas(val *[]*string) {
	if err := j.validateSetRrdatasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rrdatas",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) PutHealthCheckedTargets(value *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets) {
	if err := g.validatePutHealthCheckedTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHealthCheckedTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ResetHealthCheckedTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthCheckedTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ResetRrdatas() {
	_jsii_.InvokeVoid(
		g,
		"resetRrdatas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

