package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/drizlye0/GopherSocial/internal/store"
)

var usernames = []string{
	"shadowWolf93", "crazyPenguin88", "fastFalcon21", "blueTiger55", "silentNinja09", "luckyBear77", "darkViper12", "tinyKoala42", "mysticFox66", "wildHawk34", "happyOtter19", "fuzzyPanda08", "stealthCat73", "redDragon11", "ghostRider28", "braveEagle90", "noisyLlama62", "frozenShark07", "electricCobra25", "sleepyMoose01",
}

var titles = []string{
	"Cómo empezar con Go", "Trucos para ser más productivo", "Lo que nadie te dijo del freelance", "Errores comunes al programar", "Guía rápida de Git", "Aprende Docker en 10 minutos", "¿Vale la pena aprender Rust?", "Tips para entrevistas técnicas", "Organiza tu día como un pro", "Diseña APIs efectivas", "React vs Vue: ¿Cuál elegir?",
	"Introducción a WebAssembly", "Evita el burnout programando", "Tu primer servidor con Go", "Scrum sin complicaciones", "Manejo de errores elegante", "Automatiza con scripts Bash", "Consejos para trabajar remoto", "Cómo escribir código limpio", "Lo básico de bases de datos",
}

var contents = []string{
	"Aprende los pasos básicos para iniciarte en el lenguaje Go sin complicaciones.",
	"Te comparto algunos hábitos sencillos que aumentaron mi productividad diaria.",
	"Una reflexión personal sobre lo que implica trabajar por cuenta propia en tecnología.",
	"Descubre los fallos más comunes que cometemos al comenzar a programar.",
	"Una guía rápida y clara para usar Git sin enredos.",
	"Resumen práctico para empezar a usar Docker en tus proyectos personales.",
	"Comparo las ventajas y desventajas de aprender Rust en 2025.",
	"Prepárate mejor para entrevistas técnicas con estos consejos directos y efectivos.",
	"Te muestro cómo estructurar tu jornada para trabajar con foco y menos estrés.",
	"Buenas prácticas para diseñar APIs fáciles de usar y mantener.",
	"Analizamos las diferencias clave entre React y Vue para ayudarte a decidir.",
	"Explico qué es WebAssembly y por qué podría cambiar el desarrollo web.",
	"Comparto señales y soluciones para evitar el agotamiento en programación.",
	"Tutorial paso a paso para crear un servidor web básico con Go.",
	"Desmitificamos Scrum y te enseño cómo aplicarlo sin burocracia.",
	"Aprende a manejar errores de forma clara y segura en tu código.",
	"Automatiza tareas repetitivas con simples scripts Bash.",
	"Consejos útiles para que el trabajo remoto sea realmente eficiente.",
	"Descubre cómo escribir código legible y fácil de mantener.",
	"Un repaso simple por los conceptos clave de bases de datos relacionales.",
}

var comments = []string{
	"Muy buen artículo, justo lo que necesitaba.", "No sabía esto, gracias por compartir!", "¿Podrías profundizar más en este tema?", "Excelente explicación, todo claro.", "Me encantó cómo resumiste el concepto.", "Lo aplicaré en mi próximo proyecto.", "¡Gracias! Me ahorraste mucho tiempo.", "Esto debería enseñarse en todas partes.", "Una lectura rápida pero muy útil.", "¿Tienes algún ejemplo práctico?", "Justo estaba buscando algo así.", "Qué bueno que lo explicaste sin rodeos.", "Lo guardo en favoritos. ¡Gracias!", "Muy bien escrito, sigue así.", "Nunca lo había pensado de esa forma.", "¿Podrías recomendar más recursos?", "Esto resolvió una duda que tenía hace días.", "¡Buen trabajo! Me suscribo al blog.", "Conciso y directo, me encantó.", "Siento que aprendí algo nuevo hoy.",
}

var tags = []string{
	"go", "docker", "productividad", "freelance", "git", "rust", "entrevistas", "aprender", "api", "react", "vue", "wasm", "scrum", "bash", "remoto", "backend", "frontend", "basesdedatos", "codigo", "tutorial",
}

func Seed(s *store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := s.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Printf("error on generate user: %s", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := s.Posts.Create(ctx, post); err != nil {
			log.Printf("error on generate post: %s", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := s.Comments.Create(ctx, comment); err != nil {
			log.Printf("error on generate comment: %s", err)
			return
		}
	}

	log.Println("seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := range num {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := range num {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := range num {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
