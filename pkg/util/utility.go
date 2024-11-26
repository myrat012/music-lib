package util

func PageToLimitOffset(page, perPage int) (offset, limit int) {
	if page < 1 {
		page = 1
	}
	if perPage > 100 {
		perPage = 100
	}
	limit = perPage
	offset = perPage * (page - 1)
	return
}
