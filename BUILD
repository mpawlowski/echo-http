load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/mpawlowski/echo-http
gazelle(name = "gazelle")

go_library(
    name = "echo-http_lib",
    srcs = ["main.go"],
    importpath = "github.com/mpawlowski/echo-http",
    visibility = ["//visibility:private"],
    deps = [
        "//handler",
        "@com_github_gorilla_mux//:mux",
    ],
)

go_binary(
    name = "echo-http",
    embed = [":echo-http_lib"],
    visibility = ["//visibility:public"],
)
