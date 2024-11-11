package repositories

import (
	"github.com/rs/zerolog/log"
	"github.com/synapsis-library-management-server/microservices/categories/models"
	"gorm.io/gorm"
)

func (r *Repository) CreateCategory(createCategory *models.Category) error {
	createdCategory := r.PostgreSqlConn.Db.Create(createCategory)
	err := createdCategory.Error
	if err != nil {
		log.Error().Err(err).Msg("[CreateCategory] Repository error creating category")
		return err
	}

	return nil
}

func (r *Repository) GetCategoryById(primaryId models.CategoryPrimaryId) (models.Category, error) {
	var category models.Category
	categoryData := r.PostgreSqlConn.Db.First(&category, primaryId)
	err := categoryData.Error
	if err != nil {
		log.Error().Err(err).Msg("[GetCategoryById] Repository error retrieving category by id")
		return models.Category{}, err
	}

	return category, nil
}

func (r *Repository) GetCategoriesByFilter(filter models.Filter) ([]models.Category, int64, error) {
	var categories []models.Category
	var totalCategories int64
	var err error

	// Start a transaction
	err = r.PostgreSqlConn.Db.Transaction(func(tx *gorm.DB) error {
		// Start building the query with the transaction context (tx)
		query := tx.Model(&models.Category{})
		countQuery := tx.Model(&models.Category{})

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
					log.Warn().Msgf("[GetCategoriesByFilter] Unsupported filter operator: %s for field: %s", filterField.Operator, filterField.Field)
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
					log.Warn().Msgf("[GetCategoriesByFilter] Unknown sort order: %s for field: %s", sort.Order, sort.Field)
				}
			}
		}

		// Finds all records matching given conditions
		err = query.Find(&categories).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetCategoriesByFilter] Repository error retrieving categories by filter")
			return err
		}

		// Count total categories based on the filtered conditions
		err = countQuery.Count(&totalCategories).Error
		if err != nil {
			log.Error().Err(err).Msg("[GetCategoriesByFilter] Repository error counting total categories")
			return err
		}

		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	return categories, totalCategories, nil
}
