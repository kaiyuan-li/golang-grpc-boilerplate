Boilerplate GoLang + gRPC + Bazel
=================================

# To run the `go_binary`

```
bazel build src:main
```

# To run the greeting gRPC service

```
bazel run src/greet/server:greet_server
bazel run src/greet/client:greet_client
```

# Notes
* A `BUILD` file is needed in the root directory so `WORKSPACE` will take effect.

* For go_repository if you don't know the hash, just put in a version number. Error message will tell you what's the hash you got.
