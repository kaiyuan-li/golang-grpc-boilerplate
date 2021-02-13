load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "client_lib",
    srcs = ["greet_client.go"],
    importpath = "github.com/laofeizhu/bazel-go-grpc/src/greet/client",
    visibility = ["//visibility:private"],
    deps = [
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_x_net//context",
        "//src/proto:greeting_go_proto",
    ],
)

go_binary(
    name = "greet_client",
    embed = [":client_lib"],
    visibility = ["//visibility:public"],
)
