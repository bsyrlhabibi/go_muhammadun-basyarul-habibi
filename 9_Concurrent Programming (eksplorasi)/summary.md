
# Summary Concurrent Programing

1. Sequential, Paralel, dan Concurrent: Pemrograman sequential melibatkan pelaksanaan tugas secara berurutan. Pemrograman paralel memungkinkan tugas berjalan bersamaan, tetapi memerlukan multiple core. Sementara itu, pemrograman concurrent memungkinkan tugas-tugas berjalan bersamaan dan independen.

2. Goroutine, Channels dan Select: Goroutine adalah fungsi yang berjalan concurrent dan dapat ditentukan untuk berjalan bersamaan. Channels digunakan untuk komunikasi antar goroutine, sedangkan Select digunakan untuk mengontrol aliran komunikasi antara mereka.

3. Multiple Goroutine dan Synchronization: Dalam konteks multiple goroutine, channels memfasilitasi komunikasi antar tugas concurrent, termasuk buffer channel yang dapat menyimpan data dalam kapasitas tertentu. Selain itu, untuk mengatasi masalah race condition yang dapat terjadi ketika tugas-tugas bersaing untuk mengakses dan memodifikasi data tanpa sinkronisasi, terdapat konsep seperti waitgroup untuk koordinasi antar goroutine dan mutex untuk mengunci dan membuka kunci sumber daya bersama guna menghindari race condition. Dengan pemahaman ini, program concurrent dapat dirancang dengan efisien dan andal.