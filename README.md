# backend-social-media
Backend written in Go using gin gonic web server.
# LICENSE
Shield: [![CC BY 4.0][cc-by-shield]][cc-by]

This work is licensed under a
[Creative Commons Attribution 4.0 International License][cc-by].

[![CC BY 4.0][cc-by-image]][cc-by]

[cc-by]: http://creativecommons.org/licenses/by/4.0/
[cc-by-image]: https://i.creativecommons.org/l/by/4.0/88x31.png
[cc-by-shield]: https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg

# Requirements
- ## Go 1.19
- ## Nodejs 18+
# HOW TO RUN 1
- git clone https://github.com/conspirey/backend-social-media
- git clone https://github.com/conspirey/frontend-conspirey
- Rename frontend-conspirey to frontend
- create .env file where backend is located
- .env example 
    - MONGO={MONGO_CONNECTION_STRING}
    - ENCRYPTION_KEY={32_key_long_string}
    - ADMIN_KEY={RANDOM_STRING}
- Debugging mode runs on http://localhost:3100
  - go run main.go
  - or
  - go build main.go
  - ./main
- Release mode runs on https://localhost:3100 
  - go run main.go release
  - or
  - go build main.go
  - ./main release
- Changing ssl certificates to your own website
  - Generate ssl certificate from letsencrypt
  - change cert.pem to your own certificate
  - change keys.pem to your own private key
- Compiling frontend
  - in frontend folder do
  - npm run build
### Following these steps you can run this website succesfully
# README planned updates
- ADD releases to github actions for both frontend and backend
