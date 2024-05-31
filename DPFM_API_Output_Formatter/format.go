package dpfm_api_output_formatter

import (
	"data-platform-api-event-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Event,
			&pm.EventType,
			&pm.EventOwner,
			&pm.EventOwnerBusinessPartnerRole,
			&pm.PersonResponsible,
			&pm.ValidityStartDate,
			&pm.ValidityStartTime,
			&pm.ValidityEndDate,
			&pm.ValidityEndTime,
			&pm.OperationStartDate,
			&pm.OperationStartTime,
			&pm.OperationEndDate,
			&pm.OperationEndTime,
			&pm.Description,
			&pm.LongText,
			&pm.Introduction,
			&pm.Site,
			&pm.Project,
			&pm.WBSElement,
			&pm.Tag1,
			&pm.Tag2,
			&pm.Tag3,
			&pm.Tag4,
			&pm.DistributionProfile,
			&pm.PointConditionType,
			&pm.QuestionnaireType,
			&pm.QuestionnaireTemplate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Event:							data.Event,
			EventType:						data.EventType,
			EventOwner:						data.EventOwner,
			EventOwnerBusinessPartnerRole:	data.EventOwnerBusinessPartnerRole,
			PersonResponsible:				data.PersonResponsible,
			ValidityStartDate:				data.ValidityStartDate,
			ValidityStartTime:				data.ValidityStartTime,
			ValidityEndDate:				data.ValidityEndDate,
			ValidityEndTime:				data.ValidityEndTime,
			OperationStartDate:				data.OperationStartDate,
			OperationStartTime:				data.OperationStartTime,
			OperationEndDate:				data.OperationEndDate,
			OperationEndTime:				data.OperationEndTime,
			Description:					data.Description,
			LongText:						data.LongText,
			Introduction:					data.Introduction,
			Site:							data.Site,
			Project:						data.Project,
			WBSElement:						data.WBSElement,
			Tag1:							data.Tag1,
			Tag2:							data.Tag2,
			Tag3:							data.Tag3,
			Tag4:							data.Tag4,
			DistributionProfile:			data.DistributionProfile,
			PointConditionType:				data.PointConditionType,
			QuestionnaireType:				data.QuestionnaireType,
			QuestionnaireTemplate:			data.QuestionnaireTemplate,
			CreationDate:					data.CreationDate,
			CreationTime:					data.CreationTime,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
			CreateUser:						data.CreateUser,
			LastChangeUser:					data.LastChangeUser,
			IsReleased:						data.IsReleased,
			IsCancelled:					data.IsCancelled,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToPartner(rows *sql.Rows) (*[]Partner, error) {
	defer rows.Close()
	partner := make([]Partner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Partner{}

		err := rows.Scan(
			&pm.Event,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
			&pm.EmailAddress,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &partner, err
		}

		data := pm
		partner = append(partner, Partner{
			Event:                   data.Event,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
			EmailAddress:            data.EmailAddress,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &partner, nil
	}

	return &partner, nil
}

func ConvertToAddress(rows *sql.Rows) (*[]Address, error) {
	defer rows.Close()
	address := make([]Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.Event,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalSubRegion,
			&pm.LocalRegion,
			&pm.Country,
			&pm.GlobalRegion,
			&pm.TimeZone,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
			&pm.XCoordinate,
			&pm.YCoordinate,
			&pm.ZCoordinate,
			&pm.Site,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &address, err
		}

		data := pm
		address = append(address, Address{
			Event:       	data.Event,
			AddressID:   	data.AddressID,
			PostalCode:  	data.PostalCode,
			LocalSubRegion: data.LocalSubRegion,
			LocalRegion: 	data.LocalRegion,
			Country:     	data.Country,
			GlobalRegion: 	data.GlobalRegion,
			TimeZone:	 	data.TimeZone,
			District:    	data.District,
			StreetName:  	data.StreetName,
			CityName:    	data.CityName,
			Building:    	data.Building,
			Floor:       	data.Floor,
			Room:        	data.Room,
			XCoordinate: 	data.XCoordinate,
			YCoordinate: 	data.YCoordinate,
			ZCoordinate: 	data.ZCoordinate,
			Site:		 	data.Site,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &address, nil
	}

	return &address, nil
}

func ConvertToCampaign(rows *sql.Rows) (*[]Campaign, error) {
	defer rows.Close()
	campaign := make([]Campaign, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Campaign{}

		err := rows.Scan(
			&pm.Event,
			&pm.Campaign,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsReleased,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &campaign, err
		}

		data := pm
		campaign = append(campaign, Campaign{
			Event:					data.Event,
			Campaign:				data.Campaign,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsReleased:				data.IsReleased,
			IsCancelled:			data.IsCancelled,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &campaign, nil
	}

	return &campaign, nil
}

func ConvertToGame(rows *sql.Rows) (*[]Game, error) {
	defer rows.Close()
	game := make([]Game, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Game{}

		err := rows.Scan(
			&pm.Event,
			&pm.Game,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsReleased,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &game, err
		}

		data := pm
		game = append(game, Game{
			Event:					data.Event,
			Game:					data.Game,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsReleased:				data.IsReleased,
			IsCancelled:			data.IsCancelled,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &game, nil
	}

	return &game, nil
}

func ConvertToCounter(rows *sql.Rows) (*[]Counter, error) {
	defer rows.Close()
	counter := make([]Counter, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Counter{}

		err := rows.Scan(
			&pm.Event,
			&pm.NumberOfLikes,
			&pm.NumberOfParticipations,
			&pm.NumberOfAttendances,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &counter, err
		}

		data := pm
		counter = append(counter, Counter{
			Event:					data.Event,
			NumberOfLikes:			data.NumberOfLikes,
			NumberOfParticipations:	data.NumberOfParticipations,
			NumberOfAttendances:	data.NumberOfAttendances,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &counter, nil
	}

	return &counter, nil
}

func ConvertToParticipation(rows *sql.Rows) (*[]Participation, error) {
	defer rows.Close()
	participation := make([]Participation, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Participation{}

		err := rows.Scan(
			&pm.Event,
			&pm.Participator,
			&pm.Participation,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &participation, err
		}

		data := pm
		participation = append(participation, Participation{
			Event:								data.Event,
			Participator:						data.Participator,
			Participation:						data.Participation,
			CreationDate:						data.CreationDate,
			CreationTime:						data.CreationTime,
			LastChangeDate:						data.LastChangeDate,
			LastChangeTime:						data.LastChangeTime,
			IsCancelled:						data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &participation, nil
	}

	return &participation, nil
}

func ConvertToAttendance(rows *sql.Rows) (*[]Attendance, error) {
	defer rows.Close()
	attendance := make([]Attendance, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Attendance{}

		err := rows.Scan(
			&pm.Event,
			&pm.Attender,
			&pm.Attendance,
			&pm.Participation,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &attendance, err
		}

		data := pm
		attendance = append(attendance, Attendance{
			Event:								data.Event,
			Attender:							data.Attender,
			Attendance:							data.Attendance,
			Participation:						data.Participation,
			CreationDate:						data.CreationDate,
			CreationTime:						data.CreationTime,
			LastChangeDate:						data.LastChangeDate,
			LastChangeTime:						data.LastChangeTime,
			IsCancelled:						data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &attendance, nil
	}

	return &attendance, nil
}

func ConvertToPointTransaction(rows *sql.Rows) (*[]PointTransaction, error) {
	defer rows.Close()
	pointTransaction := make([]PointTransaction, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PointTransaction{}

		err := rows.Scan(
			&pm.Event,
			&pm.Sender,
			&pm.Receiver,
			&pm.PointConditionRecord,
			&pm.PointConditionSequentialNumber,
			&pm.PointTransaction,
			&pm.PointSymbol,
			&pm.PointTransactionType,
			&pm.PointConditionType,
			&pm.PointConditionRateValue,
			&pm.PointConditionRatio,
			&pm.PlusMinus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &pointTransaction, err
		}

		data := pm
		pointTransaction = append(pointTransaction, PointTransaction{
			Event:								data.Event,
			Sender:								data.Sender,
			Receiver:							data.Receiver,
			PointConditionRecord:				data.PointConditionRecord,
			PointConditionSequentialNumber:		data.PointConditionSequentialNumber,
			PointTransaction:					data.PointTransaction,
			PointSymbol:						data.PointSymbol,
			PointTransactionType:				data.PointTransactionType,
			PointConditionType:					data.PointConditionType,
			PointConditionRateValue:			data.PointConditionRateValue,
			PointConditionRatio:				data.PointConditionRatio,
			PlusMinus:							data.PlusMinus,
			CreationDate:						data.CreationDate,
			CreationTime:						data.CreationTime,
			LastChangeDate:						data.LastChangeDate,
			LastChangeTime:						data.LastChangeTime,
			IsCancelled:						data.IsCancelled,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &pointTransaction, nil
	}

	return &pointTransaction, nil
}

func ConvertToPointConditionElement(rows *sql.Rows) (*[]PointConditionElement, error) {
	defer rows.Close()
	pointConditionElement := make([]PointConditionElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PointConditionElement{}

		err := rows.Scan(
			&pm.Event,
			&pm.PointConditionRecord,
			&pm.PointConditionSequentialNumber,
			&pm.PointSymbol,
			&pm.Sender,
			&pm.PointTransactionType,
			&pm.PointConditionType,
			&pm.PointConditionRateValue,
			&pm.PointConditionRatio,
			&pm.PlusMinus,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsReleased,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &pointConditionElement, err
		}

		data := pm
		pointConditionElement = append(pointConditionElement, PointConditionElement{
			Event:							data.Event,
			PointConditionRecord:			data.PointConditionRecord,
			PointConditionSequentialNumber:	data.PointConditionSequentialNumber,
			PointSymbol:					data.PointSymbol,
			Sender:							data.Sender,
			PointTransactionType:			data.PointTransactionType,
			PointConditionType:				data.PointConditionType,
			PointConditionRateValue:		data.PointConditionRateValue,
			PointConditionRatio:			data.PointConditionRatio,
			PlusMinus:						data.PlusMinus,
			CreationDate:					data.CreationDate,
			LastChangeDate:					data.LastChangeDate,
			IsReleased:						data.IsReleased,
			IsCancelled:					data.IsCancelled,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &pointConditionElement, nil
	}

	return &pointConditionElement, nil
}
