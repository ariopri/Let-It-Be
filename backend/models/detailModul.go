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
			{
				ID:   3,
				Name: "kimia",
				Kelas: []Kelass{
					{
						ID:          1,
						Class:       "kimia1",
						Description: "Belajar Kimia Bersama Kimia 1",
						Video: []Video{
							{
								VideoID:    "eXijwwYrZ38",
								Judul:      "Belajar Materi Kekhasan Dan Jenis-Jenis Atom Karbon",
								Thumbnail:  "https://i.ytimg.com/vi/eXijwwYrZ38/hqdefault.jpg",
								Keterangan: "Materi Momen Kekhasan Dan Jenis-Jenis Atom Karbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "57jPesaJRzM",
								Judul:      "Belajar Materi Penggolongan Senyawa Hidrokarbon",
								Thumbnail:  "https://i.ytimg.com/vi/57jPesaJRzM/hqdefault.jpg",
								Keterangan: "Materi Penggolongan Senyawa Hidrokarbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "4sJzwo4bh60",
								Judul:      "Belajar Materri Tata Nama Senyawa Hidrokarbon",
								Thumbnail:  "https://i.ytimg.com/vi/4sJzwo4bh60/hqdefault.jpg",
								Keterangan: "Materi Materri Tata Nama Senyawa Hidrokarbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "64Ns-OU0wnQ",
								Judul:      "Belajar Materi Isomer Senyawa Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/64Ns-OU0wnQ/hqdefault.jpg",
								Keterangan: "Materi Isomer Senyawa Alkana Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "oPur-Ti5KEc",
								Judul:      "Belajar Materi Isomer Alkena dan Alkuna",
								Thumbnail:  "https://i.ytimg.com/vi/oPur-Ti5KEc/hqdefault.jpg",
								Keterangan: "Materi Isomer Alkena dan Alkuna Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AZY27IeQIOI",
								Judul:      "Belajar Materi Reaksi kimia Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/AZY27IeQIOI/hqdefault.jpg",
								Keterangan: "Materi Reaksi kimia Alkana Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          2,
						Class:       "kimia2",
						Description: "Belajar Kimia Bersama Kimia 2",
						Video: []Video{
							{
								VideoID:    "hopJg2kXiQY",
								Judul:      "Belajar Materi Sifat Alkena dan Alkuna",
								Thumbnail:  "https://i.ytimg.com/vi/hopJg2kXiQY/hqdefault.jpg",
								Keterangan: "Materi Sifat Alkena dan Alkuna Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AZY27IeQIOI",
								Judul:      "Belajar Materi Reaksi kimia Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/AZY27IeQIOI/hqdefault.jpg",
								Keterangan: "Materi Reaksi kimia Alkana Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "64Ns-OU0wnQ",
								Judul:      "Belajar Materi Isomer Senyawa Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/64Ns-OU0wnQ/hqdefault.jpg",
								Keterangan: "Materi Isomer Senyawa Alkana Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "eXijwwYrZ38",
								Judul:      "Belajar Materi Kekhasan Dan Jenis-Jenis Atom Karbon",
								Thumbnail:  "https://i.ytimg.com/vi/eXijwwYrZ38/hqdefault.jpg",
								Keterangan: "Materi Momen Kekhasan Dan Jenis-Jenis Atom Karbon Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          3,
						Class:       "kimia3",
						Description: "Belajar Kimia Bersama Kimia 3",
						Video: []Video{
							{
								VideoID:    "AZY27IeQIOI",
								Judul:      "Belajar Materi Reaksi kimia Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/AZY27IeQIOI/hqdefault.jpg",
								Keterangan: "Materi Reaksi kimia Alkana Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "hopJg2kXiQY",
								Judul:      "Belajar Materi Sifat Alkena dan Alkuna",
								Thumbnail:  "https://i.ytimg.com/vi/hopJg2kXiQY/hqdefault.jpg",
								Keterangan: "Materi Sifat Alkena dan Alkuna Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AZY27IeQIOI",
								Judul:      "Belajar Materi Reaksi kimia Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/AZY27IeQIOI/hqdefault.jpg",
								Keterangan: "Materi Reaksi kimia Alkana Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "eXijwwYrZ38",
								Judul:      "Belajar Materi Kekhasan Dan Jenis-Jenis Atom Karbon",
								Thumbnail:  "https://i.ytimg.com/vi/eXijwwYrZ38/hqdefault.jpg",
								Keterangan: "Materi Momen Kekhasan Dan Jenis-Jenis Atom Karbon Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          4,
						Class:       "kimia4",
						Description: "Belajar Kimia Bersama Kimia 4",
						Video: []Video{
							{
								VideoID:    "eXijwwYrZ38",
								Judul:      "Belajar Materi Kekhasan Dan Jenis-Jenis Atom Karbon",
								Thumbnail:  "https://i.ytimg.com/vi/eXijwwYrZ38/hqdefault.jpg",
								Keterangan: "Materi Momen Kekhasan Dan Jenis-Jenis Atom Karbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "eXijwwYrZ38",
								Judul:      "Belajar Materi Kekhasan Dan Jenis-Jenis Atom Karbon",
								Thumbnail:  "https://i.ytimg.com/vi/eXijwwYrZ38/hqdefault.jpg",
								Keterangan: "Materi Momen Kekhasan Dan Jenis-Jenis Atom Karbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "57jPesaJRzM",
								Judul:      "Belajar Materi Penggolongan Senyawa Hidrokarbon",
								Thumbnail:  "https://i.ytimg.com/vi/57jPesaJRzM/hqdefault.jpg",
								Keterangan: "Materi Penggolongan Senyawa Hidrokarbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "4sJzwo4bh60",
								Judul:      "Belajar Materri Tata Nama Senyawa Hidrokarbon",
								Thumbnail:  "https://i.ytimg.com/vi/4sJzwo4bh60/hqdefault.jpg",
								Keterangan: "Materi Materri Tata Nama Senyawa Hidrokarbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "64Ns-OU0wnQ",
								Judul:      "Belajar Materi Isomer Senyawa Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/64Ns-OU0wnQ/hqdefault.jpg",
								Keterangan: "Materi Isomer Senyawa Alkana Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          5,
						Class:       "kimia5",
						Description: "Belajar Kimia Bersama Kimia 5",
						Video: []Video{
							{
								VideoID:    "57jPesaJRzM",
								Judul:      "Belajar Materi Penggolongan Senyawa Hidrokarbon",
								Thumbnail:  "https://i.ytimg.com/vi/57jPesaJRzM/hqdefault.jpg",
								Keterangan: "Materi Penggolongan Senyawa Hidrokarbon Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "hopJg2kXiQY",
								Judul:      "Belajar Materi Sifat Alkena dan Alkuna",
								Thumbnail:  "https://i.ytimg.com/vi/hopJg2kXiQY/hqdefault.jpg",
								Keterangan: "Materi Sifat Alkena dan Alkuna Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AZY27IeQIOI",
								Judul:      "Belajar Materi Reaksi kimia Alkana",
								Thumbnail:  "https://i.ytimg.com/vi/AZY27IeQIOI/hqdefault.jpg",
								Keterangan: "Materi Reaksi kimia Alkana Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "eXijwwYrZ38",
								Judul:      "Belajar Materi Kekhasan Dan Jenis-Jenis Atom Karbon",
								Thumbnail:  "https://i.ytimg.com/vi/eXijwwYrZ38/hqdefault.jpg",
								Keterangan: "Materi Momen Kekhasan Dan Jenis-Jenis Atom Karbon Untuk Persiapan SBMPTN",
							},
						},
					},
				},
			},
			{
				ID:   4,
				Name: "biologi",
				Kelas: []Kelass{
					{
						ID:          1,
						Class:       "biologi1",
						Description: "Belajar Biologi Bersama Biologi 1",
						Video: []Video{
							{
								VideoID:    "TYfN2oW02eI",
								Judul:      "Belajar Materi Metabolisme",
								Thumbnail:  "https://i.ytimg.com/vi/TYfN2oW02eI/hqdefault.jpg",
								Keterangan: "Materi Metabolisme Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "Gk_-dkBMZk0",
								Judul:      "Belajar Materi Metabolisme bagian 2",
								Thumbnail:  "https://i.ytimg.com/vi/Gk_-dkBMZk0/hqdefault.jpg",
								Keterangan: "Materi Materri Metabolisme bagian 2 Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "WJNbwom-aXQ",
								Judul:      "Belajar Materi Metabolisme bagian 3",
								Thumbnail:  "https://i.ytimg.com/vi/WJNbwom-aXQ/hqdefault.jpg",
								Keterangan: "Materi Materri Metabolisme bagian 3 Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "F-VOURqRFgo",
								Judul:      "Belajar Materi Pertumbuhan dan perkembangan",
								Thumbnail:  "https://i.ytimg.com/vi/F-VOURqRFgo/hqdefault.jpg",
								Keterangan: "Materi Pertumbuhan dan perkembangan Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "yGkl3iTlY9c",
								Judul:      "Belajar Materi Pertumbuhan dan perkembangan bagian 2",
								Thumbnail:  "https://i.ytimg.com/vi/yGkl3iTlY9c/hqdefault.jpg",
								Keterangan: "Materi Pertumbuhan dan perkembangan bagian 2 Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          2,
						Class:       "biologi2",
						Description: "Belajar Biologi Bersama Biologi 2",
						Video: []Video{
							{
								VideoID:    "7HEAvF7-b1k",
								Judul:      "Belajar Materi Metabolisme",
								Thumbnail:  "https://i.ytimg.com/vi/7HEAvF7-b1k/hqdefault.jpg",
								Keterangan: "Materi Metabolisme Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "OBcNgttpjsk",
								Judul:      "Belajar Materi Mutasi",
								Thumbnail:  "https://i.ytimg.com/vi/OBcNgttpjsk/hqdefault.jpg",
								Keterangan: "Materi Mutasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "fGWsYylYG4g",
								Judul:      "Belajar Materi Evolusi",
								Thumbnail:  "https://i.ytimg.com/vi/fGWsYylYG4g/hqdefault.jpg",
								Keterangan: "Materi Evolusi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "XSS18DC8abE",
								Judul:      "Belajar Materi Bioteknologi",
								Thumbnail:  "https://i.ytimg.com/vi/XSS18DC8abE/hqdefault.jpg",
								Keterangan: "Materi Bioteknol Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "Zq65RvpmXSk",
								Judul:      "Belajar Materi Pembelahan sel",
								Thumbnail:  "https://i.ytimg.com/vi/Zq65RvpmXSk/hqdefault.jpg",
								Keterangan: "Materi Pembelahan sel Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          3,
						Class:       "biologi3",
						Description: "Belajar Biologi Bersama Biologi 3",
						Video: []Video{
							{
								VideoID:    "XSS18DC8abE",
								Judul:      "Belajar Materi Bioteknologi",
								Thumbnail:  "https://i.ytimg.com/vi/XSS18DC8abE/hqdefault.jpg",
								Keterangan: "Materi Bioteknol Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "7HEAvF7-b1k",
								Judul:      "Belajar Materi Metabolisme",
								Thumbnail:  "https://i.ytimg.com/vi/7HEAvF7-b1k/hqdefault.jpg",
								Keterangan: "Materi Metabolisme Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "OBcNgttpjsk",
								Judul:      "Belajar Materi Mutasi",
								Thumbnail:  "https://i.ytimg.com/vi/OBcNgttpjsk/hqdefault.jpg",
								Keterangan: "Materi Mutasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "fGWsYylYG4g",
								Judul:      "Belajar Materi Evolusi",
								Thumbnail:  "https://i.ytimg.com/vi/fGWsYylYG4g/hqdefault.jpg",
								Keterangan: "Materi Evolusi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "Zq65RvpmXSk",
								Judul:      "Belajar Materi Pembelahan sel",
								Thumbnail:  "https://i.ytimg.com/vi/Zq65RvpmXSk/hqdefault.jpg",
								Keterangan: "Materi Pembelahan sel Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          4,
						Class:       "biologi4",
						Description: "Belajar Biologi Bersama Biologi 4",
						Video: []Video{
							{
								VideoID:    "Zq65RvpmXSk",
								Judul:      "Belajar Materi Pembelahan sel",
								Thumbnail:  "https://i.ytimg.com/vi/Zq65RvpmXSk/hqdefault.jpg",
								Keterangan: "Materi Pembelahan sel Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "XSS18DC8abE",
								Judul:      "Belajar Materi Bioteknologi",
								Thumbnail:  "https://i.ytimg.com/vi/XSS18DC8abE/hqdefault.jpg",
								Keterangan: "Materi Bioteknol Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "7HEAvF7-b1k",
								Judul:      "Belajar Materi Metabolisme",
								Thumbnail:  "https://i.ytimg.com/vi/7HEAvF7-b1k/hqdefault.jpg",
								Keterangan: "Materi Metabolisme Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "OBcNgttpjsk",
								Judul:      "Belajar Materi Mutasi",
								Thumbnail:  "https://i.ytimg.com/vi/OBcNgttpjsk/hqdefault.jpg",
								Keterangan: "Materi Mutasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "fGWsYylYG4g",
								Judul:      "Belajar Materi Evolusi",
								Thumbnail:  "https://i.ytimg.com/vi/fGWsYylYG4g/hqdefault.jpg",
								Keterangan: "Materi Evolusi Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          5,
						Class:       "biologi5",
						Description: "Belajar Biologi Bersama Biologi 5",
						Video: []Video{
							{
								VideoID:    "XSS18DC8abE",
								Judul:      "Belajar Materi Bioteknologi",
								Thumbnail:  "https://i.ytimg.com/vi/XSS18DC8abE/hqdefault.jpg",
								Keterangan: "Materi Bioteknol Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "7HEAvF7-b1k",
								Judul:      "Belajar Materi Metabolisme",
								Thumbnail:  "https://i.ytimg.com/vi/7HEAvF7-b1k/hqdefault.jpg",
								Keterangan: "Materi Metabolisme Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "OBcNgttpjsk",
								Judul:      "Belajar Materi Mutasi",
								Thumbnail:  "https://i.ytimg.com/vi/OBcNgttpjsk/hqdefault.jpg",
								Keterangan: "Materi Mutasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "fGWsYylYG4g",
								Judul:      "Belajar Materi Evolusi",
								Thumbnail:  "https://i.ytimg.com/vi/fGWsYylYG4g/hqdefault.jpg",
								Keterangan: "Materi Evolusi Untuk Persiapan SBMPTN",
							},
						},
					},
				},
			},
			{
				ID:   5,
				Name: "indonesia",
				Kelas: []Kelass{
					{
						ID:          1,
						Class:       "indonesia1",
						Description: "Belajar Indonesia Bersama Indonesia 1",
						Video: []Video{
							{
								VideoID:    "v_fibdXgmRQ",
								Judul:      "Belajar Materi Teks Laporan Hasil Observasi",
								Thumbnail:  "https://i.ytimg.com/vi/v_fibdXgmRQ/hqdefault.jpg",
								Keterangan: "Materi Teks Laporan Hasil Observasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Puisi",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Puisi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "l5_OYUdgN1Y",
								Judul:      "Belajar Materi Biografi",
								Thumbnail:  "https://i.ytimg.com/vi/l5_OYUdgN1Y/hqdefault.jpg",
								Keterangan: "Materi Biografi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "i4ozP9VCC8s",
								Judul:      "Belajar Materi Hikayat",
								Thumbnail:  "https://i.ytimg.com/vi/i4ozP9VCC8s/hqdefault.jpg",
								Keterangan: "Materi Hikayat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UkEFkPNnyro",
								Judul:      "Belajar Materi Anekdot",
								Thumbnail:  "https://i.ytimg.com/vi/UkEFkPNnyro/hqdefault.jpg",
								Keterangan: "Materi Anekdot Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          2,
						Class:       "indonesia2",
						Description: "Belajar Indonesia Bersama Indonesia 2",
						Video: []Video{
							{
								VideoID:    "UkEFkPNnyro",
								Judul:      "Belajar Materi Anekdot",
								Thumbnail:  "https://i.ytimg.com/vi/UkEFkPNnyro/hqdefault.jpg",
								Keterangan: "Materi Anekdot Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "v_fibdXgmRQ",
								Judul:      "Belajar Materi Teks Laporan Hasil Observasi",
								Thumbnail:  "https://i.ytimg.com/vi/v_fibdXgmRQ/hqdefault.jpg",
								Keterangan: "Materi Teks Laporan Hasil Observasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Puisi",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Puisi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "l5_OYUdgN1Y",
								Judul:      "Belajar Materi Biografi",
								Thumbnail:  "https://i.ytimg.com/vi/l5_OYUdgN1Y/hqdefault.jpg",
								Keterangan: "Materi Biografi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "i4ozP9VCC8s",
								Judul:      "Belajar Materi Hikayat",
								Thumbnail:  "https://i.ytimg.com/vi/i4ozP9VCC8s/hqdefault.jpg",
								Keterangan: "Materi Hikayat Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          3,
						Class:       "indonesia3",
						Description: "Belajar Indonesia Bersama Indonesia 3",
						Video: []Video{
							{
								VideoID:    "l5_OYUdgN1Y",
								Judul:      "Belajar Materi Biografi",
								Thumbnail:  "https://i.ytimg.com/vi/l5_OYUdgN1Y/hqdefault.jpg",
								Keterangan: "Materi Biografi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "i4ozP9VCC8s",
								Judul:      "Belajar Materi Hikayat",
								Thumbnail:  "https://i.ytimg.com/vi/i4ozP9VCC8s/hqdefault.jpg",
								Keterangan: "Materi Hikayat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UkEFkPNnyro",
								Judul:      "Belajar Materi Anekdot",
								Thumbnail:  "https://i.ytimg.com/vi/UkEFkPNnyro/hqdefault.jpg",
								Keterangan: "Materi Anekdot Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "v_fibdXgmRQ",
								Judul:      "Belajar Materi Teks Laporan Hasil Observasi",
								Thumbnail:  "https://i.ytimg.com/vi/v_fibdXgmRQ/hqdefault.jpg",
								Keterangan: "Materi Teks Laporan Hasil Observasi Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          4,
						Class:       "indonesia4",
						Description: "Belajar Indonesia Bersama Indonesia 4",
						Video: []Video{
							{
								VideoID:    "v_fibdXgmRQ",
								Judul:      "Belajar Materi Teks Laporan Hasil Observasi",
								Thumbnail:  "https://i.ytimg.com/vi/v_fibdXgmRQ/hqdefault.jpg",
								Keterangan: "Materi Teks Laporan Hasil Observasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "l5_OYUdgN1Y",
								Judul:      "Belajar Materi Biografi",
								Thumbnail:  "https://i.ytimg.com/vi/l5_OYUdgN1Y/hqdefault.jpg",
								Keterangan: "Materi Biografi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "i4ozP9VCC8s",
								Judul:      "Belajar Materi Hikayat",
								Thumbnail:  "https://i.ytimg.com/vi/i4ozP9VCC8s/hqdefault.jpg",
								Keterangan: "Materi Hikayat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UkEFkPNnyro",
								Judul:      "Belajar Materi Anekdot",
								Thumbnail:  "https://i.ytimg.com/vi/UkEFkPNnyro/hqdefault.jpg",
								Keterangan: "Materi Anekdot Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Puisi",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Puisi Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          5,
						Class:       "indonesia5",
						Description: "Belajar Indonesia Bersama Indonesia 5",
						Video: []Video{
							{
								VideoID:    "UkEFkPNnyro",
								Judul:      "Belajar Materi Anekdot",
								Thumbnail:  "https://i.ytimg.com/vi/UkEFkPNnyro/hqdefault.jpg",
								Keterangan: "Materi Anekdot Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Puisi",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Puisi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "v_fibdXgmRQ",
								Judul:      "Belajar Materi Teks Laporan Hasil Observasi",
								Thumbnail:  "https://i.ytimg.com/vi/v_fibdXgmRQ/hqdefault.jpg",
								Keterangan: "Materi Teks Laporan Hasil Observasi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "l5_OYUdgN1Y",
								Judul:      "Belajar Materi Biografi",
								Thumbnail:  "https://i.ytimg.com/vi/l5_OYUdgN1Y/hqdefault.jpg",
								Keterangan: "Materi Biografi Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "i4ozP9VCC8s",
								Judul:      "Belajar Materi Hikayat",
								Thumbnail:  "https://i.ytimg.com/vi/i4ozP9VCC8s/hqdefault.jpg",
								Keterangan: "Materi Hikayat Untuk Persiapan SBMPTN",
							},
						},
					},
				},
			},
			{
				ID:   6,
				Name: "inggris",
				Kelas: []Kelass{
					{
						ID:          1,
						Class:       "inggris1",
						Description: "Belajar Inggris Bersama Inggris 1",
						Video: []Video{
							{
								VideoID:    "UMMBQwsNeB8",
								Judul:      "Belajar Materi Offers and Suggestions",
								Thumbnail:  "https://i.ytimg.com/vi/UMMBQwsNeB8/hqdefault.jpg",
								Keterangan: "Materi Offers and Suggestions Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AmMxJcmPnwU",
								Judul:      "Belajar Materi Cause and Effect",
								Thumbnail:  "https://i.ytimg.com/vi/AmMxJcmPnwU/hqdefault.jpg",
								Keterangan: "Materi Cause and  Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Past Perfect Continous Tense",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Past Perfect Continuous Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "o8UmwSi0lkE",
								Judul:      "Belajar Materi Future Perfect Tense",
								Thumbnail:  "https://i.ytimg.com/vi/o8UmwSi0lkE/hqdefault.jpg",
								Keterangan: "Materi Future Perfect Tense Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          2,
						Class:       "inggris2",
						Description: "Belajar Inggris Bersama Inggris 2",
						Video: []Video{
							{
								VideoID:    "o8UmwSi0lkE",
								Judul:      "Belajar Materi Future Perfect Tense",
								Thumbnail:  "https://i.ytimg.com/vi/o8UmwSi0lkE/hqdefault.jpg",
								Keterangan: "Materi Future Perfect Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UMMBQwsNeB8",
								Judul:      "Belajar Materi Offers and Suggestions",
								Thumbnail:  "https://i.ytimg.com/vi/UMMBQwsNeB8/hqdefault.jpg",
								Keterangan: "Materi Offers and Suggestions Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AmMxJcmPnwU",
								Judul:      "Belajar Materi Cause and Effect",
								Thumbnail:  "https://i.ytimg.com/vi/AmMxJcmPnwU/hqdefault.jpg",
								Keterangan: "Materi Cause and  Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Past Perfect Continous Tense",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Past Perfect Continuous Tense Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          3,
						Class:       "inggris3",
						Description: "Belajar Inggris Bersama Inggris 3",
						Video: []Video{
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Past Perfect Continous Tense",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Past Perfect Continuous Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "o8UmwSi0lkE",
								Judul:      "Belajar Materi Future Perfect Tense",
								Thumbnail:  "https://i.ytimg.com/vi/o8UmwSi0lkE/hqdefault.jpg",
								Keterangan: "Materi Future Perfect Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UMMBQwsNeB8",
								Judul:      "Belajar Materi Offers and Suggestions",
								Thumbnail:  "https://i.ytimg.com/vi/UMMBQwsNeB8/hqdefault.jpg",
								Keterangan: "Materi Offers and Suggestions Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AmMxJcmPnwU",
								Judul:      "Belajar Materi Cause and Effect",
								Thumbnail:  "https://i.ytimg.com/vi/AmMxJcmPnwU/hqdefault.jpg",
								Keterangan: "Materi Cause and  Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          4,
						Class:       "inggris4",
						Description: "Belajar Inggris Bersama Inggris 4",
						Video: []Video{
							{
								VideoID:    "AmMxJcmPnwU",
								Judul:      "Belajar Materi Cause and Effect",
								Thumbnail:  "https://i.ytimg.com/vi/AmMxJcmPnwU/hqdefault.jpg",
								Keterangan: "Materi Cause and  Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "o8UmwSi0lkE",
								Judul:      "Belajar Materi Future Perfect Tense",
								Thumbnail:  "https://i.ytimg.com/vi/o8UmwSi0lkE/hqdefault.jpg",
								Keterangan: "Materi Future Perfect Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UMMBQwsNeB8",
								Judul:      "Belajar Materi Offers and Suggestions",
								Thumbnail:  "https://i.ytimg.com/vi/UMMBQwsNeB8/hqdefault.jpg",
								Keterangan: "Materi Offers and Suggestions Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "AmMxJcmPnwU",
								Judul:      "Belajar Materi Cause and Effect",
								Thumbnail:  "https://i.ytimg.com/vi/AmMxJcmPnwU/hqdefault.jpg",
								Keterangan: "Materi Cause and  Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Past Perfect Continous Tense",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Past Perfect Continuous Tense Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          3,
						Class:       "inggris3",
						Description: "Belajar Inggris Bersama Inggris 3",
						Video: []Video{
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Past Perfect Continous Tense",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Past Perfect Continuous Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "o8UmwSi0lkE",
								Judul:      "Belajar Materi Future Perfect Tense",
								Thumbnail:  "https://i.ytimg.com/vi/o8UmwSi0lkE/hqdefault.jpg",
								Keterangan: "Materi Future Perfect Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UMMBQwsNeB8",
								Judul:      "Belajar Materi Offers and Suggestions",
								Thumbnail:  "https://i.ytimg.com/vi/UMMBQwsNeB8/hqdefault.jpg",
								Keterangan: "Materi Offers and Suggestions Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          4,
						Class:       "inggris4",
						Description: "Belajar Inggris Bersama Inggris 4",
						Video: []Video{
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Past Perfect Continous Tense",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Past Perfect Continuous Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "o8UmwSi0lkE",
								Judul:      "Belajar Materi Future Perfect Tense",
								Thumbnail:  "https://i.ytimg.com/vi/o8UmwSi0lkE/hqdefault.jpg",
								Keterangan: "Materi Future Perfect Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UMMBQwsNeB8",
								Judul:      "Belajar Materi Offers and Suggestions",
								Thumbnail:  "https://i.ytimg.com/vi/UMMBQwsNeB8/hqdefault.jpg",
								Keterangan: "Materi Offers and Suggestions Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
						},
					},
					{
						ID:          5,
						Class:       "inggris5",
						Description: "Belajar Inggris Bersama Inggris 5",
						Video: []Video{
							{
								VideoID:    "3ewvPXPgZf8",
								Judul:      "Belajar Materi Past Perfect Continous Tense",
								Thumbnail:  "https://i.ytimg.com/vi/3ewvPXPgZf8/hqdefault.jpg",
								Keterangan: "Materi Past Perfect Continuous Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "o8UmwSi0lkE",
								Judul:      "Belajar Materi Future Perfect Tense",
								Thumbnail:  "https://i.ytimg.com/vi/o8UmwSi0lkE/hqdefault.jpg",
								Keterangan: "Materi Future Perfect Tense Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "UMMBQwsNeB8",
								Judul:      "Belajar Materi Offers and Suggestions",
								Thumbnail:  "https://i.ytimg.com/vi/UMMBQwsNeB8/hqdefault.jpg",
								Keterangan: "Materi Offers and Suggestions Untuk Persiapan SBMPTN",
							},
							{
								VideoID:    "m2YFr6R94Xw",
								Judul:      "Belajar Materi The Owl and the Pussy Cat",
								Thumbnail:  "https://i.ytimg.com/vi/m2YFr6R94Xw/hqdefault.jpg",
								Keterangan: "Materi The Owl and the Pussy Cat Untuk Persiapan SBMPTN",
							},
						},
					},
				},
			},
		},
	},
}
