package repositories

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/authors/models"
	"gorm.io/gorm"
)

func (r *Repository) CreateAuthor(createAuthor *models.Author) error {
	createdAuthor := r.PostgreSqlConn.Db.Create(createAuthor)
	err := createdAuthor.Error
	if err != nil {
		log.Error().Err(err).Msg("[CreateAuthor] Repository error creating author")
		return err
	}

	return nil
}

func (r *Repository) GetAuthorsByFilter(filter models.Filter) ([]models.Author, int64, error) {
	var authors []models.Author
	var totalAuthors int64
	var err error

	// Start a transaction
	err = r.PostgreSqlConn.Db.Transaction(func(tx *gorm.DB) error {
		// Start building the query with the transaction context (tx)
		query := tx.Model(&models.Author{})
		countQuery := tx.Model(&models.Author{})

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
					log.Warn().Msgf("[GetAuthorsByFilter] Unsupported filter operator: %s for field: %s", filterField.Operator, filterField.Field)
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
					log.Warn().Msgf("[GetAuthorsByFilter] Unknown sort order: %s for field: %s", sort.Order, sort.Field)
				}
			}
		}

		// Finds all records matching given conditions
		err = query.Find(&authors).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetAuthorsByFilter] Repository error retrieving authors by filter")
			return err
		}

		// Count total authors based on the filtered conditions
		err = countQuery.Count(&totalAuthors).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetAuthorsByFilter] Repository error counting total authors")
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return authors, totalAuthors, nil
}
