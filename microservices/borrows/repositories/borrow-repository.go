package repositories

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/borrows/models"
	"gorm.io/gorm"
)

func (r *Repository) CreateBorrow(createBorrow *models.Borrow) error {
	createdBorrow := r.PostgreSqlConn.Db.Create(createBorrow)
	err := createdBorrow.Error
	if err != nil {
		log.Error().Err(err).Msg("[CreateBorrow] Repository error creating borrow")
		return err
	}

	return nil
}

func (r *Repository) GetBorrowById(primaryId models.BorrowPrimaryId) (models.Borrow, error) {
	var borrow models.Borrow
	borrowData := r.PostgreSqlConn.Db.First(&borrow, primaryId)
	err := borrowData.Error
	if err != nil {
		log.Error().Err(err).Msg("[GetBorrowById] Repository error retrieving borrow by id")
		return models.Borrow{}, err
	}

	return borrow, nil
}

func (r *Repository) UpdateBorrow(primaryId models.BorrowPrimaryId, updateData *models.Borrow) error {
	var borrow models.Borrow
	updatedBorrow := r.PostgreSqlConn.Db.Model(&borrow).Where("id = ?", primaryId.Id).Updates(updateData)
	err := updatedBorrow.Error
	if err != nil {
		log.Error().Err(err).Msg("[UpdateBorrowById] Repository error updating borrow by id")
		return err
	}

	return nil
}

func (r *Repository) GetBorrowsByFilter(filter models.Filter) ([]models.Borrow, int64, error) {
	var borrows []models.Borrow
	var totalBorrows int64
	var err error

	// Start a transaction
	err = r.PostgreSqlConn.Db.Transaction(func(tx *gorm.DB) error {
		// Start building the query with the transaction context (tx)
		query := tx.Model(&models.Borrow{})
		countQuery := tx.Model(&models.Borrow{})

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
					log.Warn().Msgf("[GetBorrowsByFilter] Unsupported filter operator: %s for field: %s", filterField.Operator, filterField.Field)
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
					log.Warn().Msgf("[GetBorrowsByFilter] Unknown sort order: %s for field: %s", sort.Order, sort.Field)
				}
			}
		}

		// Finds all records matching given conditions
		err = query.Find(&borrows).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetBorrowsByFilter] Repository error retrieving borrows by filter")
			return err
		}

		// Count total borrows based on the filtered conditions
		err = countQuery.Count(&totalBorrows).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetBorrowsByFilter] Repository error counting total borrows")
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return borrows, totalBorrows, nil
}
