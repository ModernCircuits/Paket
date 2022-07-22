#######
 Paket
#######

|License| |Go Reference| |Test| |Coverage| |Pre-Commit|

.. warning::

   **Work in progress.**

Cross-platform meta installer generator. Primarily targets Windows
`InnoSetup` & macOS `pkgbuild/productbuild`.

*************
 Quick Start
*************

Application
===========

.. code:: hcl

   name       = "Example App"
   vendor     = "Modern Circuits"
   version    = "0.1.0"
   license    = "LICENSE.txt"
   identifier = "com.modern-circuits.example-app"

   installer "macOS" "macos-pkg" {
      artifact "App" {
         payload     = "macOS/${project.name}.app"
         destination = "/Application"
      }
   }

   installer "Windows" "innosetup" {
      artifact "App" {
         payload     = "Windows/${project.name}.exe"
         destination = "{commonpf64}/${project.vendor}"
      }
   }

Audio Plugin
============

.. code:: hcl

   name       = "Example Effect"
   vendor     = "Modern Circuits"
   version    = "0.1.0"
   license    = "LICENSE.txt"
   identifier = "com.modern-circuits.example-effect"

   installer "macOS" "macos-pkg" {
      artifact "AU" {
         name        = "Audio Unit"
         version     = "0.1.1"
         payload     = "macOS/AU/${project.name}.component"
         destination = "/Library/Audio/Plug-Ins/AU"
      }

      artifact "VST3" {
         payload     = "macOS/VST3/${project.name}.vst3"
         destination = "/Library/Audio/Plug-Ins/VST3"
      }

      artifact "CLAP" {
         payload     = "macOS/CLAP/${project.name}.clap"
         destination = "/Library/Audio/Plug-Ins/CLAP"
      }
   }

   installer "Windows" "innosetup" {
      artifact "VST3" {
         payload     = "Windows/VST3/${project.name}.vst3"
         destination = "{commoncf64}/VST3"
      }

      artifact "CLAP" {
         payload     = "Windows/CLAP/${project.name}.clap"
         destination = "{commoncf64}/CLAP"
      }
   }

.. |License| image:: https://img.shields.io/badge/License-Boost_1.0-lightblue.svg
   :target: https://github.com/ModernCircuits/Paket/blob/main/LICENSE.txt

.. |Go Reference| image:: https://pkg.go.dev/badge/github.com/moderncircuits/paket.svg
   :target: https://pkg.go.dev/github.com/moderncircuits/paket

.. |Test| image:: https://github.com/ModernCircuits/Paket/actions/workflows/test.yml/badge.svg
   :target: https://github.com/ModernCircuits/Paket/actions/workflows/test.yml

.. |Coverage| image:: https://codecov.io/gh/ModernCircuits/Paket/branch/main/graph/badge.svg?token=S8XON74JQU
   :target: https://codecov.io/gh/ModernCircuits/Paket

.. |Pre-Commit| image:: https://github.com/ModernCircuits/Paket/actions/workflows/pre-commit.yml/badge.svg
   :target: https://github.com/ModernCircuits/Paket/actions/workflows/pre-commit.yml
