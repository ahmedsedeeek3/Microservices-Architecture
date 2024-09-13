package loadbalancer

import "math/rand"

func LoadBalance(instances []string) string {
	if len(instances) == 0 {
		return ""
	}
	return instances[rand.Intn(len(instances))]
}