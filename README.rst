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
