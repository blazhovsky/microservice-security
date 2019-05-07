# Microservice Security Shared library

[![Build](https://travis-ci.com/blazhovsky/microservice-security.svg?token=UB5yzsLHNSbtjSYrGbWf&branch=master)](https://travis-ci.com/blazhovsky/microservice-security)
[![Test Coverage](https://api.codeclimate.com/v1/badges/c29a5eb496f7da156e83/test_coverage)](https://codeclimate.com/repos/59e72654b82c7d02b2001889/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/c29a5eb496f7da156e83/maintainability)](https://codeclimate.com/repos/59e72654b82c7d02b2001889/maintainability)

This library contains functions that are commonly used by all microservices for
setting up the security.
Also exposes functions to set up different security mechanisms for securing the
microservices.

There are two important packages:
 * auth - defines the standard Auth object and offers functions for manipulating the request context for setting/getting the Auth object.
 * chain - defines the SecurityChain and function for creating new security middleware and registering it with the security chain.

# Manipulating the security Auth

Auth object can be manipulated in the request context with helpers of the ```auth``` package.
The package exposes function for checking, setting and getting the Auth object from the context.

To get the current Auth object from the context, you can use ```auth.GetAuth(context.Context)``` helper.

```go

authObj := auth.GetAuth(ctx)

if authObj == nil {
  // The context contains no Auth
}

```
## Using it with Goa generated service actions

Generated Goa actions provide an action context structure (for example GetUserContext)
which contain the Context generated by the middleware chain.

To extract the Auth from the context, you can use the following pattern:

```go
// Get runs the get action.
func (c *UserController) Get(ctx *app.GetUserContext) error {
  // ctx contains the context.Context returned by the middleware chain
  authObj := auth.GetAuth(ctx.Context)

  // if protected by the security chain, the Auth is guaranteed to be in the context.

  userID := authObj.UserID
  roles := authObj.Roles

  // do something with the userID or roles
}
```

# Checking for Auth

To check for created authentication in the context, you can use ```auth.HasAuth(context.Context)``` helper.

```go
// if context is provided as ctx

if auth.HasAuth(ctx) {
  // There is an authentication in the context
}else{
  // The context does not contain auth object
}
```

# Setting the Auth in the context

As context.Context is immutable itself, setting an auth object comes down to
creating new context that inherits from a parent context and returns a new context
with the value of the Auth. You can use ```auth.SetAuth(parentContext context.Context, authObj auth.Auth) context.Context``` helper.
This helper returns the new context that inherits from the parent and adds new the Auth as value.

```go
// Goa middleware
func SomeCustomGoaMiddleware(hnd goa.Handler) goa.Handler {
  return func (parentContext context.Context, rw http.ResponseWriter, req *http.Request) error {
    // generate Auth in some way - JWT for example
    authObj := checkRequestAndGenerateAuthFromJWT(ctx, req)

    // then we want to set it in context
    return hnd(auth.SetAuth(parentContext, authObj), rw, req)
  }
}
```

# SecurityContext

This structure holds the value of the Authentication and any possible errors in
the request context.
The API for manipulating the security context consists of:

 * ```GetSecurityContext``` - returns the SecurityContext from the current request context. If the context does not
 contain a security context yet, a nil is returned.
 * ```GetSecurityErrors``` - returns the SecurityErrors map from the security context.
 * ```SetSecurityError``` - sets an error for a particular security mechanism. There is only one error per mechanism.


# Security Chain

SecurityChain is a standard chain of processing of the incoming http request. It is intended to be
used as a middleware in Goa infrastructure (although it can execute by its own using the standard http Handlers by Go).

SecurityChain is an interface residing in the ```chain``` package. This package also provides other helper functions
related to creating the chain, wrapping it in goa.Midlleware or wrapping a Goa middleware in the chain itself.

The chain is composed of a list of middleware functions of type ```chain.SecurityChainMiddleware```.
The signature of this function is:

```go
func (context.Context, http.ResponseWriter, *http.Request) (context.Context, http.ResponseWriter, error)
```
The chain executes the middleware functions in the order they are registered. The input context and ResponseWriter for
each middleware is the output of the previous middleware.
The resulting context, ResponseWriter and Request are returned back by the ```SecurityChain.Execute(...)``` method.

## Creating a new security chain

To create new security chain, you can use the factory function provided in the ```chain``` package:

```go
securityChain := chain.NewSecurityChain()

```

## Adding middleware to a security chain

To create a security chain that will actually do some processing you need to add some middlewares to it.

```go
// in Goa's microservice main.go file:

// Assuming we have implemented the following middleware functions:
func  JWTMiddleware(ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, http.ResponseWriter, error)  {
  // Processes a JWT token and creates Auth object based on it
}

func  OAuth2Middleware(ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, http.ResponseWriter, error)  {
  // Processes an OAuth2 access token and creates Auth object based on it
}

func  SAMLMiddleware(ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, http.ResponseWriter, error)  {
  // Processes SAML token and creates Auth object based on it
}

func  CheckAuthMiddleware(ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, http.ResponseWriter, error)  {
  // check for existence of Auth object in context
}

func  ACLMiddleware(ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, http.ResponseWriter, error)  {
  // Processes the Auth object and the request against the ACL policies
}

func main() {
  // We want to set up a security chain that first will attmpt to create
  // authentication based on JWT, OAuth2 or SAML, then will check the Auth object
  // (if created) against an ACL policy.

  securityChain := chain.NewSecurityChain().  // create new security chain
    AddMiddleware(JWTMiddleware).        // 1. Attempt JWT
    AddMiddleware(OAuth2Middleware).     // 2. Attempt OAuth2
    AddMiddleware(SAMLMiddleware).       // 3. Attempt SAML
    AddMiddleware(CheckAuthMiddleware).  // Check if an Auth object has been created in steps 1-3
    AddMiddleware(ACLMiddleware)         // If Auth was created, check with the ACL policy

    // Goa generated service setup

    // Create service
  	service := goa.New("user")

  	// Mount middleware
  	service.Use(middleware.RequestID())
  	service.Use(middleware.LogRequest(true))
  	service.Use(middleware.ErrorHandler(service, true))
  	service.Use(middleware.Recover())

    // Attach the SecurityChain as Goa Middleware
    service.Use(chain.AsGoaMiddleware(securityChain))

    // continue initialization here
}

```

## Writing security middleware functions with chain.MiddlewareBuilder

A ```chain.MiddlewareBuilder``` is used to build a ```chain.SecurityChainMiddleware```.
This is useful in cases when we need to have some prior initialiation for the security
middleware.
The  ```chain.MiddlewareBuilder``` type is a function with signature:
```go
func () SecurityChainMiddleware
```

A simple example to illustrate the usage of a builder that requires prior initialization
is in the case of a JWT middleware. A JWT requires a shared secret key between the service
and the token issuer to verify the token signature.
So when registering a JWT security middleware we need to pass the secret key somehow,
without having to hard code it in the source code or access it as a global constant.

One way to do it is to read it from a file (or a shared key-value store) and then
pass it to a builder in a closure.

Let's assume we have a function that loads the secret key from some kind of persistence (file, store etc)
called ```loadJWTSecret() string```. The we can use a MiddlewareBuilder to build a
SecurityChainMiddleware with the shared secret:

```go

func JWTSecurityBuilder() SecurityChainMiddleware {
  secretKey := loadJWTSecret() // load the secret and trap it in the function closure

  // now return the SecurityChainMiddleware that can access the secretKey
  return func (ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, http.ResponseWriter, error) {
    // the secret key is available here and we can use it to verify the JWT token
  }
}

```

Then, you can build the middleware by executing the MiddlewareBuilder:

```go
securityChain.AddMiddleware(JWTSecurityBuilder()) // build the JWT security middleware with the shared secret key

```


## Security Mechanism types

In real-world scenarios, setting up a security with shared secret keys would require a more complicated
code and possibly lots of other configuration parameters (key store URL, key file path etc).

The procedure for initializing a JWT/OAuth2/SAML security middleware would be the same across the services,
only the configuration parameters would change.

Instead of calling the MiddlewareBuilder directly, you can delegate that job to the SecurityChain itself. You can
tell the chain to use a certain type of security (by name) and let it decide when to initialize it and to build the
security middleware itself.

In order to do so, you'll need to register the security type:

```go
chain.NewSecurity("JWT", JWTSecurityBuilder)
```

**NewSecurity** registers a security builder for a specified security type in the global Midlleware registry.
Once registered, it can later be used with the security chain, without passing the actual middleware or builder to the chain by
using ```SecurityChain.AddMiddlewareType(name string)```

Using the example above for the JWT security, we can create a "JWT" security type and use it in the following way:

```go
// in a jwt.go file

func JWTSecurityBuilder() SecurityChainMiddleware {
  secretKey := loadJWTSecret() // load the secret and trap it in the function closure

  // now return the SecurityChainMiddleware that can access the secretKey
  return func (ctx context.Context, rw http.ResponseWriter, req *http.Request) (context.Context, http.ResponseWriter, error) {
    // the secret key is available here and we can use it to verify the JWT token
  }
}

// register it with the middleware registry

func init(){
  chain.NewSecurity("JWT", JWTSecurityBuilder)
}

```

then, in **main.go**:

```go

func main(){
  sc := chain.NewSecurityChain() // create the security chain
  sc.AddMiddlewareType("JWT") // add JWT security to it

  // continue initializing
}


```

# Allowing access to public resources

To allow access to a public resource you need to specify an ignore pattern to the security chain.
The patterns are regular expressions to which the request Path is matched against. If the path
matches the ignore pattern regexp, then the chain does not execute and the handling of the
request is passed down the next middlewares.

Here's an example on how to add ignore pattern to the security chain:

```Go

func main(){
  securityChain := chain.NewSecurityChain()
  // Add the ignore patterns:
  // - everything under /public/
  // - CSS, Javascript and HTML files

  // Note that the order does not matter
  securityChain.AddIgnorePattern("/public/.+")
  securityChain.AddIgnorePattern(".+\\.js")
  securityChain.AddIgnorePattern(".+\\.css")
  securityChain.AddIgnorePattern(".+\\.html")

  // Add your middleware functions next
  securityChain.AddMiddleware(MiddlewareOne)
  securityChain.AddMiddleware(MiddlewareTwo)

  // Note that the order of adding patterns and middleware functions does not matter,
  // you can add the middleware first, then the ignore patterns.

}

```

If you are setting up the chain using the provided builders in package "flow", then
you can add the ignore patterns in the configuration. The configuration property name
is ```ignorePatterns``` and accepts and array of strings.


# Setting up a security for a microservice

The easier way to set up a security is to use the ```flow``` package and the helper ```flow.NewSecurityFromConfig()```.

In the microservice main file, you first need to load the microservice configuration.
The pass the configuration to the helper and create new SecurityChain.

Finally you need to add the security chain as a middleware to the service itself.

```go
func main(){

  // 1. Load the configuration
  conf, err := conf.LoadConfiguration("config.json")
  if err {
    // We have a problem loading the configuration
    panic(err)
  }

  // 2. Create a security chain
  securityChain, cleanup, err := flow.NewSecurityFromConfig(conf)
  if err != nil {
    // There was a problem setting up the security
    panic(err)
  }

  defer cleanup()

  // ... Goa service and controllers setup

  // 3. Finally add the security chain to the service as a middleware.
  service.Use(chain.AsGoaMiddleware(securityChain))

}

```

Here is an example of the configuration file:

```json
{
  "service":{
    "name": "user-microservice",
    "port": 8081,
    "virtual_host": "user.services.jormungandr.org",
    "hosts": ["localhost", "user.services.jormungandr.org"],
    "weight": 10,
    "slots": 100
  },
  "gatewayUrl": "http://localhost:8000",
  "security":{
    "keysDir": "keys",
    "ignorePatterns": ["/public-resource/.+", ".+\\.js", ".+\\.css", ".+\\.html"],
    "jwt":{
      "name": "JWTSecurity",
      "description": "JWT security middleware",
      "tokenUrl": "http://localhost:8000/jwt"
    },
    "saml":{
      "certFile": "keys/user-service.cert",
      "keyFile": "keys/user-service.key",
      "identityProviderUrl": "http://localhost:8000/saml/idp",
      "userServiceUrl": "http://localhost:8000/user",
      "registrationServiceUrl": "http://localhost:8000/user/register"
    },
    "oauth2":{
      "description": "OAuth2 security middleware",
      "tokenUrl": "https://localhost:8000/oauth2/token",
      "authorizeUrl": "https://localhost:8000/oauth2/authorize"
    }
  },
  "database":{
    "host": "127.0.0.1:27017",
    "database": "users",
    "user": "restapi",
    "pass": "restapi"
  }
}
```

Note that if you don't want to use any of the JWT, OAuth2 or SAML security middlewares,
you can omit the approriate subsections ("jwt", "oauth2" or "saml") from the "security" section.

## Contributing

For contributing to this repository or its documentation, see the [Contributing guidelines](CONTRIBUTING.md).