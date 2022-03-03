package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/go-kumparan/internal/domain/dto"
	"math"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (h *Handler) initBackendRoutes(api *gin.RouterGroup) {
	backend := api.Group("/backend")
	{
		backend.POST("/binary-decimal", h.binaryDecimal)
		backend.POST("/decimal-binary", h.decimalBinary)
		backend.POST("/palindrome", h.palindrome)
	}
}

func (h *Handler) binaryDecimal(context *gin.Context) {
	var binaryToDecimalDTO dto.BinaryToDecimalDTO

	err := context.BindJSON(&binaryToDecimalDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	s := strings.Split(binaryToDecimalDTO.Value, "")

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	result := 0
	for index, item := range s {

		temp, _ := strconv.Atoi(item)

		if temp == 1 {

			if index == 0 {
				result += 1
			} else {
				result += int(math.Pow(2, float64(index)))
			}

		}

	}

	var data interface{}
	data = result
	successResponse(context, data)
}

func (h *Handler) decimalBinary(context *gin.Context) {
	var decimalToBinaryDTO dto.DecimalToBinaryDTO

	err := context.BindJSON(&decimalToBinaryDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	var binary [6]int

	i := 0
	value := decimalToBinaryDTO.Value

	for value > 0 {
		binary[i] = value % 2
		value = int(value / 2)
		i++
	}

	result := ""
	j := i - 1
	for j >= 0 {
		result += strconv.Itoa(binary[j])
		j--
	}

	var data interface{}
	data = result
	successResponse(context, data)
}

func (h *Handler) palindrome(context *gin.Context) {
	var palindromeDTO dto.PalindromeDTO

	err := context.BindJSON(&palindromeDTO)
	if err != nil {
		errorResponse(context, http.StatusBadRequest, "invalid input body")
		return
	}

	s := palindromeDTO.Value
	stringLength := utf8.RuneCountInString(s)
	isPalindromeMatrix := make([][]int, stringLength)
	for i := range isPalindromeMatrix {
		isPalindromeMatrix[i] = make([]int, stringLength)
	}

	for i, outer := range isPalindromeMatrix {
		for j := range outer {
			if i == j {
				isPalindromeMatrix[i][j] = 1
			}

		}
	}

	palindromeLengthMatrix := make([][]int, stringLength)
	for i := range palindromeLengthMatrix {
		palindromeLengthMatrix[i] = make([]int, stringLength)
	}

	for i, outer := range palindromeLengthMatrix {
		for j := range outer {
			if i == j {
				palindromeLengthMatrix[i][j] = 1
			}

		}
	}

	for len := 2; len <= stringLength; len++ {
		for i := 0; i <= stringLength-len; i++ {
			j := i + len - 1

			if s[i] == s[j] {
				if len == 2 {
					isPalindromeMatrix[i][j] = 1
					palindromeLengthMatrix[i][j] = 2
				} else {
					if isPalindromeMatrix[i+1][j-1] == 1 {
						isPalindromeMatrix[i][j] = 1
						palindromeLengthMatrix[i][j] = 2 + palindromeLengthMatrix[i+1][j-1]
					} else {
						isPalindromeMatrix[i][j] = -1
						palindromeLengthMatrix[i][j] = int(math.Max(float64(palindromeLengthMatrix[i+1][j]), float64(palindromeLengthMatrix[i][j-1])))
					}

				}

			} else {
				isPalindromeMatrix[i][j] = -1
			}

		}
	}

	max_row_index := 0
	max_column_index := 0
	max := 0

	for i, outer := range palindromeLengthMatrix {
		for j := range outer {
			if palindromeLengthMatrix[i][j] > max && isPalindromeMatrix[i][j] == 1 {
				max = palindromeLengthMatrix[i][j]
				max_row_index = i
				max_column_index = j
			}
		}
	}

	var data interface{}
	data = s[max_row_index : max_column_index+1]
	successResponse(context, data)
}
