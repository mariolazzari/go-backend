package main

import (
	"testing"
)

func MainTest(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck", func(t *testing.T) {
			nt := &NormalTruck{id: "T1", cargo: 10}
			et := &ElectrictTruck{id: "eT2", cargo: 20}

			err := processTruck(nt)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			err = processTruck(et)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatal("Normal cargo should be 0")
			}
		})
	})
}
