package types

import (
	"testing"
)

func TestFilterOperatorAddOrNot(t *testing.T) {
	tests := []struct {
		name     string
		operator FilterOperator
		expected bool
	}{
		{
			name:     "FilterOperatorLike should not add operator field",
			operator: FilterOperatorLike,
			expected: false,
		},
		{
			name:     "FilterOperatorFree should not add operator field",
			operator: FilterOperatorFree,
			expected: false,
		},
		{
			name:     "Empty operator should not add operator field",
			operator: FilterOperator(""),
			expected: false,
		},
		{
			name:     "FilterOperatorGreater should add operator field",
			operator: FilterOperatorGreater,
			expected: true,
		},
		{
			name:     "FilterOperatorEqual should add operator field",
			operator: FilterOperatorEqual,
			expected: true,
		},
		{
			name:     "FilterOperatorNotEqual should add operator field",
			operator: FilterOperatorNotEqual,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operator.AddOrNot()
			if result != tt.expected {
				t.Errorf("FilterOperator.AddOrNot() = %v, expected %v for operator %s", result, tt.expected, string(tt.operator))
			}
		})
	}
}

func TestFilterOperatorLabel(t *testing.T) {
	tests := []struct {
		name     string
		operator FilterOperator
		expected string
	}{
		{
			name:     "FilterOperatorLike should return empty label",
			operator: FilterOperatorLike,
			expected: "",
		},
		{
			name:     "FilterOperatorGreater should return > label",
			operator: FilterOperatorGreater,
			expected: ">",
		},
		{
			name:     "FilterOperatorEqual should return = label",
			operator: FilterOperatorEqual,
			expected: "=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := string(tt.operator.Label())
			if result != tt.expected {
				t.Errorf("FilterOperator.Label() = %v, expected %v for operator %s", result, tt.expected, string(tt.operator))
			}
		})
	}
}
