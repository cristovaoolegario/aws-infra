package route

import "testing"

func TestLoadPositions(t *testing.T) {
	t.Run("Should return error When there is no route ID", func(t *testing.T) {
		route := Route{
			ID:        "",
			ClientID:  "",
			Positions: []Position{},
		}

		err := route.LoadPositions()

		if err == nil {
			t.Errorf("Expected an error")
		}
		if err.Error() != "route id not informed" {
			t.Errorf("Expected an error: 'route id not informed'")
		}
	})

	t.Run("Should return error when there is no destination file to open", func(t *testing.T) {
		route := Route{
			ID:        "5",
			ClientID:  "1",
			Positions: []Position{},
		}

		err := route.LoadPositions()

		if err == nil {
			t.Errorf("Expected an error")
		}
		if err.Error() != "open destinations/5.txt: no such file or directory" {
			t.Errorf("Expected an error: 'open destinations/5.txt: no such file or directory', got: %v", err.Error())
		}
	})

	t.Run("Should read file and load positions properly", func(t *testing.T) {
		route := Route{
			ID:        "1",
			ClientID:  "1",
			Positions: []Position{},
		}

		err := route.LoadPositions()

		if err != nil {
			t.Errorf("An Unexpected error occurs: %v", err.Error())
		}

	})
}

func TestExportJsonPositions(t *testing.T) {
	//TODO
}
