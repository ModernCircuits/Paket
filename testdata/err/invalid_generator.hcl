name       = "Plugin Template"
vendor     = "Modern Circuits"
version    = "0.1.0"
identifier = "com.modern-circuits.plugin-template"

installer "macOS" "macos-magic" {
  artifact "AU" {
    payload     = "macOS/AU/${project.name}.component"
    destination = "/Library/Audio/Plug-Ins/AU"
  }
}
