package user

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Handler interface {
	Get(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	Update(*fiber.Ctx) error
}

type handler struct {
	service Service
}

var _ Handler = handler{}

func NewHandler(service Service) Handler {
	return handler{service: service}
}

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func (h handler) Get(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	span := tracer.StartSpan("web.request/test123lelelelelelelele", tracer.ResourceName("/Get"))
	defer span.Finish()
	log.Printf("my log message %v", span)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}
	model, err := h.service.Get(c.UserContext(), uint(id))
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}
	return c.Status(200).JSON(Response{Data: model})
}

func (h handler) Create(c *fiber.Ctx) error {
	model := Model{}
	err := c.BodyParser(&model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})

	}
	_, err = h.service.Create(c.UserContext(), model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}
	return c.SendStatus(201)
}

func (h handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()}) //TODO change error
	}
	err = h.service.Delete(c.UserContext(), uint(id))
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}
	return c.SendStatus(200)
}

func (h handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()}) //TODO change error
	}
	model := Model{}
	err = c.BodyParser(&model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}
	err = h.service.Update(c.UserContext(), uint(id), model)
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}
	return c.SendStatus(200)
}
