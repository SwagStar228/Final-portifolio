package main

import (
	"html/template"
	"net/http"
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
			{"Предмет 1", []Link{
				{"Ссылка 1", "https://example.com/course1/subject1/link1"},
			},
			}},
	},
	"2": {
		Name: "Семестр 2",
		Subjects: []Subject{
			{"Практика Учебная (эксплуатационная)", []Link{
				{"Портфолио по эксплуатационной практике", "https://github.com/SwagStar228/SwagStar228.AET_pratic.io"},
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
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
            max-width: 400px;
            text-align: center;
        }
        .header {
            font-size: 24px;
            font-weight: bold;
            margin-bottom: 20px;
        }
        .button {
            width: 100%;
            padding: 10px;
            margin: 5px 0;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background 0.3s;
            background-color: #007BFF;
            color: white;
        }
        .button:hover {
            background-color: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">Портфолио Ахмедова Эдгара Тимуровича студента 4 курса ИВТ</div>
        {{range $index, $course := .Courses}}
            <button class="button" onclick="location.href='/course/{{$index}}'">{{$course.Name}}</button>
        {{end}}
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
        }
        .button {
            width: 100%;
            padding: 10px;
            margin: 5px 0;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background 0.3s;
            background-color: #007BFF;
            color: white;
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
        }
        .button {
            width: 100%;
            padding: 10px;
            margin: 5px 0;
            font-size: 16px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            transition: background 0.3s;
            background-color: #007BFF;
            color: white;
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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mainTemplate.Execute(w, struct {
			Courses map[string]Course
		}{Courses: courses})
	})

	http.HandleFunc("/course/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) >= 3 {
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
			}
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
