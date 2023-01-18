package models

type Classses struct {
	ID      int64   `json:"id"`
	Teacher string  `json:"teacher"`
	Rating  float64 `json:"rating"`
	Link    string  `json:"link"`
}

type Subjects struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Image       string     `json:"image"`
	Description string     `json:"description"`
	Color       string     `json:"color"`
	Classes     []Classses `json:"classes"`
}

type klas struct {
	Subjects []Subjects `json:"subjects"`
}

var Kelas = []klas{
	{
		Subjects: []Subjects{{
			ID:          1,
			Name:        "matematika",
			Image:       "https://posi.id/wp-content/uploads/2022/09/Matematika-Adalah-scaled.jpg",
			Description: "Belajar Matematika",
			Color:       "linear(to-b, rgba(97, 210, 242, 1), rgba(97, 210, 242, 1))",
			Classes: []Classses{
				{
					ID:      1,
					Teacher: "Matematika 1",
					Rating:  4.9,
					Link:    "/modul/matematika/1",
				},
				{
					ID:      2,
					Teacher: "Matematika 2",
					Rating:  4.9,
					Link:    "/modul/matematika/2",
				},
				{
					ID:      3,
					Teacher: "Matematika 3",
					Rating:  4.9,
					Link:    "/modul/matematika/3",
				},
				{
					ID:      4,
					Teacher: "Matematika 4",
					Rating:  4.9,
					Link:    "/modul/matematika/4",
				},
				{
					ID:      5,
					Teacher: "Matematika 5",
					Rating:  4.9,
					Link:    "/modul/matematika/5",
				},
			},
		},
			{
				ID:          2,
				Name:        "fisika",
				Image:       "https://cdnwpedutorenews.gramedia.net/wp-content/uploads/2022/11/07022828/fisika.jpg",
				Description: "Belajar Fisika",
				Color:       "linear(to-b, rgba(255, 155, 41, 1), rgba(255, 155, 41, 1))",
				Classes: []Classses{
					{
						ID:      1,
						Teacher: "Fisika 1",
						Rating:  4.9,
						Link:    "/modul/fisika/1",
					},
					{
						ID:      2,
						Teacher: "Fisika 2",
						Rating:  4.9,
						Link:    "/modul/fisika/2",
					},
					{
						ID:      3,
						Teacher: "Fisika 3",
						Rating:  4.9,
						Link:    "/modul/fisika/3",
					},
					{
						ID:      4,
						Teacher: "Fisika 4",
						Rating:  4.9,
						Link:    "/modul/fisika/4",
					},
					{
						ID:      5,
						Teacher: "Fisika 5",
						Rating:  4.9,
						Link:    "/modul/fisika/5",
					},
				},
			},
			{
				ID:          3,
				Name:        "kimia",
				Image:       "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQu2Ih5qTe8KPlpvahkDsZLFl9YLRSdVud2Bw&usqp=CAU",
				Description: "Belajar Kimia",
				Color:       "linear(to-b, rgba(255, 234, 0, 1)	, rgba(255, 234, 0, 1)	)",
				Classes: []Classses{
					{
						ID:      1,
						Teacher: "Kimia 1",
						Rating:  4.9,
						Link:    "/modul/kimia/1",
					},
					{
						ID:      2,
						Teacher: "Kimia 2",
						Rating:  4.9,
						Link:    "/modul/kimia/2",
					},
					{
						ID:      3,
						Teacher: "Kimia 3",
						Rating:  4.9,
						Link:    "/modul/kimia/3",
					},
					{
						ID:      4,
						Teacher: "Kimia 4",
						Rating:  4.9,
						Link:    "/modul/kimia/4",
					},
					{
						ID:      5,
						Teacher: "Kimia 5",
						Rating:  4.9,
						Link:    "/modul/kimia/5",
					},
				},
			},
			{
				ID:          4,
				Name:        "biologi",
				Image:       "https://rimbakita.com/wp-content/uploads/2020/09/biologi.jpg",
				Description: "Belajar Biologi",
				Color:       "linear(to-b, rgba(13, 215, 95, 1), ,rgba(13, 215, 95, 1))",
				Classes: []Classses{
					{
						ID:      1,
						Teacher: "Biologi 1",
						Rating:  4.9,
						Link:    "/modul/biologi/1",
					},
					{
						ID:      2,
						Teacher: "Biologi 2",
						Rating:  4.9,
						Link:    "/modul/biologi/2",
					},
					{
						ID:      3,
						Teacher: "Biologi 3",
						Rating:  4.9,
						Link:    "/modul/biologi/3",
					},
					{
						ID:      4,
						Teacher: "Biologi 4",
						Rating:  4.9,
						Link:    "/modul/biologi/4",
					},
					{
						ID:      5,
						Teacher: "Biologi 5",
						Rating:  4.9,
						Link:    "/modul/biologi/5",
					},
				},
			},
			{
				ID:          5,
				Name:        "indonesia",
				Image:       "data:image/jpeg;base64,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",
				Description: "Belajar bahasa Indonesia",
				Classes: []Classses{
					{
						ID:      1,
						Teacher: "Bahasa Indonesia 1",
						Rating:  4.9,
						Link:    "/modul/indonesia/1",
					},
					{
						ID:      2,
						Teacher: "Bahasa Indonesia 2",
						Rating:  4.9,
						Link:    "/modul/indonesia/2",
					},
					{
						ID:      3,
						Teacher: "Bahasa Indonesia 3",
						Rating:  4.9,
						Link:    "/modul/indonesia/3",
					},
					{
						ID:      4,
						Teacher: "Bahasa Indonesia 4",
						Rating:  4.9,
						Link:    "/modul/indonesia/4",
					},
					{
						ID:      5,
						Teacher: "Bahasa Indonesia 5",
						Rating:  4.9,
						Link:    "/modul/indonesia/5",
					},
				},
			},
			{
				ID:          6,
				Name:        "inggris",
				Image:       "https://www.stkippgriponorogo.ac.id/wp-content/uploads/2019/09/Belajar-Bahasa-Inggris-01-Finansialku.jpg",
				Description: "Belajar bahasa Inggris",
				Classes: []Classses{
					{
						ID:      1,
						Teacher: "Bahasa Inggris 1",
						Rating:  4.9,
						Link:    "/modul/inggris/1",
					},
					{
						ID:      2,
						Teacher: "Bahasa Inggris 2",
						Rating:  4.9,
						Link:    "/modul/inggris/2",
					},
					{
						ID:      3,
						Teacher: "Bahasa Inggris 3",
						Rating:  4.9,
						Link:    "/modul/inggris/3",
					},
					{
						ID:      4,
						Teacher: "Bahasa Inggris 4",
						Rating:  4.9,
						Link:    "/modul/inggris/4",
					},
					{
						ID:      5,
						Teacher: "Bahasa Inggris 5",
						Rating:  4.9,
						Link:    "/modul/inggris/5",
					},
				},
			}},
	},
}
