package constant

import "errors"

var (
	ErrDatabase                      	= errors.New("database error")
	ErrArticleIDNotNumber               = errors.New("param id is not a number")
	ErrArticleNotFound                  = errors.New("article record not found")
	ErrArticleAddFormatIncorrect        = errors.New("collects id null")
	ErrArticleAddCollectsDuplicate      = errors.New("collects id duplicate")
	ErrArticleAddCollectsRecordNotFound = errors.New("collects record not found")
	ErrArticleDelIDIncorrect            = errors.New("crticle id negative or not found")
	ErrArticleDelDeleted                = errors.New("crticle has been deleted (or database error)")
)