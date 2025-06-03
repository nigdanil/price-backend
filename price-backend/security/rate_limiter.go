package security

import (
	"sync"
	"time"
)

type LoginAttemptInfo struct {
	Attempts    int
	LastAttempt time.Time
	LockedUntil time.Time
}

var (
	loginAttempts = make(map[string]*LoginAttemptInfo)
	mu            sync.Mutex
)

// CheckAndRecordLoginAttempt проверяет и записывает попытку входа
func CheckAndRecordLoginAttempt(ip string) (blocked bool, retryAfter time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	info, exists := loginAttempts[ip]

	if exists && now.Before(info.LockedUntil) {
		return true, time.Until(info.LockedUntil)
	}

	if info == nil {
		info = &LoginAttemptInfo{}
		loginAttempts[ip] = info
	}

	info.Attempts++
	info.LastAttempt = now

	switch info.Attempts {
	case 1:
		info.LockedUntil = now.Add(1 * time.Minute)
	case 2:
		info.LockedUntil = now.Add(20 * time.Minute)
	default:
		info.LockedUntil = now.Add(1 * time.Hour)
	}

	return false, 0
}

// ClearLoginAttempts удаляет запись после успешного входа
func ClearLoginAttempts(ip string) {
	mu.Lock()
	defer mu.Unlock()
	delete(loginAttempts, ip)
}
