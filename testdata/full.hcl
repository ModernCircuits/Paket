name       = "Plugin Template"
vendor     = "Modern Circuits"
version    = "0.1.0"
license    = "LICENSE.txt"
identifier = "com.modern-circuits.plugin-template"

installer "macOS" {
  welcome    = "LICENSE.txt"
  conclusion = "LICENSE.txt"

  component "AU" {
    name        = "Audio Unit"
    payload     = "macOS/AU/Plugin Template.component"
    destination = "/Library/Audio/Plug-Ins/AU"
  }

  component "VST3" {
    payload     = "macOS/VST3/Plugin Template.vst3"
    destination = "/Library/Audio/Plug-Ins/VST3"
  }

  component "CLAP" {
    payload     = "macOS/CLAP/Plugin Template.clap"
    destination = "/Library/Audio/Plug-Ins/CLAP"
  }
}

installer "Windows" {
  uuid = "ASDFGH-JMKJH-GFDSDFGHJKL"

  component "VST3" {
    payload     = "Windows/VST3/Plugin Template.vst3"
    destination = "${windows.commoncf64}/VST3"
  }

  component "CLAP" {
    payload     = "Windows/CLAP/Plugin Template.clap"
    destination = "${windows.commoncf64}/CLAP"
  }
}
