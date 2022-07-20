<h1 align="center">
 Lectronic App
</h1>

<p align="center">
    <img src="https://res.cloudinary.com/dlyp1s66j/image/upload/v1658298984/Logo_zjryrz.png" width="200px" alt="Lectronic" />

</p>

<p align="center">
    <a href="/#" target="blank">View Demo</a>
  · <a href="https://github.com/Irsad99/Lectronic-Apps/issues">Report Bug</a>
  · <a href="https://github.com/Irsad99/Lectronic-Apps/pulls">Request Feature</a>
</p>


## Built with

**Backend:**

![golang](https://img.shields.io/badge/Go-100000?style=for-the-badge&logo=Go&logoColor=white&labelColor=51DEF0&color=51DEF0)&nbsp;
![postgresql](https://img.shields.io/badge/PostgreSQL-100000?style=for-the-badge&logo=PostgreSQL&logoColor=white&labelColor=3A7373&color=384A5F)&nbsp;
![jwt](https://img.shields.io/badge/JWT-100000?style=for-the-badge&logo=JSONWebTokens&logoColor=white&labelColor=000000&color=000000)&nbsp;

**Deployed On:**

![heroku](https://img.shields.io/badge/heroku-100000?style=for-the-badge&logo=Heroku&logoColor=white&labelColor=3C8932&color=3C8932)&nbsp;

## Description about project
Lectronic App is an application for selling electronic devices. users can order products and make payments through the third application, namely Midtrans. and users can manage profiles. For the admin role, you can manage product and user. This application is built with Golang using the gorilla/mux package for routing. The databases used in this application are PostgreSQL

## Installation Steps

1. Clone the repository

   ```bash
    https://github.com/Irsad99/Lectronic-Apps
    ```

2. Install dependencies

   ```bash
   go mod tidy
   ```

3. Environtment Variable

   #### Variable site
   ```bash
   BASE_URL=https://myelectronic.herokuapp.com
   APP_PORT=8080
   JWT_KEYS=my - lectronic - key
   ```
   
   #### Database
   ```bash
   DB_HOST=
   DB_NAME=
   DB_USER=
   DB_PASS=
   ```
   #### Config SMTP Gmail
   ```bash
   CONFIG_SMTP_HOST=smtp.gmail.com
   CONFIG_SMTP_PORT=587
   CONFIG_AUTH_EMAIL=youremail@gmail.com
   CONFIG_AUTH_PASSWORD=yourpassword
   ```
   #### Cloudinary Config
   ```bash
   CLOUD_NAME=
   CLOUD_KEY=
   CLOUD_SECRET=
   ```
   #### Midtrans Config
   ```bash
   MIDTRANS_CLIENT=
   MIDTRANS_SERVER=
   ```

4. Run the app

   ```bash
   go run main.go server
   ```

## Frontend Repository
Frontend Lectronic [here](https://github.com/Irsad99/FE-Lectronics-App)

### Contributors
<a href = "https://github.com/Irsad99/FE-Lectronics-App/graphs/contributors">
  <img src="https://avatars.githubusercontent.com/u/80185253?s=60&amp;v=4" class="avatar avatar-user" alt="calvinrahmat" width="38" height="38">
  <img src="https://avatars.githubusercontent.com/u/76877980?s=60&amp;v=4" class="avatar avatar-user" alt="Gustiana882" width="38" height="38">
  <img src="https://avatars.githubusercontent.com/u/38394430?s=60&amp;v=4" class="avatar avatar-user" alt="firyal-salsa" width="38" height="38">
</a>

<hr>
<p align="center">
Developed with ❤️ in Indonesia 	🇮🇩
</p>