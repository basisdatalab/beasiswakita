package scholarship

import (
	"time"

	"github.com/basisdatalab/beasiswakita"
)

func CreateScholarship(s beasiswakita.Scholarship) (beasiswakita.Scholarship, error) {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	err := beasiswakita.Transaction.Insert(&s)
	if err != nil {
		return s, err
	}

	return s, nil
}
