#solidfire-go
The solidfire-go library is a golang sdk for interacting with a NetApp/SolidFire Cluster.  This particular repo is a derivative of the auto-generated golang sdk package found on netapps ngage bitbucket repository.  What this derivative does is strip out the auto-gen code and just provides the resultant sdk itself.  In addition this library also provides a "methods" package with some wrappers around the basic and most common volume operations a user might be interested in.

The sdk itself isn't that difficult to use (as long as you have an API reference guide and a Golang IDE with auto-completion), but the methods package wraps soem of the basics into very easy to consume functions and also can serve as a good example on how to use the SDK.
