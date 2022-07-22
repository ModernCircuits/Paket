name       = "Plugin Template"
vendor     = "Modern Circuits"
version    = "0.1.0"
identifier = "com.modern-circuits.plugin-template"

installer "macOS" {
  artifact "AU" {
    payload     = "macOS/AU/Plugin Template.artifact"
    destination = "/Library/Audio/Plug-Ins/AU"
  }
}
