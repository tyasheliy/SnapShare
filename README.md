this repository has moved to [gitlab](https://gitlab.com/tyasheliy/snap-share)

# SnapShare

Fullstack web application for fast file sharing.

## Table of contents

* <a href="#about-the-project">About the project</a>
* <a href="#tech-stack">Tech stack</a>
* <a href="#getting-started">Getting started</a>

## About the project

### Main idea

The main idea of ​​this project is to create a simple service for sharing user files. A distinctive feature of the service should be its ease of use, absolutely free use and at the same time ensuring sufficient security for this type of service.

### Architecture

Applicataion is built on microservice architecture. There are 4 microservices:

* **Identity service** authenticates and authorizes users.
* **Share service** handles users' files uploading and downloading.
* **SPA client** provides a web graphical interface for users.
* **Gateway** is a API gateway for proxying requests.

## Tech stack

### Identity service

* C#
* ASP NET Core
* Entity Framework Core
* PostgreSQL
* Redis

### Share service

* Go
* Echo
* Viper
* Redis

### SPA client

* JavaScript
* VueJS
* Nginx

### Gateway

* Nginx

## Getting started

### Prerequisites

You need to have only [Docker](https://www.docker.com/get-started/) to be installed on your system to run the application.

### Installation

1. Clone repository
```
git clone https://github.com/tyasheliy/SnapShare
```
2. Start Docker Engine if you haven't yet
3. Run the application
```
make run
```
