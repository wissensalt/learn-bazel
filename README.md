# learn-bazel

## Deps:

- [Golang](https://go.dev/)
- [Bazel](https://bazel.build/)
- [Bazelisk](https://github.com/bazelbuild/bazelisk)
- [Gazelle](https://github.com/bazelbuild/bazel-gazelle)

## Install Bazel
```bash
brew install bazel
```

## Install Bazelisk

```bash
go get github.com/bazelbuild/bazelisk
```

## Install Gazelle

```bash
go get github.com/bazelbuild/bazel-gazelle
```

## Build Binary

```bash
  bazel build //...
```

## Run Binary for app-one

```bash
  bazel run //apps/app-one
```

## Run Binary for app-two
```bash
  bazel run //apps/app-two
```

## Recommendations
### Install [Go Version Manager](https://github.com/moovweb/gvm)
```bash
  bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```