project_name = "Plugin Template"
project_version = "0.1.0"
bundle_id = "com.modern-circuits.plugin-template"
windows_uuid = "ASDFGH-JMKJH-GFDSDFGHJKL"

installer "macOS"  {
    component "AU" {
        name = "Audio Unit (64 Bit)"
        payload = "macOS/AU/Plugin Template.component"
        install_path = "/Library/Audio/Plug-Ins/AU"
        version = "0.1.0"
    }

    component "VST3" {
        name = "VST3 (64 Bit)"
        payload = "macOS/VST3/Plugin Template.vst3"
        install_path = "/Library/Audio/Plug-Ins/VST3"
        version = "0.1.0"
    }

    component "CLAP" {
        name = "CLAP (64 Bit)"
        payload = "macOS/CLAP/Plugin Template.clap"
        install_path = "/Library/Audio/Plug-Ins/CLAP"
        version = "0.1.0"
    }
}

installer "Windows"  {
    component "VST3" {
        name = "VST3 (64 Bit)"
        payload = "Windows/VST3/Plugin Template.vst3"
        install_path = "/Library/Audio/Plug-Ins/VST3"
        version = "0.1.0"
    }

    component "CLAP" {
        name = "CLAP (64 Bit)"
        payload = "Windows/CLAP/Plugin Template.clap"
        install_path = "/Library/Audio/Plug-Ins/CLAP"
        version = "0.1.0"
    }
}