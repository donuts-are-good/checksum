![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# checksum

checksum and file comparison tool

## usage
here's how to use checksum:

```bash
checksum <file_or_directory_path> [file_or_directory_path]
```

if you only give it one path, `checksum` will calculate the `md5`, `sha1`, `sha-256`, and `sha-512` checksums of the file or all files in the directory recursively.

if two paths are provided, checksum will calculate the checksums for each path and compare them. both paths should be either files or directories.

## examples

### calculating checksums for a file

consider you want to calculate checksums for a file `myfile.txt`.

#### command:

```bash
checksum myfile.txt
```

#### output:

```bash
calculating checksums for myfile.txt...
md5      5eb63bbbe01eeed093cb22bb8f5acdc3
sha-1    2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
sha-256  6dcd4ce23d88e2ee95838f7b014b6284f4a620e5f0a5f5f7170bcea25de41d2a
sha-512  9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3aef53fb3cf4dbaaa6
```

### calculating checksums for a directory

consider you want to calculate checksums for all files within a directory `/var/logs`.

#### command:

```bash
checksum /var/logs
```

#### output:

```
calculating checksums for /var/logs...
md5      68b329da9893e34099c7d8ad5cb9c940
sha-1    5ba93c9db0cff93f52b521d7420e43f6eda2784f
sha-256  e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
sha-512  cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e
```

### comparing two files

consider you want to compare two files file1.txt and file2.txt.

#### command:

```bash
checksum file1.txt file2.txt
```

#### output:

```
comparing checksums for file1.txt and file2.txt...
checksums do not match
```

### comparing two directories
consider you want to compare two directories `/var/logs` and `/var/backup_logs`.

#### command:

```bash
checksum /var/logs /var/backup_logs
```

#### output:

```
comparing checksums for /var/logs and /var/backup_logs...
checksums match
```

## license

MIT License 2023 donuts-are-good, for more info see license.md
