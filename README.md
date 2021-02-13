Boilerplate GoLang + gRPC + Bazel
=================================

# Run `build_cleaner`

```
bazel run //:gazelle
```


# To run the `go_binary`

```
bazel build src:main
```

# Notes
* A `BUILD` file is needed in the root directory so `WORKSPACE` will take effect.

* For go_repository if you don't know the hash, just put in a version number. Error message will tell you what's the hash you got.
