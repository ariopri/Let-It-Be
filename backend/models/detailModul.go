package models

type Video struct {
	VideoID    string `json:"videoid"`
	Judul      string `json:"judul"`
	Thumbnail  string `json:"thumbnail"`
	Keterangan string `json:"keterangan"`
}

type Kelass struct {
	ID          int     `json:"id"`
	Class       string  `json:"class"`
	Description string  `json:"description"`
	Video       []Video `json:"video"`
}

type Pelajaran struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Kelas []Kelass `json:"kelas"`
}

type Detail struct {
	Pelajaran []Pelajaran `json:"pelajaran"`
}

var DetailModul = []Detail{
	{
		Pelajaran: []Pelajaran{
			{
				ID:   1,
				Name: "matematika",
				Kelas: []Kelass{
					{
						ID:          1,
						Class:       "matematika1",
						Description: "Belajar Matematika Bersama Matematika 1",
						Video: []Video{
							{
								VideoID:    "caNhv2VNJ40",
								Judul:      "Belajar Materi Aljabar Linear",
								Thumbnail:  "https://i.ytimg.com/vi/caNhv2VNJ40/hqdefault.jpg",
								Keterangan: "Materi Aljabar Linear Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "zNzuGfXldU0",
								Judul:      "Belajar Materi Ekspoential dan Logaritma",
								Thumbnail:  "https://i.ytimg.com/vi/zNzuGfXldU0/hqdefault.jpg",
								Keterangan: "Materi Ekspoential dan Logaritma Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "ERCGYJiw2l4",
								Judul:      "Belajar Materi Trigonometri",
								Thumbnail:  "https://i.ytimg.com/vi/ERCGYJiw2l4/hqdefault.jpg",
								Keterangan: "Materi Trigonometri Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "HXcoifPIvc0",
								Judul:      "Belajar Materi Integral",
								Thumbnail:  "https://i.ytimg.com/vi/HXcoifPIvc0/hqdefault.jpg",
								Keterangan: "Materi Integral Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "HXcoifPIvc0",
								Judul:      "Belajar Materi Persamaan Diferensial",
								Thumbnail:  "https://i.ytimg.com/vi/HXcoifPIvc0/hqdefault.jpg",
								Keterangan: "Materi Persamaan Diferensial Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          2,
						Class:       "matematika2",
						Description: "Belajar Matematika Bersama Matematika 2",
						Video: []Video{
							{
								VideoID:    "e8uAx8ULEhY",
								Judul:      "Belajar Materi Pembuktian Deret Bilangan",
								Thumbnail:  "https://i.ytimg.com/vi/e8uAx8ULEhY/hqdefault.jpg",
								Keterangan: "Materi Pembuktian Deret Bilangan Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "ZfxPFHBEmvg",
								Judul:      "Belajar Materi Pembuktian Keterbagian",
								Thumbnail:  "https://i.ytimg.com/vi/ZfxPFHBEmvg/hqdefault.jpg",
								Keterangan: "Materi Pembuktian Keterbagian Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "15uY57754tQ",
								Judul:      "Belajar Materi Pertidaksamaan",
								Thumbnail:  "https://i.ytimg.com/vi/15uY57754tQ/hqdefault.jpg",
								Keterangan: "Materi Pertidaksamaan Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "c5NBcC2xMx8",
								Judul:      "Belajar Materi Sistem Program Linear",
								Thumbnail:  "https://i.ytimg.com/vi/c5NBcC2xMx8/hqdefault.jpg",
								Keterangan: "Materi Sistem Program Linear Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          3,
						Class:       "matematika3",
						Description: "Belajar Matematika Bersama Matematika 3",
						Video: []Video{
							{
								VideoID:    "HqssIxIg7T4",
								Judul:      "Belajar Materi Matriks",
								Thumbnail:  "https://i.ytimg.com/vi/HqssIxIg7T4/hqdefault.jpg",
								Keterangan: "Materi Matriks Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "2mIQgRSsGyc",
								Judul:      "Belajar Materi Determinan",
								Thumbnail:  "https://i.ytimg.com/vi/2mIQgRSsGyc/hqdefault.jpg",
								Keterangan: "Materi Determinan Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "KWNkmpEpVBY",
								Judul:      "Belajar Materi Minor, Kofaktor dan Adjoin Matriks",
								Thumbnail:  "https://i.ytimg.com/vi/KWNkmpEpVBY/hqdefault.jpg",
								Keterangan: "Materi Minor, Kofaktor dan Adjoin Matriks Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          4,
						Class:       "matematika4",
						Description: "Belajar Matematika Bersama Matematika 4",
						Video: []Video{
							{
								VideoID:    "CzrlEdZdhB0",
								Judul:      "Belajar Materi Invers Matriks",
								Thumbnail:  "https://i.ytimg.com/vi/CzrlEdZdhB0/hqdefault.jpg",
								Keterangan: "Materi Invers Matriks Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "iojzS-0DI_g",
								Judul:      "Belajar Materi Translasi",
								Thumbnail:  "https://i.ytimg.com/vi/iojzS-0DI_g/hqdefault.jpg",
								Keterangan: "Materi Translasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "lBFnXBvfvnY",
								Judul:      "Belajar Materi Rotasi",
								Thumbnail:  "https://i.ytimg.com/vi/lBFnXBvfvnY/hqdefault.jpg",
								Keterangan: "Materi Rotasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "uMxYXmfCLgA",
								Judul:      "Belajar Materi Dilatasi",
								Thumbnail:  "https://i.ytimg.com/vi/uMxYXmfCLgA/hqdefault.jpg",
								Keterangan: "Materi Dilatasi Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          5,
						Class:       "matematika5",
						Description: "Belajar Matematika Bersama Matematika 5",
						Video: []Video{
							{
								VideoID:    "KBX6B1-ULF0",
								Judul:      "Belajar Materi Transformasi Matriks",
								Thumbnail:  "https://i.ytimg.com/vi/KBX6B1-ULF0/hqdefault.jpg",
								Keterangan: "Materi Transformasi Matriks Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "r-4fG7LB6Hw",
								Judul:      "Belajar Materi Deret Aritmatika",
								Thumbnail:  "https://i.ytimg.com/vi/r-4fG7LB6Hw/hqdefault.jpg",
								Keterangan: "Materi Deret Aritmatika Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "",
								Judul:      "Belajar Materi Baris Geometri",
								Thumbnail:  "https://i.ytimg.com/vi/6ZQYQZ1ZQ0A/hqdefault.jpg",
								Keterangan: "Materi Baris Geometri Untuk Persiapan SBMPTN",
							},
						},
					},
				},
			},
			{
				ID:   2,
				Name: "fisika",
				Kelas: []Kelass{
					{
						ID:          1,
						Class:       "fisika1",
						Description: "Belajar Fisika Bersama Fisika 1",
						Video: []Video{
							{
								VideoID:    "UKK6AISk8vU",
								Judul:      "Belajar Materi Momen Inersia",
								Thumbnail:  "https://i.ytimg.com/vi/UKK6AISk8vU/hqdefault.jpg",
								Keterangan: "Materi Momen Inersia Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "FiWPgbzHMGQ",
								Judul:      "Belajar Materi Momentum Sudut",
								Thumbnail:  "https://i.ytimg.com/vi/FiWPgbzHMGQ/hqdefault.jpg",
								Keterangan: "Materi Momentum Sudut Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "HDzVLyJhDnI",
								Judul:      "Belajar Materi Dinamika Rotasi",
								Thumbnail:  "https://i.ytimg.com/vi/HDzVLyJhDnI/hqdefault.jpg",
								Keterangan: "Materi Dinamika Rotasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "QD4sd2KAg1Q",
								Judul:      "Belajar Materi Energi Kinetik",
								Thumbnail:  "https://i.ytimg.com/vi/QD4sd2KAg1Q/hqdefault.jpg",
								Keterangan: "Materi Energi Kinetik Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "8mA7d0OSTzk",
								Judul:      "Belajar Materi Keseimbangan",
								Thumbnail:  "https://i.ytimg.com/vi/8mA7d0OSTzk/hqdefault.jpg",
								Keterangan: "Materi Keseimbangan Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          2,
						Class:       "fisika2",
						Description: "Belajar Fisika Bersama Fisika 2",
						Video: []Video{
							{
								VideoID:    "nTX_lAUJ4AQ",
								Judul:      "Belajar Matieri Fluida Statis",
								Thumbnail:  "https://i.ytimg.com/vi/nTX_lAUJ4AQ/hqdefault.jpg",
								Keterangan: "Materi Fluida Statis Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "BI9RjIICQzY",
								Judul:      "Belajar Materi Fluida Dinamis",
								Thumbnail:  "https://i.ytimg.com/vi/BI9RjIICQzY/hqdefault.jpg",
								Keterangan: "Materi Fluida Dinamis Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "U2IRS3pnJDo",
								Judul:      "Belajar Materi Elastisisitas",
								Thumbnail:  "https://i.ytimg.com/vi/U2IRS3pnJDo/hqdefault.jpg",
								Keterangan: "Materi Elastisisitas Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          3,
						Class:       "fisika3",
						Description: "Belajar Fisika Bersama Fisika 3",
						Video: []Video{
							{
								VideoID:    "QD4sd2KAg1Q",
								Judul:      "Belajar Materi Energi Kinetik",
								Thumbnail:  "https://i.ytimg.com/vi/QD4sd2KAg1Q/hqdefault.jpg",
								Keterangan: "Materi Energi Kinetik Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "8mA7d0OSTzk",
								Judul:      "Belajar Materi Energi Dinamis",
								Thumbnail:  "https://i.ytimg.com/vi/8mA7d0OSTzk/hqdefault.jpg",
								Keterangan: "Materi Energi Dinamis Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "nTX_lAUJ4AQ",
								Judul:      "Belajar Materi Fluida Statis",
								Thumbnail:  "https://i.ytimg.com/vi/nTX_lAUJ4AQ/hqdefault.jpg",
								Keterangan: "Materi Fluida Statis Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "BI9RjIICQzY",
								Judul:      "Belajar Materi Fluida Dinamis",
								Thumbnail:  "https://i.ytimg.com/vi/BI9RjIICQzY/hqdefault.jpg",
								Keterangan: "Materi Fluida Dinamis Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          4,
						Class:       "fisika4",
						Description: "Belajar Fisika Bersama Fisika 4",
						Video: []Video{
							{
								VideoID:    "UKK6AIS",
								Judul:      "Belajar Materi Momen Inersia",
								Thumbnail:  "https://i.ytimg.com/vi/UKK6AIS/hqdefault.jpg",
								Keterangan: "Materi Momen Inersia Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "FiWPgbzHMGQ",
								Judul:      "Belajar Materi Momentum Sudut",
								Thumbnail:  "https://i.ytimg.com/vi/FiWPgbzHMGQ/hqdefault.jpg",
								Keterangan: "Materi Momentum Sudut Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "HDzVLyJhDnI",
								Judul:      "Belajar Materi Dinamika Rotasi",
								Thumbnail:  "https://i.ytimg.com/vi/HDzVLyJhDnI/hqdefault.jpg",
								Keterangan: "Materi Dinamika Rotasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "QD4sd2KAg1Q",
								Judul:      "Belajar Materi Energi Kinetik",
								Thumbnail:  "https://i.ytimg.com/vi/QD4sd2KAg1Q/hqdefault.jpg",
								Keterangan: "Materi Energi Kinetik Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          5,
						Class:       "fisika5",
						Description: "Belajar Fisika Bersama Fisika 5",
						Video: []Video{
							{
								VideoID:    "UKK6AIS",
								Judul:      "Belajar Materi Momen Inersia",
								Thumbnail:  "https://i.ytimg.com/vi/UKK6AIS/hqdefault.jpg",
								Keterangan: "Materi Momen Inersia Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "FiWPgbzHMGQ",
								Judul:      "Belajar Materi Momentum Sudut",
								Thumbnail:  "https://i.ytimg.com/vi/FiWPgbzHMGQ/hqdefault.jpg",
								Keterangan: "Materi Momentum Sudut Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "HDzVLyJhDnI",
								Judul:      "Belajar Materi Dinamika Rotasi",
								Thumbnail:  "https://i.ytimg.com/vi/HDzVLyJhDnI/hqdefault.jpg",
								Keterangan: "Materi Dinamika Rotasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "QD4sd2KAg1Q",
								Judul:      "Belajar Materi Energi Kinetik",
								Thumbnail:  "https://i.ytimg.com/vi/QD4sd2KAg1Q/hqdefault.jpg",
								Keterangan: "Materi Energi Kinetik Untuk Persiapan SBMPTN",
							},
						},
					},
				},
			},
		},
	},
}
