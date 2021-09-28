package sample

func SliceSample(d1, d2 [][]string) ([][]string, error) {
	r := [][]string{}
	for _, v := range d1 {
		id := val(v, 0)
		for _, v2 := range d2 {
			if val(v2, 0) == id {
				if len(v2) > 1 {
					r = append(r, append(v, v2[1:]...))
				}
				break
			}
		}
	}

	return r, nil
}

func MapSample(d1, d2 [][]string) ([][]string, error) {
	m := map[string][]string{}
	for _, v := range d2 {
		if len(v) > 1 {
			m[val(v, 0)] = v[1:]
		}
	}

	r := [][]string{}
	for _, v := range d1 {
		id := val(v, 0)
		r = append(r, append(v, m[id]...))
	}

	return r, nil
}

func val(l []string, i int) string {
	if len(l) > i {
		return l[i]
	}
	return ""
}
