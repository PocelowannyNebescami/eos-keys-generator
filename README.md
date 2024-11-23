# EOS key generator

HTMX website to get EOS key-pair (private and public key).

## Setup

Project dependecies:
- [golang](https://go.dev/doc/install)
- [templ](https://templ.guide/quick-start/installation)
- [tailwindcss](https://tailwindcss.com/docs/installation)

### Build

To build invoke make at the project root

```sh
make
```

or add target to perform specific step

```sh
# run templ
make templ
# run tailwindcss
make tailwind
# remove binary
make clean
```

#### live reload

[Air](https://github.com/air-verse/air?tab=readme-ov-file#installation) is used
to rebuild the project, once changes were introduced. To start air, run it
inside project root

```sh
air
```
