name       = "Plugin Template"
vendor     = "Modern Circuits"
version    = "0.1.0"
license    = "LICENSE.txt"
identifier = "com.modern-circuits.plugin-template"

installer "Windows" "innosetup" {
  uuid = "ASDFGH-JMKJH-GFDSDFGHJKL"

  artifact "VST3" {
    payload     = "Windows/VST3/${project.name}.vst3"
    destination = "{commoncf64}/VST3"
  }

  artifact "CLAP" {
    payload     = "Windows/CLAP/${project.name}.clap"
    destination = "{commoncf64}/CLAP"
  }
}
