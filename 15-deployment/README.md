# Deployment

## GCP
* Buat Instance (Compute Engine): https://cloud.google.com/compute/docs/instances/create-start-instance
* Generate SSH untuk connect ke Instance GCP: https://cloud.google.com/compute/docs/connect/create-ssh-keys#windows-10-or-later
* Menambahkan SSH ke Instance: https://cloud.google.com/compute/docs/connect/add-ssh-keys

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

## Transfer file ke server menggunakan SCP
Jika ingin melakukan transfer 1 folder, tambahkan `-r`

Untuk AWS, silakan ganti `namafilessh` dengan lokasi tempat file `.pem` berada.
```bash
scp -i </direktori/namafilessh> </direktori/nama-file-yang-ditransfer> <username-server>@<public-ipv4>:/home/<username>

# example
scp -i ~/.ssh/sshgcpalta main.go fakhry@18.10.1.20:/home/fakhry

scp -i ~/.ssh/sshgcpalta -r ./15-deployment fakhry@18.10.1.20:/home/fakhry
```