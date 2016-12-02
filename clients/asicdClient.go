package clients

import (
	"asicdServices"
	"fmt"
	"models/objects"
	"utils/dbutils"
)

const (
	BREAKOUT_MODE_ATTR_OFFSET = 11
)

func (clnt *ASICDClient) PostUpdateProcessing(dbObj, obj objects.ConfigObj, attrSet []bool, dbHdl *dbutils.DBUtil) error {
	if clnt.ClientHdl != nil {
		switch obj.(type) {
		case objects.Port:
			if attrSet[BREAKOUT_MODE_ATTR_OFFSET] {
				objList, err := clnt.ClientHdl.GetAllPortsWithDirtyCache()
				if err != nil {
					fmt.Println("Asicd post processing failed")
					return err
				}
				for _, val := range objList {
					var dbObj objects.Port
					obj := (*asicdServices.Port)(val)
					objects.ConvertThriftToasicdPortObj(obj, &dbObj)
					dbHdl.StoreObjectInDb(dbObj)
				}
			}
		default:
			fmt.Println("Post update processing not implemented for this config object")
		}
	}
	return nil
}
