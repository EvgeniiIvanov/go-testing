package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Keep users in map
type mockMapStorage struct {
	users map[int]User
}

// Implement Save for users in map
func (m *mockMapStorage) Save(u User) error {
	if m.users == nil {
		m.users = make(map[int]User)
	}
	if _, ok := m.users[u.ID]; ok {
		return ErrAlreadyExists
	}
	m.users[u.ID] = u
	return nil
}

// Implement ByID for users in map
func (m *mockMapStorage) ByID(id int) (User, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return User{}, ErrNotFound
}

// Using mock package
type mockUserStorage struct {
	mock.Mock
}

func (m *mockUserStorage) Save(u User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *mockUserStorage) ByID(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

// Tests for Get with map
func TestGetUser(t *testing.T) {
	mock := &mockMapStorage{
		users: map[int]User{
			1: {ID: 1, Name: "Din"},
			2: {ID: 2, Name: "Sam"},
		},
	}

	service := New(mock)

	t.Run("Return Sam by id", func(t *testing.T) {
		user, err := service.Get(2)
		expected := User{ID: 2, Name: "Sam"}
		require.NoError(t, err)
		assert.Equal(t, expected, user)
	})

	t.Run("Somebody not from this show", func(t *testing.T) {
		_, err := service.Get(15)
		require.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}

// Tests for Add with map
func TestAddUser(t *testing.T) {
	t.Run("Add main another main character", func(t *testing.T) {
		mock := &mockMapStorage{
			users: map[int]User{
				1: {ID: 1, Name: "Din"},
				2: {ID: 2, Name: "Sam"},
			},
		}
		service := New(mock)

		newUser := User{ID: 3, Name: "Cas"}
		err := service.Add(newUser)
		require.NoError(t, err)
		// check if user was added
		user, _ := service.Get(3)
		assert.Equal(t, newUser, user)
	})

	t.Run("Replace Din with Squidward", func(t *testing.T) {
		mock := &mockMapStorage{
			users: map[int]User{
				1: {ID: 1, Name: "Din"},
			},
		}
		service := New(mock)

		newUser := User{ID: 1, Name: "Squidward"}
		err := service.Add(newUser)
		require.Error(t, err)
		assert.ErrorIs(t, err, ErrAlreadyExists)
	})

}

func TestByIDMethodWithMock(t *testing.T) {

	cases := []struct {
		name      string
		id        int
		mockSetup func(m *mockUserStorage)
		expected  User
		wantErr   error
	}{
		{
			name: "Just get Bobby Singer by ID",
			id:   1,
			mockSetup: func(m *mockUserStorage) {
				m.On("ByID", 1).Return(User{ID: 1, Name: "Bob"}, nil)
			},
			expected: User{ID: 1, Name: "Bob"},
			wantErr:  nil,
		},
		{
			name: "Try to get SpongeBob by ID",
			id:   999,
			mockSetup: func(m *mockUserStorage) {
				m.On("ByID", 999).Return(User{}, ErrNotFound)
			},
			expected: User{},
			wantErr:  ErrNotFound,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			m := new(mockUserStorage)
			c.mockSetup(m)

			service := New(m)
			user, err := service.Get(c.id)
			if c.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, c.wantErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, c.expected, user)
			}
			m.AssertExpectations(t)
		})
	}
}

func TestSaveMethodWithMock(t *testing.T) {
	cases := []struct {
		name      string
		id        int
		mockSetup func(m *mockUserStorage)
		user      User
		wantErr   error
	}{
		{
			name: "Save Sam Winchester",
			id:   1,
			mockSetup: func(m *mockUserStorage) {
				m.On("Save", User{ID: 1, Name: "Sam"}).Return(nil)
			},
			user:    User{ID: 1, Name: "Sam"},
			wantErr: nil,
		},
		{
			name: "Save Patrick Star instead of Din",
			id:   2,
			mockSetup: func(m *mockUserStorage) {
				m.On("Save", User{ID: 2, Name: "Patrick"}).Return(ErrAlreadyExists)
			},
			user:    User{ID: 2, Name: "Patrick"},
			wantErr: ErrAlreadyExists,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			m := new(mockUserStorage)
			c.mockSetup(m)

			service := New(m)
			err := service.Add(c.user)
			if c.wantErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, c.wantErr)
			} else {
				require.NoError(t, err)
			}
			m.AssertExpectations(t)
		})
	}

}
