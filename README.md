secret
======

secret is a utility to encrypt and decrypt small amounts of data like credentials so that they can be checked in along with the rest of the source code.

## As a Library

As a library to easily read and write secret data.

```
import "github.com/savaki/secret"
```

## Secret-Tool

In addition to a library, secret also comes with a command line tool, secret-tool, to simplify managing secrets.

#### Building 

To build secret-tool, execute the following commands:

```
cd secret-tool
go build
```

This will create a native binary, secret-tool, in the secret-tool directory

#### Create a New Secret

```
secret-tool --create secret.dat
```

creates a new secret and stores the data in base64 encoded form in the file you've specified, in this case, secret.dat

#### Encrypt Data

```
matt$ secret-tool --filename sample.dat --encrypt
input: hello world
+wqyl2Z6NkE3o7NpKYok+VMY1WoyH6zhrg0z
```

In this example, we encrypt the string, hello world, using the secret data stored in the provided file, sample.dat.  The resulting encoded string is output to the screen.

#### Exec

The last mode that secret-tool can be used in is to exec other commands.  Here's an example of how it's used:

```
secret-tool --filename sample.dat --exec /your/command/here
```

Here's what happens under the hood:

1. secret-tool loads the secret information from the file provided, sample.dat
2. secret-tool scans the environment params, looking for a param name {NAME}_CIPHER
3. secret-tool will (a) assume that each param named {NAME}_CIPHER points to the base64-encoded cipher param and (b) proceed to decrypt the data and store the result in an environment variable named {NAME}
4. secret-tool will then exec the command specified, /your/command/here

This is useful in a number of scenarios:

* placing environment variables into env and executing a script.  For example, suppose we want to execute a goose db migration using DB_USERNAME and DB_PASSWORD environment variables.  We could configure the following values:
   
   ```
   DB_USERNAME_CIPHER = "base64-encoded-cipher-text-of-username"
   DB_PASSWORD_CIPHER = "base64-encoded-cipher-text-of-password"
   ```


## Notes

Secret utilizes the golang secret port [link](http://godoc.org/code.google.com/p/go.crypto/nacl/secretbox) which was based on the original [NaCl project](http://nacl.cr.yp.to/secretbox.html)