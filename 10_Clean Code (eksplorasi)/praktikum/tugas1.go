package main

// User adalah tipe data untuk merepresentasikan pengguna dengan ID, Username, dan Password.
// Menggunakan string untuk tipe data Username dan Password
type User struct {
	ID       int
	Username string
	Password string
}

type UserService struct {
	users []User
}

// GetAllUsers mengembalikan salinan seluruh daftar pengguna.
func (u *UserService) GetAllUsers() []User {
	// Buat salinan daftar pengguna agar tidak dapat diubah.
	usersCopy := make([]User, len(u.users))
	copy(usersCopy, u.users)
	return usersCopy
}

// GetUserByID mencari pengguna berdasarkan ID dan mengembalikan pengguna serta status apakah pengguna ditemukan.
func (u *UserService) GetUserByID(id int) (User, bool) {
	for _, r := range u.users {
		if id == r.ID {
			return r, true
		}
	}
	// Mengembalikan pengguna default
	return User{}, false
}
