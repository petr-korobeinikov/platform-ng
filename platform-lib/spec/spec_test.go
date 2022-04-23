package spec_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/pkorobeinikov/platform-ng/platform-lib/spec"
)

func TestLoad(t *testing.T) {
	t.Run(`positive`, func(t *testing.T) {
		tests := []struct {
			Name     string
			File     string
			Expected *Spec
		}{
			{
				Name: "minimal",
				File: "testdata/" + File,
				Expected: &Spec{
					Name:      "wordcounter",
					Namespace: "wordcounting",
				},
			},
			{
				Name: "with environment",
				File: "testdata/platform-env.yaml",
				Expected: &Spec{
					Name:      "wordcounter",
					Namespace: "wordcounting",
					Environment: map[string]map[string]string{
						"_": {
							"FOO":          "foo",
							"NAP_DURATION": "1s",
						},
						"local": {
							"FOO": "foo local override",
						},
						"staging": {
							"FOO": "foo staging override",
						},
					},
				},
			},
			{
				Name: "with environment and component",
				File: "testdata/platform-env-component.yaml",
				Expected: &Spec{
					Name:      "wordcounter",
					Namespace: "wordcounting",
					Environment: map[string]map[string]string{
						"_": {
							"FOO":          "foo",
							"NAP_DURATION": "1s",
						},
						"local": {
							"FOO": "foo local override",
						},
						"staging": {
							"FOO": "foo staging override",
						},
					},
					Component: []*Component{
						{
							Name: "master",
							Type: "postgres",
						},
						{
							Name: "olap",
							Type: "postgres",
						},
						{
							Name: "key-value-storage",
							Type: "redis",
						},
						{
							Name: "document-storage",
							Type: "mongodb",
						},
					},
				},
			},
			{
				Name: "with environment, component and task",
				File: "testdata/platform-env-component-task.yaml",
				Expected: &Spec{
					Name:      "wordcounter",
					Namespace: "wordcounting",
					Environment: map[string]map[string]string{
						"_": {
							"FOO":          "foo",
							"NAP_DURATION": "1s",
						},
						"local": {
							"FOO": "foo local override",
						},
						"staging": {
							"FOO": "foo staging override",
						},
					},
					Component: []*Component{
						{
							Name: "master",
							Type: "postgres",
						},
						{
							Name: "olap",
							Type: "postgres",
						},
						{
							Name: "key-value-storage",
							Type: "redis",
						},
						{
							Name: "document-storage",
							Type: "mongodb",
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Name, func(t *testing.T) {
				actual, err := Load(tt.File)

				assert.NoError(t, err)
				assert.Equal(t, tt.Expected, actual)
			})
		}
	})
}
