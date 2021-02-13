load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "greeting_proto",
    srcs = ["greeting.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "greeting_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "greeting",
    proto = ":greeting_proto",
    visibility = ["//visibility:public"],
)
