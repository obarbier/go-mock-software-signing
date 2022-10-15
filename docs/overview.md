core :
	- Definition:
		key and user management + authorization server
	- Technology stack
		basic authentication
		    user management will follow basic authentication
		token authentication:
		    key management will be done using token based authentication
		swagger v2 specs
	- requirements
		authentication/authorization (V1-Done)
		    - ACL based authorization (see [policy acl](./policy_acl.md))
		secure communication between services
		    - implementation should follow security best practice
		        * input validation
		        * https encryption
		        * database security
		        * etc.
		support links/HATEOAS
		    * application should return link to get additional information about  resources
		80%  hermetic test coverage
		    * this include database test
		    * this include http server test
		    * this include functional test
		    * fuzz testing
		logging framework
		    * ability to configure log output + set log level
		    * develop best practice for logging
		        * https://www.dataset.com/blog/the-10-commandments-of-logging/
		key wrapping
	-API
		/users
		/key
sign-data
	- Definition:
		sign data using key from different place (kms , db, yubikey, etc)
	- Technology stack
		api key authentication
		nats publisher
		swagger v2 specs
	- requirements
		authentication/authorization (hmac based authentications)
		secure communication between services
		support links/HATEOAS
		80% hermetic test coverage
		enable https
		logging framework
		rate limit sign request (quotas implementation)
	-API	
		/sign
transparency
	- Definition:
		signature transparency logs
	- Technology stack
		nats subscriber
		trillian backend
		grpc
		logging framework
		

===== generator =====
swagger init spec \
  --title "core" \
  --description "key/user management and authorization server" \
  --version 1.0.0 \
  --scheme http \
  --consumes application/github.com.obarbier.custom-app.core.v1+json \
  --produces application/github.com.obarbier.custom-app.core.v1+json

TODO: user api
	1) write swagger file 
	2) how to include hateoas
