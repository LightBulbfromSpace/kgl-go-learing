# golang-learning

Exercises made during the course [Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/).  Names of folders match the chapters of course. Inside folders are contained version of code reflecting progress during each lesson.

Упражнения, сделанные во время прохождения [Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/). Названия папок соответствуют разделам курса. Внутри папок содержатся версии кода, отражающие прогресс в течение урока.

# About me

Student of mathematical and information direction. Interested in backend development.

Студент математического и информационного направления. Заинтересован бэкенд-разработке.

# Learned lessons
{{range .}}
### {{.Package}} - {{.TestCoverage}}
{{.Description}}{{range .Functions}}
<details>
  <summary><code>{{.Interface}}</code></summary>
{{range .DocLines}}
    {{.}}{{end}}
</details>
{{end}}{{end}}