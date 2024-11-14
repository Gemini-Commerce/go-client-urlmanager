# \BasicOperationsAPI

All URIs are relative to *https://urlmanager.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UrlManagerChangeUrlRewriteRequestPath**](BasicOperationsAPI.md#UrlManagerChangeUrlRewriteRequestPath) | **Post** /urlmanager/change_url_rewrite_request_path | Change Url Rewrite Request Path
[**UrlManagerChangeUrlRewriteRequestPath2**](BasicOperationsAPI.md#UrlManagerChangeUrlRewriteRequestPath2) | **Post** /urlmanager.UrlManager/ChangeUrlRewriteRequestPath | Change Url Rewrite Request Path
[**UrlManagerCreateUrlRewrite**](BasicOperationsAPI.md#UrlManagerCreateUrlRewrite) | **Post** /urlmanager/create_url_rewrite | Create Url Rewrite
[**UrlManagerCreateUrlRewrite2**](BasicOperationsAPI.md#UrlManagerCreateUrlRewrite2) | **Post** /urlmanager.UrlManager/CreateUrlRewrite | Create Url Rewrite
[**UrlManagerDeleteUrlRewrite**](BasicOperationsAPI.md#UrlManagerDeleteUrlRewrite) | **Post** /urlmanager/delete_url_rewrite | Delete Url Rewrite
[**UrlManagerDeleteUrlRewrite2**](BasicOperationsAPI.md#UrlManagerDeleteUrlRewrite2) | **Post** /urlmanager.UrlManager/DeleteUrlRewrite | Delete Url Rewrite
[**UrlManagerGetUrlRewrite**](BasicOperationsAPI.md#UrlManagerGetUrlRewrite) | **Post** /urlmanager/get_url_rewrite | Get Url Rewrite
[**UrlManagerGetUrlRewrite2**](BasicOperationsAPI.md#UrlManagerGetUrlRewrite2) | **Post** /urlmanager.UrlManager/GetUrlRewrite | Get Url Rewrite
[**UrlManagerListUrlRewrites**](BasicOperationsAPI.md#UrlManagerListUrlRewrites) | **Post** /urlmanager/list_url_rewrites | List Url Rewrites
[**UrlManagerListUrlRewrites2**](BasicOperationsAPI.md#UrlManagerListUrlRewrites2) | **Post** /urlmanager.UrlManager/ListUrlRewrites | List Url Rewrites
[**UrlManagerListUrlRewritesByTargetPaths**](BasicOperationsAPI.md#UrlManagerListUrlRewritesByTargetPaths) | **Post** /urlmanager/list_url_rewrites_by_target_paths | List Url Rewrites By Target Paths
[**UrlManagerListUrlRewritesByTargetPaths2**](BasicOperationsAPI.md#UrlManagerListUrlRewritesByTargetPaths2) | **Post** /urlmanager.UrlManager/ListUrlRewritesByTargetPaths | List Url Rewrites By Target Paths
[**UrlManagerResolveUrlRewrite**](BasicOperationsAPI.md#UrlManagerResolveUrlRewrite) | **Post** /urlmanager/resolve_url_rewrite | Resolve Url Rewrite
[**UrlManagerResolveUrlRewrite2**](BasicOperationsAPI.md#UrlManagerResolveUrlRewrite2) | **Post** /urlmanager.UrlManager/ResolveUrlRewrite | Resolve Url Rewrite



## UrlManagerChangeUrlRewriteRequestPath

> map[string]interface{} UrlManagerChangeUrlRewriteRequestPath(ctx).Body(body).Execute()

Change Url Rewrite Request Path



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerChangeUrlRewriteRequestPathRequest() // UrlmanagerChangeUrlRewriteRequestPathRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerChangeUrlRewriteRequestPath(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerChangeUrlRewriteRequestPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerChangeUrlRewriteRequestPath`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerChangeUrlRewriteRequestPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerChangeUrlRewriteRequestPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerChangeUrlRewriteRequestPathRequest**](UrlmanagerChangeUrlRewriteRequestPathRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerChangeUrlRewriteRequestPath2

> map[string]interface{} UrlManagerChangeUrlRewriteRequestPath2(ctx).Body(body).Execute()

Change Url Rewrite Request Path



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerChangeUrlRewriteRequestPathRequest() // UrlmanagerChangeUrlRewriteRequestPathRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerChangeUrlRewriteRequestPath2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerChangeUrlRewriteRequestPath2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerChangeUrlRewriteRequestPath2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerChangeUrlRewriteRequestPath2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerChangeUrlRewriteRequestPath2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerChangeUrlRewriteRequestPathRequest**](UrlmanagerChangeUrlRewriteRequestPathRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerCreateUrlRewrite

> UrlmanagerUrlRewrite UrlManagerCreateUrlRewrite(ctx).Body(body).Execute()

Create Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerCreateUrlRewriteRequest() // UrlmanagerCreateUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerCreateUrlRewrite(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerCreateUrlRewrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerCreateUrlRewrite`: UrlmanagerUrlRewrite
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerCreateUrlRewrite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerCreateUrlRewriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerCreateUrlRewriteRequest**](UrlmanagerCreateUrlRewriteRequest.md) |  | 

### Return type

[**UrlmanagerUrlRewrite**](UrlmanagerUrlRewrite.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerCreateUrlRewrite2

> UrlmanagerUrlRewrite UrlManagerCreateUrlRewrite2(ctx).Body(body).Execute()

Create Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerCreateUrlRewriteRequest() // UrlmanagerCreateUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerCreateUrlRewrite2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerCreateUrlRewrite2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerCreateUrlRewrite2`: UrlmanagerUrlRewrite
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerCreateUrlRewrite2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerCreateUrlRewrite2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerCreateUrlRewriteRequest**](UrlmanagerCreateUrlRewriteRequest.md) |  | 

### Return type

[**UrlmanagerUrlRewrite**](UrlmanagerUrlRewrite.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerDeleteUrlRewrite

> map[string]interface{} UrlManagerDeleteUrlRewrite(ctx).Body(body).Execute()

Delete Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerDeleteUrlRewriteRequest() // UrlmanagerDeleteUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerDeleteUrlRewrite(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerDeleteUrlRewrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerDeleteUrlRewrite`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerDeleteUrlRewrite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerDeleteUrlRewriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerDeleteUrlRewriteRequest**](UrlmanagerDeleteUrlRewriteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerDeleteUrlRewrite2

> map[string]interface{} UrlManagerDeleteUrlRewrite2(ctx).Body(body).Execute()

Delete Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerDeleteUrlRewriteRequest() // UrlmanagerDeleteUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerDeleteUrlRewrite2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerDeleteUrlRewrite2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerDeleteUrlRewrite2`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerDeleteUrlRewrite2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerDeleteUrlRewrite2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerDeleteUrlRewriteRequest**](UrlmanagerDeleteUrlRewriteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerGetUrlRewrite

> UrlmanagerUrlRewrite UrlManagerGetUrlRewrite(ctx).Body(body).Execute()

Get Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerGetUrlRewriteRequest() // UrlmanagerGetUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerGetUrlRewrite(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerGetUrlRewrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerGetUrlRewrite`: UrlmanagerUrlRewrite
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerGetUrlRewrite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerGetUrlRewriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerGetUrlRewriteRequest**](UrlmanagerGetUrlRewriteRequest.md) |  | 

### Return type

[**UrlmanagerUrlRewrite**](UrlmanagerUrlRewrite.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerGetUrlRewrite2

> UrlmanagerUrlRewrite UrlManagerGetUrlRewrite2(ctx).Body(body).Execute()

Get Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerGetUrlRewriteRequest() // UrlmanagerGetUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerGetUrlRewrite2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerGetUrlRewrite2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerGetUrlRewrite2`: UrlmanagerUrlRewrite
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerGetUrlRewrite2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerGetUrlRewrite2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerGetUrlRewriteRequest**](UrlmanagerGetUrlRewriteRequest.md) |  | 

### Return type

[**UrlmanagerUrlRewrite**](UrlmanagerUrlRewrite.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerListUrlRewrites

> UrlmanagerListUrlRewritesResponse UrlManagerListUrlRewrites(ctx).Body(body).Execute()

List Url Rewrites



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerListUrlRewritesRequest() // UrlmanagerListUrlRewritesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerListUrlRewrites(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerListUrlRewrites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerListUrlRewrites`: UrlmanagerListUrlRewritesResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerListUrlRewrites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerListUrlRewritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerListUrlRewritesRequest**](UrlmanagerListUrlRewritesRequest.md) |  | 

### Return type

[**UrlmanagerListUrlRewritesResponse**](UrlmanagerListUrlRewritesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerListUrlRewrites2

> UrlmanagerListUrlRewritesResponse UrlManagerListUrlRewrites2(ctx).Body(body).Execute()

List Url Rewrites



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerListUrlRewritesRequest() // UrlmanagerListUrlRewritesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerListUrlRewrites2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerListUrlRewrites2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerListUrlRewrites2`: UrlmanagerListUrlRewritesResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerListUrlRewrites2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerListUrlRewrites2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerListUrlRewritesRequest**](UrlmanagerListUrlRewritesRequest.md) |  | 

### Return type

[**UrlmanagerListUrlRewritesResponse**](UrlmanagerListUrlRewritesResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerListUrlRewritesByTargetPaths

> UrlmanagerListUrlRewritesByTargetPathsRequest UrlManagerListUrlRewritesByTargetPaths(ctx).Body(body).Execute()

List Url Rewrites By Target Paths



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerListUrlRewritesByTargetPathsRequest() // UrlmanagerListUrlRewritesByTargetPathsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerListUrlRewritesByTargetPaths(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerListUrlRewritesByTargetPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerListUrlRewritesByTargetPaths`: UrlmanagerListUrlRewritesByTargetPathsRequest
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerListUrlRewritesByTargetPaths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerListUrlRewritesByTargetPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerListUrlRewritesByTargetPathsRequest**](UrlmanagerListUrlRewritesByTargetPathsRequest.md) |  | 

### Return type

[**UrlmanagerListUrlRewritesByTargetPathsRequest**](UrlmanagerListUrlRewritesByTargetPathsRequest.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerListUrlRewritesByTargetPaths2

> UrlmanagerListUrlRewritesByTargetPathsRequest UrlManagerListUrlRewritesByTargetPaths2(ctx).Body(body).Execute()

List Url Rewrites By Target Paths



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerListUrlRewritesByTargetPathsRequest() // UrlmanagerListUrlRewritesByTargetPathsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerListUrlRewritesByTargetPaths2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerListUrlRewritesByTargetPaths2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerListUrlRewritesByTargetPaths2`: UrlmanagerListUrlRewritesByTargetPathsRequest
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerListUrlRewritesByTargetPaths2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerListUrlRewritesByTargetPaths2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerListUrlRewritesByTargetPathsRequest**](UrlmanagerListUrlRewritesByTargetPathsRequest.md) |  | 

### Return type

[**UrlmanagerListUrlRewritesByTargetPathsRequest**](UrlmanagerListUrlRewritesByTargetPathsRequest.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerResolveUrlRewrite

> UrlmanagerUrlRewrite UrlManagerResolveUrlRewrite(ctx).Body(body).Execute()

Resolve Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerResolveUrlRewriteRequest() // UrlmanagerResolveUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerResolveUrlRewrite(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerResolveUrlRewrite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerResolveUrlRewrite`: UrlmanagerUrlRewrite
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerResolveUrlRewrite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerResolveUrlRewriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerResolveUrlRewriteRequest**](UrlmanagerResolveUrlRewriteRequest.md) |  | 

### Return type

[**UrlmanagerUrlRewrite**](UrlmanagerUrlRewrite.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UrlManagerResolveUrlRewrite2

> UrlmanagerUrlRewrite UrlManagerResolveUrlRewrite2(ctx).Body(body).Execute()

Resolve Url Rewrite



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-urlmanager"
)

func main() {
	body := *openapiclient.NewUrlmanagerResolveUrlRewriteRequest() // UrlmanagerResolveUrlRewriteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicOperationsAPI.UrlManagerResolveUrlRewrite2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicOperationsAPI.UrlManagerResolveUrlRewrite2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UrlManagerResolveUrlRewrite2`: UrlmanagerUrlRewrite
	fmt.Fprintf(os.Stdout, "Response from `BasicOperationsAPI.UrlManagerResolveUrlRewrite2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUrlManagerResolveUrlRewrite2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UrlmanagerResolveUrlRewriteRequest**](UrlmanagerResolveUrlRewriteRequest.md) |  | 

### Return type

[**UrlmanagerUrlRewrite**](UrlmanagerUrlRewrite.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

