# Changelog

## 0.0.307-4 (2020/08/07)

* Discord 0.0.307
* Portapps 2.5.0

## 0.0.306-3 (2020/03/04)

* Discord 0.0.306

## 0.0.305-2 (2019/12/15)

* Portapps 1.31.0
* Add cleanup config

## 0.0.305-1 (2019/03/13)

* Discord 0.0.305
* Portapps 1.20.2

## 0.0.304-27 (2019/01/18)

* Discord 0.0.304

## 0.0.302-26 (2019/01/15)

* Discord 0.0.302

## 0.0.301-25 (2019/01/10)

* Remove shortcut on close (Issue #13)

## 0.0.301-24 (2018/10/24)

* Fix Windows desktop notifications
* Go 1.11
* Use [go mod](https://golang.org/cmd/go/#hdr-Module_maintenance) instead of `dep`

## 0.0.301-23 (2018/05/23)

* Discord : 0.0.301
* Update bootstrap script for multiInstance

## 0.0.300-22 (2018/02/11)

* Move ia32/x64 to win32/win64 for arch def
* Remove nupkg file

## 0.0.300-21 (2018/02/11)

* Disable host updates (Issue #10)

## 0.0.300-20 (2018/02/08)

* Ability to pass custom args to the portable process

## 0.0.300-19 (2018/01/13)

* Rebuild electron.asar to push data directly in `data` subfolder
* No need to override USERPROFILE environment variable anymore
* Allow multiple instances (Issue #6)

> :warning: **UPGRADE NOTES**
> * Move everything in `data\AppData\Roaming\discord\*` to `data` folder and remove folder `data\AppData`.

## 0.0.300-18 (2018/01/12)

* Discord : 0.0.300

## 0.0.299-17 (2017/12/14)

* Discord : 0.0.299

## 0.0.298-16 (2017/11/21)

* Move app to a subfolder
* Reduce dependencies to avoid heuristic detection
* Add UPX compression

> :warning: **UPGRADE NOTES**
> * Delete all files and folders in root folder except `data` folder.

## 0.0.298-15 (2017/11/19)

* Move repository to [Portapps](https://github.com/portapps) org

## 0.0.298-14 (2017/11/11)

* New logger
* Switch to [Golang Dep](https://github.com/golang/dep) as dependency manager
* Go 1.9.1

## 0.0.298-13 (2017/09/04)

* Override USERPROFILE env var instead of using symlink to APPDATA to store data
* Do not migrate old data folder from APPDATA
* Reduce dependencies and system calls to avoid heuristic detection

> :warning: **UPGRADE NOTES**
> * Move the content of `data\*` in `data\AppData\Roaming\discord\`
> * Remove symlink `%APPDATA%\discord`

## 0.0.298-12 (2017/08/26)

* Discord voice error (Issue #4)
* Heuristic detection (Issue #3)
* Go 1.9
* Add support guidelines

## 0.0.298-10 (2017/08/09)

* Discord : 0.0.298

## 0.0.297-7 (2017/07/20)

* Admin privileges not required for Setup (Issue #1)
* Small refactoring

## 0.0.297-4 (2017/05/14)

* Provide the nupkg file in the release

## 0.0.297-2 (2017/05/14)

* Move log
* Update go libs

## 0.0.297-1 (2017/05/14)

* Initial version
