# Paket

[![License](https://img.shields.io/badge/License-Boost_1.0-lightblue.svg)](https://github.com/ModernCircuits/Paket/blob/main/LICENSE.txt)
[![codecov](https://codecov.io/gh/ModernCircuits/Paket/branch/main/graph/badge.svg?token=S8XON74JQU)](https://codecov.io/gh/ModernCircuits/Paket)
[![Test](https://github.com/ModernCircuits/Paket/actions/workflows/test.yml/badge.svg)](https://github.com/ModernCircuits/Paket/actions/workflows/test.yml)
[![Pre-Commit Hooks](https://github.com/ModernCircuits/Paket/actions/workflows/pre-commit.yml/badge.svg)](https://github.com/ModernCircuits/Paket/actions/workflows/pre-commit.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/moderncircuits/paket.svg)](https://pkg.go.dev/github.com/moderncircuits/paket)
[![Go Report Card](https://goreportcard.com/badge/github.com/moderncircuits/paket)](https://goreportcard.com/report/github.com/moderncircuits/paket)

> :warning: **Work in progress.**

Cross-platform meta installer generator. Primarily targets Windows `InnoSetup` & macOS `pkgbuild/productbuild`.

## Quick Start

### Application

```hcl
name       = "Example App"
vendor     = "Modern Circuits"
version    = "0.1.0"
license    = "LICENSE.txt"
identifier = "com.modern-circuits.example-app"

installer "macOS" "macos-pkg" {
  artifact "App" {
    payload     = "macOS/${project.name}.app"
    destination = "/Application"
  }
}

installer "Windows" "innosetup" {
  artifact "App" {
    payload     = "Windows/${project.name}.exe"
    destination = "{commonpf64}/${project.vendor}"
  }
}
```

### Audio Effect

```hcl
name       = "Example Effect"
vendor     = "Modern Circuits"
version    = "0.1.0"
license    = "LICENSE.txt"
identifier = "com.modern-circuits.example-effect"

installer "macOS" "macos-pkg" {
  artifact "AU" {
    name        = "Audio Unit"
    version     = "0.1.1"
    payload     = "macOS/AU/${project.name}.component"
    destination = "/Library/Audio/Plug-Ins/AU"
  }

  artifact "VST3" {
    payload     = "macOS/VST3/${project.name}.vst3"
    destination = "/Library/Audio/Plug-Ins/VST3"
  }

  artifact "CLAP" {
    payload     = "macOS/CLAP/${project.name}.clap"
    destination = "/Library/Audio/Plug-Ins/CLAP"
  }
}

installer "Windows" "innosetup" {
  artifact "VST3" {
    payload     = "Windows/VST3/${project.name}.vst3"
    destination = "{commoncf64}/VST3"
  }

  artifact "CLAP" {
    payload     = "Windows/CLAP/${project.name}.clap"
    destination = "{commoncf64}/CLAP"
  }
}
```

## Resources

### Text Formats

- [github.com/sajari/docconv](https://github.com/sajari/docconv)
- [github.com/gomarkdown/markdown](https://github.com/gomarkdown/markdown)

### macOS

- [github.com/open-eid/osx-installer/blob/master/distribution.xml](https://github.com/open-eid/osx-installer/blob/master/distribution.xml)
- [github.com/nodejs/node/blob/main/tools/macos-installer/productbuild](https://github.com/nodejs/node/blob/main/tools/macos-installer/productbuild)
- [keith.github.io/xcode-man-pages/productbuild.1.html](https://keith.github.io/xcode-man-pages/productbuild.1.html)
