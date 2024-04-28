package routes

import (
	"math/rand"
	"net/http"
	"squeeze-auction/internal/errors"
	"squeeze-auction/internal/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (r *Routes) RegisterItemsRoutes(g *echo.Group) {
	g.GET("/", r.getItemsList)
	g.POST("/", r.createItem)
	g.GET("/:id", r.getItem)
	g.PUT("/:id", r.updateItem)
	g.DELETE("/:id", r.deleteItem)
}

// getItemsList retrieves a list of items.
// Returns a JSON array of item details.
// @summary Get Items List
// @tags Items
// @description Retrieves a list of all available items.
// @accept json
// @produce json
// @success 200 {array} models.Item "List of items"
// @failure 500 {object} errors.Error "Internal server error"
// @router /items/ [get]
func (r *Routes) getItemsList(c echo.Context) error {
	r.Log.Info("Get items list")
	items := []*models.Item{
		{
			ID:           rand.Intn(999_999) + 1,
			Title:        "Item 1",
			InitialPrice: 100,
			Description:  "Description of item 1",
		},
		{
			ID:           rand.Intn(999_999) + 1,
			Title:        "Item 2",
			InitialPrice: 200,
			Description:  "Description of item 2",
		},
	}

	return c.JSON(http.StatusOK, &items)
}

// createItem creates a new item based on the provided details.
// It returns the created item details or an error if the creation fails.
// @summary Create Item
// @tags Items
// @description Creates a new item based on the provided request body.
// @accept json
// @produce json
// @param item body models.Item true "Item details"
// @success 201 {object} models.Item "Item created successfully"
// @failure 400 {object} errors.Error "Bad Request: Failed to parse request body or validation failed"
// @router /items/ [post]
func (r *Routes) createItem(c echo.Context) error {
	r.Log.Infof("Creating item ...")

	req := new(models.Item)

	err := c.Bind(req)
	if err != nil {
		r.Log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, errors.NewError("INCORRECT_REQUEST_BODY", "Failed to parse request body"))
	}

	err = req.Validate()
	if err != nil {
		r.Log.Error("validation failed: ", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.ValidationFailedErr.Code, err.Error()))
	}

	req.ID = rand.Intn(999_999) + 1
	r.Log.Infof("Item created: %+v", req)

	return c.JSON(http.StatusCreated, &req)
}

// getItem retrieves an item by its ID.
// Returns a JSON object of the item details if found.
// @summary Get Item
// @tags Items
// @description Retrieves the details of an item by its ID.
// @accept json
// @produce json
// @param id path int true "ID of the item to retrieve"
// @success 200 {object} models.Item "Item retrieved successfully"
// @failure 404 {object} errors.Error "Item not found"
// @router /items/{id} [get]
func (r *Routes) getItem(c echo.Context) error {
	r.Log.Infof("Get item with id: %s", c.Param("id"))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.Log.Error("failed to convert id to int!!!", err)
		return c.JSON(http.StatusBadRequest, errors.NewError("INVALID_ID", "Invalid ID"))
	}

	return c.JSON(http.StatusOK, &models.Item{
		ID:           id,
		Title:        "Item 1",
		InitialPrice: 100,
		Description:  "Description of item 1",
	})
}

// updateItem updates an item by its ID based on the provided details.
// Returns the updated item details or an error if the update fails.
// @summary Update Item
// @tags Items
// @description Updates the details of an item based on the provided request body and item ID.
// @accept json
// @produce json
// @param id path int true "ID of the item to update"
// @param item body models.Item true "Updated item details"
// @success 200 {object} models.Item "Item updated successfully"
// @failure 400 {object} errors.Error "Bad Request: Failed to parse request body or validation failed"
// @router /items/{id} [put]
func (r *Routes) updateItem(c echo.Context) error {
	r.Log.Infof("Update item with id: %s", c.Param("id"))

	req := new(models.Item)

	err := c.Bind(req)
	if err != nil {
		r.Log.Error("failed to parse request body", err)
		return c.JSON(http.StatusBadRequest, errors.NewError("INCORRECT_REQUEST_BODY", "Failed to parse request body"))
	}

	err = req.Validate()
	if err != nil {
		r.Log.Error("validation failed: ", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.ValidationFailedErr.Code, err.Error()))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.Log.Error("failed to convert id to int!!!", err)
		return c.JSON(http.StatusBadRequest, errors.NewError("INVALID_ID", "Invalid ID"))
	}
	req.ID = id

	return c.JSON(http.StatusOK, &req)
}

// deleteItem deletes an item by its ID.
// Returns no content if the deletion is successful, or an error otherwise.
// @summary Delete Item
// @tags Items
// @description Deletes an item by its ID.
// @accept json
// @produce json
// @param id path int true "ID of the item to delete"
// @success 204 "Item deleted successfully"
// @failure 400 {object} errors.Error "Bad Request: Invalid ID"
// @router /items/{id} [delete]
func (r *Routes) deleteItem(c echo.Context) error {
	r.Log.Infof("Delete item with id: %s", c.Param("id"))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.Log.Error("failed to convert id to int!!!", err)
		return c.JSON(http.StatusBadRequest, errors.NewError("INVALID_ID", "Invalid ID"))
	}
	r.Log.Infof("Item with id %d deleted", id)

	return c.NoContent(http.StatusNoContent)
}
