package routes

import (
    "github.com/gofiber/fiber/v2"
)

var UseJobSyncRouter = func(app *fiber.App) {
    app.Get("/jobsync", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Get all invoices from specific creator ID")
        // TODO: Fetch from DB logic
    })
    app.Post("/jobsync", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Invoice created")
        // TODO: Write to DB logic
    })
    app.Get("/jobsync/{id}", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Get invoice from creator by ID")
        // TODO: Fetch by ID logic
    })
    app.Put("/jobsync/{id}", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Updated invoice with ID of _")
        // TODO: Update by ID logic
    })
    app.Delete("/jobsync/{id}", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Deleted invoice with ID of _")
        // TODO: Delete by ID logic
    })
}

