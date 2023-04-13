package arrays

func AppendArray[T any](target []T, src []T) []T {
	if target != nil && src != nil && len(target) > 0 && len(src) > 0 {
		for _, tem := range src {
			target = append(target, tem)
		}
	}
	return target
}
