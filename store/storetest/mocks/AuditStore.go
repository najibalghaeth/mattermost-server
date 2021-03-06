// Code generated by mockery v1.0.0

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// AuditStore is an autogenerated mock type for the AuditStore type
type AuditStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: user_id, offset, limit
func (_m *AuditStore) Get(user_id string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(user_id, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(user_id, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteBatch provides a mock function with given fields: endTime, limit
func (_m *AuditStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	ret := _m.Called(endTime, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int64, int64) store.StoreChannel); ok {
		r0 = rf(endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteByUser provides a mock function with given fields: userId
func (_m *AuditStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	ret := _m.Called(userId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: audit
func (_m *AuditStore) Save(audit *model.Audit) store.StoreChannel {
	ret := _m.Called(audit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Audit) store.StoreChannel); ok {
		r0 = rf(audit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
