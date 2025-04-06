# Changelog

## 1.0.9188-21 (2025/04/06)

* Discord 1.0.9188
* Portapps 3.16.0

## 1.0.9156-20 (2024/08/03)

* Discord 1.0.9156 (no more 32-bit release)

## 1.0.9055-19 (2024/08/03)

* Discord 1.0.9055
* Portapps 3.11.0

## 1.0.9035-18 (2024/03/10)

* Discord 1.0.9035
* Portapps 3.10.0

## 1.0.9028-17 (2023/12/24)

* Discord 1.0.9028
* Portapps 3.9.0

## 1.0.9022-16 (2023/11/03)

* Discord 1.0.9022
* Portapps 3.8.0

## 1.0.9010-15 (2023/02/12)

* Discord 1.0.9010
* Portapps 3.7.0

## 1.0.9007-14 (2022/11/20)

* Discord 1.0.9007
* Bump github.com/kevinburke/go-bindata from 3.23.0+incompatible to 3.24.0+incompatible

## 1.0.9006-13 (2022/08/13)

* Discord 1.0.9006

## 1.0.9005-12 (2022/07/15)

* Discord 1.0.9005
* Portapps 3.6.0

## 1.0.9004-11 (2022/03/06)

* Discord 1.0.9004
* Portapps 3.4.0

## 1.0.9003-10 (2021/10/27)

* Discord 1.0.9003

## 1.0.9002-9 (2021/09/01)

* Discord 1.0.9002
* Portapps 3.3.1

## 0.0.309-8 (2021/03/09)

* Cleanup reg key (#57)
* Portapps 3.3.0

## 0.0.309-7 (2020/12/26)

* Discord 0.0.309
* Cleanup more folders (#51)
* Portapps 3.1.0

## 0.0.308-6 (2020/09/21)

* Discord 0.0.308

## 0.0.307-5 (2020/09/06)

* Portapps 2.6.0

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
