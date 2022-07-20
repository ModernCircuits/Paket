name = "Plugin Template"
vendor = "Modern Circuits"
version = "0.1.0"
identifier = "com.modern-circuits.plugin-template"

installer "macOS"  {
    component "AU" {
        payload = "macOS/AU/Plugin Template.component"
        destination = "/Library/Audio/Plug-Ins/AU"
    }

}

