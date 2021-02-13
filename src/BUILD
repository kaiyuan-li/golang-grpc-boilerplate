load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    embed = [":src_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "src_lib",
    srcs = ["main.go"],
    importpath = "github.com/laofeizhu/bazel-go-grpc/src",
    visibility = ["//visibility:private"],
)
