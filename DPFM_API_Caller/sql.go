package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-event-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-event-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var address *[]dpfm_api_output_formatter.Address
	var partner *[]dpfm_api_output_formatter.Partner
	var campaign *[]dpfm_api_output_formatter.Campaign
	var game *[]dpfm_api_output_formatter.Game
	var counter *[]dpfm_api_output_formatter.Counter
	var participation *[]dpfm_api_output_formatter.Participation
	var attendance *[]dpfm_api_output_formatter.Attendance
	var pointTransaction *[]dpfm_api_output_formatter.PointTransaction
	var pointConditionElement *[]dpfm_api_output_formatter.PointConditionElement

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "Headers":
			func() {
				header = c.Headers(mtx, input, output, errs, log)
			}()
		case "HeadersByEvents":
			func() {
				header = c.HeadersByEvents(mtx, input, output, errs, log)
			}()
		case "HeadersBySite":
			func() {
				header = c.HeadersBySite(mtx, input, output, errs, log)
			}()
		case "Partner":
			func() {
				partner = c.Partner(mtx, input, output, errs, log)
			}()
		case "Partners":
			func() {
				partner = c.Partners(mtx, input, output, errs, log)
			}()
		case "Address":
			func() {
				address = c.Address(mtx, input, output, errs, log)
			}()
		case "Addresses":
			func() {
				address = c.Addresses(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalSubRegion":
			func() {
				address = c.AddressesByLocalSubRegion(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalSubRegions":
			func() {
				address = c.AddressesByLocalSubRegions(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalRegion":
			func() {
				address = c.AddressesByLocalRegion(mtx, input, output, errs, log)
			}()
		case "AddressesByLocalRegions":
			func() {
				address = c.AddressesByLocalRegions(mtx, input, output, errs, log)
			}()
		case "Campaign":
			func() {
				campaign = c.Campaign(mtx, input, output, errs, log)
			}()
		case "Campaigns":
			func() {
				campaign = c.Campaigns(mtx, input, output, errs, log)
			}()
		case "Game":
			func() {
				game = c.Game(mtx, input, output, errs, log)
			}()
		case "Games":
			func() {
				game = c.Games(mtx, input, output, errs, log)
			}()
		case "Counter":
			func() {
				counter = c.Counter(mtx, input, output, errs, log)
			}()
		case "CountersByEvents":
			func() {
				counter = c.CountersByEvents(mtx, input, output, errs, log)
			}()
		case "Participation":
			func() {
				participation = c.Participation(mtx, input, output, errs, log)
			}()
		case "Participations":
			func() {
				participation = c.Participations(mtx, input, output, errs, log)
			}()
		case "Attendance":
			func() {
				attendance = c.Attendance(mtx, input, output, errs, log)
			}()
		case "Attendances":
			func() {
				attendance = c.Attendances(mtx, input, output, errs, log)
			}()
		case "PointTransaction":
			func() {
				pointTransaction = c.PointTransaction(mtx, input, output, errs, log)
			}()
		case "PointTransactions":
			func() {
				pointTransaction = c.PointTransactions(mtx, input, output, errs, log)
			}()
		case "PointConditionElement":
			func() {
				pointConditionElement = c.PointConditionElement(mtx, input, output, errs, log)
			}()
		case "PointConditionElements":
			func() {
				pointConditionElement = c.PointConditionElements(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                header,
		Partner:               partner,
		Address:               address,
		Campaign:              campaign,
		Game:                  game,
		Counter:			   counter,
		Participation:		   participation,
		Attendance:			   attendance,
		PointTransaction:      pointTransaction,
		PointConditionElement: pointConditionElement,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.Event = %d", input.Header.Event)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.IsReleased ASC, header.Event ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Headers(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.IsReleased ASC, header.Event ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByEvents(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	log.Info("HeadersByEvents")
	in := ""

	for iHeader, vHeader := range input.Headers {
		event := vHeader.Event
		if iHeader == 0 {
			in = fmt.Sprintf(
				"( '%d' )",
				event,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%d' )",
			in,
			event,
		)
	}

	where := fmt.Sprintf(" WHERE ( Event ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.IsReleased ASC, header.Event ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySite(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Site = %d", input.Header.Site)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.IsReleased ASC, header.Event ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partner(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	event := input.Header.Event
	partner := input.Header.Partner

	cnt := 0
	for _, v := range partner {
		args = append(args, event, v.PartnerFunction, v.BusinessPartner)
		cnt++
	}
	repeat := strings.Repeat("(?,?,?),", cnt-1) + "(?,?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_partner_data
		WHERE (Event, PartnerFunction, BusinessPartner) IN ( `+repeat+` ) 
		ORDER BY Event ASC, PartnerFunction ASC, BusinessPartner ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Partners(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	var args []interface{}
	event := input.Header.Event
	partner := input.Header.Partner

	cnt := 0
	for _, _ = range partner {
		args = append(args, event)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_partner_data
		WHERE (Event) IN ( `+repeat+` ) 
		ORDER BY Event ASC, PartnerFunction ASC, BusinessPartner ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartner(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Address(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	address := input.Header.Address

	cnt := 0
	for _, v := range address {
		args = append(args, v.Event, v.Event)
		cnt++
	}
	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_address_data
		WHERE (Event, AddressID) IN ( `+repeat+` ) 
		ORDER BY Event ASC, AddressID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Addresses(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	var args []interface{}
	event := input.Header.Event
	address := input.Header.Address

	cnt := 0
	for _, _ = range address {
		args = append(args, event)
		cnt++
	}
	repeat := strings.Repeat("(?),", cnt-1) + "(?)"

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_address_data
		WHERE (Event) IN ( `+repeat+` ) 
		ORDER BY Event ASC, AddressID ASC;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalSubRegion(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	where := "WHERE 1 = 1"

	if input.Header.Address[0].LocalSubRegion != nil {
		where = fmt.Sprintf("WHERE address.LocalSubRegion = \"%s\"", *input.Header.Address[0].LocalSubRegion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_address_data AS address
		` + where + ` ORDER BY address.LocalSubRegion ASC, address.Event ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalSubRegions(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {

	log.Info("AddressesByLocalSubRegions")
	in := ""

	for iAddress, vAddress := range input.Header.Address {
		localSubRegion := vAddress.LocalSubRegion
		if iAddress == 0 {
			in = fmt.Sprintf(
				"( '%s' )",
				localSubRegion,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%s' )",
			in,
			localSubRegion,
		)
	}

	where := fmt.Sprintf(" WHERE ( LocalSubRegion ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_address_data AS address
		` + where + ` ORDER BY address.LocalSubRegion ASC, address.Event ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalRegion(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {
	where := "WHERE 1 = 1"

	if input.Header.Address[0].LocalRegion != nil {
		where = fmt.Sprintf("WHERE address.LocalRegion = \"%s\"", *input.Header.Address[0].LocalRegion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_address_data AS address
		` + where + ` ORDER BY address.LocalRegion ASC, address.Event ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) AddressesByLocalRegions(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Address {

	log.Info("AddressesByLocalRegions")
	in := ""

	for iAddress, vAddress := range input.Header.Address {
		localRegion := vAddress.LocalRegion
		if iAddress == 0 {
			in = fmt.Sprintf(
				"( '%s' )",
				localRegion,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%s' )",
			in,
			localRegion,
		)
	}

	where := fmt.Sprintf(" WHERE ( LocalRegion ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_address_data AS address
		` + where + ` ORDER BY address.LocalRegion ASC, address.Event ASC, address.AddressID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAddress(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Campaign(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Campaign {
	var args []interface{}
	where := fmt.Sprintf("WHERE Event = %d ", input.Header.Event)

	campaignIDs := ""
	for _, v := range input.Header.Campaign {
		campaignIDs = fmt.Sprintf("%s, %d", campaignIDs, v.Campaign)
	}

	if len(campaignIDs) != 0 {
		where = fmt.Sprintf("%s\nAND Campaign IN ( %s ) ", where, campaignIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_campaign_data
		`+where+` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, IsReleased ASC, Event ASC, Campaign ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCampaign(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Campaigns(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Campaign {
	campaign := &dpfm_api_input_reader.Campaign{}
	if len(input.Header.Campaign) > 0 {
		campaign = &input.Header.Campaign[0]
	}
	where := "WHERE 1 = 1"

	if input.Header.Event != 0 {
		where = fmt.Sprintf("WHERE Event = %d", input.Header.Event)
	}

	if campaign != nil {
		if campaign.IsReleased != nil {
			where = fmt.Sprintf("%s\nAND campaign.IsReleased = %v", where, *campaign.IsReleased)
		}
		if campaign.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND campaign.IsCancelled = %v", where, *campaign.IsCancelled)
		}
		if campaign.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND campaign.IsMarkedForDeletion = %v", where, *campaign.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_campaign_data as campaign
		` + where + ` ORDER BY campaign.IsMarkedForDeletion ASC, campaign.IsCancelled ASC, campaign.IsReleased ASC, campaign.Event ASC, campaign.Campaign ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCampaign(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Game(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Game {
	var args []interface{}
	where := fmt.Sprintf("WHERE Event = %d ", input.Header.Event)

	gameIDs := ""
	for _, v := range input.Header.Game {
		gameIDs = fmt.Sprintf("%s, %d", gameIDs, v.Game)
	}

	if len(gameIDs) != 0 {
		where = fmt.Sprintf("%s\nAND Game IN ( %s ) ", where, gameIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_game_data
		`+where+` ORDER BY IsMarkedForDeletion ASC, IsCancelled ASC, IsReleased ASC, Event ASC, Game ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGame(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Games(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Game {
	game := &dpfm_api_input_reader.Game{}
	if len(input.Header.Game) > 0 {
		game = &input.Header.Game[0]
	}
	where := "WHERE 1 = 1"

	if input.Header.Event != 0 {
		where = fmt.Sprintf("WHERE Event = %d", input.Header.Event)
	}

	if game != nil {
		if game.IsReleased != nil {
			where = fmt.Sprintf("%s\nAND game.IsReleased = %v", where, *game.IsReleased)
		}
		if game.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND game.IsCancelled = %v", where, *game.IsCancelled)
		}
		if game.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND game.IsMarkedForDeletion = %v", where, *game.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_game_data as game
		` + where + ` ORDER BY game.IsMarkedForDeletion ASC, game.IsCancelled ASC, game.IsReleased ASC, game.Event ASC, game.Game ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGame(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Counter(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Counter {
	var args []interface{}
	where := fmt.Sprintf("WHERE Event = %d ", input.Header.Event)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_counter_data AS counter
		` + where + ` ORDER BY counter.IsMarkedForDeletion ASC, counter.IsCancelled ASC, counter.IsReleased ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCounter(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) CountersByEvents(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Counter {
	log.Info("CountersByEvents")
	in := ""

	for iCounter, vCounter := range input.Headers {
		event := vCounter.Event
		if iCounter == 0 {
			in = fmt.Sprintf(
				"( '%d' )",
				event,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%d' )",
			in,
			event,
		)
	}

	where := fmt.Sprintf(" WHERE ( Event ) IN ( %s ) ", in)

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_counter_data AS counter
		` + where + ` ORDER BY counter.IsMarkedForDeletion ASC, counter.IsCancelled ASC, counter.IsReleased ASC, counter.Event ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCounter(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Participation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Participation {
	var args []interface{}
	where := fmt.Sprintf("WHERE Event = %d", input.Header.Event)

	participationIDs := ""
	for _, v := range input.Header.Participation {
		participationIDs = fmt.Sprintf("%s, %d", participationIDs, v.Participator)
	}

	if len(participationIDs) != 0 {
		where = fmt.Sprintf("%s\nAND Participator IN ( %s ) ", where, participationIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_participation_data
		`+where+` ORDER BY IsCancelled ASC, Event ASC, Participator ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToParticipation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Participations(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Participation {
	participation := &dpfm_api_input_reader.Participation{}
	if len(input.Header.Participation) > 0 {
		participation = &input.Header.Participation[0]
	}
	where := "WHERE 1 = 1"

	if input.Header.Event != 0 {
		where = fmt.Sprintf("WHERE Event = %d", input.Header.Event)
	}

	if participation != nil {
		if participation.IsReleased != nil {
			where = fmt.Sprintf("%s\nAND participation.IsReleased = %v", where, *participation.IsReleased)
		}
		if participation.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND participation.IsCancelled = %v", where, *participation.IsCancelled)
		}
		if participation.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND participation.IsMarkedForDeletion = %v", where, *participation.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_participation_data as participation
		` + where + ` ORDER BY participation.IsMarkedForDeletion ASC, participation.IsCancelled ASC, participation.IsReleased ASC, participation.Event ASC, participation.Participation ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToParticipation(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Attendance(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Attendance {
	var args []interface{}
	where := fmt.Sprintf("WHERE Event = %d", input.Header.Event)

	attendanceIDs := ""
	for _, v := range input.Header.Attendance {
		attendanceIDs = fmt.Sprintf("%s, %d", attendanceIDs, v.Attender)
	}

	if len(attendanceIDs) != 0 {
		where = fmt.Sprintf("%s\nAND Attender IN ( %s ) ", where, attendanceIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_attendance_data
		`+where+` ORDER BY IsCancelled ASC, Event ASC, Attender ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAttendance(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Attendances(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Attendance {
	attendance := &dpfm_api_input_reader.Attendance{}
	if len(input.Header.Attendance) > 0 {
		attendance = &input.Header.Attendance[0]
	}
	where := "WHERE 1 = 1"

	if input.Header.Event != 0 {
		where = fmt.Sprintf("WHERE Event = %d", input.Header.Event)
	}

	if attendance != nil {
		if attendance.IsReleased != nil {
			where = fmt.Sprintf("%s\nAND attendance.IsReleased = %v", where, *attendance.IsReleased)
		}
		if attendance.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND attendance.IsCancelled = %v", where, *attendance.IsCancelled)
		}
		if attendance.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND attendance.IsMarkedForDeletion = %v", where, *attendance.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_attendance_data as attendance
		` + where + ` ORDER BY attendance.IsMarkedForDeletion ASC, attendance.IsCancelled ASC, attendance.IsReleased ASC, attendance.Event ASC, attendance.Attendance ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToAttendance(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PointTransaction(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PointTransaction {
	var args []interface{}
	where := fmt.Sprintf("WHERE Event = %d", input.Header.Event)

	pointTransactionIDs := ""
	for _, v := range input.Header.PointTransaction {
		pointTransactionIDs = fmt.Sprintf("%s, %d, %d, %d, %d", pointTransactionIDs, v.Sender, v.Receiver, v.PointConditionRecord, v.PointConditionSequentialNumber)
	}

	if len(pointTransactionIDs) != 0 {
		where = fmt.Sprintf("%s\nAND PointTransaction IN ( %s ) ", where, pointTransactionIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_point_transaction_data
		`+where+` ORDER BY IsCancelled ASC, Event ASC, Sender ASC, Receiver ASC, PointConditionRecord ASC, PointConditionSequentialNumber ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPointTransaction(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PointTransactions(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PointTransaction {
	pointTransaction := &dpfm_api_input_reader.PointTransaction{}
	if len(input.Header.PointTransaction) > 0 {
		pointTransaction = &input.Header.PointTransaction[0]
	}
	where := "WHERE 1 = 1"

	if input.Header.Event != 0 {
		where = fmt.Sprintf("WHERE Event = %d", input.Header.Event)
	}

	if pointTransaction != nil {
		if pointTransaction.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND pointTransaction.IsCancelled = %v", where, *pointTransaction.IsCancelled)
		}
	}

	rows, err := c.db.Query(
		`SELECT	*
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_point_transaction_data as pointTransaction
		` + where + ` ORDER BY pointTransaction.IsCancelled ASC, pointTransaction.Event ASC, pointTransaction.Sender ASC, pointTransaction.Receiver ASC, pointTransaction.PointConditionRecord ASC, pointTransaction.PointConditionSequentialNumber ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPointTransaction(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PointConditionElement(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PointConditionElement {
	var args []interface{}
	where := fmt.Sprintf("WHERE Event = %d", input.Header.Event)

	pointConditionElementIDs := ""
	for _, v := range input.Header.PointConditionElement {
		pointConditionElementIDs = fmt.Sprintf("%s, %d, %d", pointConditionElementIDs, v.PointConditionRecord, v.PointConditionSequentialNumber)
	}

	if len(pointConditionElementIDs) != 0 {
		where = fmt.Sprintf("%s\nAND PointConditionElement IN ( %s ) ", where, pointConditionElementIDs[1:])
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_point_condition_element_data
		`+where+` ORDER BY IsMakedForDeletion ASC, IsCancelled ASC, IsReleased ASC, Event ASC, PointConditionRecord ASC, PointConditionSequentialNumber ASC ;`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPointConditionElement(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PointConditionElements(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PointConditionElement {
	pointConditionElement := &dpfm_api_input_reader.PointConditionElement{}
	if len(input.Header.PointConditionElement) > 0 {
		pointConditionElement = &input.Header.PointConditionElement[0]
	}
	where := "WHERE 1 = 1"

	if input.Header.Event != 0 {
		where = fmt.Sprintf("WHERE Event = %d", input.Header.Event)
	}

	if pointConditionElement != nil {
		if pointConditionElement.IsReleased != nil {
			where = fmt.Sprintf("%s\nAND pointConditionElement.IsReleased = %v", where, *pointConditionElement.IsReleased)
		}
	}

	if pointConditionElement != nil {
		if pointConditionElement.IsCancelled != nil {
			where = fmt.Sprintf("%s\nAND pointConditionElement.IsCancelled = %v", where, *pointConditionElement.IsCancelled)
		}
	}

	if pointConditionElement != nil {
		if pointConditionElement.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND pointConditionElement.IsMarkedForDeletion = %v", where, *pointConditionElement.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT	*
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_event_point_condition_element_data as pointConditionElement
		` + where + ` ORDER BY pointConditionElement.IsMarkedForDeletion ASC, pointConditionElement.IsCancelled ASC, pointConditionElement.IsReleased ASC, pointConditionElement.Event ASC, pointConditionElement.PointConditionRecord ASC, pointConditionElement.PointConditionSequentialNumber ASC ;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPointConditionElement(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
