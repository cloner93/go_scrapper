# 📌 GoByExample Scraper  

A simple web scraper built with [Colly](https://github.com/gocolly/colly) to extract and save Go code examples from [GoByExample](https://gobyexample.com/).  

## ✨ Features  
- Scrapes all Go example pages from `gobyexample.com`.  
- Extracts both documentation and code snippets.  
- Saves extracted content as `.go` files in a structured directory.  
- Ensures existing files are not overwritten.  
- Uses proper file naming conventions.  

## 🏗️ Installation  

Ensure you have **Go** installed. If not, download it from [golang.org](https://go.dev/).  

Clone the repository:  
```sh
git clone https://github.com/YOUR_USERNAME/gobyexample-scraper.git
cd gobyexample-scraper
```

Install dependencies:  
```sh
go mod tidy
```

## 🚀 Usage  

Run the scraper:  
```sh
go run main.go
```

This will:  
1. Visit `https://gobyexample.com/` and scrape all example links.  
2. Extract documentation and code.  
3. Save them as `.go` files inside the `example_dir/` folder.  

Example file structure after running the script:  
```
example_dir/
├── Hello-World.go
├── Variables.go
├── Constants.go
...
```

## ⚙️ Configuration  

You can change the save directory by modifying the `saveDir` constant in `main.go`:  
```go
const saveDir = "my_examples/"
```

## 📌 Dependencies  

- [Colly](https://github.com/gocolly/colly) – Fast and flexible web scraping framework.  
- [golang.org/x/text](https://pkg.go.dev/golang.org/x/text) – Handles proper text casing.  

Install dependencies manually with:  
```sh
go get github.com/gocolly/colly
go get golang.org/x/text
```

## 📜 License  

This project is licensed under the **MIT License**. Feel free to modify and share!  

