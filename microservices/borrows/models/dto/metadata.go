package dto

type PaginationResponse struct {
	TotalPages   int `json:"totalPages,omitempty"`
	CurrentPage  int `json:"currentPage"`
	NextPage     int `json:"nextPage,omitempty"`
	PreviousPage int `json:"previousPage,omitempty"`
}
