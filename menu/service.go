package menu

import (
	"errors"
	"fmt"
	"rest-api-coffee-server/database"
	"strconv"
)

//Для наглядности написал сразу тут
var DrinkNotFound = errors.New("напиток не найден")
var InvalidID = errors.New("неверный формат ID")

func GetAll() ([]DrinkList, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке списка напитков: %w", err)
	}

	items := make([]DrinkList, 0, len(records))
	for _, record := range records {
		items = append(items, DrinkList{
			ID:    record.ID,
			Name:  record.Name,
			Price: record.Price,
		})
	}

	return items, nil
}

func AvailableDrinks() ([]DrinkList, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке напитков: %w", err)
	}

	items := make([]DrinkList, 0, len(records))
	for _, record := range records {
		if record.InStock == true {
			items = append(items, DrinkList{
				ID:    record.ID,
				Name:  record.Name,
				Price: record.Price,
			})
		}
	}

	return items, nil
}

func GetByID(id int) (*Drink, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке списка напитков: %w", err)
	}

	for _, record := range records {
		if record.ID == id {
			return &Drink{
				ID:               record.ID,
				Name:             record.Name,
				Price:            record.Price,
				InStock:          record.InStock,
				ContainsCaffeine: record.ContainsCaffeine,
				Volume:           record.Volume,
				Description:      record.Description,
			}, nil
		}
	}

	return nil, DrinkNotFound
}

func generateID(records []database.Record) int {
	if len(records) == 0 {
		return 1
	}

	maxID := 0
	for _, record := range records {
		if record.ID > maxID {
			maxID = record.ID
		}
	}

	return maxID + 1
}

func Add(req DrinkCreate) (*Drink, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке списка напитков: %w", err)
	}

	newID := generateID(records)

	drink := &Drink{
		ID:               newID,
		Name:             req.Name,
		Price:            req.Price,
		InStock:          req.InStock,
		ContainsCaffeine: req.ContainsCaffeine,
		Volume:           req.Volume,
		Description:      req.Description,
	}

	newDrink := database.Record{
		ID:               drink.ID,
		Name:             drink.Name,
		Price:            drink.Price,
		InStock:          drink.InStock,
		ContainsCaffeine: drink.ContainsCaffeine,
		Volume:           drink.Volume,
		Description:      drink.Description,
	}
	records = append(records, newDrink)

	if err := database.SaveRecords(records); err != nil {
		return nil, fmt.Errorf("ошибка при сохранении напитка: %w", err)
	}

	return drink, nil
}

func Delete(id string) error {
	//--------Ошибка 1--------
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("%w:%s", InvalidID, id)
	}
	//------------------------

	records, err := database.LoadRecords()
	if err != nil {
		return fmt.Errorf("ошибка при загрузке списка напитков:%w", err)
	}

	found := false
	var newRecords []database.Record
	for _, record := range records {
		if record.ID != idInt {
			newRecords = append(newRecords, record)
		} else {
			found = true
		}
	}

	if !found {
		return DrinkNotFound
	}

	if err := database.SaveRecords(newRecords); err != nil {
		return fmt.Errorf("ошибка при сохранении напитков:%w", err)
	}

	return nil
}

func Update(req DrinkUpdate) (*Drink, error) {
	records, err := database.LoadRecords()
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке списка напитков:%w", err)
	}

	found := false
	for i, record := range records {
		if record.ID == req.ID {
			records[i].Name = req.Name
			records[i].Price = req.Price
			records[i].InStock = req.InStock
			records[i].ContainsCaffeine = req.ContainsCaffeine
			records[i].Volume = req.Volume
			records[i].Description = req.Description
			found = true
			break
		}
	}

	if !found {
		return nil, DrinkNotFound
	}

	if err := database.SaveRecords(records); err != nil {
		return nil, fmt.Errorf("ошибка при сохранении напитка:%w", err)
	}

	return &Drink{
		ID:               req.ID,
		Name:             req.Name,
		Price:            req.Price,
		InStock:          req.InStock,
		ContainsCaffeine: req.ContainsCaffeine,
		Volume:           req.Volume,
		Description:      req.Description,
	}, nil
}
