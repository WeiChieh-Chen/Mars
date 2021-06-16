package main

import (
	"github.com/Yimismi/sql2go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type (
	Input struct {
		SQL string
	}

	Output struct {
		Code   int
		Result string
	}
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("", sqlToStruct)

	log.Fatal(r.Run("127.0.0.1:36988"))
}

func sqlToStruct(g *gin.Context) {
	var i Input

	if err := g.ShouldBind(&i); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	sql2go.GoXormTmp = `
	{{- range .Tables -}}
		type {{TableMapper .Name}} struct {
		{{$table := .}}
			{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{ColMapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
		{{end}}
		}
	{{end}}
	`
	args := sql2go.NewConvertArgs().SetGenJson(true).SetGenXorm(true)

	code, err := sql2go.FromSql(i.SQL, args)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusOK, Output{Code: 999, Result: string(code)})
}
