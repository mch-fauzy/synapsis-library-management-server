package pagination

import (
	"math"

	"github.com/synapsis-library-management-server/microservices/authors/models/dto"
)

// CalculatePaginationMetadata calculates pagination metadata
func CalculatePaginationMetadata(page, pageSize, totalItems int64) dto.PaginationResponse {
	totalPages := int64(0)
	if pageSize > 0 {
		totalPages = int64(math.Ceil(float64(totalItems) / float64(pageSize)))
	}

	nextPage := page + 1
	previousPage := page - 1
	if nextPage > totalPages {
		nextPage = 0
	}
	if previousPage < 1 {
		previousPage = 0
	}

	return dto.PaginationResponse{
		TotalPages:   int(totalPages),
		CurrentPage:  int(page),
		NextPage:     int(nextPage),
		PreviousPage: int(previousPage),
	}
}
