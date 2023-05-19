package datacloudflarerulesets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/cloudflare/datacloudflarerulesets/internal"
)

type DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference interface {
	cdktf.ComplexObject
	AutomaticHttpsRewrites() cdktf.IResolvable
	Autominify() DataCloudflareRulesetsRulesetsRulesActionParametersAutominifyList
	Bic() cdktf.IResolvable
	BrowserTtl() DataCloudflareRulesetsRulesetsRulesActionParametersBrowserTtlList
	Cache() cdktf.IResolvable
	CacheKey() DataCloudflareRulesetsRulesetsRulesActionParametersCacheKeyList
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
	Content() *string
	ContentType() *string
	CookieFields() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisableApps() cdktf.IResolvable
	DisableRailgun() cdktf.IResolvable
	DisableZaraz() cdktf.IResolvable
	EdgeTtl() DataCloudflareRulesetsRulesetsRulesActionParametersEdgeTtlList
	EmailObfuscation() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	FromList() DataCloudflareRulesetsRulesetsRulesActionParametersFromListList
	FromValue() DataCloudflareRulesetsRulesetsRulesActionParametersFromValueList
	Headers() DataCloudflareRulesetsRulesetsRulesActionParametersHeadersList
	HostHeader() *string
	HotlinkProtection() cdktf.IResolvable
	Id() *string
	Increment() *float64
	InternalValue() *DataCloudflareRulesetsRulesetsRulesActionParameters
	SetInternalValue(val *DataCloudflareRulesetsRulesetsRulesActionParameters)
	MatchedData() DataCloudflareRulesetsRulesetsRulesActionParametersMatchedDataList
	Mirage() cdktf.IResolvable
	OpportunisticEncryption() cdktf.IResolvable
	Origin() DataCloudflareRulesetsRulesetsRulesActionParametersOriginList
	OriginErrorPagePassthru() cdktf.IResolvable
	Overrides() DataCloudflareRulesetsRulesetsRulesActionParametersOverridesList
	Phases() *[]*string
	Polish() *string
	Products() *[]*string
	RequestFields() *[]*string
	RespectStrongEtags() cdktf.IResolvable
	Response() DataCloudflareRulesetsRulesetsRulesActionParametersResponseList
	ResponseFields() *[]*string
	RocketLoader() cdktf.IResolvable
	Rules() cdktf.StringMap
	Ruleset() *string
	Rulesets() *[]*string
	SecurityLevel() *string
	ServerSideExcludes() cdktf.IResolvable
	ServeStale() DataCloudflareRulesetsRulesetsRulesActionParametersServeStaleList
	Sni() DataCloudflareRulesetsRulesetsRulesActionParametersSniList
	Ssl() *string
	StatusCode() *float64
	Sxg() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() DataCloudflareRulesetsRulesetsRulesActionParametersUriList
	Version() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference
type jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) AutomaticHttpsRewrites() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"automaticHttpsRewrites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Autominify() DataCloudflareRulesetsRulesetsRulesActionParametersAutominifyList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersAutominifyList
	_jsii_.Get(
		j,
		"autominify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Bic() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"bic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) BrowserTtl() DataCloudflareRulesetsRulesetsRulesActionParametersBrowserTtlList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersBrowserTtlList
	_jsii_.Get(
		j,
		"browserTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Cache() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) CacheKey() DataCloudflareRulesetsRulesetsRulesActionParametersCacheKeyList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersCacheKeyList
	_jsii_.Get(
		j,
		"cacheKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) CookieFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cookieFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) DisableApps() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) DisableRailgun() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableRailgun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) DisableZaraz() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disableZaraz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) EdgeTtl() DataCloudflareRulesetsRulesetsRulesActionParametersEdgeTtlList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersEdgeTtlList
	_jsii_.Get(
		j,
		"edgeTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) EmailObfuscation() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"emailObfuscation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) FromList() DataCloudflareRulesetsRulesetsRulesActionParametersFromListList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersFromListList
	_jsii_.Get(
		j,
		"fromList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) FromValue() DataCloudflareRulesetsRulesetsRulesActionParametersFromValueList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersFromValueList
	_jsii_.Get(
		j,
		"fromValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Headers() DataCloudflareRulesetsRulesetsRulesActionParametersHeadersList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) HostHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) HotlinkProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hotlinkProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Increment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"increment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) InternalValue() *DataCloudflareRulesetsRulesetsRulesActionParameters {
	var returns *DataCloudflareRulesetsRulesetsRulesActionParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) MatchedData() DataCloudflareRulesetsRulesetsRulesActionParametersMatchedDataList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersMatchedDataList
	_jsii_.Get(
		j,
		"matchedData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Mirage() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mirage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) OpportunisticEncryption() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"opportunisticEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Origin() DataCloudflareRulesetsRulesetsRulesActionParametersOriginList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersOriginList
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) OriginErrorPagePassthru() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"originErrorPagePassthru",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Overrides() DataCloudflareRulesetsRulesetsRulesActionParametersOverridesList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersOverridesList
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Phases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Polish() *string {
	var returns *string
	_jsii_.Get(
		j,
		"polish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Products() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"products",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) RequestFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) RespectStrongEtags() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"respectStrongEtags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Response() DataCloudflareRulesetsRulesetsRulesActionParametersResponseList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersResponseList
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ResponseFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) RocketLoader() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rocketLoader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Rules() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Ruleset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Rulesets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rulesets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ServerSideExcludes() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"serverSideExcludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ServeStale() DataCloudflareRulesetsRulesetsRulesActionParametersServeStaleList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersServeStaleList
	_jsii_.Get(
		j,
		"serveStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Sni() DataCloudflareRulesetsRulesetsRulesActionParametersSniList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersSniList
	_jsii_.Get(
		j,
		"sni",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Ssl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) StatusCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Sxg() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sxg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Uri() DataCloudflareRulesetsRulesetsRulesActionParametersUriList {
	var returns DataCloudflareRulesetsRulesetsRulesActionParametersUriList
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewDataCloudflareRulesetsRulesetsRulesActionParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareRulesetsRulesetsRulesActionParametersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference{}

	_jsii_.Create(
		"cloudflare.dataCloudflareRulesets.DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataCloudflareRulesetsRulesetsRulesActionParametersOutputReference_Override(d DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cloudflare.dataCloudflareRulesets.DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference)SetInternalValue(val *DataCloudflareRulesetsRulesetsRulesActionParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareRulesetsRulesetsRulesActionParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

