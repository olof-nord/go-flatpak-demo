# go-flatpak-demo

A demo go application with Flatpak packaging

## Dependencies

This demo requires `flatpak` and `flatpak-builder` to be installed locally.
To build and run the go files directly, `go` needs to be installed as well.

Additionally, the following Flatpak SDKs should be installed:
```
flatpak install flathub org.freedesktop.Sdk//20.08
flatpak install org.freedesktop.Platform//20.08
flatpak install org.freedesktop.Platform//20.08
flatpak install org.freedesktop.Sdk.Extension.golang//20.08
```

## Build

`flatpak-builder build com.github.go-flatpak-demo.yml --force-clean --verbose`

## Install

`flatpak-builder --user --install build com.github.go-flatpak-demo.yml --force-clean --verbose`

## Run

To run the program with go directly:
`go run .`

To run the installed Flatpak application:
`flatpak run com.github.go-flatpak-demo`

If everything works as expected, "Hello World" should be shown in both cases.