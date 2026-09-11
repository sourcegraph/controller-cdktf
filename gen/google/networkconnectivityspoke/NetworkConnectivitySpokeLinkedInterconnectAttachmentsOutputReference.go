package networkconnectivityspoke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google/networkconnectivityspoke/internal"
)

type NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference interface {
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
	ExcludeExportRanges() *[]*string
	SetExcludeExportRanges(val *[]*string)
	ExcludeExportRangesInput() *[]*string
	ExcludeImportRanges() *[]*string
	SetExcludeImportRanges(val *[]*string)
	ExcludeImportRangesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeExportRanges() *[]*string
	SetIncludeExportRanges(val *[]*string)
	IncludeExportRangesInput() *[]*string
	IncludeImportRanges() *[]*string
	SetIncludeImportRanges(val *[]*string)
	IncludeImportRangesInput() *[]*string
	InternalValue() *NetworkConnectivitySpokeLinkedInterconnectAttachments
	SetInternalValue(val *NetworkConnectivitySpokeLinkedInterconnectAttachments)
	SiteToSiteDataTransfer() interface{}
	SetSiteToSiteDataTransfer(val interface{})
	SiteToSiteDataTransferInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Uris() *[]*string
	SetUris(val *[]*string)
	UrisInput() *[]*string
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
	ResetExcludeExportRanges()
	ResetExcludeImportRanges()
	ResetIncludeExportRanges()
	ResetIncludeImportRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference
type jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ExcludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ExcludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ExcludeImportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeImportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ExcludeImportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeImportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) IncludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) IncludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) IncludeImportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeImportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) IncludeImportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeImportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) InternalValue() *NetworkConnectivitySpokeLinkedInterconnectAttachments {
	var returns *NetworkConnectivitySpokeLinkedInterconnectAttachments
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) SiteToSiteDataTransfer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"siteToSiteDataTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) SiteToSiteDataTransferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"siteToSiteDataTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) Uris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) UrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urisInput",
		&returns,
	)
	return returns
}


func NewNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google.networkConnectivitySpoke.NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference_Override(n NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google.networkConnectivitySpoke.NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetExcludeExportRanges(val *[]*string) {
	if err := j.validateSetExcludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeExportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetExcludeImportRanges(val *[]*string) {
	if err := j.validateSetExcludeImportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeImportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetIncludeExportRanges(val *[]*string) {
	if err := j.validateSetIncludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetIncludeImportRanges(val *[]*string) {
	if err := j.validateSetIncludeImportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeImportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetInternalValue(val *NetworkConnectivitySpokeLinkedInterconnectAttachments) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetSiteToSiteDataTransfer(val interface{}) {
	if err := j.validateSetSiteToSiteDataTransferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteToSiteDataTransfer",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetUris(val *[]*string) {
	if err := j.validateSetUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uris",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ResetExcludeExportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeExportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ResetExcludeImportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeImportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ResetIncludeExportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludeExportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ResetIncludeImportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludeImportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

