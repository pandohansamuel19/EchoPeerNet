
// Representasi dari data request dan response, dan tentunya dibuat dalam struct dengan field sesuai kebutuhan request response

// Utilities:
// 1. create_request.go: Representasi data request ketika mau insert data ke database
// 2. update_request.go: Representasi data request ketika mau update data di database
// 3. response.go: Representasi data response yang diberikan ketika selesai memproses aksi dari database
// 4. web_reponse.go: Representasi data standard response yang diberikan ke client ketika selesai memproses request

// Beda antara response.go dan web_response.go adalah dalam penggunaannya dimana response.go ketika selesai memproses aksi DB , sedangkan web_response.go ketika selesai memproses semua request

package web