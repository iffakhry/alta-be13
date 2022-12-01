# Deployment

## GCP
* Buat Instance (Compute Engine): https://cloud.google.com/compute/docs/instances/create-start-instance
* Generate SSH untuk connect ke Instance GCP: https://cloud.google.com/compute/docs/connect/create-ssh-keys#windows-10-or-later
* Menambahkan SSH ke Instance: https://cloud.google.com/compute/docs/connect/add-ssh-keys
* Setup Firewall Rule: https://cloud.google.com/vpc/docs/using-firewalls

Note saat melakukan generate SSH di windows, jika nama user windows nya ada spasi, maka tambahkan tanda "" di url tempat penyimpanan file ssh nya.
contoh "c:/Users/Fakhry Firdaus/.ssh/sshgcpalta"

## Connect to Compute Engine / EC2
```bash
ssh -i </directory/namafilessh> <username-server>@<public-ipv4>

# example:
[GCP] ssh -i ~/.ssh/sshgcp fakhry@18.0.1.2

[AWS] ssh -i ~/.ssh/file.pem ubuntu@18.0.1.2
```
Note: untuk nama username di GCP, biasanya sesuai dengan nama alamat email kita. misal fakhry@gmail.com
maka username instance kita adalah fakhry.
Untuk AWS, jika anda menggunakan ubuntu. maka username nya adalah ubuntu.

Jika ada error saat akses server via ssh (permission denied). coba ubah hak akses file .pem nya
```bash
chmod 400 namafile.pem
```

## Update dan Upgrade Instance
Saat pertama kali masuk ke Instance server, jalankan perintah berikut (Cukup jalankan sekali saja.)
```bash
sudo apt-get update
sudo apt-get upgrade
```

## Transfer file ke server menggunakan SCP
Jika ingin melakukan transfer 1 folder, tambahkan `-r`

Untuk AWS, silakan ganti `namafilessh` dengan lokasi tempat file `.pem` berada.

Jalankan perintah berikut di terminal local komputer kita.
```bash
scp -i </direktori/namafilessh> </direktori/nama-file-yang-ditransfer> <username-server>@<public-ipv4>:/home/<username>

# example
scp -i ~/.ssh/sshgcpalta main.go fakhry@18.10.1.20:/home/fakhry

scp -i ~/.ssh/sshgcpalta -r ./15-deployment fakhry@18.10.1.20:/home/fakhry
```

## Install GO in server
```bash
wget https://go.dev/dl/go1.19.3.linux-amd64.tar.gz

sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz

nano .profile 
#atau 
nano .bashrc

# tambahkan baris dibawah ini
export PATH=$PATH:/usr/local/go/bin

# reload
source .bashrc
#atau
source .profile

#cek apakah sudah keinstall
go version
```

## Install MySQL client di server
```bash
sudo apt install mysql-client-core-8.0
```
#### Cara Connect ke DB MySQL di Server
[GCP] Klik instance DB yang ingin digunakan. db host bisa dilihat di dashboard SQL instance tsb

[AWS] Klik instance DB yang ingin digunakan. db host bisa dilihat dibagian DB Endpoint.
```bash
mysql -h <db-host> -P <db-port> -u <username> -p

# ketika anda menekan enter, maka terminal akan meminta input password. anda harus memasukkan password db anda.
```

# Docker

## Install Docker di Ubuntu Server
```bash
sudo apt install docker.io
```

## Jika terjadi error (Permission Denied) saat Run Docker di Ubuntu
```bash
sudo usermod -a -G docker ubuntu

or

sudo chmod 777 /var/run/docker.sock
```

## Build Docker Image
```bash
docker build -t <image-name>:<tag> .

# example:
docker build -t api-images:latest .
```

## Show Docker Images
```bash
docker images

docker images list
```

## Delete Docker Image
```bash
docker rmi <image-id>
#or
docker rmi <image-name>

# example:
docker rmi api-images
```

## Create Docker Container from Image
Note: -d digunakan agar app berjalan di background
```bash
docker run -d
-p <host-port>:<container-port>
-e <env-name>=<env-value>
-e <env-name>=<env-value>
-v <host-volume>:<container-volume>
--name <container-name> <image-name>:<tag>

# example:
docker run -p 80:8000 --name apiContainer api-images:latest
```

## Show Container
```bash
# melihat container yang sedang running
docker ps

# melihat seluruh container, termasuk yang sedang stop
docker ps -a
```

## Start/Stop Container
```bash
docker stop <container-name>

docker start <container-name>
```

## Remove Container
```bash
docker rm <container-name>

docker rm <container-id>

# example
docker rm apiContainer
```

## Docker Logs Container
melihat logs dari container. berguna untuk tracing ketika terjadi error di aplikasi/container.
```bash
docker logs <container-name>
```

## Login Docker Hub dan Push Image ke Docker Hub
```bash
docker login -u <username>

docker build -t <username-dockerhub>/<image-name>:<tag> .

docker push <username-dockerhub>/<image-name>
```

## Pull Image dari Container Registry
```bash
docker pull <image-name>
```

# Notes CICD

## Clone Repository
* Lakukan clone github repo di server.
```bash
git clone <url-git-repo>

# example
git clone https://github.com/iffakhry/alta-be11-cicd.git
```

## Setup Github Action
* Buat folder `.github/workflows`
* Buat file `deploy.yml`
* Push ke branch `main`

referensi: https://github.com/appleboy/ssh-action

```bash
name: Deploy to GCP

on:
  push:
    branches:
      - main

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - name: Connect to server using SSH
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.HOST }}
          username: ${{ secrets.USERNAME }}
          key: ${{ secrets.KEY }}
          port: ${{ secrets.PORT }}
          script: |
            cd /home/fakhry/alta-be13-gcp
            git pull origin main
            docker stop beContainer
            docker rm beContainer
            docker build -t be-images:latest .
            docker run -d -p 80:8080 -e SERVER_PORT=${{ secrets.SERVER_PORT }} -e DB_USERNAME=${{ secrets.DB_USERNAME }} -e DB_PASSWORD=${{ secrets.DB_PASSWORD }} -e DB_HOST=${{ secrets.DB_HOST }} -e DB_PORT=${{ secrets.DB_PORT }} -e DB_NAME=${{ secrets.DB_NAME }} --name beContainer be-images:latest
```

## Konfigurasi SECRET variables
```bash
HOST --> IP Public v4 Server
KEY --> isi ssh private key(gcp) atau file .pem(aws)
PORT --> 22
USERNAME --> ubuntu (aws) atau username (gcp)

DB_HOST --> (rds)db endpoint, (gcp)ip database server
DB_NAME --> database name
DB_USERNAME --> username database. by default root(gcp) atau admin(aws)
DB_PASSWORD
DB_PORT --> 3306

SERVER_PORT 
```
