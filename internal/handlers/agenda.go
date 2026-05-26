package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mosquilab/internal/models"
	"mosquilab/internal/repository"
)

type AgendaHandler struct {
	events *repository.EventRepo
}

func NewAgendaHandler(events *repository.EventRepo) *AgendaHandler {
	return &AgendaHandler{events: events}
}

// GET /api/agenda — public
func (h *AgendaHandler) ListPublic(c *gin.Context) {
	events, err := h.events.ListFuture()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// GET /api/admin/agenda
func (h *AgendaHandler) ListAll(c *gin.Context) {
	events, err := h.events.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

type eventRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	EventDate   string `json:"event_date" binding:"required"`
	EventTime   string `json:"event_time" binding:"required"`
	EventType   string `json:"event_type"`
}

// POST /api/admin/agenda
func (h *AgendaHandler) Create(c *gin.Context) {
	var req eventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ev := &models.Event{
		Title: req.Title, Description: req.Description, Location: req.Location,
		EventDate: req.EventDate, EventTime: req.EventTime, EventType: req.EventType,
	}
	created, err := h.events.Create(ev)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// PUT /api/admin/agenda/:id
func (h *AgendaHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	var req eventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ev := &models.Event{
		Title: req.Title, Description: req.Description, Location: req.Location,
		EventDate: req.EventDate, EventTime: req.EventTime, EventType: req.EventType,
	}
	updated, err := h.events.Update(id, ev)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DELETE /api/admin/agenda/:id
func (h *AgendaHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	if err := h.events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "evento excluído"})
}
