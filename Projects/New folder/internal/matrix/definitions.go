package matrix

type ProjectType string
type Complexity string
type Language string
type Framework string

const (
	WebApp  ProjectType = "Web Application"
	CLI     ProjectType = "CLI Tool"
	Backend ProjectType = "Backend Service"

	Minimal    Complexity = "Minimal (MVP)"
	Standard   Complexity = "Standard (Clean Architecture)"
	Enterprise Complexity = "Enterprise (Microservices Ready)"

	Go     Language = "Go"
	JS     Language = "JavaScript"
	Python Language = "Python"

	// Go Frameworks
	Gin   Framework = "Gin"
	Echo  Framework = "Echo"
	Fiber Framework = "Fiber"

	// JS Frameworks
	Express Framework = "Express"
	Fastify Framework = "Fastify"

	// Python Frameworks
	Flask   Framework = "Flask"
	FastAPI Framework = "FastAPI"
	Django  Framework = "Django"
)

type FileTemplate struct {
	Path    string
	Content string
}

type ProjectMatrix struct {
	Files []FileTemplate
}

func GetMatrix(lang Language, fw Framework, pt ProjectType, c Complexity) ProjectMatrix {
	var files []FileTemplate

	switch lang {
	case Go:
		files = getGoMatrix(fw, pt, c)
	case JS:
		files = getJSMatrix(fw, pt, c)
	case Python:
		files = getPythonMatrix(fw, pt, c)
	}

	return ProjectMatrix{Files: files}
}

func getGoMatrix(fw Framework, pt ProjectType, c Complexity) []FileTemplate {
	var files []FileTemplate
	mainContent := ""

	switch fw {
	case Gin:
		mainContent = `package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}`
	case Echo:
		mainContent = `package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}`
	case Fiber:
		mainContent = `package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()
    app.Get("/", func(c *fiber.Context) error {
        return c.SendString("Hello, Fiber!")
    })
    app.Listen(":3000")
}`
	default:
		mainContent = `package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}`
	}

	files = append(files, FileTemplate{Path: "main.go", Content: mainContent})
	files = append(files, FileTemplate{Path: "go.mod", Content: "module app\n\ngo 1.21"})

	if c == Standard || c == Enterprise {
		files = append(files, FileTemplate{Path: "internal/repository/repo.go", Content: "package repository"})
		files = append(files, FileTemplate{Path: "internal/service/service.go", Content: "package service"})
	}

	return files
}

func getJSMatrix(fw Framework, pt ProjectType, c Complexity) []FileTemplate {
	var files []FileTemplate
	mainContent := ""

	switch fw {
	case Express:
		mainContent = `const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
  res.send('Hello Express!')
})

app.listen(port, () => {
  console.log(` + "`Example app listening on port ${port}`" + `)
})`
	case Fastify:
		mainContent = `const fastify = require('fastify')({ logger: true })

fastify.get('/', async (request, reply) => {
  return { hello: 'world' }
})

const start = async () => {
  try {
    await fastify.listen({ port: 3000 })
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }
}
start()`
	}

	files = append(files, FileTemplate{Path: "index.js", Content: mainContent})
	files = append(files, FileTemplate{Path: "package.json", Content: `{"name": "app", "version": "1.0.0", "scripts": {"start": "node index.js"}}`})

	return files
}

func getPythonMatrix(fw Framework, pt ProjectType, c Complexity) []FileTemplate {
	var files []FileTemplate
	mainContent := ""

	switch fw {
	case Flask:
		mainContent = `from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, Flask!'

if __name__ == '__main__':
    app.run()`
	case FastAPI:
		mainContent = `from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello FastAPI"}
`
	case Django:
		mainContent = `# Django project entry point
import os
import sys

def main():
    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'app.settings')
    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError("Couldn't import Django.") from exc
    execute_from_command_line(sys.argv)

if __name__ == '__main__':
    main()`
	}

	files = append(files, FileTemplate{Path: "app.py", Content: mainContent})
	files = append(files, FileTemplate{Path: "requirements.txt", Content: string(fw)})

	return files
}
