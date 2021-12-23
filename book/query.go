package book

import (
	"asb_microservice_with_golang/config"
	"asb_microservice_with_golang/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	table          = "book"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll Book
func GetAll(ctx context.Context, data map[string]string) ([]models.Book, error) {
	is_first := false
	is_init := true
	var book []models.Book
	param_query := ``

	// scanning
	for key, val := range data {
		if val == "" {
			delete(data, key)
		} else {
			is_first = true
		}
	}
	if is_first == true {
		param_query += `where`
	}

	for key, val := range data {
		if key != "sortByTitle" {
			if (key == "minyear") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					release_year >= %d
				`, val)
				is_init = false
			} else if (key == "minyear") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and release_year >= %d
				`, val)
			}

			if (key == "maxyear") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					release_year <= %d
				`, val)
				is_init = false
			} else if (key == "maxyear") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and release_year <= %d
				`, val)
			}

			if (key == "maxpage") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					total_page <= %d
				`, val)
				is_init = false
			} else if (key == "maxpage") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and total_page <= %d
				`, val)
			}
			if (key == "minpage") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					total_page >= %d
				`, val)
				is_init = false
			} else if (key == "minpage") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and total_page >= %d
				`, val)
			}

			if (key == "title") && (is_init == true) {
				param_query += fmt.Sprintf(`
					%v = '%v'
				`, key, val)
				is_init = false
			} else if (key == "title") && (is_init == false) {
				param_query += fmt.Sprintf(`
					and %v = '%v'
				`, key, val)
			}
		}
	}
	// order by query params
	if data["sortbytitle"] != "" {
		param_query += fmt.Sprintf(`
		order by title %v
		`, data["sortbytitle"])
	}

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	// parsing hasil olahan
	query := `
		SELECT 
			*
		FROM 
			book
	
	`
	query += param_query
	queryText := fmt.Sprintf(query)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var buku models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&buku.ID,
			&buku.Title,
			&buku.Description,
			&buku.ImageUrl,
			&buku.ReleaseYear,
			&buku.Price,
			&buku.TotalPage,
			&buku.Thickness,
			&createdAt,
			&updatedAt,
			&buku.CategoryID,
			&buku.TokoID,
		); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		buku.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		buku.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		book = append(book, buku)
	}

	return book, nil
}

func GetBookFilterByCategory(ctx context.Context, id string, data map[string]string) ([]models.Book, error) {

	var book []models.Book
	is_init := false
	param_query := ``

	// scanning
	for key, val := range data {
		if val == "" {
			delete(data, key)
		}
	}
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	for key, val := range data {
		if key != "sortByTitle" {
			if (key == "minyear") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					release_year >= %d
				`, val)
				is_init = false
			} else if (key == "minyear") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and release_year >= %d
				`, val)
			}

			if (key == "maxyear") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					release_year <= %d
				`, val)
				is_init = false
			} else if (key == "maxyear") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and release_year <= %d
				`, val)
			}

			if (key == "maxpage") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					total_page <= %d
				`, val)
				is_init = false
			} else if (key == "maxpage") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and total_page <= %d
				`, val)
			}
			if (key == "minpage") && (is_init == true) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
					total_page >= %d
				`, val)
				is_init = false
			} else if (key == "minpage") && (is_init == false) {
				var val, _ = strconv.Atoi(val)
				param_query += fmt.Sprintf(`
				and total_page >= %d
				`, val)
			}

			if (key == "title") && (is_init == true) {
				param_query += fmt.Sprintf(`
					%v = '%v'
				`, key, val)
				is_init = false
			} else if (key == "title") && (is_init == false) {
				param_query += fmt.Sprintf(`
					and %v = '%v'
				`, key, val)
			}
		}
	}
	// order by query params
	if data["sortbytitle"] != "" {
		param_query += fmt.Sprintf(`
		order by title %v
		`, data["sortbytitle"])
	}

	query := `
		SELECT 
			*
		FROM 
			book 
		WHERE 
			category_id = "%s"
	`
	queryText := fmt.Sprintf(query, id)
	queryText += param_query
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var buku models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&buku.ID,
			&buku.Title,
			&buku.Description,
			&buku.ImageUrl,
			&buku.ReleaseYear,
			&buku.Price,
			&buku.TotalPage,
			&buku.Thickness,
			&createdAt,
			&updatedAt,
			&buku.CategoryID,
		); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		buku.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		buku.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		book = append(book, buku)
	}

	return book, nil
}

// Insert book
func Insert(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf(`
		INSERT INTO %v
			(title, description, image_url, release_year,
			price, total_page, thickness, 
			created_at, updated_at,
			category_id)
		values
			('%v', '%v', '%v', %d,
			'%v', %d, '%v',
			NOW(), NOW() ,
			%d)
		`,
		table,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update Book
func Update(ctx context.Context, book models.Book, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	// konversi thickness
	if book.TotalPage <= 100 {
		book.Thickness = "tipis"
	} else if (book.TotalPage > 100) && (book.TotalPage <= 200) {
		book.Thickness = "sedang"
	} else if book.TotalPage > 200 {
		book.Thickness = "tebal"
	}

	queryText := fmt.Sprintf(
		` UPDATE %v
		set
			title ='%s',
			description = '%s',
			image_url = '%s',
			release_year = %d,
			price = '%s',
			total_page = %d,
			thickness = '%s',
			updated_at = NOW(),
			category_id = %d
		where
			id = %s`,
		table,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		id,
	)
	// fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete Book
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ditemukan")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
