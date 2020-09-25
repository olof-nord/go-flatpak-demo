# go-flatpak-demo

A demo go application with Flatpak packaging

This requires `flatpak` and `flatpak-builder` to be installed locally.

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

`flatpak run com.github.go-flatpak-demo`

If everything works as expected, "Hello World" should be shown.