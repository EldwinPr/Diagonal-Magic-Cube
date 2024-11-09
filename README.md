Cara pakai:
1. buka terminal di root folder
2. run "go run main.go"
3. kalau display/cubes masih kosong "y", kalau udah ada "n". kalau mau buat ulang cube juga bisa "y".
4. buka di localhost:8080

Cara Load dan Save:
1. Save:
    1. Generate Results terlebih dahulu
    2. Copy/move konfigurasi yang diinginkan di display/cubes/
    3. Simpan Konfigurasi di tempat lain
2. Load:
    1. Run server
    2. Buka Display/Cubes
    3. Masukan konfigurasi yang disimpan ke folder yang sesuai
    4. Overwrite jika sudah ada. Jangan generate new result lagi kalau udah masukin yang baru, nanti ke overwrite
    Catatan: kalau gagal coba hapus cache, terus refresh

Singkatan folder:
- SAHC: Steepest Ascent Hill Climb
- HCWSM: Hill Climb With Sideways Move
- RRHC: Random Restart Hill Climb
- SHC: Stochastic Hill Climb
- SA: Simulated Annealing
- GA: Genetic Algorithm

catatan:
- Simulated Annealing lama karena json-nya gede
- disaranin buat disable web caching
- kalau mau hapus result bisa "go run main.go clear"