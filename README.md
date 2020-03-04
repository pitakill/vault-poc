### Quick start

```
$ vault server -config=config.hcl # start server
$ vault operator unseal # Unseal Vault with the provided key in base64 in the json file
$ cd code && go run *.go
```
