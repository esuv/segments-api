package segment

import "segments-api/internal/controller/rest/dto"

func (it Segment) ToDTO() dto.SegmentDTO {
	return dto.SegmentDTO{
		ID:   it.ID,
		Name: it.Name,
	}
}

func ToModel(dto dto.SegmentDTO) Segment {
	return Segment{
		ID:   dto.ID,
		Name: dto.Name,
	}
}

func ToDTOs(models []Segment) (result []dto.SegmentDTO) {
	for _, model := range models {
		result = append(result, model.ToDTO())
	}

	return result
}
