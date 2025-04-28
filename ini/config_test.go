package ini

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "初始化成功",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConfig("D:\\golang_project\\ranbao\\config\\config.yaml")
			fmt.Println(Conf)
		})
	}
}
