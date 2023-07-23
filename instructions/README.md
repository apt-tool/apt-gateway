# FTP

![](https://img.shields.io/badge/Language-Golang-blue)
![](https://img.shields.io/badge/Version-v0.1-blue)

This is file transfer part of our system. In this app we upload our attack scripts,
we execute attack scripts, and we download the results. This system provides an interface
in order to manage scripts and logs files from core app.

## Environment Variables

- ```HTTP_PORT``` which is the system port
- ```ACCESS_KEY``` which is used to block download access for files
- ```PRIVATE_KEY``` which is used to give execute script system call enable only for core
