#!/bin/bash

# Mengambil argumen pertama (nama folder utama) dan waktu saat ini
# folder_name="$1 at $(date '+%c')"
# mkdir "$folder_name"

# Mengambil argumen pertama (nama folder utama) dan waktu saat ini dengan zona waktu WIB
folder_name="$1 at $(date '+%a %b %d %T %Z %Y')"

# Membuat direktori
mkdir "$folder_name"

# Membuat struktur folder yang sesuai
mkdir -p "$folder_name/about_me/personal"
mkdir -p "$folder_name/about_me/professional"
mkdir -p "$folder_name/my_friends"
mkdir -p "$folder_name/my_system_info"

# Mengisi file facebook.txt dan linkedin.txt dengan URL
echo "https://www.facebook.com/$2" > "$folder_name/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$folder_name/about_me/professional/linkedin.txt"

# Menggunakan curl untuk mengunduh daftar teman dari URL dan menyimpannya dalam file list_of_my_friends.txt
curl -s "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" -o "$folder_name/my_friends/list_of_my_friends.txt"

# Mengisi file about_this_laptop.txt dengan informasi uname dan nama user
echo "My username: $(whoami)" > "$folder_name/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$folder_name/my_system_info/about_this_laptop.txt"

# Mengisi file internet_connection.txt dengan hasil ping ke google.com sebanyak 3 kali
echo "Connection to google:" > "$folder_name/my_system_info/internet_connection.txt"
ping -c 3 google.com >> "$folder_name/my_system_info/internet_connection.txt"

# Menampilkan struktur folder
tree "$folder_name"

# Menampilkan pesan berhasil
echo "Folder dan file telah dibuat di $folder_name"
