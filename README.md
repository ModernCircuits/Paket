# Modern Circuits - Paket

[![License](https://img.shields.io/badge/License-Boost_1.0-lightblue.svg)](https://github.com/ModernCircuits/Paket/blob/main/LICENSE.txt)
[![Test](https://github.com/ModernCircuits/Paket/actions/workflows/test.yml/badge.svg)](https://github.com/ModernCircuits/Paket/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/ModernCircuits/Paket/branch/main/graph/badge.svg?token=S8XON74JQU)](https://codecov.io/gh/ModernCircuits/Paket)
[![Pre-Commit Hooks](https://github.com/ModernCircuits/Paket/actions/workflows/pre-commit.yml/badge.svg)](https://github.com/ModernCircuits/Paket/actions/workflows/pre-commit.yml)

Cross-platform installer generator. Primarily targets Windows `InnoSetup` & macOS `pkgbuild/productbuild`. Work in progress.

## Quick Start

## Application

```hcl
name       = "Example App"
vendor     = "Modern Circuits"
version    = "0.1.0"
license    = "LICENSE.txt"
identifier = "com.modern-circuits.example-app"

installer "macOS" {
  component "App" {
    payload     = "macOS/Example App.app"
    destination = "/Application"
  }
}

installer "InnoSetup" {
  component "App" {
    payload     = "Windows/Example App.exe"
    destination = "{commoncf64}/commonpf64/Modern Circuits"
  }
}
```

## Audio Effect

```hcl
name       = "Example Effect"
vendor     = "Modern Circuits"
version    = "0.1.0"
license    = "LICENSE.txt"
identifier = "com.modern-circuits.example-effect"

installer "macOS" {
  component "AU" {
    name        = "Audio Unit"
    version     = "0.1.1"
    payload     = "macOS/AU/Example Effect.component"
    destination = "/Library/Audio/Plug-Ins/AU"
  }

  component "VST3" {
    payload     = "macOS/VST3/Example Effect.vst3"
    destination = "/Library/Audio/Plug-Ins/VST3"
  }

  component "CLAP" {
    payload     = "macOS/CLAP/Example Effect.clap"
    destination = "/Library/Audio/Plug-Ins/CLAP"
  }
}

installer "InnoSetup" {
  component "VST3" {
    payload     = "Windows/VST3/Example Effect.vst3"
    destination = "{commoncf64}/VST3"
  }

  component "CLAP" {
    payload     = "Windows/CLAP/Example Effect.clap"
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
