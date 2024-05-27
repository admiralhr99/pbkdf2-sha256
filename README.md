# pbkdf2-sha256
This repository contains a Go program that concurrently hashes passwords from a file using PBKDF2 with SHA-256 and compares them to a given hash.

It is useful in CTF competitions when you find a hash like this and want to perform a rainbow table attack on the provided password list file.

You can install it with the following command:
```bash
go install https://github.com/admiralhr99/pbkdf2-sha256@latest
```

Please pay attention to the hardcoded values for the name of the password list file, salt, iterations, and expected hash.

