package unit

import (
	"sync"
	"time"
)

// VisionLink описывает одну запись "кому предоставляется обзор"
type VisionLink struct {
	TargetID int       // ID юнита, получающего обзор
	Expires  time.Time // Время, когда ссылка протухнет
}

// VisionLinkManager потокобезопасный менеджер ссылок обзора
type VisionLinkManager struct {
	mu    sync.RWMutex
	links []VisionLink
}

// NewVisionLinkManager создает новый менеджер
func NewVisionLinkManager() *VisionLinkManager {
	return &VisionLinkManager{
		links: make([]VisionLink, 0, 8),
	}
}

func (u *Unit) VisionLinkHasPlayer(targetID int) bool {
	u.mx.Lock()
	defer u.mx.Unlock()

	if u.visionLinkManager == nil {
		u.visionLinkManager = NewVisionLinkManager()
	}

	return u.visionLinkManager.Has(targetID)
}

func (u *Unit) VisionLinkAddPlayer(targetID int, ttl time.Duration) {
	u.mx.Lock()
	defer u.mx.Unlock()

	if u.visionLinkManager == nil {
		u.visionLinkManager = NewVisionLinkManager()
	}

	u.visionLinkManager.Add(targetID, ttl)
}

// Has проверяет, есть ли активная (не протухшая) ссылка на targetID.
// Попутно удаляет протухшие записи для этого ID.
func (m *VisionLinkManager) Has(targetID int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	for i := 0; i < len(m.links); {
		link := m.links[i]

		if link.TargetID == targetID {
			if link.Expires.After(now) {
				// Нашли валидную запись
				return true
			}
			// Запись протухла — удаляем её и продолжаем поиск (вдруг есть ещё)
			m.links = append(m.links[:i], m.links[i+1:]...)
			continue
		}

		// Попутно чистим любые протухшие записи
		if !link.Expires.After(now) {
			m.links = append(m.links[:i], m.links[i+1:]...)
			continue
		}

		i++
	}

	return false
}

// Add добавляет или обновляет ссылку на targetID с указанным TTL (временем жизни).
// Если ссылка уже существует — её таймаут обновляется.
func (m *VisionLinkManager) Add(targetID int, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	expires := time.Now().Add(ttl)

	// Ищем существующую запись для обновления
	for i := range m.links {
		if m.links[i].TargetID == targetID {
			m.links[i].Expires = expires
			return
		}
	}

	// Не нашли — добавляем новую
	m.links = append(m.links, VisionLink{
		TargetID: targetID,
		Expires:  expires,
	})
}

// Remove удаляет ссылку на targetID (если она есть)
func (m *VisionLinkManager) Remove(targetID int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i := 0; i < len(m.links); {
		if m.links[i].TargetID == targetID {
			m.links = append(m.links[:i], m.links[i+1:]...)
			// Не return'им — вдруг есть дубликаты (маловероятно, но безопасно)
			continue
		}
		i++
	}
}

// GetActiveIDs возвращает слайс всех активных (не протухших) ID.
func (m *VisionLinkManager) GetActiveIDs() []int {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	result := make([]int, 0, len(m.links))
	alive := make([]VisionLink, 0, len(m.links))

	for _, link := range m.links {
		if link.Expires.After(now) {
			result = append(result, link.TargetID)
			alive = append(alive, link)
		}
	}

	// Сжимаем слайс, удаляя все протухшие за одним проход
	m.links = alive
	return result
}

// Cleanup удаляет все протухшие записи.
func (m *VisionLinkManager) Cleanup() {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	alive := make([]VisionLink, 0, len(m.links))
	for _, link := range m.links {
		if link.Expires.After(now) {
			alive = append(alive, link)
		}
	}
	m.links = alive
}

// Count возвращает количество записей (включая потенциально протухшие)
func (m *VisionLinkManager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.links)
}
