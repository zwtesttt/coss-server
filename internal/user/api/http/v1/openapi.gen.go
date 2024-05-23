// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 修改用户信息
	// (PUT /api/v1/user)
	UpdateUser(c *gin.Context)
	// 用户激活
	// (GET /api/v1/user/activate)
	UserActivate(c *gin.Context, params UserActivateParams)
	// 修改用户头像
	// (PUT /api/v1/user/avatar)
	UpdateUserAvatar(c *gin.Context)
	// 修改用户密钥包
	// (PUT /api/v1/user/bundle)
	UpdateUserBundle(c *gin.Context)
	// 获取该用户当前登录的所有客户端
	// (GET /api/v1/user/clients)
	GetUserLoginClients(c *gin.Context)
	// 发送激活邮件
	// (POST /api/v1/user/email/verification)
	UserEmailVerification(c *gin.Context)
	// 用户登录
	// (POST /api/v1/user/login)
	UserLogin(c *gin.Context)
	// 退出登录
	// (POST /api/v1/user/logout)
	UserLogout(c *gin.Context)
	// 修改密码
	// (PUT /api/v1/user/password)
	UpdateUserPassword(c *gin.Context)
	// 设置用户pgp公钥
	// (POST /api/v1/user/public_key)
	SetUserPublicKey(c *gin.Context)
	// 重置用户公钥
	// (PUT /api/v1/user/public_key)
	ResetUserPublicKey(c *gin.Context)
	// 用户注册
	// (POST /api/v1/user/register)
	UserRegister(c *gin.Context)
	// 搜索用户
	// (GET /api/v1/user/search)
	SearchUser(c *gin.Context, params SearchUserParams)
	// 获取服务端pgp公钥
	// (GET /api/v1/user/system/public_key)
	GetPGPPublicKey(c *gin.Context)
	// 获取用户信息
	// (GET /api/v1/user/{id})
	GetUser(c *gin.Context, id string)
	// 获取用户密钥包
	// (GET /api/v1/user/{id}/bundle)
	GetUserBundle(c *gin.Context, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateUser(c)
}

// UserActivate operation middleware
func (siw *ServerInterfaceWrapper) UserActivate(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params UserActivateParams

	// ------------- Required query parameter "user_id" -------------

	if paramValue := c.Query("user_id"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument user_id is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "user_id", c.Request.URL.Query(), &params.UserId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter user_id: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "key" -------------

	if paramValue := c.Query("key"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument key is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "key", c.Request.URL.Query(), &params.Key)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter key: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UserActivate(c, params)
}

// UpdateUserAvatar operation middleware
func (siw *ServerInterfaceWrapper) UpdateUserAvatar(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateUserAvatar(c)
}

// UpdateUserBundle operation middleware
func (siw *ServerInterfaceWrapper) UpdateUserBundle(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateUserBundle(c)
}

// GetUserLoginClients operation middleware
func (siw *ServerInterfaceWrapper) GetUserLoginClients(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserLoginClients(c)
}

// UserEmailVerification operation middleware
func (siw *ServerInterfaceWrapper) UserEmailVerification(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UserEmailVerification(c)
}

// UserLogin operation middleware
func (siw *ServerInterfaceWrapper) UserLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UserLogin(c)
}

// UserLogout operation middleware
func (siw *ServerInterfaceWrapper) UserLogout(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UserLogout(c)
}

// UpdateUserPassword operation middleware
func (siw *ServerInterfaceWrapper) UpdateUserPassword(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateUserPassword(c)
}

// SetUserPublicKey operation middleware
func (siw *ServerInterfaceWrapper) SetUserPublicKey(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SetUserPublicKey(c)
}

// ResetUserPublicKey operation middleware
func (siw *ServerInterfaceWrapper) ResetUserPublicKey(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ResetUserPublicKey(c)
}

// UserRegister operation middleware
func (siw *ServerInterfaceWrapper) UserRegister(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UserRegister(c)
}

// SearchUser operation middleware
func (siw *ServerInterfaceWrapper) SearchUser(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchUserParams

	// ------------- Required query parameter "email" -------------

	if paramValue := c.Query("email"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument email is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "email", c.Request.URL.Query(), &params.Email)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter email: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SearchUser(c, params)
}

// GetPGPPublicKey operation middleware
func (siw *ServerInterfaceWrapper) GetPGPPublicKey(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPGPPublicKey(c)
}

// GetUser operation middleware
func (siw *ServerInterfaceWrapper) GetUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUser(c, id)
}

// GetUserBundle operation middleware
func (siw *ServerInterfaceWrapper) GetUserBundle(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserBundle(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.PUT(options.BaseURL+"/api/v1/user", wrapper.UpdateUser)
	router.GET(options.BaseURL+"/api/v1/user/activate", wrapper.UserActivate)
	router.PUT(options.BaseURL+"/api/v1/user/avatar", wrapper.UpdateUserAvatar)
	router.PUT(options.BaseURL+"/api/v1/user/bundle", wrapper.UpdateUserBundle)
	router.GET(options.BaseURL+"/api/v1/user/clients", wrapper.GetUserLoginClients)
	router.POST(options.BaseURL+"/api/v1/user/email/verification", wrapper.UserEmailVerification)
	router.POST(options.BaseURL+"/api/v1/user/login", wrapper.UserLogin)
	router.POST(options.BaseURL+"/api/v1/user/logout", wrapper.UserLogout)
	router.PUT(options.BaseURL+"/api/v1/user/password", wrapper.UpdateUserPassword)
	router.POST(options.BaseURL+"/api/v1/user/public_key", wrapper.SetUserPublicKey)
	router.PUT(options.BaseURL+"/api/v1/user/public_key", wrapper.ResetUserPublicKey)
	router.POST(options.BaseURL+"/api/v1/user/register", wrapper.UserRegister)
	router.GET(options.BaseURL+"/api/v1/user/search", wrapper.SearchUser)
	router.GET(options.BaseURL+"/api/v1/user/system/public_key", wrapper.GetPGPPublicKey)
	router.GET(options.BaseURL+"/api/v1/user/:id", wrapper.GetUser)
	router.GET(options.BaseURL+"/api/v1/user/:id/bundle", wrapper.GetUserBundle)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"IDIAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
