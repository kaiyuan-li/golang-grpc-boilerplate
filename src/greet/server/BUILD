load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "greet_server",
    embed = [":greet_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "greet_lib",
    srcs = ["greet_server.go"],
    importpath = "github.com/laofeizhu/bazel-go-grpc/src/greet",
    visibility = ["//visibility:private"],
    deps = [
        "//src/proto:greeting_go_proto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
