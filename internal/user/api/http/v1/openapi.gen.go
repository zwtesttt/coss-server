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
	// 用户激活
	// (GET /api/v1/user/activate)
	UserActivate(c *gin.Context, params UserActivateParams)
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
	// 获取服务端pgp公钥
	// (GET /api/v1/user/system/public_key)
	GetPGPPublicKey(c *gin.Context)
	// 获取用户信息
	// (GET /api/v1/user/{id})
	GetUser(c *gin.Context, id string)
	// 修改用户信息
	// (PUT /api/v1/user/{id})
	UpdateUser(c *gin.Context, id string)
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

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UserRegister(c)
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

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

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

	siw.Handler.UpdateUser(c, id)
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

	router.GET(options.BaseURL+"/api/v1/user/activate", wrapper.UserActivate)
	router.PUT(options.BaseURL+"/api/v1/user/bundle", wrapper.UpdateUserBundle)
	router.GET(options.BaseURL+"/api/v1/user/clients", wrapper.GetUserLoginClients)
	router.POST(options.BaseURL+"/api/v1/user/email/verification", wrapper.UserEmailVerification)
	router.POST(options.BaseURL+"/api/v1/user/login", wrapper.UserLogin)
	router.POST(options.BaseURL+"/api/v1/user/logout", wrapper.UserLogout)
	router.PUT(options.BaseURL+"/api/v1/user/password", wrapper.UpdateUserPassword)
	router.POST(options.BaseURL+"/api/v1/user/public_key", wrapper.SetUserPublicKey)
	router.PUT(options.BaseURL+"/api/v1/user/public_key", wrapper.ResetUserPublicKey)
	router.POST(options.BaseURL+"/api/v1/user/register", wrapper.UserRegister)
	router.GET(options.BaseURL+"/api/v1/user/system/public_key", wrapper.GetPGPPublicKey)
	router.GET(options.BaseURL+"/api/v1/user/:id", wrapper.GetUser)
	router.PUT(options.BaseURL+"/api/v1/user/:id", wrapper.UpdateUser)
	router.GET(options.BaseURL+"/api/v1/user/:id/bundle", wrapper.GetUserBundle)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xaW1MbyRX+K6pOHpIqIQF2XCk9xc7uupzsA7XO5cGhVK1RI3qZmZ7t6QEUl6okFtvg",
	"NYYkrC+EDfYG25S9lnB8AUNs/xiYGelfpHpaGo00PbohXypvGql1Tvf5zved091zGShEM4iOdGaC1GVg",
	"KtNIg97HCYqmEEW6grzHP5qIXmSQWd4T0i0NpC6Nxsfi4/FT8dOTccDyBgIpgHWGcoiCQhwYlBiIMiwM",
	"EAPp6YxF9TScYoimKYJZrOf4T1lkKhQbDBMdpIBzp2KvPbT/W7TXKrU7V+y1m/bKc3dx4y868L1kCFER",
	"1LmXCLtphjWUJhbjDsJzo0iDdCbwm8kon00hDkysIp2ldcLwFFagmFXHSf5wt3Z4p1bccLceSCdZ8L8i",
	"ma+RwrgXHs8L+hThplsDBWchgzQ8tTiYH8mREf7liDmDjRHiTQaqIwbhK6MgxaiF+DCiYYY0g+VBagqq",
	"JirEgUJMM42zwzaLNIhVaRhVaLK0SnJY96DgY6YI1SATQJw5DUI50+NUuHHPrm5pGUSHa1lHc+ksmsUK",
	"EpMPrK2BZx/GsDKjQ7H4YUbdaOXmzymaAinws2STzMk6k5NBGnt5r3oZnTZ9LofJYeKcDplFhz7vptNW",
	"No2mYs7mY3frgXv9lVMsxcZSMefpv+39/frzeCpW/fGJ+7Dkru/ETqVi9t5/7KV7tbvbsdOpWG29ZJc3",
	"xEiPfj2JE0PqsFdnmYgOnWIeZt9YmKIsSF3yfTSY58c0DO1khOp8ydP6typGOguLj+J9n8ZGGCS7/KOz",
	"tOc+qWCjSbChBC5L8awfulan1fJbe/uat+D34FGYlPt0nx3a//puyH6FbkEWdurePbTffO/cflW7/QLE",
	"hyVoUSklTYyLSKGInbP0rIrCmWF6v6Yz/s9t81/fcZb27MrV2t8f2DeuDDluAWrJ3A45P8IB4tqFFIti",
	"lr/IhVWE5ByCFNGzFpvmTxnv6YsGcL/78x84O73RvHh4vzanOc2YAQrcMK53AbKFOZsr9vX7sbMTF2LO",
	"6trx23+Kr6vl7WqldFQsOc937Ks3joql4/3HxwcH1ZeLzq27bvm+u3bVfbpsX9+qfvvmqLjA3WLGYQvZ",
	"BXEwi6gpnI4lRhOjjaYKGhikwKnEaOIUiAMDsmlv0Ulo4OTsWJJDkoQKw7OQefmQQyxyGe+KzotD4Nml",
	"nkhdyIKUl3RnGxa4Cwo1xBA1QeqS3NCFzwAPGEiBbyxE8yAORHkN6GJTLQW+ohBKWpRCPNTUedMUORzh",
	"Zwbl+/IxyQebBtFNkTLjo6OezBKd1fUXGoZabzOTX5ui12zaa+Vg/zTwuVNPsk6tgt+SygjARVMmWEtr",
	"9vUtQRBL0yDNh0FnMGc2aheY5ENbcqipKIYlyaDjd2Vn/XVYXtpSychChvgK6volQEImO0ey+ROE/ISy",
	"116/W81NSuPcml2FE2ZQVyClAQ7iWhc+j5RBybs0ydO7CXsUUl3wFy2HGSkh1Zt79uqtauVB3fKbf9jL",
	"KyL73I1FZ7nobC77vUkoMc4j1tb0iF7pBBHl1cLshUzBTqtJKUgpzMuA6GuhgwDUbyS7AOd1oMlZRFv2",
	"yQYxJSDaq3+rFUtCEmoL5ePDV9Jq8Dk3+aegxWHx2N+o+p1Vo4HuTFgx6pMgajiGUvWVhroLlP5WV46e",
	"yBeRKaL3sF+/cte3nKW1avld7Xb5+HDbXb4hOo0wqh4PToAkmoeaUVff5lYBwIwyNs57k0Y3T2YQn+18",
	"/q+nf3XG3yKJ5uA3dSMJhWher2Gac4RyM42PwpahQsZThPfaxAwfY3XYqwT5VN9DbCza65Xj/aJz71q1",
	"ctX96aGsQLdOX27Vtyc+eGNj7vrO8cFN5+YOx3v5O2fzwH1ScTcW3eeH7uGWOJCSOfS5IEV5Y9Fdf2k/",
	"XRO5Y2/u2j8UZVaaIYwsiu69kvSffow7JRoP3uvn9uruLzAxj4olqGcpwdmjYmkOZY6KpWkLziGcSCR+",
	"2SOJA1MO7jnbAAhM70OwvseGrP/2Swzoifr1U1I592vFon3twLcmZTc3MCyhbiFYZ1ibQz8JfQ5GapDy",
	"3BbpLrgF6dehcfZZGNUvTzQ5MRwAFaJPYaqlO+jD/XK1vO3PTIPzXyI9x0MyNv7rONCw3ng+I1EPomY7",
	"2HZuPxrUcAejt3YHM9qWsC1Tb5GjUNBkGf2BtgJiqYPvAJrC3yWDrYyKlTTfUkeqT7X81n1TFopm5Az7",
	"yk9ib96azRdFkz/hGfy9v0U/eS63TrEzuIGx/cuRbqlqP518RFjaIh6XK0Pt2or/74iIfoXM9xVThWSR",
	"9Mao5xY9PiguceH8Q1SLTnkk7ZZ8COXnN30d+vRyfhNOgoFqliyVuvCeohw2mbi467TfEIeb0p7jq4aJ",
	"j1a0+muqawtlt/wseK4fmdrBO0NpVO68dB/tDvl8/QS9/MCZ3Ymmkoa9pwrZnuQihaqVPefZgr264Hy/",
	"C94v0SM2ceFd8sc7pu2V2m0k7EJqM28ypLXV9A7neeIGwn1Sia7q5xGbOD/RWoA+AUXuJepRaxz81E4a",
	"ry6YXMbZQhcYxFKP3913SpFnp/1ezRiQTTdvTPq8lJn8SPvpcDgGx6otqL01ZcHT8wg8mpu2Dw/JMIps",
	"880mKd+2X9jfrsqkMfLdpcHLZftbLm1Y7D92io/cp2/ttRXZX+svj7TVGu8Ezl7dC5Z5Y5roqDcN+Vgn",
	"SeHEO+nlT1Tmy/QpcP3XVaair//qSuXf/f3/6FXLSxlddOtEt3dRgW7H0LNJZxuhtahaf5shleTgJoLn",
	"7NxD/f/tIHw+i2ieTWM9F4MZYrFYXdbQPENUh+pnRJG8KPYF1rMxPlojlOPc4t2cg7kcoglMPGFovhgA",
	"CpOF/wUAAP//vM7CuuoqAAA=",
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
