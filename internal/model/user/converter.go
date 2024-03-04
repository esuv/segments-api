package user

import "segments-api/internal/controller/rest/dto"

func (it User) ToDTO() dto.UserDTO {
	return dto.UserDTO{}
}

func ToDTOs(models []User) (result []dto.UserDTO) {
	for _, model := range models {
		result = append(result, model.ToDTO())
	}

	return result
}
