name = "Plugin Template"
version = "0.1.0"
identifier = "com.modern-circuits.plugin-template"
windows_uuid = "ASDFGH-JMKJH-GFDSDFGHJKL"

installer "macOS"  {
    component "AU" {
        name = "Audio Unit (64 Bit)"
        payload_path = "macOS/AU/Plugin Template.component"
        install_path = "/Library/Audio/Plug-Ins/AU"
        version = "0.1.0"
    }

    component "VST3" {
        name = "VST3 (64 Bit)"
        payload_path = "macOS/VST3/Plugin Template.vst3"
        install_path = "/Library/Audio/Plug-Ins/VST3"
        version = "0.1.0"
    }

    component "CLAP" {
        name = "CLAP (64 Bit)"
        payload_path = "macOS/CLAP/Plugin Template.clap"
        install_path = "/Library/Audio/Plug-Ins/CLAP"
        version = "0.1.0"
    }
}

installer "Windows"  {
    component "VST3" {
        name = "VST3 (64 Bit)"
        payload_path = "Windows/VST3/Plugin Template.vst3"
        install_path = "/Library/Audio/Plug-Ins/VST3"
        version = "0.1.0"
    }

    component "CLAP" {
        name = "CLAP (64 Bit)"
        payload_path = "Windows/CLAP/Plugin Template.clap"
        install_path = "/Library/Audio/Plug-Ins/CLAP"
        version = "0.1.0"
    }
}