package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Link struct {
	Name string
	URL  string
}

type Subject struct {
	Name  string
	Links []Link
}

type Course struct {
	Name     string
	Subjects []Subject
}

var courses = map[string]Course{
	"1": {
		Name: "Семестр 1",
		Subjects: []Subject{
			{"Информатика", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/1r-x-5tgyyYmD--Td4DN1poNxBCXclu_M/view?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/file/d/1NF4nG6D42Fus6wqYBLiamOVxlHF4eVZF/view?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1CmZR5jZpLfBElwYSxVo2GmGxH1dHhmeq/view?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/file/d/1BAopMtUug905t2ZgjXojuxl8YeD4-oEa/view?usp=drive_link"},
				{"Лабораторная работа 5", "https://drive.google.com/file/d/1jItA7gA4PN2vdhtNDZPbexDhNyOnnIMe/view?usp=drive_link"},
				{"Лабораторная работа 6", "https://drive.google.com/file/d/1i4ZrxKMWGHn0scMpXKWf0bBwY4OE7rk-/view?usp=drive_link"},
				{"Лабораторная работа 7", "https://drive.google.com/file/d/1-bRBCVYQiybC3PDMd8i6as1p6GjrEPhe/view?usp=drive_link"},
				{"Лабораторная работа 8", "https://drive.google.com/file/d/1Pb8C2foPeJ0R2Z9PPzYlvYMDWcQHgN_W/view?usp=drive_link"},
				{"Лабораторная работа 9", "https://drive.google.com/file/d/15Yq_kVtcKQLFyli7uFj9oNoVGU8DMFQc/view?usp=drive_link"},
				{"Лабораторная работа 10", "https://drive.google.com/file/d/1NOXSx-qxM3UU64IAZBQeAaed-0JG6eho/view?usp=drive_link"},
				{"Лабораторная работа 11", "https://drive.google.com/file/d/1J0zlKlG0IuMmhRJ-NvQa0p6zVsFnSQ2H/view?usp=drive_link"},
				{"Лабораторная работа 12", "https://drive.google.com/file/d/1tk4uGB_9FI62REooGZv1Tmtf5CJSGfHI/view?usp=drive_link"},
				{"Самостоятельная работа 1", "https://drive.google.com/file/d/1Lis4r8BJJh5f58mz-Unqu7lUxmJP-oNv/view?usp=drive_link"},
				{"Самостоятельная работа 3", "https://drive.google.com/file/d/1QfCbecdcR1CXVVKwteJg7V_jK5N7fLKb/view?usp=drive_link"},
				{"Самостоятельная работа 5", "https://drive.google.com/file/d/1npb71_36DKQ_3XDWbcxDsi_dooardEho/view?usp=drive_link"},
				{"Самостоятельная работа 7.1", "https://drive.google.com/file/d/1Hx1QLhdla9g0Me_-9amJKygn1DaWUj2z/view?usp=drive_link"},
				{"Самостоятельная работа 7.2", "https://drive.google.com/file/d/133Qku3YjuAP1kSXbQSqIz3VCcyf4Aj_D/view?usp=drive_link"},
				{"Самостоятельная работа 8", "https://drive.google.com/file/d/1i5XVNZdIEMvHidcCwG-RYP8IjQh6Cyu7/view?usp=drive_link"},
			}},
			{"Информационные технологии в математике", []Link{
				{"ЛР Матрицы", "https://docs.google.com/spreadsheets/d/1G7bK8ORINRnT8xFhaNo6YXllgwVsT5ad/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"ЛР Графики", "https://docs.google.com/spreadsheets/d/1aoFsTbURaPnDHZSFs1h1qIGMO4FCi3eZ/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Вариативная работа 1.1", "https://drive.google.com/file/d/1Zgm95pxlpp22ZgcNNH-HSzzXnA2k_wAQ/view?usp=drive_link"},
				{"Вариативная работа 1.2", "https://drive.google.com/file/d/11EPFw8u6Rw93_fzxNv-DNy043U0JFLXF/view?usp=drive_link"},
				{"Вариативная работа 2", "https://drive.google.com/file/d/12H65T5oWPJT9C7mio-rloMAsW08nvdk1/view?usp=drive_link"},
				{"Вариативная работа 3.1", "https://drive.google.com/file/d/1rhJGqwDU8NykX8DGTuGaxidqUwHyX_kE/view?usp=drive_link"},
				{"Вариативная работа 3.2", "https://drive.google.com/file/d/1jigTE6Gy0JXDnl6Twlr6oqJr5t0qzOz2/view?usp=drive_link"},
				{"Вариативная работа 4", "https://drive.google.com/file/d/1NfVu2AuQmaAG0qBF5VjR7ncroqFdp78g/view?usp=drive_link"},
				{"Вариативная работа 5", "https://drive.google.com/file/d/1WJpCE4lgrGYm8X5TQ4CVj4ZZDOmm5H3U/view?usp=drive_link"},
				{"Вариативная работа 6.1", "https://drive.google.com/file/d/19OftX473r0cKoQQfL3UsuVHXCk7MbaQF/view?usp=drive_link"},
				{"Вариативная работа 6.2", "https://drive.google.com/file/d/1W0C-SQ6gNzq2F63cYfBgQVIDNLWQ-Lo2/view?usp=drive_link"},
				{"Все ИСР", "https://drive.google.com/drive/folders/1Iwt7jr3ov5FAY5GCHwAKCVKMj6HcbcJO?usp=drive_link"},
			}},
			{"Основы компьютерной алгебры", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/drive/folders/1hOZfRuzFwlAdvzZ5CE9tICFsSK06lBOL?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/drive/folders/1R9UsrZgg_y4TWw7BWdV7Gjt3zC_0J0Nv?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/drive/folders/1toGPKVmOhUHEAO4U33MklOWlw1i1AIiJ?usp=drive_link"},
				{"Вариативная самостоятельная работа 1", "https://drive.google.com/file/d/1C6V8XH3J5zDTWLUDWW8AgKPexXVo0dW2/view?usp=drive_link"},
				{"Вариативная самостоятельная работа 2", "https://drive.google.com/drive/folders/1py6bo2wRDGGHUCyyh9fLOG3IY2SjydKh?usp=drive_link"},
				{"Вариативная самостоятельная работа 3", "https://drive.google.com/drive/folders/1IBpxOGBc_92vDUYkx_MC2-_C5pLzblhV?usp=drive_link"},
				{"Инвариативная самостоятельная работа 1", "https://drive.google.com/file/d/10wyigpCOx_PjONrudtDNVIbAxgd-R2C6/view?usp=drive_link"},
				{"Инвариативная самостоятельная работа 3", "https://drive.google.com/drive/folders/14o5EFvzuKBqRXmrn32is1LGb1pSD13b_?usp=drive_link"},
			}},
		},
	},
	"2": {
		Name: "Семестр 2",
		Subjects: []Subject{
			{"Практика Учебная (эксплуатационная)", []Link{
				{"Портфолио по эксплуатационной практике", "https://github.com/SwagStar228/SwagStar228.AET_pratic.io"},
			}},
			{"Программирование", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/1WvSt8oi-Y2dmjnO10dVyHwOp22GMJrCt/view?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1I-gV8WCwJxUyqqU9iXXpjREPS1ziY2YH/view?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/file/d/1Nj_SHMniABotO9NeGz5IcsxBfPupZKvp/view?usp=drive_link"},
				{"Лабораторная работа 5.1", "https://drive.google.com/file/d/1gesPFJmMApv6rHbilFda_dLDe4ALPtYG/view?usp=drive_link"},
				{"Лабораторная работа 5.2", "https://drive.google.com/file/d/1DmqWXobjK4HZipZ6ZeeFczgjr4oQ0oj-/view?usp=drive_link"},
				{"Лабораторная работа 6", "https://drive.google.com/file/d/170_qr91FtD0vJJQYZLU1BojmpcPP45mm/view?usp=drive_link"},
				{"Лабораторная работа 7", "https://drive.google.com/file/d/1Ad4lfAQJm-2k6jnpOnMukg6FGaqci-5O/view?usp=drive_link"},
				{"Лабораторная работа 8", "https://drive.google.com/file/d/1SnRxkMMRUVWy6mR6Oe-57yHJbFTmuoP2/view?usp=drive_link"},
				{"Лабораторная работа 9", "https://drive.google.com/file/d/1NKXqe-6R1TC6f-tQypS753KCiY2mPCvb/view?usp=drive_link"},
			}},
		},
	},
	"3": {
		Name: "Семестр 3",
		Subjects: []Subject{
			{"Практика производственная (технологическая (проектно-технологическая))", []Link{
				{"Задание на производственную практику", "https://drive.google.com/file/d/1_GOBgwTnDmPYdfYzK4FPfOnLQN67dIGa/view?usp=drive_link"},
				{"Отчет о прохождении производственной практики", "https://drive.google.com/file/d/1lk75ZP1L2hTfQu6AdGKnLXYZpa3w-gDK/view?usp=sharing"},
			}},
			{"Анализ данных", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/drive/folders/1jyOkAKMaTCf0miKDqlnkZuFe-zPz08Yq?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/drive/folders/1_rVGLYu7Tq1nlv9V20aXQSKjkpCv4w_x?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1sdxv0HaS6T_9KuvBxfQ3tvW82bLWT85y/view?usp=drive_link"},
				{"Лабораторная работа 4.1", "https://drive.google.com/drive/folders/1eHKnkatkwQV7LtN8GNW1S5m8-sgDv8SC?usp=drive_link"},
				{"Лабораторная работа 4.2", "https://drive.google.com/drive/folders/1BcJ9XXb0eJLAuXSjJi_SHM74GZEBdn4K?usp=drive_link"},
				{"Лабораторная работа 5", "https://drive.google.com/drive/folders/1NDY65J5rOrnUTMtChLw7tBNNnnzggblj?usp=drive_link"},
				{"Лабораторная работа 8", "https://drive.google.com/drive/folders/17aIefqWkeYjyczE7M2VPgXib79Q8FHux?usp=drive_link"},
				{"Самостоятельная работ 2", "https://drive.google.com/drive/folders/1Mc5vnsZil9HPHlWpQPgTvRJSQ4no0Gbi?usp=drive_link"},
				{"Самостоятельная работ 3", "https://drive.google.com/file/d/1h9CSBXX3SW7Ctnm0GqccCLwqSp9wkQ1k/view?usp=drive_link"},
				{"Самостоятельная работ 4", "https://drive.google.com/file/d/1NU0uPAL-0YsE2f3FbSspM2KbpGpIYR09/view?usp=drive_link"},
			}},
			{"Физика", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/1UCoVTqoO7sFV3TJ6lp1-CiziODOb6rH3/view?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/file/d/1pJwJnvpOuR-ViUT7K7BZIsQC7Ra6wtw7/view?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1s2olgS1yTl_cvbawI-RBfsJUh3mErtmq/view?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/file/d/1xRd9IX40LFcnqd_ZvfkmpjWWJCeO_wEp/view?usp=drive_link"},
				{"Лабораторная работа 5", "https://drive.google.com/file/d/1DykDirz2cnNh4zCUDUNxuzRPPwf9YKCc/view?usp=drive_link"},
			}},
			{"Вычислительная математика", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/1XvgiEdNGOx_b82tvn3TgCgihfqVcwPFE/view?usp=drive_link"},
				{"Файлы для ЛР 1", "https://drive.google.com/file/d/1BOWX8k77V3Wp_kitgFZRKLdxYVP1kObc/view?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/file/d/1pheMi9MF16A35CIXNebJ7g3B2ZY-_old/view?usp=drive_link"},
				{"Файлы для ЛР 2", "https://drive.google.com/file/d/1ObbEF5QHD7jy6tBCuGz65m7w8FRYD0ZE/view?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1vzhwagD9LWzCEQGl71dkJxxd90r7BVMg/view?usp=drive_link"},
				{"Файлы для ЛР 3", "https://drive.google.com/file/d/1Vcm0LBmjDcJqmmgWzed4CbbjL-4duw3K/view?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/file/d/1U8UwWR3BckA9FTDhKWHBL8Qy61XAzSvh/view?usp=drive_link"},
				{"Файлы для ЛР 4", "https://drive.google.com/file/d/1uQTKMQh7HeVkl5qXdu03oZWBcK31U8vG/view?usp=drive_link"},
				{"Лабораторная работа 5", "https://drive.google.com/file/d/1Ga89fuATCtKoqBAfalFHCQi1hRU18dnh/view?usp=drive_link"},
				{"Файлы для ЛР 5", "https://drive.google.com/file/d/1ARB2PGf9jbPAPSlX49d_xPPPT4ydyfra/view?usp=drive_link"},
				{"Лабораторная работа 6", "https://drive.google.com/file/d/1EVBy5zoDUDqS6KNc7mf_4S1z6sZFIXwi/view?usp=drive_link"},
				{"Файлы для ЛР 6", "https://drive.google.com/file/d/1_531q_Vm8AKkAAtmQ3vTPkkBxq1SjexY/view?usp=drive_link"},
			}},
		},
	},
	"4": {
		Name: "Семестр 4",
		Subjects: []Subject{
			{"Вычислительная техника", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/drive/folders/1IbL8WzqJy9PlKagsU6cA10NrMztV858m?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/drive/folders/1N4UeHUE-UUzO86Th_LW1t_kN4U4nnD-Y?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/drive/folders/18Z0Hua_JMZM8NaIhTuz8LU3Lur8haoaI?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/drive/folders/11BWt94LeTJL2wT8xA_eK5HvuYQQdvt-G?usp=drive_link"},
				{"Лабораторная работа 5", "https://drive.google.com/drive/folders/1TmkxdTYlahXodLarDzWdnxlZU8PNn3fw?usp=drive_link"},
				{"Лабораторная работа 6", "https://drive.google.com/drive/folders/1o7HUxkVB8akcVT08ltp1WUjE0GKzsn93?usp=drive_link"},
			}},
			{"Базы данных", []Link{
				{"Лабораторная работа 1", "https://docs.google.com/document/d/1tW2u6BbvLhvM9QGzMbFAuUf46_6n3g_F/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Лабораторная работа 2", "https://docs.google.com/document/d/1hGA1NMVNwqR5yfwhQnWBbjPpoKi-KOdG/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Лабораторная работа 3", "https://docs.google.com/document/d/1BGyPyLYdwXpWbCgN6TvEXpOro_cT0fMC/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Лабораторная работа 4", "https://docs.google.com/document/d/1bsvft8gocEN6hipylh-qgJtV-wtWpdvZ/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Лабораторная работа 5", "https://docs.google.com/spreadsheets/d/1P7mg0BzdWwZ06e2l5zhwBOazOuR08APa/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
			{"Генерация текста на основе языковых моделей", []Link{
				{"Курсовая работа", "https://docs.google.com/document/d/1zW8xjDIojivSMCTUvvwNsmSkGQwJJW_q/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Презентация к курсовой работе", "https://docs.google.com/presentation/d/1R3pmFKc_y-oR2HhRhlorWKLXrQ6Ot7sk/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Скринкаст к курсовой работе", "https://drive.google.com/file/d/1it1SSYVwl9PKzmWPX0-uNJgkXt_tDlyD/view?usp=drive_link"},
				{"Код для Триграмм", "https://drive.google.com/file/d/13D1DNPlqZ440xQC8VWT57M5oDJAh2qwv/view?usp=drive_link"},
				{"Код для бигграмм", "https://drive.google.com/file/d/11olxlaCBFvDocdbXXMwTT3WuELUBSIX4/view?usp=drive_link"},
			}},
			{"Операционные системы", []Link{
				{"Лабараторные работы 1-6 Linux", "https://docs.google.com/document/d/1paXKm5fTPiJHC3IO3GzhR2ytln0pBcpc/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Лабараторные работы 1-7 Virtual Box", "https://docs.google.com/document/d/1G_6-xk0KhESp9DpVPEx6CWzsqEqSwxVL/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
		},
	},
	"5": {
		Name: "Семестр 5",
		Subjects: []Subject{
			{"Компьютерная графика", []Link{
				{"Анимационное изображение", "https://drive.google.com/file/d/14FSS5wMAGfiYFwCZEbFFMK632D0SL4c9/view?usp=sharing"},
				{"Векторное изображение", "https://drive.google.com/file/d/16-KW8rieAL0--MAHMTU_RVbouQ2vGt6k/view?usp=sharing"},
				{"Растровое изображение", "https://drive.google.com/file/d/1Qb4g_u-CwChtLisSDBE9MxvR0sHc_Nyd/view?usp=sharing"},
			}},
			{"Программный проект по разработке информационной системы для организации в сфере услуг аренды техники", []Link{
				{"Курсовая работа", "https://docs.google.com/document/d/1l2LrMm_VkforlfH9Ece-FYsHjRQqHKtE/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
			{"Основы бизнес-информатики", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/1p3bqJ2WaI1AYfFA_m7rTAVw2WeVY9yJo/view?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/file/d/1DRMZMlU9YD9CPGD1jAjd_j5dOM1_m2fi/view?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1N31JC1f3tGFOrs0au6h5fDCsZxwFB2Vv/view?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/file/d/1kYrfxqvPecIETOBFYDbh0HdoXhfL2JeR/view?usp=drive_link"},
				{"Лабораторная работа 5", "https://drive.google.com/file/d/1tx1H1_p9gOJqAUEK6t17eOr6l0osnp6i/view?usp=drive_link"},
				{"Лабораторная работа 6", "https://drive.google.com/file/d/11zQQai2G2G9_8OCBLaMTJddbT5VoPUkj/view?usp=drive_link"},
				{"Лабораторная работа 7", "https://drive.google.com/file/d/1fLdLOpFt3aSeFHRO8zza-W5q0cB1p0rO/view?usp=drive_link"},
				{"Лабораторная работа 8", "https://drive.google.com/file/d/1tGbVU60tv_4570hMEN_mcYC5YJyO5sb2/view?usp=drive_link"},
				{"Лабораторная работа 9", "https://drive.google.com/file/d/1Cd6ZhvPUJh6ofGVzaHjDuR7ON7yhlloK/view?usp=drive_link"},
				{"Лабораторная работа 10", "https://drive.google.com/file/d/1MwVYI9k06UXf_3VjSttLM7SQSinVIanJ/view?usp=drive_link"},
			}},
		},
	},
	"6": {
		Name: "Семестр 6",
		Subjects: []Subject{
			{"Математические основы компьютерной графики", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/1_f5m2BvraUw6REIMj26FgkQR36XgIgZf/view?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/file/d/1Dentl_aU3pGNnAZts07Cc9PGEKXICDPU/view?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1ayNGsqnLeFVVwArUAyeftfdc1RvvK1_c/view?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/file/d/1D0MHYBk8AgqyTgyM9boAn9WrzGv5pGJA/view?usp=drive_link"},
			}},
			{"Прикладные информационные технологии", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/1WCBrnzO-rmT0f5N_UnTFlZXEzBXW5Uhh/view?usp=drive_link"},
				{"Лабораторная работа 2", "https://drive.google.com/file/d/1JOfOqgDC0Yn8Wv1pASDUtTnaE7EPDhGG/view?usp=drive_link"},
				{"Лабораторная работа 3", "https://drive.google.com/file/d/1AZ58PZp2mFPF65nmoD6mlre0qqrlZ1Cs/view?usp=drive_link"},
				{"Лабораторная работа 4", "https://drive.google.com/file/d/1TdfUSOakg0FfLnEK2xarTbKAyDi6VbQ1/view?usp=drive_link"},
				{"Лабораторная работа 5", "https://drive.google.com/file/d/1KE-toUVqCP3sxKSBYwvyzpzs4x9xQYgM/view?usp=drive_link"},
			}},
			{"Основы электронного управления", []Link{
				{"Итоговый отчет", "https://docs.google.com/presentation/d/1zGmyyarMAtSgMqdGU4mL_W80gShwvH-GVVeKcRMML_s/edit?usp=drive_link"},
			}},
			{"Программирование", []Link{
				{"Итоговый отчет", "https://docs.google.com/document/d/1OQUjMDdb1SiHPl1QluDVSDeK2Zaa9UXWt3C63ubxhCI/edit"},
			}},
			{"Иностранный язык (английский)", []Link{
				{"Эссе", "https://docs.google.com/document/d/1Tf3qIjb5InCwWMPGKFBI8p-mVgOk8gWV/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
		},
	},
	"7": {
		Name: "Семестр 7",
		Subjects: []Subject{
			{"Электротехника", []Link{
				{"Лабораторная работа 1", "https://docs.google.com/document/d/1LlEZW8B9kYiTcK-WX8TtWt4YEMsL6rQc/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Лабораторная работа 2", "https://docs.google.com/document/d/1TzguE-WcJvZhgKaQPGas_NNHHnFO1sls/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Лабораторная работа 3", "https://docs.google.com/document/d/1CovJSqBgi5Ni60Kak6ZzCJyCkBUlrAGU/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
			{"Практика Производственная (научно-исследовательская работа)", []Link{
				{"Задание на производственную практику", "https://drive.google.com/file/d/11Zt5QNBLn3DbWwZOR2yD5SVZjyGpMLEn/view?usp=drive_link"},
				{"Отчет о прохождении производственной практики", "https://drive.google.com/file/d/1pKkp47svhKVVHQZ71yHT5GHHm3TiUVEv/view?usp=drive_link"},
				{"Курсовая работа", "https://drive.google.com/file/d/1nEqCifj58pPwmq6H_C2Lraxs9XJdXK5-/view?usp=drive_link"},
			}},
			{"Инструментальные средства программирования", []Link{
				{"Лабораторная работа 1", "https://colab.research.google.com/drive/1QD7CO83gdyp1TVLL7PB4X171lFgtnqFJ?usp=drive_link"},
				{"Лабораторная работа 2", "https://colab.research.google.com/drive/1uzeJj5tT8naxBPnWIaEPZczYFt-CdpUD?usp=drive_link"},
				{"Лабораторная работа 3", "https://colab.research.google.com/drive/13D6orSrgiRyvrrHpyuP7qR0xhKqQD3p-?usp=drive_link"},
				{"Лабораторная работа 4-5", "https://colab.research.google.com/drive/1Vn0yRC2qDVElIPx68cJFQSPJ93mIpzo6?usp=drive_link"},
				{"Лабораторная работа 6", "https://colab.research.google.com/drive/1w00P6XVekGf7aOoRNM9vGz2NTS1tVwLm?usp=drive_link"},
				{"Лабораторная работа 7", "https://colab.research.google.com/drive/1tzfzimBK9_P8bnv0Kmg2PF4kEVUo8Kik?usp=drive_link"},
			}},
		},
	},
	"8": {
		Name: "Семестр 8",
		Subjects: []Subject{
			{"Языки написания спецификаций", []Link{
				{"Лабораторная работа 1", "https://drive.google.com/file/d/11Kh3qKOd7c3Wdwjj6i4bBa_6SSplVHk7/view?usp=drive_link"},
				{"Диаграмма последовательности", "https://drive.google.com/file/d/1Y8MMG0gqx5K3mFJsAJgSOADxvH_1blyC/view?usp=drive_link"},
				{"Презентация для диаграммы последовательности", "https://docs.google.com/presentation/d/1qZ-VlcerhB813fa1b5q_rnR0DQTJmfSJZ7_nDi_kWIc/edit?usp=drive_link"},
			}},
			{"Мировые информационные ресурсы и цифровые библиотеки", []Link{
				{"Вариативная работа 1", "https://drive.google.com/file/d/1np5h62fm5twE8RwgcAIT1O_nuLUmCiQn/view?usp=drive_link"},
				{"Вариативная работа 2", "https://drive.google.com/file/d/16GwpY4z2TmC3qqNHImmsmyHUhLBeq7Vl/view?usp=drive_link"},
				{"Вариативная работа 3", "https://drive.google.com/file/d/1QJyoLF0bJik-TSwHGMCk04OnYKIVsWnr/view?usp=drive_link"},
				{"Вариативная работа 4", "https://drive.google.com/file/d/1Y6xBKEQO6eri9oW_IHuiywEne47lL4JU/view?usp=drive_link"},
				{"Вариативная работа 5", "https://docs.google.com/document/d/11FztS-nzuNmmthhAMxMgkpBVaCpszZN8/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Вариативная работа 6", "https://docs.google.com/document/d/1rEy6okjtMFArCld-V5Xz9BoxYlrRmjIX/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Вариативная работа 9", "https://docs.google.com/document/d/1zvSEAgqMIWrb3nyxRfErmNFQQiUcsrt1/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Вариативная работа 10", "https://drive.google.com/file/d/1b2xpId0TmCpsVVSy8pWvmBDZMSC1u3ji/view?usp=drive_link"},
				{"Вариативная работа 11", "https://docs.google.com/document/d/1eAMkEwL0Q-3mwPcOhWjYew2DZbZ-fbAJ/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 1", "https://docs.google.com/document/d/1MSUPpGvCOLk0ilv8uwOxikwRDXqWemYS/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 4", "https://docs.google.com/document/d/13ViBHKVFyWP0Lraftq8fPMDGVPqcdBOG/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 5", "https://drive.google.com/file/d/1zDGO2eJBzOiyAS_O_4TsADXHiGMtAxVi/view?usp=drive_link"},
				{"Инвариативная работа 6", "https://docs.google.com/document/d/1ugbTevUdTw2S3rKiVBE7mc37pjKpD3fB/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 7", "https://docs.google.com/document/d/1Tq3qYapzTgAvNH6DSfvKr9LaLpKelRge/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 8", "https://docs.google.com/document/d/1y6I-DAohUyLAKQ3f20TjUVq_eK0PF0GY/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 9", "https://docs.google.com/document/d/1I8HBYY94fio7K1XgBu_5rNLvQffpqxwd/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 10", "https://docs.google.com/document/d/1BE0_hsj7M3r4uSUBP3T10tpiELOKOLJ_/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 11", "https://docs.google.com/document/d/1RHUG-wdb9jYxyW5ekl-PoE_9BXf1JL1t/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
			{"Социальные и профессиональные вопросы информатики и ИТ", []Link{
				{"Вариативная работа 1", "https://docs.google.com/document/d/1N0X9yXnXiV40ISSJyDGuPAj2Ejr8cn-M/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Вариативная работа 3", "https://docs.google.com/document/d/1kJP_oAD8ipj_FAXE_Oozx8i_HpOPP9NS/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Вариативная работа 5", "https://docs.google.com/document/d/1X_q1sf7Ieu4AoO6zjmBsDhbW3s66CWSv/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Вариативная работа 6", "https://docs.google.com/document/d/1_HFd-2DAAjeYbsJkH8h8BSsRiCk4XG90/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 1", "https://docs.google.com/document/d/118hXv_hyga9Vtb66BGru391QVDL87Cc1/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 2", "https://docs.google.com/document/d/10k78LI_ib92wW5-bKpX1w2h0-fvp8mwo/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 3", "https://docs.google.com/document/d/1mj9UtIwt4i1FZLX1g0KXICCSR5DBFvNg/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 4", "https://drive.google.com/file/d/1PBcvn3cIi5vw5HGGI2gwMKs33LtqrZM6/view?usp=drive_link"},
				{"Инвариативная работа 5", "https://docs.google.com/document/d/179CUOJAHmuKc90z86GS73z7aAKs6LGv8/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
				{"Инвариативная работа 6", "https://docs.google.com/document/d/1CRDwnVXc0ArdKC8KTHQj11uPRwJvGYqf/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
			{"Особенности профеcсиональной иноязычной коммуникации", []Link{
				{"Эссе", "https://docs.google.com/document/d/1NzmVdkJq4MZjjtC1eyh-itpS9AlLSgob/edit?usp=drive_link&ouid=110155577049603111308&rtpof=true&sd=true"},
			}},
		},
	},
}

var mainTemplate = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Курсы</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background-color: #f0f0f0;
            margin: 0;
        }
        .container {
            position: relative; /* Добавлено для позиционирования ссылки */
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
            max-width: 400px;
            text-align: center;
            display: flex;
            flex-direction: column;
        }
        .header {
            font-size: 24px;
            font-weight: bold;
            margin-bottom: 20px;
        }
        .button {
            flex: 1;
            padding: 10px;
            margin: 5px 0;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background 0.3s;
            background-color: #007BFF;
            color: white;
            display: block;
        }
        .button:hover {
            background-color: #0056b3;
        }
        .code-link {
            position: absolute; /* Позиционирование в правом верхнем углу */
            top: 10px;
            right: 10px;
            font-size: 14px;
        }
        .code-link a {
            color: #007BFF;
            text-decoration: none;
        }
        .code-link a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">Портфолио Ахмедова Эдгара Тимуровича студента 4 курса ИВТ</div>
        {{range $index, $course := .Courses}}
            <button class="button" onclick="location.href='/course/{{$index}}'">{{$course.Name}}</button>
        {{end}}
        <div class="code-link">
            <a href="https://github.com/your-repo/your-project" target="_blank">Посмотреть код</a>
        </div>
    </div>
</body>
</html>
`))

var courseTemplate = template.Must(template.New("course").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Course.Name}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background-color: #f0f0f0;
        }
        .container {
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
            display: flex;
            flex-direction: column;
        }
        .button {
            flex: 1;
            padding: 10px;
            margin: 5px 0;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background 0.3s;
            background-color: #007BFF;
            color: white;
            display: block;
        }
        .button:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        {{range $index, $subject := .Course.Subjects}}
            <button class="button" onclick="location.href='/course/{{$.CourseNumber}}/subject/{{$index}}'">{{$subject.Name}}</button>
        {{end}}
        <button class="button" onclick="location.href='/'">Назад</button>
    </div>
</body>
</html>
`))

var subjectTemplate = template.Must(template.New("subject").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Subject.Name}}</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            background-color: #f0f0f0;
        }
        .container {
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
            display: flex;
            flex-direction: column;
        }
        .button {
            flex: 1;
            padding: 10px;
            margin: 5px 0;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background 0.3s;
            background-color: #007BFF;
            color: white;
            display: block;
        }
        .button:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        {{range $link := .Subject.Links}}
            <button class="button" onclick="location.href='{{$link.URL}}'">{{$link.Name}}</button>
        {{end}}
        <button class="button" onclick="location.href='/course/{{.CourseNumber}}'">Назад</button>
    </div>
</body>
</html>
`))

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		mainTemplate.Execute(w, struct {
			Courses map[string]Course
		}{Courses: courses})
		return
	}

	parts := strings.Split(path, "/")
	if len(parts) >= 3 && parts[1] == "course" {
		courseNumber := parts[2]
		course, exists := courses[courseNumber]
		if !exists {
			http.NotFound(w, r)
			return
		}
		if len(parts) == 3 {
			courseTemplate.Execute(w, struct {
				CourseNumber string
				Course       Course
			}{CourseNumber: courseNumber, Course: course})
		} else if len(parts) == 5 && parts[3] == "subject" {
			subjectIndex, err := strconv.Atoi(parts[4])
			if err != nil || subjectIndex < 0 || subjectIndex >= len(course.Subjects) {
				http.NotFound(w, r)
				return
			}
			subject := course.Subjects[subjectIndex]
			subjectTemplate.Execute(w, struct {
				CourseNumber string
				Subject      Subject
			}{CourseNumber: courseNumber, Subject: subject})
		} else {
			http.NotFound(w, r)
		}
		return
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", Handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
