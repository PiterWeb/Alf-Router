package alf

import (
	"testing"
)

func TestCreateRouter(t *testing.T) {

	routes := []Route{
		{
			Path: "/index",
			Handle: func(ctx *Ctx) error {
				return nil
			},
			Method: "get",
			Children: []Route{
				{
					Path: "/nested",
					Handle: func(ctx *Ctx) error {
						return nil
					},
					Method: "GET",
					Children: []Route{
						{
							Path: "/morenested",
							Handle: func(ctx *Ctx) error {
								return nil
							},
							Method: "get",
						},
					},
				},
				{
					Path: "/othernested",
					Handle: func(ctx *Ctx) error {
						return nil
					},
					Method: "get",
				},
			},
		},
		{
			Path: "/get",
			Handle: func(ctx *Ctx) error {
				_, err := ctx.WriteString("GET")
				return err
			},
			Method: "get",
		}, {
			Path: "/post",
			Handle: func(ctx *Ctx) error {
				return nil
			},
			Method: "post",
		},
	}

	for i := 0; i < 50; i++ {
		methodRoutes := CreateRouter(routes)

		for m, v := range methodRoutes {

			t.Log(m, v)

		}

	}

}
