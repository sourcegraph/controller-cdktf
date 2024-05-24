package googlenetappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlenetappvolume/internal"
)

type GoogleNetappVolumeExportPolicyRulesOutputReference interface {
	cdktf.ComplexObject
	AccessType() *string
	SetAccessType(val *string)
	AccessTypeInput() *string
	AllowedClients() *string
	SetAllowedClients(val *string)
	AllowedClientsInput() *string
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
	HasRootAccess() *string
	SetHasRootAccess(val *string)
	HasRootAccessInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Kerberos5IReadOnly() interface{}
	SetKerberos5IReadOnly(val interface{})
	Kerberos5IReadOnlyInput() interface{}
	Kerberos5IReadWrite() interface{}
	SetKerberos5IReadWrite(val interface{})
	Kerberos5IReadWriteInput() interface{}
	Kerberos5PReadOnly() interface{}
	SetKerberos5PReadOnly(val interface{})
	Kerberos5PReadOnlyInput() interface{}
	Kerberos5PReadWrite() interface{}
	SetKerberos5PReadWrite(val interface{})
	Kerberos5PReadWriteInput() interface{}
	Kerberos5ReadOnly() interface{}
	SetKerberos5ReadOnly(val interface{})
	Kerberos5ReadOnlyInput() interface{}
	Kerberos5ReadWrite() interface{}
	SetKerberos5ReadWrite(val interface{})
	Kerberos5ReadWriteInput() interface{}
	Nfsv3() interface{}
	SetNfsv3(val interface{})
	Nfsv3Input() interface{}
	Nfsv4() interface{}
	SetNfsv4(val interface{})
	Nfsv4Input() interface{}
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
	ResetAccessType()
	ResetAllowedClients()
	ResetHasRootAccess()
	ResetKerberos5IReadOnly()
	ResetKerberos5IReadWrite()
	ResetKerberos5PReadOnly()
	ResetKerberos5PReadWrite()
	ResetKerberos5ReadOnly()
	ResetKerberos5ReadWrite()
	ResetNfsv3()
	ResetNfsv4()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetappVolumeExportPolicyRulesOutputReference
type jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) AccessTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) AllowedClients() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) AllowedClientsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedClientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) HasRootAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hasRootAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) HasRootAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hasRootAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5IReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5IReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5PReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5PReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Kerberos5ReadWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberos5ReadWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Nfsv3() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Nfsv3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Nfsv4() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Nfsv4Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nfsv4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetappVolumeExportPolicyRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleNetappVolumeExportPolicyRulesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetappVolumeExportPolicyRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeExportPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleNetappVolumeExportPolicyRulesOutputReference_Override(g GoogleNetappVolumeExportPolicyRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeExportPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetAccessType(val *string) {
	if err := j.validateSetAccessTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessType",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetAllowedClients(val *string) {
	if err := j.validateSetAllowedClientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedClients",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetHasRootAccess(val *string) {
	if err := j.validateSetHasRootAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasRootAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetKerberos5IReadOnly(val interface{}) {
	if err := j.validateSetKerberos5IReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5IReadOnly",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetKerberos5IReadWrite(val interface{}) {
	if err := j.validateSetKerberos5IReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5IReadWrite",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetKerberos5PReadOnly(val interface{}) {
	if err := j.validateSetKerberos5PReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5PReadOnly",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetKerberos5PReadWrite(val interface{}) {
	if err := j.validateSetKerberos5PReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5PReadWrite",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetKerberos5ReadOnly(val interface{}) {
	if err := j.validateSetKerberos5ReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5ReadOnly",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetKerberos5ReadWrite(val interface{}) {
	if err := j.validateSetKerberos5ReadWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberos5ReadWrite",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetNfsv3(val interface{}) {
	if err := j.validateSetNfsv3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv3",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetNfsv4(val interface{}) {
	if err := j.validateSetNfsv4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nfsv4",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetAccessType() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetAllowedClients() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedClients",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetHasRootAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetHasRootAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetKerberos5IReadOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberos5IReadOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetKerberos5IReadWrite() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberos5IReadWrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetKerberos5PReadOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberos5PReadOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetKerberos5PReadWrite() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberos5PReadWrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetKerberos5ReadOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberos5ReadOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetKerberos5ReadWrite() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberos5ReadWrite",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetNfsv3() {
	_jsii_.InvokeVoid(
		g,
		"resetNfsv3",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ResetNfsv4() {
	_jsii_.InvokeVoid(
		g,
		"resetNfsv4",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetappVolumeExportPolicyRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

