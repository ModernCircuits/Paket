name       = "Plugin Template"
vendor     = "Modern Circuits"
version    = "0.1.0"
license    = "LICENSE.txt"
identifier = "com.modern-circuits.plugin-template"

installer "macOS" {
  welcome    = "LICENSE.txt"
  conclusion = "LICENSE.txt"

  artifact "AU" {
    name        = "Audio Unit"
    version     = "0.1.1"
    payload     = "macOS/AU/Plugin Template.artifact"
    destination = "/Library/Audio/Plug-Ins/AU"
  }

  artifact "VST3" {
    payload     = "macOS/VST3/Plugin Template.vst3"
    destination = "/Library/Audio/Plug-Ins/VST3"
  }

  artifact "CLAP" {
    payload     = "macOS/CLAP/Plugin Template.clap"
    destination = "/Library/Audio/Plug-Ins/CLAP"
  }
}

installer "InnoSetup" {
  uuid = "ASDFGH-JMKJH-GFDSDFGHJKL"

  artifact "VST3" {
    payload     = "Windows/VST3/Plugin Template.vst3"
    destination = "{commoncf64}/VST3"
  }

  artifact "CLAP" {
    payload     = "Windows/CLAP/Plugin Template.clap"
    destination = "{commoncf64}/CLAP"
  }
}
