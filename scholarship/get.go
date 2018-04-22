package scholarship

import (
	"fmt"

	"github.com/harkce/beasiswakita"
)

func GetScholarships(filter map[string]string, limit int, offset int) ([]beasiswakita.Scholarship, int, error) {
	var scholarships []beasiswakita.Scholarship

	query := fmt.Sprintf("select * from scholarships where name LIKE '%%%s%%'", filter["keywords"])
	countQuery := fmt.Sprintf("select count(id) from scholarships where name LIKE '%%%s%%'", filter["keywords"])

	if filter["start_date"] != "" {
		query = fmt.Sprintf("%s and start_date >= STR_TO_DATE('%s', '%%Y-%%m-%%d')", query, filter["start_date"])
		countQuery = fmt.Sprintf("%s and start_date >= STR_TO_DATE('%s', '%%Y-%%m-%%d')", countQuery, filter["start_date"])
	}

	if filter["end_date"] != "" {
		query = fmt.Sprintf("%s and end_date >= STR_TO_DATE('%s', '%%Y-%%m-%%d')", query, filter["end_date"])
		countQuery = fmt.Sprintf("%s and end_date >= STR_TO_DATE('%s', '%%Y-%%m-%%d')", countQuery, filter["start_date"])
	}

	if filter["organization_id"] != "" {
		query = fmt.Sprintf("%s and organization_id = %s", query, filter["organization_id"])
		countQuery = fmt.Sprintf("%s and organization_id = %s", countQuery, filter["organization_id"])
	}

	count, err := beasiswakita.DbMap.SelectInt(countQuery)
	if err != nil {
		return scholarships, 0, err
	}

	if filter["sort"] == "" {
		filter["sort"] = "desc"
	}
	query = fmt.Sprintf("%s order by created_at %s", query, filter["sort"])

	if limit != 0 || offset != 0 {
		query = fmt.Sprintf("%s limit %d offset %d", query, limit, offset)
	}

	_, err = beasiswakita.DbMap.Select(&scholarships, query)
	if err != nil {
		fmt.Println("errselect")
		return scholarships, 0, err
	}

	return scholarships, int(count), nil
}
