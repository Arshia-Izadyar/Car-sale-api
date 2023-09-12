package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

type IpRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  sync.RWMutex
	r   rate.Limit
	b   int
}

func NewIpRateLimiter(r rate.Limit, b int) *IpRateLimiter {
	i := IpRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  sync.RWMutex{},
		r:   r,
		b:   b,
	}
	return &i
}

func (i *IpRateLimiter) AddIp(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter
	return limiter
}

func (i *IpRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exits := i.ips[ip]
	if !exits {
		i.mu.Unlock()
		return i.AddIp(ip)
	}
	i.mu.Unlock()
	return limiter
}
