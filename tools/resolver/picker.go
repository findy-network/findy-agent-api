package resolver

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/findy-network/findy-agent-api/resolver"
	"github.com/findy-network/findy-agent-api/tools/data"
	"github.com/lainio/err2"
)

const MAX_PATCH_SIZE = 100

func parseCursor(cursor string) (int64, error) {
	plain, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, errors.New(resolver.ErrorCursorInvalid)
	}

	parts := strings.Split(string(plain), ":")
	if len(parts) != 2 {
		return 0, errors.New(resolver.ErrorCursorInvalid)
	}

	value, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, errors.New(resolver.ErrorCursorInvalid)
	}

	return value, nil
}

func validateFirstAndLast(first, last *int) error {
	if first == nil && last == nil {
		return errors.New(resolver.ErrorFirstLastMissing)
	}
	if (first != nil && (*first < 1 || *first > MAX_PATCH_SIZE)) ||
		(last != nil && (*last < 1 || *last > MAX_PATCH_SIZE)) {
		return errors.New(resolver.ErrorFirstLastInvalid)
	}
	return nil
}

func pick(
	items data.Items,
	after *string, before *string,
	first *int, last *int) (afterIndex int, beforeIndex int, err error) {
	defer err2.Return(&err)

	afterIndex = 0
	beforeIndex = items.Count() - 1

	err2.Check(validateFirstAndLast(first, last))

	if after != nil || before != nil {
		var afterVal, beforeVal int64
		if after != nil {
			afterVal, err = parseCursor(*after)
			err2.Check(err)
		}
		if before != nil {
			beforeVal, err = parseCursor(*before)
			err2.Check(err)
		}
		for index := 0; index < items.Count(); index++ {
			created := items.CreatedForIndex(index)
			if afterVal > 0 && created <= afterVal {
				afterIndex = index + 1
			}
			if beforeVal > 0 && created < beforeVal {
				beforeIndex = index
			}
			if (beforeVal > 0 && created > beforeVal) ||
				(beforeVal == 0 && created > afterVal) {
				break
			}
		}
	}

	if first != nil {
		afterPlusFirst := afterIndex + (*first - 1)
		if beforeIndex > afterPlusFirst {
			beforeIndex = afterPlusFirst
		}
	} else if last != nil {
		beforeMinusLast := beforeIndex - (*last - 1)
		if afterIndex < beforeMinusLast {
			afterIndex = beforeMinusLast
		}
	}
	return afterIndex, beforeIndex + 1, nil
}
