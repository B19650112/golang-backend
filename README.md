Installation Software
---------------------
- Install Go,     link download: https://golang.org/dl/
- Install Git,    link download: https://git-scm.com/downloads
- Install VS Code link download: https://code.visualstudio.com/Download


Preparation before run golang-backend
-------------------------------------
We assume already setup Golang environment and XAMPP version 7.3.1 or latest already installed<br/>
Sample environment Golang like this:<br/>
![Untitled](https://user-images.githubusercontent.com/49906150/57184891-ab7ddb00-6eec-11e9-978a-e694700a016b.png) <br/>

Look at the GOBIN, GOPATH and GOROOT settings, if have not already exist, press the New button to created<br/>

Create workspace folder :
-------------------------
- md C:\GoWeb        --> we asssume
- cd C:\GoWeb
- md bin
- md pkg
- md src
- cd src
- md golang-backend
- cd golang-backend
- copy all files above (include all folder)

- Open Visual Studio Code
- Open Folder golang-backend
- Select View, select Terminal

The screen below will appear:

Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

PS C:\Goweb\src\golang-backend


Install Driver Mysql and Gin Gonic
----------------------------------
Type from Terminal:
- go get -u github.com/go-sql-driver/mysql
- go get -u github.com/gin-gonic/gin


Restore Database
----------------
- Run XAMPP as administrator <br/>
  From windows panel (select XAMPP Control Panel, click right, select more, select Run as administrator)

- Click button Start on Mysql and Apache
- Click button Shell

The screen below will appear:

Setting environment for using XAMPP for Windows.<br/>
DELL@DESKTOP-36P79TD c:\xampp<br/>
#<br/>

- on position #, type: mysql -u root -p <br/>
                 Enter password:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;skip <br/>

- on position MariaDB [(none)]> type: create database mysqldata;
- on position MariaDB [(none)]> type: exit

- on position #, type: mysql -u root -p mysqldata < c:\Goweb\src\golang-backend\mysqldata.sql <br/>
                 Enter password:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;skip <br/>
                 type: exit <br/>


Build and Run Program
---------------------
Run Command Prompt, type :
- cd C:\Goweb\src\golang-backend
- go build
- golang-backend.exe
<br/>
<br/>
<br/>
&nbsp;&nbsp;&nbsp;i'm sorry, my english is not good.
