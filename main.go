package main
import(
	"log"
	"github.com/gofiber/fiber/v2"
	"time"
	"encoding/json"
	"os"
)
type todo struct {
	ID          int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}
var todos []todo
var nextID int = 1
func saveTodos() {
	file, _ := json.MarshalIndent(todos, "", " ")
	_ = os.WriteFile("todos.json", file, 0644)
}
func loadTodos() {
	if _, err := os.Stat("todos.json"); err == nil {
		file, _ := os.ReadFile("todos.json")
		json.Unmarshal(file, &todos)
		if len(todos) > 0 {
			nextID = todos[len(todos)-1].ID + 1
		}
	}
}
func main() {
	app := fiber.New()
	loadTodos()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API running!")
	})
	app.Post("/todos", func(c *fiber.Ctx) error {
		var input struct {
			Title string `json:"title"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
		}
		todo := todo{
			ID:        nextID,
			Title:     input.Title,
			Completed: false,
			CreatedAt: time.Now(),
		}
		todos = append(todos, todo)
		saveTodos()
		nextID++
		return c.Status(fiber.StatusCreated).JSON(todo)
	})
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})
	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		for i, t := range todos {
			if t.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				saveTodos()
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	})
	app.Patch("/todos/:id", func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		found := false
		for i := range todos {
			isTarget := todos[i].ID == id
			todos[i].Completed = isTarget 
			if isTarget {
				found = true
			}
		}
		if !found {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
				
		}
		saveTodos()
		return c.JSON(fiber.Map{"message": "Todo updated successfully"})
	})
	app.Get("/todos/count", func(c *fiber.Ctx) error {
		total := len(todos)
		completed := 0
		for _, t := range todos {
			if t.Completed {
				completed++
			}
		}
		return c.JSON(fiber.Map{
			"total":     total,
			"completed": completed,
		})
	})
	log.Fatal(app.Listen(":3000"))
}