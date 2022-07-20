name = "Plugin Template"
version = "0.1.0"
license = "LICENSE.txt"
identifier = "com.modern-circuits.plugin-template"

installer "macOS"  {
    welcome = "LICENSE.txt"
    conclusion = "LICENSE.txt"

    component "AU" {
        name = "Audio Unit"
        payload_path = "macOS/AU/Plugin Template.component"
        install_path = "/Library/Audio/Plug-Ins/AU"
    }

    component "VST3" {
        payload_path = "macOS/VST3/Plugin Template.vst3"
        install_path = "/Library/Audio/Plug-Ins/VST3"
    }

    component "CLAP" {
        payload_path = "macOS/CLAP/Plugin Template.clap"
        install_path = "/Library/Audio/Plug-Ins/CLAP"
    }
}

installer "Windows"  {
    uuid = "ASDFGH-JMKJH-GFDSDFGHJKL"

    component "VST3" {
        payload_path = "Windows/VST3/Plugin Template.vst3"
        install_path = "${windows.commoncf64}/VST3"
    }

    component "CLAP" {
        payload_path = "Windows/CLAP/Plugin Template.clap"
        install_path = "${windows.commoncf64}/CLAP"
    }
}