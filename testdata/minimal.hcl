// SPDX-License-Identifier: BSL-1.0

name       = "Plugin Template"
vendor     = "Modern Circuits"
version    = "0.1.0"
identifier = "com.modern-circuits.plugin-template"

installer "macOS" "macos-pkg" {
  artifact "AU" {
    payload     = "macOS/AU/${project.name}.component"
    destination = "/Library/Audio/Plug-Ins/AU"
  }
}
