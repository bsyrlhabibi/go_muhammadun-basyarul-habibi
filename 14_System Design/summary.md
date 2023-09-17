
# Summary System Design

1. Representasi Sistem: Desain sistem sering kali menggunakan simbol-simbol seperti Flowchart, Use Case, ERD, dan HLA (High-Level Architecture) untuk menggambarkan informasi terkait sistem. HLA memberikan pandangan menyeluruh tentang sistem, mencakup teknologi dan bahasa pemrograman yang digunakan. Distributed system menjadi penting ketika transaksi tinggi, dan ini melibatkan beberapa server dengan karakteristik seperti skalabilitas, keandalan, ketersediaan, efisiensi, dan kemudahan pengelolaan.

2. Manajemen Sistem Pihak Ketiga: Job/Work Queue digunakan untuk mengelola sistem pihak ketiga, dengan load balancing menjadi faktor penting dalam sistem terdistribusi. Microservices memecah layanan di dalam sistem untuk meningkatkan fleksibilitas.

3. SQL dan NoSQL: SQL digunakan dalam basis data relasional dengan entitas yang berhubungan, sementara NoSQL lebih dinamis dan tidak terstruktur. Cache digunakan untuk penyimpanan sementara data, dan indexing membantu mengakses data dengan lebih efisien.