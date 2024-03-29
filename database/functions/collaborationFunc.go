package functions

import (
	"konsultanku-app/database"
	"konsultanku-app/models"

	"github.com/google/uuid"
)

func SendOffer(idMse string, idTeam uuid.UUID) (models.Collaboration, error) {
	collaborationID := uuid.New()
	collaboration := models.Collaboration{
		ID:     collaborationID,
		MseID:  idMse,
		TeamID: idTeam,
	}
	if err := database.DB.Create(&collaboration).Error; err != nil {
		return collaboration, err
	}
	return collaboration, nil
}

func AnyOffer(idTeam string) ([]models.Collaboration, error) {

	var collaborations []models.Collaboration
	if err := database.DB.Where("team_id = ?", idTeam).Find(&collaborations).Error; err != nil {
		return collaborations, err
	}
	return collaborations, nil
}

func AcceptOffer(idCollaboration string, teamDecision bool) (models.Collaboration, error) {

	var collaboration models.Collaboration
	if err := database.DB.Where("id = ?", idCollaboration).Find(&collaboration).Error; err != nil {
		return collaboration, err
	}

	collaboration.IsCollaboration = teamDecision
	if err := database.DB.Save(&collaboration).Error; err != nil {
		return collaboration, err
	}

	return collaboration, nil
}
