
# Summary Middleware

1. Middleware adalah entitas yang memfasilitasi komunikasi antara klien dan server dalam proses server. Contoh third-party middleware termasuk negroni, echo, interpose, dan alice. Misalnya, dalam framework echo, terdapat dua jenis middleware, yaitu Echo#pre() dan Echo#Use().

2. Salah satu jenis middleware yang umum digunakan adalah Log Middleware, yang berfungsi untuk mencatat informasi yang terjadi dalam server atau permintaan. Selain itu, ada juga metode otentikasi dasar (Basic Authentication) yang melibatkan pengiriman data username dan password melalui header permintaan API.

3. JSON Web Token (JWT), digunakan untuk meningkatkan tingkat keamanan dalam proses identifikasi dan otentikasi pengguna dalam komunikasi antara klien dan server. Dengan penggunaan JWT, sistem menjadi lebih aman dan dapat diandalkan dalam menjaga integritas data dan keamanan komunikasi.