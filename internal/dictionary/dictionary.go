package dictionary

import "errors"

func GetDefinition(word string) (string, error) {
	definitions := map[string]string{
		"apple":  "木々に成る果物で、通常赤色か緑色",
		"banana": "房で成る曲がった長い果物",
		"orange": "堅い鮮やかな黄赤色の皮を持った柑橘類",
		"grape":  "小さな丸い果実で、多くの場合赤や緑や紫色",
		"peach":  "甘くて柔らかな果肉を持つ果物",
		"kiwi":   "小さな茶色の果物で、中に緑色の果肉があります",
	}

	definition, ok := definitions[word]
	if !ok {
		return "", errors.New("Definition not found")
	}

	return definition, nil
}
