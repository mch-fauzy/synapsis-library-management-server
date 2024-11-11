package repositories

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/books/models"
	"gorm.io/gorm"
)

func (r *Repository) CreateBook(createBook *models.Book) error {
	createdBook := r.PostgreSqlConn.Db.Create(createBook)
	err := createdBook.Error
	if err != nil {
		log.Error().Err(err).Msg("[CreateBook] Repository error creating book")
		return err
	}

	return nil
}

func (r *Repository) GetBooksByFilter(filter models.Filter) ([]models.Book, int64, error) {
	var books []models.Book
	var totalBooks int64
	var err error

	// Start a transaction
	err = r.PostgreSqlConn.Db.Transaction(func(tx *gorm.DB) error {
		// Start building the query with the transaction context (tx)
		query := tx.Model(&models.Book{})
		countQuery := tx.Model(&models.Book{})

		// Handle select specific fields
		if len(filter.SelectFields) > 0 {
			query = query.Select(filter.SelectFields)
			countQuery = countQuery.Select(filter.SelectFields)
		}

		// Handle filter fields (where conditions)
		if len(filter.FilterFields) > 0 {
			for _, filterField := range filter.FilterFields {
				switch filterField.Operator {
				case models.OperatorEqual:
					query = query.Where(filterField.Field+" = ?", filterField.Value)
					countQuery = countQuery.Where(filterField.Field+" = ?", filterField.Value)
				case models.OperatorBetween:
					values, ok := filterField.Value.([]interface{})
					if ok && len(values) == 2 {
						query = query.Where(filterField.Field+" BETWEEN ? AND ?", values[0], values[1])
						countQuery = countQuery.Where(filterField.Field+" BETWEEN ? AND ?", values[0], values[1])
					}
				case models.OperatorIn:
					query = query.Where(filterField.Field+" IN ?", filterField.Value)
					countQuery = countQuery.Where(filterField.Field+" IN ?", filterField.Value)
				case models.OperatorIsNull:
					query = query.Where(filterField.Field + " IS NULL")
					countQuery = countQuery.Where(filterField.Field + " IS NULL")
				case models.OperatorNot:
					query = query.Where(filterField.Field+" != ?", filterField.Value)
					countQuery = countQuery.Where(filterField.Field+" != ?", filterField.Value)
				default:
					log.Warn().Msgf("[GetBooksByFilter] Unsupported filter operator: %s for field: %s", filterField.Operator, filterField.Field)
				}
			}
		}

		// Handle pagination
		if filter.Pagination.Page > 0 && filter.Pagination.PageSize > 0 {
			offset := (filter.Pagination.Page - 1) * filter.Pagination.PageSize
			query = query.Offset(offset).Limit(filter.Pagination.PageSize)
		}

		// Handle sorting
		if len(filter.Sorts) > 0 {
			for _, sort := range filter.Sorts {
				switch sort.Order {
				case models.SortAsc:
					query = query.Order(sort.Field + " asc")
				case models.SortDesc:
					query = query.Order(sort.Field + " desc")
				default:
					log.Warn().Msgf("[GetBooksByFilter] Unknown sort order: %s for field: %s", sort.Order, sort.Field)
				}
			}
		}

		// Finds all records matching given conditions
		err = query.Find(&books).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetBooksByFilter] Repository error retrieving books by filter")
			return err
		}

		// Count total books based on the filtered conditions
		err = countQuery.Count(&totalBooks).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetBooksByFilter] Repository error counting total books")
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return books, totalBooks, nil
}
