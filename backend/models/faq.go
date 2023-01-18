package models

type FAQ struct {
	Tanya string `json:"tanya"`
	Jawab string `json:"jawab"`
	Judul string `json:"judul,omitempty"`
	ID    string `json:"id"`
}

var FAQList = []FAQ{
	{
		Tanya: "Apa itu Let It Be?",
		Jawab: "Let It Be adalah sebuah aplikasi yang dibuat untuk memudahkan para siswa dalam melakukan pembelajaran secara online guna membantu dalam persiapan Ujian Masuk Perguruan Tinggi.",
		Judul: "Segala pertanyaan yang kerap ditanyakan kepada kami, kami rangkum di dalam komponen FAQ, agar mempermudah pekerjaan tim kami dan menambah pemahaman anda tentang platform Let It Be ini.",
		ID:    "1",
	},
	{
		Tanya: "Bagaimana cara menggunakan Let It Be",
		Jawab: "Untuk menggunakan Let It Be, kamu harus mendaftar terlebih dahulu. Setelah mendaftar, kamu bisa langsung login dan mulai belajar.",
		ID:    "2",
	},
	{
		Tanya: "Kursus apa saja yang ada di Let It Be?",
		Jawab: "Kursus yang ada di Let It Be adalah kursus untuk persiapan Ujian Masuk Perguruan Tinggi",
		ID:    "3",
	},
	{
		Tanya: "Bagaimana cara mendaftar di Let It Be?",
		Jawab: "Kamu bisa mendaftar di Let It Be dengan mengisi form pendaftaran yang ada di website Let It Be.",
		ID:    "4",
	},
}
