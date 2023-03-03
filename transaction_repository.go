package repository

import (
	"database/sql"
	"fmt"

	"example.com/m/dto"
	"example.com/m/entity"
)

type TransactionRepo interface {
	FetchMachineUtil(machineRequest dto.MachineTransactionRequest) (*[]entity.Machine, error)
}

type transactionRepoImpl struct {
	db *sql.DB
}

type TransactionRepositoryConfig struct {
	DB *sql.DB
}

func NewTransactionRepository(c TransactionRepositoryConfig) TransactionRepo {
	return &transactionRepoImpl{
		db: c.DB,
	}
}

func (t *transactionRepoImpl) FetchMachineUtil(assetReq dto.MachineTransactionRequest) (*[]entity.Machine, error) {
	var machines []entity.Machine

	fmt.Println("Hello Hello")
	fmt.Println(assetReq)
	fmt.Println(assetReq.TransactionDate)

	rows, err := t.db.Query(
		`with query1 as (
			select * from plants p
			left join sites s  
			on p.site_id  = s.site_id 
			where s.site_name = $1
		), query2 as (
			select * from departements d 
			left join query1
			on query1.plant_id = d.plant_id 
			where query1.plant_name = $2
		), query3 as (
			select * from work_centers wc
			left join query2
			on query2.departement_id =  wc.departement_id 
			where query2.departement_name = $3
		), query4 as (
			select * from work_stations ws 
			left join query3
			on query3.work_center_id = ws.work_center_id 
			where query3.work_center_name = $4
		), query5 as (
			select * from assets 
			left join query4
			on query4.workstation_id = assets.workstation_id
			where query4.workstation_name = $5
		), query6 as (
			select * from transactions
			join query5
			on query5.asset_id = transactions.asset_id 
		) select transaction_id, transaction_time, asset_image_path, asset_name from query6
		where date_trunc('day', query6.transaction_time) = $6
		order by asset_name asc, query6.transaction_time asc`,
		assetReq.SiteName,
		assetReq.PlantName,
		assetReq.DepartementName,
		assetReq.WorkCenterName,
		assetReq.WorkstationName,
		assetReq.TransactionDate)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		machine := new(entity.Machine)
		err = rows.Scan(&machine.Asset_Image, &machine.Asset_Name, &machine.TransactionID, &machine.Transaction_Time)
		machines = append(machines, *machine)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return &machines, nil
}
