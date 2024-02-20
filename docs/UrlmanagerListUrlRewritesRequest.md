# # UrlmanagerListUrlRewritesRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** | Required.  | [optional]
**Filter**| [**ListUrlRewritesRequestFilter**](ListUrlRewritesRequestFilter.md) |   | [optional]
**PageSize**| **int64** | The maximum number of url rewrites to return. The service may return fewer than this value. If unspecified, at most 10 url rewrites will be returned. The maximum value is 200; values above 200 will be coerced to 200.  | [optional]
**PageToken**| **string** | A page token, received from a previous &#x60;ListUrlRewrites&#x60; call. Provide this to retrieve the subsequent page.   When paginating, all other parameters provided to &#x60;ListUrlRewrites&#x60; must match the call that provided the page token.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

