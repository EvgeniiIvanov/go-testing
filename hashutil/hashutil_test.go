package hashutil

import "testing"

func TestHashSHA256(t *testing.T) {
	cases := []struct {
		name, in, want string
	}{
		{"simple ascii symbols", "yes", "8a798890fe93817163b10b5f7bd2ca4d25d84c52739a645a889c173eee7d9d3d"},
		{"more simple ascii symbols", "Longer than before", "ced248d507f38c549c94688a64e7dad731202d9ff344d16e44bf3d44a70c0a58"},
		{"simple unicode symbols", "Реликт", "b67340a83c963a2eb2077e994285ea7e884de3bc345dff8189e6050708e82298"},
		{"More simple unicode symbols", "ἀπὸ μηχανῆς θεός", "8b75deb7a72be09487feea808c24f845a1ad3f34f01741f1fde40d844f8e1499"},
		{"empty string", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}

	for _, c := range cases {
		// copy "c" (from ranges cases) into new var "c" in this scope
		// to protect from race condition
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			got := HashSHA256(c.in)
			if got != c.want {
				t.Errorf("SHA256(%s) = %q, but we got %q", c.in, got, c.want)
			}
		})
	}
}
