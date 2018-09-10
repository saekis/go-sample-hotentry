package validator

import "errors"

func Validate(ll int, c string) error {
	if ll < 0 {
		return errors.New("lowwerlimit can be allowed only integer")
	}

	for _, v := range []string{"all", "general", "social", "economics", "life", "knowlegde", "it", "fun", "entertaiment", "game"} {
		if c == v {
			return nil
		}
	}
	return errors.New("category not found")
}
