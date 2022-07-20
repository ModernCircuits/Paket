name = "Plugin Template"
version = "0.1.0"
identifier = "com.modern-circuits.plugin-template"
windows_uuid = "ASDFGH-JMKJH-GFDSDFGHJKL"

installer "macOS"  {
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
    component "VST3" {
        payload_path = "Windows/VST3/Plugin Template.vst3"
        install_path = "/Library/Audio/Plug-Ins/VST3"
    }

    component "CLAP" {
        payload_path = "Windows/CLAP/Plugin Template.clap"
        install_path = "/Library/Audio/Plug-Ins/CLAP"
    }
}